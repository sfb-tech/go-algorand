// Copyright (C) 2019-2020 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"io"
	"log"

	"github.com/algorand/go-algorand/config"
	"github.com/algorand/go-algorand/data/basics"
	"github.com/algorand/go-algorand/data/transactions"
	"github.com/algorand/go-algorand/data/transactions/logic"
	"github.com/algorand/go-algorand/protocol"
)

func protoFromString(protoString string) (name string, proto config.ConsensusParams, err error) {
	if len(protoString) == 0 || protoString == "current" {
		name = string(protocol.ConsensusCurrentVersion)
		proto = config.Consensus[protocol.ConsensusCurrentVersion]
	} else {
		var ok bool
		proto, ok = config.Consensus[protocol.ConsensusVersion(protoString)]
		if !ok {
			err = fmt.Errorf("Unknown protocol %s", protoString)
			return
		}
		name = protoString
	}

	return
}

// txnGroupFromParams validates DebugParams.TxnBlob
// DebugParams.TxnBlob parsed as JSON object, JSON array or MessagePack array of transactions.SignedTxn.
// The function returns ready to use txnGroup or an error
func txnGroupFromParams(dp *DebugParams) (txnGroup []transactions.SignedTxn, err error) {
	if len(dp.TxnBlob) == 0 {
		txnGroup = append(txnGroup, transactions.SignedTxn{})
		return
	}

	var data []byte = dp.TxnBlob

	// 1. Attempt json - a single transaction
	var txn transactions.SignedTxn
	err = protocol.DecodeJSON(data, &txn)
	if err == nil {
		txnGroup = append(txnGroup, txn)
		return
	}

	// 2. Attempt json - array of transactions
	err = protocol.DecodeJSON(data, &txnGroup)
	if err == nil {
		return
	}

	// 3. Attempt msgp - array of transactions
	dec := protocol.NewDecoderBytes(data)
	for {
		var txn transactions.SignedTxn
		err = dec.Decode(&txn)
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			break
		}
		txnGroup = append(txnGroup, txn)
	}

	return
}

// balanceRecordsFromParams attempts to parse DebugParams.BalanceBlob as
// JSON object, JSON array or MessagePack array of basics.BalanceRecord
func balanceRecordsFromParams(dp *DebugParams) (records []basics.BalanceRecord, err error) {
	if len(dp.BalanceBlob) == 0 {
		return
	}

	var data []byte = dp.BalanceBlob

	// 1. Attempt json - a single record
	var record basics.BalanceRecord
	err = protocol.DecodeJSON(data, &record)
	if err == nil {
		records = append(records, record)
		return
	}

	// 2. Attempt json - a array of records
	err = protocol.DecodeJSON(data, &records)
	if err == nil {
		return
	}

	// 2. Attempt msgp - a array of records
	dec := protocol.NewDecoderBytes(data)
	for {
		var record basics.BalanceRecord
		err = dec.Decode(&record)
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			break
		}
		records = append(records, record)
	}

	return
}

// evaluation is a description of a single debugger run
type evaluation struct {
	program      []byte
	source       string
	offsetToLine map[int]int
	name         string
	groupIndex   int
	eval         func(program []byte, ep logic.EvalParams) (bool, error)
	ledger       *localLedger
}

// LocalRunner runs local eval
type LocalRunner struct {
	debugger  *Debugger
	proto     config.ConsensusParams
	protoName string
	txnGroup  []transactions.SignedTxn
	runs      []evaluation
}

// MakeLocalRunner creates LocalRunner
func MakeLocalRunner(debugger *Debugger) *LocalRunner {
	r := new(LocalRunner)
	r.debugger = debugger
	return r
}

// Setup validates input params
func (r *LocalRunner) Setup(dp *DebugParams) (err error) {
	r.protoName, r.proto, err = protoFromString(dp.Proto)
	if err != nil {
		return
	}
	log.Printf("Using proto: %s", r.protoName)

	r.txnGroup, err = txnGroupFromParams(dp)
	if err != nil {
		return
	}

	records, err := balanceRecordsFromParams(dp)
	if err != nil {
		return
	}

	balances := make(map[basics.Address]basics.AccountData)
	for _, record := range records {
		balances[record.Addr] = record.AccountData
	}

	// if program(s) specified then run from it
	if len(dp.ProgramBlobs) > 0 {
		if len(r.txnGroup) == 1 && dp.GroupIndex != 0 {
			err = fmt.Errorf("invalid group index %d for a single transaction", dp.GroupIndex)
			return
		}
		if len(r.txnGroup) > 0 && dp.GroupIndex >= len(r.txnGroup) {
			err = fmt.Errorf("invalid group index %d for a txn in a transaction group of %d", dp.GroupIndex, len(r.txnGroup))
			return
		}
		groupIndex := dp.GroupIndex
		ledger := &localLedger{
			round:      dp.Round,
			balances:   balances,
			txnGroup:   r.txnGroup,
			groupIndex: groupIndex,
		}

		var eval func(program []byte, ep logic.EvalParams) (bool, error)
		switch dp.RunMode {
		case "signature":
			eval = logic.Eval
		case "application":
			eval = func(program []byte, ep logic.EvalParams) (bool, error) {
				pass, _, err := logic.EvalStateful(program, ep)
				return pass, err
			}
		default:
			err = fmt.Errorf("unknown run mode")
			return
		}

		r.runs = make([]evaluation, len(dp.ProgramBlobs))
		for i, data := range dp.ProgramBlobs {
			r.runs[i].program = data
			if IsTextFile(data) {
				source := string(data)
				program, offsets, err := logic.AssembleStringWithVersionEx(source, r.proto.LogicSigVersion)
				if err != nil {
					return err
				}
				r.runs[i].program = program
				r.runs[i].offsetToLine = offsets
				r.runs[i].source = source
			}
			r.runs[i].groupIndex = groupIndex
			r.runs[i].ledger = ledger
			r.runs[i].eval = eval
			r.runs[i].name = dp.ProgramNames[i]
		}
		return nil
	}

	r.runs = nil
	// otherwise, if no program(s) set, check transactions for TEAL programs
	for gi, stxn := range r.txnGroup {
		// make a new ledger per possible execution since it requires a current group index
		ledger := localLedger{
			round:      dp.Round,
			balances:   balances,
			txnGroup:   r.txnGroup,
			groupIndex: gi,
		}
		if len(stxn.Lsig.Logic) > 0 {
			run := evaluation{
				program:    stxn.Lsig.Logic,
				groupIndex: gi,
				eval:       logic.Eval,
				ledger:     &ledger,
			}
			r.runs = append(r.runs, run)
		} else if stxn.Txn.Type == protocol.ApplicationCallTx {
			eval := func(program []byte, ep logic.EvalParams) (bool, error) {
				pass, _, err := logic.EvalStateful(program, ep)
				return pass, err
			}
			appIdx := stxn.Txn.ApplicationID
			if appIdx == 0 { // app create, use ApprovalProgram from the transaction
				if len(stxn.Txn.ApprovalProgram) > 0 {
					run := evaluation{
						program:    stxn.Txn.ApprovalProgram,
						groupIndex: gi,
						eval:       eval,
						ledger:     &ledger,
					}
					r.runs = append(r.runs, run)
				}
			} else {
				// attempt to find this appIdx in balance records provided
				// and error if it is not there
				found := false
				for _, rec := range records {
					for a, ap := range rec.AppParams {
						if a == appIdx {
							var program []byte
							if stxn.Txn.OnCompletion == transactions.ClearStateOC {
								program = ap.ClearStateProgram
							} else {
								program = ap.ApprovalProgram
							}
							if len(program) == 0 {
								err = fmt.Errorf("empty program found for app idx %d", appIdx)
								return
							}
							run := evaluation{
								program:    program,
								groupIndex: gi,
								eval:       eval,
								ledger:     &ledger,
							}
							r.runs = append(r.runs, run)
							found = true
							break
						}
					}
				}
				if !found {
					err = fmt.Errorf("no program found for app idx %d", appIdx)
					return
				}
			}
		}
	}

	if len(r.runs) == 0 {
		err = fmt.Errorf("no programs found in transactions")
	}

	return
}

// RunAll runs all the programs
func (r *LocalRunner) RunAll() error {
	if len(r.runs) < 1 {
		return fmt.Errorf("no program to debug")
	}

	for _, run := range r.runs {
		r.debugger.SaveProgram(run.name, run.program, run.source, run.offsetToLine)

		ep := logic.EvalParams{
			Proto:      &r.proto,
			Debugger:   r.debugger,
			Txn:        &r.txnGroup[groupIndex],
			TxnGroup:   r.txnGroup,
			GroupIndex: run.groupIndex,
			Ledger:     run.ledger,
		}

		_, err := run.eval(run.program, ep)
		if err != nil {
			return err
		}
	}
	return nil
}

// Run starts the first program in list
func (r *LocalRunner) Run() (bool, error) {
	if len(r.runs) < 1 {
		return false, fmt.Errorf("no program to debug")
	}

	run := r.runs[0]

	ep := logic.EvalParams{
		Proto:      &r.proto,
		Txn:        &r.txnGroup[groupIndex],
		TxnGroup:   r.txnGroup,
		GroupIndex: run.groupIndex,
		Ledger:     run.ledger,
	}

	// Workaround for Go's nil/empty interfaces nil check after nil assignment, i.e.
	// r.debugger = nil
	// ep.Debugger = r.debugger
	// if ep.Debugger != nil // FALSE
	if r.debugger != nil {
		r.debugger.SaveProgram(run.name, run.program, run.source, run.offsetToLine)
		ep.Debugger = r.debugger
	}

	return run.eval(run.program, ep)
}
