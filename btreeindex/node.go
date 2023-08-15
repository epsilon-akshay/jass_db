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
