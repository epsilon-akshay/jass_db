package btreeindex

import "encoding/binary"

const LeafNode = uint16(1)
const InternalNode = uint16(0)

type BPlusTreeNode struct {
	data []byte
}

func (node BPlusTreeNode) Type() uint16 {
	return binary.BigEndian.Uint16(node.data[0:2])
}

func (node BPlusTreeNode) NumberOfElem() uint16 {
	return binary.BigEndian.Uint16(node.data[2:4])
}

func (node BPlusTreeNode) SetType(typeOfNode uint16) {
	binary.BigEndian.PutUint16(node.data[0:2], typeOfNode)
}

func (node BPlusTreeNode) SetNumberOfElem(numberOfElem uint16) {
	binary.BigEndian.PutUint16(node.data[2:4], numberOfElem)
}

//might need to make it to get til lend
func (node BPlusTreeNode) ChildPointer(idx uint16) uint64 {
	idxPointer := node.data[TypeOfNodeSize+NumberOfNodeElementsSize+idx*MinPointerToNodeSize : TypeOfNodeSize+NumberOfNodeElementsSize+idx*MinPointerToNodeSize+9]
	return binary.BigEndian.Uint64(idxPointer)
}

func (node BPlusTreeNode) SetChildPointer(idx uint16, val uint64) {
	idxPointer := node.data[TypeOfNodeSize+NumberOfNodeElementsSize+idx*MinPointerToNodeSize : TypeOfNodeSize+NumberOfNodeElementsSize+idx*MinPointerToNodeSize+9]
	binary.BigEndian.PutUint64(idxPointer, val)
}

func (node BPlusTreeNode) OffsetPos(idx uint16) uint16 {
	if idx == 0 {
		return 0
	}
	pos := TypeOfNodeSize + NumberOfNodeElementsSize + node.NumberOfElem()*MinPointerToNodeSize + (idx-1)*2
	return binary.BigEndian.Uint16(node.data[pos : pos+2])
}

func (node BPlusTreeNode) SetOffsetPos(idx uint16, offsetAdd uint16) {
	if idx == 0 {
		return
	}
	pos := TypeOfNodeSize + NumberOfNodeElementsSize + node.NumberOfElem()*MinPointerToNodeSize + (idx-1)*2
	binary.BigEndian.PutUint16(node.data[pos:pos+2], offsetAdd)
}
