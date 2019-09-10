# Transaction Execution Approval Language (TEAL)

TEAL is a bytecode based stack language that executes inside Algorand transactions to check the parameters of the transaction and approve the transaction as if by a signature.

TEAL programs should be short, at most 1000 bytes including all constants and operations, and run fast as they are run in-line along with signature checking, transaction balance rule checking, and other checks during block assembly and validation.

## The Stack

The stack starts empty and contains values of either uint64 or bytes. (`bytes` implemented in Go as a []byte slice)

The maximum stack depth is currently 1000.

## Constants

Constants are loaded into the environment into storage separate from the stack. They can then be pushed onto the stack by referring to the type and index. This makes for efficient re-use of byte constants used for account addresses, etc.

The assembler will hide most of this, allowing simple use of `int 1234` and `byte 0xcafed00d`. These constants will automatically get assembled into int and byte pages of constants, de-duplicated, and operations to load them from constant storage space inserted.

Constants are loaded into the environment by two opcodes, `intcblock` and `bytecblock`. Both of these use [proto-buf style variable length unsigned int](https://developers.google.com/protocol-buffers/docs/encoding#varint). The `intcblock` opcode is followed by a varuint specifying the length of the array and then that number of varuint. The `bytecblock` opcode is followed by a varuint array length then that number of pairs of (varuint, bytes) length prefixed byte strings. This should efficiently load 32 and 64 byte constants which will be common as addresses, hashes, and signatures.

Constants are pushed onto the stack by `intc`, `intc_[0123]`, `bytec`, and `bytec_[0123]`. The assembler will typically handle converting `int N` or `byte N` into the appropriate constant-offset opcode or opcode followed by index byte.

## Operations

Most operations work with only one type of argument, uint64 or bytes, and panic if the wrong type value is on the stack.

This summary is supplemented by more detail in the [opcodes document](TEAL_opcodes.md).

Some operations 'panic' and immediately end execution of the program. A transaction checked a program that panics is not valid.

### Arithmetic

For one-argument ops, `X` is the last element on the stack, which is typically replaced by a new value.

For two-argument ops, `A` is the previous element on the stack and `B` is the last element on the stack. These typically result in popping A and B from the stack and pushing the result.

@@ Arithmetic.md @@

### Loading Values

Opcodes for getting data onto the stack.

Some of these have immediate data in the byte or bytes after the opcode.

@@ Loading_Values.md @@

**Transaction Fields**

@@ txn_fields.md @@

**Global Fields**

@@ global_fields.md @@


### Flow Control

@@ Flow_Control.md @@

# Assembler Syntax

The assembler parses line by line. Ops that just use the stack appear on a line by themselves. Ops that take arguments are the op and then whitespace and then any argument or arguments.

"`//`" prefixes a line comment.

## Constants and Pseudo-Ops

A few pseudo-ops simplify writing code. `int` and `byte` and `addr` followed by a constant record the constant to a `intcblock` or `bytecblock` at the beginning of code and insert an `intc` or `bytec` reference where the instruction appears to load that value. `addr` parses an Algorand account address base32 and converts it to a regular bytes constant.

`byte` constants are:
```
byte base64 AAAA...
byte b64 AAAA...
byte base64(AAAA...)
byte b64(AAAA...)
byte base32 AAAA...
byte b32 AAAA...
byte base32(AAAA...)
byte b32(AAAA...)
byte 0x0123456789abcdef...
```

`int` constants may be `0x` prefixed for hex, `0` prefixed for octal, or decimal numbers.

`intcblock` may be explictily assembled. It will conflict with the assembler gathering `int` pseudo-ops into a `intcblock` program prefix, but may be used in code only has explicit `intc` references. `intcblock` should be followed by space separated int constants all on one line.

`bytecblock` may be explicitly assembled. It will conflict with the assembler if there are any `byte` pseudo-ops but may be used if only explicit `bytec` references are used. `bytecblock` should be followed with byte contants all on one line, either 'encoding value' pairs (`b64 AAA...`) or 0x prefix or function-style values (`base64(...)`).

## Labels and Branches

A label is defined by any string not some other op or keyword and ending in ':'. A label can be an argument (without the trailing ':') to a branch instruction.

Example:
```
int 1
bnz safe
err
safe:
pop
```

## Encoding and Versioning

A program starts with a varuint declaring the version of the compiled code. Any addition, removal, or change of opcode behavior increments the version. For the most part opcode behavior should not change, addition will be infrequent (not likely more often than every three months and less often as the language matures), and removal should be very rare.

For version 1, subsequent bytes after the varuint are program opcode bytes. Future versions could put other metadata following the version identifier.