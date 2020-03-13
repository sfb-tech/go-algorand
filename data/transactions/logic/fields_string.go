// Code generated by "stringer -type=TxnField,GlobalField,AssetParamsField,AssetHoldingField -output=fields_string.go"; DO NOT EDIT.

package logic

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Sender-0]
	_ = x[Fee-1]
	_ = x[FirstValid-2]
	_ = x[FirstValidTime-3]
	_ = x[LastValid-4]
	_ = x[Note-5]
	_ = x[Lease-6]
	_ = x[Receiver-7]
	_ = x[Amount-8]
	_ = x[CloseRemainderTo-9]
	_ = x[VotePK-10]
	_ = x[SelectionPK-11]
	_ = x[VoteFirst-12]
	_ = x[VoteLast-13]
	_ = x[VoteKeyDilution-14]
	_ = x[Type-15]
	_ = x[TypeEnum-16]
	_ = x[XferAsset-17]
	_ = x[AssetAmount-18]
	_ = x[AssetSender-19]
	_ = x[AssetReceiver-20]
	_ = x[AssetCloseTo-21]
	_ = x[GroupIndex-22]
	_ = x[TxID-23]
	_ = x[OnCompletion-24]
	_ = x[ApplicationArgs-25]
	_ = x[Accounts-26]
	_ = x[invalidTxnField-27]
}

const _TxnField_name = "SenderFeeFirstValidFirstValidTimeLastValidNoteLeaseReceiverAmountCloseRemainderToVotePKSelectionPKVoteFirstVoteLastVoteKeyDilutionTypeTypeEnumXferAssetAssetAmountAssetSenderAssetReceiverAssetCloseToGroupIndexTxIDOnCompletionApplicationArgsAccountsinvalidTxnField"

var _TxnField_index = [...]uint16{0, 6, 9, 19, 33, 42, 46, 51, 59, 65, 81, 87, 98, 107, 115, 130, 134, 142, 151, 162, 173, 186, 198, 208, 212, 224, 239, 247, 262}

func (i TxnField) String() string {
	if i < 0 || i >= TxnField(len(_TxnField_index)-1) {
		return "TxnField(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TxnField_name[_TxnField_index[i]:_TxnField_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MinTxnFee-0]
	_ = x[MinBalance-1]
	_ = x[MaxTxnLife-2]
	_ = x[ZeroAddress-3]
	_ = x[GroupSize-4]
	_ = x[invalidGlobalField-5]
}

const _GlobalField_name = "MinTxnFeeMinBalanceMaxTxnLifeZeroAddressGroupSizeinvalidGlobalField"

var _GlobalField_index = [...]uint8{0, 9, 19, 29, 40, 49, 67}

func (i GlobalField) String() string {
	if i < 0 || i >= GlobalField(len(_GlobalField_index)-1) {
		return "GlobalField(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _GlobalField_name[_GlobalField_index[i]:_GlobalField_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AssetParamsTotal-0]
	_ = x[AssetParamsDecimals-1]
	_ = x[AssetParamsDefaultFrozen-2]
	_ = x[AssetParamsUnitName-3]
	_ = x[AssetParamsAssetName-4]
	_ = x[AssetParamsURL-5]
	_ = x[AssetParamsMetadataHash-6]
	_ = x[AssetParamsManager-7]
	_ = x[AssetParamsReserve-8]
	_ = x[AssetParamsFreeze-9]
	_ = x[AssetParamsClawback-10]
	_ = x[invalidAssetParamsField-11]
}

const _AssetParamsField_name = "AssetParamsTotalAssetParamsDecimalsAssetParamsDefaultFrozenAssetParamsUnitNameAssetParamsAssetNameAssetParamsURLAssetParamsMetadataHashAssetParamsManagerAssetParamsReserveAssetParamsFreezeAssetParamsClawbackinvalidAssetParamsField"

var _AssetParamsField_index = [...]uint8{0, 16, 35, 59, 78, 98, 112, 135, 153, 171, 188, 207, 230}

func (i AssetParamsField) String() string {
	if i < 0 || i >= AssetParamsField(len(_AssetParamsField_index)-1) {
		return "AssetParamsField(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AssetParamsField_name[_AssetParamsField_index[i]:_AssetParamsField_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AssetHoldingAmount-0]
	_ = x[AssetHoldingFrozen-1]
	_ = x[invalidAssetHoldingField-2]
}

const _AssetHoldingField_name = "AssetHoldingAmountAssetHoldingFrozeninvalidAssetHoldingField"

var _AssetHoldingField_index = [...]uint8{0, 18, 36, 60}

func (i AssetHoldingField) String() string {
	if i < 0 || i >= AssetHoldingField(len(_AssetHoldingField_index)-1) {
		return "AssetHoldingField(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AssetHoldingField_name[_AssetHoldingField_index[i]:_AssetHoldingField_index[i+1]]
}