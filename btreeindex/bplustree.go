package btreeindex

import (
	"bytes"
	"encoding/binary"
)

type btree struct {
	root uint64
}

func NewBTree() btree {
	return btree{}
}

func (bt btree) GetNodeFromDiskPage(pointer uint64) BPlusTreeNode {
	return BPlusTreeNode{}
}

func (bt btree) NewNode(node BPlusTreeNode) uint64 {
	return 0
}

func (bt btree) GetMaxKeyValueSize() int {
	minKeyValue := PageSize - TypeOfNodeSize + NumberOfNodeElementsSize + KeyLengthSize + ValueLengthSize + MinPointerToNodeSize + OffsetMinSize
	return minKeyValue
}

func FindFirstKeyLessThanOrEqualInNode(node BPlusTreeNode, key []byte) (idx uint16, isExact bool) {
	var i = 1
	for i = 1; i < int(node.NumberOfElem()); i++ {
		uintIdx := uint16(i)
		if bytes.Compare(node.Key(uintIdx), key) == 0 {
			return node.GetKV(uintIdx), true
		}
		if bytes.Compare(node.Key(uintIdx), key) == -1 {
			return node.GetKV(uintIdx), false
		}
	}
	return node.GetKV(uint16(0)), false
}

//creates a new node and copies all the elements from the first node
func CreateCopyOfLeafWithNewKVValue(oldNode BPlusTreeNode, newNode BPlusTreeNode, idxToBeInsertedAt uint16, key []byte, value []byte) {
	newTotalElem := oldNode.NumberOfElem() + 1

	newNode = BPlusTreeNode{}
	newNode.SetType(LeafNode)
	newNode.SetNumberOfElem(newTotalElem)

	for i := uint16(0); i < idxToBeInsertedAt; i++ {
		newNode.SetChildPointer(i, oldNode.ChildPointer(i))
	}

	dstBegin := newNode.OffsetPos(0)
	srcBegin := oldNode.OffsetPos(0)

	for i := uint16(1); i <= idxToBeInsertedAt; i++ {
		//TODO: can be replaced with get offset at idx and then using it
		offset := dstBegin + oldNode.OffsetPos(i) - srcBegin
		newNode.SetOffsetPos(i, offset)
	}

	begin := newNode.GetKV(0)
	end := oldNode.GetKV(idxToBeInsertedAt)
	copy(newNode.data[newNode.GetKV(0):], oldNode.data[begin:end])

	newNode.SetChildPointer(idxToBeInsertedAt, 0)
	pos := newNode.GetKV(idxToBeInsertedAt)
	binary.LittleEndian.PutUint16(newNode.data[pos+0:], uint16(len(key)))
	binary.LittleEndian.PutUint16(newNode.data[pos+2:], uint16(len(value)))
	copy(newNode.data[pos+4:], key)
	copy(newNode.data[pos+4+uint16(len(key)):], value)
	// the offset of the next key
	newNode.SetOffsetPos(idxToBeInsertedAt+1, newNode.OffsetPos(idxToBeInsertedAt)+4+uint16((len(key)+len(value))))

	remaining := oldNode.NumberOfElem() - idxToBeInsertedAt
	destNew := idxToBeInsertedAt + 1
	srcOld := idxToBeInsertedAt
	for i := uint16(0); i < remaining; i++ {
		newNode.SetChildPointer(destNew+i, oldNode.ChildPointer(srcOld+i))
	}

	dstBegin = newNode.OffsetPos(destNew)
	srcBegin = oldNode.OffsetPos(srcOld)

	for i := uint16(1); i <= idxToBeInsertedAt; i++ {
		//TODO: can be replaced with get offset at idx and then using it
		offset := dstBegin + oldNode.OffsetPos(srcOld+i) - srcBegin
		newNode.SetOffsetPos(destNew+i, offset)
	}

	begin = newNode.GetKV(srcOld)
	end = oldNode.GetKV(srcOld + remaining)
	copy(newNode.data[newNode.GetKV(destNew):], oldNode.data[begin:end])

}
