# Opcodes


## err
- Opcode: 0x00 
- Pops: None
- Pushes: None
- Error. Panic immediately. This is primarily a fencepost against accidental zero bytes getting compiled into programs.

## sha256
- Opcode: 0x01 
- Pops: *... stack*, []byte
- Pushes: []byte
- SHA256 hash of value, yields [32]byte

## keccak256
- Opcode: 0x02 
- Pops: *... stack*, []byte
- Pushes: []byte
- Keccac256 hash of value, yields [32]byte

## sha512_256
- Opcode: 0x03 
- Pops: *... stack*, []byte
- Pushes: []byte
- SHA512_256 hash of value, yields [32]byte

## +
- Opcode: 0x08 
- Pops: *... stack*, {uint64 A}, {uint64 B}
- Pushes: uint64
- A plus B. Panic on overflow.

## -
- Opcode: 0x09 
- Pops: *... stack*, {uint64 A}, {uint64 B}
- Pushes: uint64
- A minus B. Panic if B > A.

## /
- Opcode: 0x0a 
- Pops: *... stack*, {uint64 A}, {uint64 B}
- Pushes: uint64
- A divided by B. Panic if B == 0.

## *
- Opcode: 0x0b 
- Pops: *... stack*, {uint64 A}, {uint64 B}
- Pushes: uint64
- A times B. Panic on overflow.

## <
- Opcode: 0x0c 
- Pops: *... stack*, {uint64 A}, {uint64 B}
- Pushes: uint64
- A less than B => {0 or 1}

## >
- Opcode: 0x0d 
- Pops: *... stack*, {uint64 A}, {uint64 B}
- Pushes: uint64
- A greater than B => {0 or 1}

## <=
- Opcode: 0x0e 
- Pops: *... stack*, {uint64 A}, {uint64 B}
- Pushes: uint64
- A less than or equal to B => {0 or 1}

## >=
- Opcode: 0x0f 
- Pops: *... stack*, {uint64 A}, {uint64 B}
- Pushes: uint64
- A greater than or equal to B => {0 or 1}

## &&
- Opcode: 0x10 
- Pops: *... stack*, {uint64 A}, {uint64 B}
- Pushes: uint64
- A is not zero and B is not zero => {0 or 1}

## ||
- Opcode: 0x11 
- Pops: *... stack*, {uint64 A}, {uint64 B}
- Pushes: uint64
- A is not zero or B is not zero => {0 or 1}

## ==
- Opcode: 0x12 
- Pops: *... stack*, {any A}, {any B}
- Pushes: uint64
- A is equal to B => {0 or 1}

## !=
- Opcode: 0x13 
- Pops: *... stack*, {any A}, {any B}
- Pushes: uint64
- A is not equal to B => {0 or 1}

## !
- Opcode: 0x14 
- Pops: *... stack*, uint64
- Pushes: uint64
- X == 0 yields 1; else 0

## len
- Opcode: 0x15 
- Pops: *... stack*, []byte
- Pushes: uint64
- yields length of byte value

## btoi
- Opcode: 0x17 
- Pops: *... stack*, []byte
- Pushes: uint64
- converts bytes as big endian to uint64

## %
- Opcode: 0x18 
- Pops: *... stack*, {uint64 A}, {uint64 B}
- Pushes: uint64
- A modulo B. Panic if B == 0.

## |
- Opcode: 0x19 
- Pops: *... stack*, {uint64 A}, {uint64 B}
- Pushes: uint64
- A bitwise-or B

## &
- Opcode: 0x1a 
- Pops: *... stack*, {uint64 A}, {uint64 B}
- Pushes: uint64
- A bitwise-and B

## ^
- Opcode: 0x1b 
- Pops: *... stack*, {uint64 A}, {uint64 B}
- Pushes: uint64
- A bitwise-xor B

## ~
- Opcode: 0x1c 
- Pops: *... stack*, uint64
- Pushes: uint64
- bitwise invert value

## intcblock
- Opcode: 0x20 {varuint length} [{varuint value}, ...]
- Pops: None
- Pushes: None
- load block of uint64 constants

`intcblock` loads following program bytes into an array of integer constants in the evaluator. These integer constants can be referred to by `intc` and `intc_*` which will push the value onto the stack.

## intc
- Opcode: 0x21 {uint8 int constant index}
- Pops: None
- Pushes: uint64
- push value from uint64 constants to stack by index into constants

## intc_0
- Opcode: 0x22 
- Pops: None
- Pushes: uint64
- push uint64 constant 0 to stack

## intc_1
- Opcode: 0x23 
- Pops: None
- Pushes: uint64
- push uint64 constant 1 to stack

## intc_2
- Opcode: 0x24 
- Pops: None
- Pushes: uint64
- push uint64 constant 2 to stack

## intc_3
- Opcode: 0x25 
- Pops: None
- Pushes: uint64
- push uint64 constant 3 to stack

## bytecblock
- Opcode: 0x26 {varuint length} [({varuint value length} bytes), ...]
- Pops: None
- Pushes: None
- load block of byte-array constants

`bytecblock` loads the following program bytes into an array of byte string constants in the evaluator. These constants can be referred to by `bytec` and `bytec_*` which will push the value onto the stack.

## bytec
- Opcode: 0x27 {uint8 byte constant index}
- Pops: None
- Pushes: []byte
- push bytes constant to stack by index into constants

## bytec_0
- Opcode: 0x28 
- Pops: None
- Pushes: []byte
- push bytes constant 0 to stack

## bytec_1
- Opcode: 0x29 
- Pops: None
- Pushes: []byte
- push bytes constant 1 to stack

## bytec_2
- Opcode: 0x2a 
- Pops: None
- Pushes: []byte
- push bytes constant 2 to stack

## bytec_3
- Opcode: 0x2b 
- Pops: None
- Pushes: []byte
- push bytes constant 3 to stack

## arg
- Opcode: 0x2c {uint8 arg index N}
- Pops: None
- Pushes: []byte
- push LogicSig.Args[N] value to stack by index

## arg_0
- Opcode: 0x2d 
- Pops: None
- Pushes: []byte
- push LogicSig.Args[0] to stack

## arg_1
- Opcode: 0x2e 
- Pops: None
- Pushes: []byte
- push LogicSig.Args[1] to stack

## arg_2
- Opcode: 0x2f 
- Pops: None
- Pushes: []byte
- push LogicSig.Args[2] to stack

## arg_3
- Opcode: 0x30 
- Pops: None
- Pushes: []byte
- push LogicSig.Args[3] to stack

## txn
- Opcode: 0x31 {uint8 transaction field index}
- Pops: None
- Pushes: any
- push field from current transaction to stack

`txn` Fields:

| Index | Name | Type |
| --- | --- | --- |
| 0 | Sender | []byte |
| 1 | Fee | uint64 |
| 2 | FirstValid | uint64 |
| 3 | LastValid | uint64 |
| 4 | Note | []byte |
| 5 | Receiver | []byte |
| 6 | Amount | uint64 |
| 7 | CloseRemainderTo | []byte |
| 8 | VotePK | []byte |
| 9 | SelectionPK | []byte |
| 10 | VoteFirst | uint64 |
| 11 | VoteLast | uint64 |
| 12 | VoteKeyDilution | uint64 |


## global
- Opcode: 0x32 {uint8 global field index}
- Pops: None
- Pushes: any
- push value from globals to stack

`global` Fields:

| Index | Name | Type |
| --- | --- | --- |
| 0 | Round | uint64 |
| 1 | MinTxnFee | uint64 |
| 2 | MinBalance | uint64 |
| 3 | MaxTxnLife | uint64 |
| 4 | TimeStamp | uint64 |


## bnz
- Opcode: 0x40 {0..0x7fff forward branch offset, big endian}
- Pops: *... stack*, uint64
- Pushes: None
- branch if value is not zero

for a bnz instruction at `pc`, if the last element of the stack is not zero then branch to instruction at `pc + 3 + N`, else procede to next instruction at `pc + 3`

## pop
- Opcode: 0x48 
- Pops: *... stack*, any
- Pushes: None
- discard value from stack

## dup
- Opcode: 0x49 
- Pops: None
- Pushes: any
- duplicate last value on stack