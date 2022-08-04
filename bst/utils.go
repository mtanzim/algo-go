package bst

func TreeFromNodes(treeEntries []BSTContent) *BST {

	if len(treeEntries) == 0 {
		return NewBST(nil)
	}

	var bst *BST
	for i, entries := range treeEntries {
		if i == 0 {
			node := NewNode(entries.Key, entries.Value, nil, nil)
			bst = NewBST(node)
			continue
		}
		bst.Put(entries.Key, entries.Value)
	}
	return bst
}
