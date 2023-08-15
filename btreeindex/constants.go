package btreeindex

const PageSize = 4096
const TypeOfNodeSize = 2
const NumberOfNodeElementsSize = 2
const KeyLengthSize = 2
const ValueLengthSize = 2
const MinPointerToNodeSize = 8
const OffsetMinSize = 2

const minKeyPLusValueSize = TypeOfNodeSize + NumberOfNodeElementsSize + KeyLengthSize + ValueLengthSize + MinPointerToNodeSize + OffsetMinSize
