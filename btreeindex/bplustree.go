package btreeindex

type btree struct {
	root uint64
}

func NewBTree() btree {
	return btree{}
}

func (bt btree) getNodeFromDiskPage(pointer uint64) bPlusTreeNode {
	return bPlusTreeNode{}
}

func (bt btree) newNode(node bPlusTreeNode) uint64 {
	return 0
}

//test function
func (bt btree) GetMaxKeyValueSize() int {
	minKeyValue := PageSize - TypeOfNodeSize + NumberOfNodeElementsSize + KeyLengthSize + ValueLengthSize + MinPointerToNodeSize + OffsetMinSize
	return minKeyValue
}

// this is a project used for learning for me and the community hence we try to make sure to fit 1 kv in a single page.
