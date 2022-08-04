package bst

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type BSTContent struct {
	Key   int
	Value interface{}
}

func (n *BSTContent) String() string {
	return fmt.Sprintf("key: %d, value: %v", n.Key, n.Value)
}

type NodeBST struct {
	*BSTContent

	left  *NodeBST
	right *NodeBST
}

func (n *NodeBST) Left() *NodeBST {
	return n.left
}

func (n *NodeBST) Right() *NodeBST {
	return n.right
}

type HeightMap = map[int]int

type bstMeta struct {
	heightMap HeightMap
}

type BST struct {
	root *NodeBST
	*bstMeta
}

func NewNode(k int, v interface{}, l, r *NodeBST) *NodeBST {
	return &NodeBST{
		BSTContent: &BSTContent{
			Key:   k,
			Value: v,
		},
		left:  l,
		right: r,
	}
}

func NewBST(node *NodeBST) *BST {
	meta := &bstMeta{make(map[int]int)}
	return &BST{node, meta}
}

func (b *BST) Root() *NodeBST {
	return b.root
}

func (b *BST) String() string {
	lst := b.Inorder()
	var sb strings.Builder
	for _, v := range lst {
		sb.WriteString(v.String())
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (b *BST) Put(k int, v interface{}) {
	b.root = b.put(k, v, b.root)
}

func (b *BST) Get(k int) (interface{}, error) {
	return b.get(k, b.root)
}

func (b *BST) GetNode(k int) (*NodeBST, error) {
	return b.getNode(k, b.root)
}

func (b *BST) get(k int, curNode *NodeBST) (interface{}, error) {
	node, err := b.getNode(k, curNode)
	if err != nil {
		return nil, err
	}
	return node.Value, nil
}

func (b *BST) getNode(k int, curNode *NodeBST) (*NodeBST, error) {
	if curNode == nil {
		return nil, errors.New("key not found in tree")

	}
	if curNode.Key == k {
		return curNode, nil
	}
	if k < curNode.Key {
		return b.getNode(k, curNode.left)
	}
	if k > curNode.Key {
		return b.getNode(k, curNode.right)
	}
	return nil, errors.New("key not found in tree")
}

func (b *BST) Inorder() []*BSTContent {
	return b.inorder(b.root, []*BSTContent{})
}

func (b *BST) Validate() bool {
	orderedNodes := b.Inorder()
	isValid := true
	for i := 1; i < len(orderedNodes); i++ {
		if orderedNodes[i-1].Key > orderedNodes[i].Key {
			return false
		}
	}
	return isValid
}

func (b *BST) inorder(node *NodeBST, lst []*BSTContent) []*BSTContent {
	if node == nil {
		return lst
	}
	lst = b.inorder(node.left, lst)
	lst = append(lst, node.BSTContent)
	lst = b.inorder(node.right, lst)
	return lst
}

func (b *BST) put(k int, v interface{}, node *NodeBST) *NodeBST {

	if node == nil {
		return NewNode(k, v, nil, nil)
	}

	if k == node.Key {
		node.Value = v
	}
	if k > node.Key {
		node.right = b.put(k, v, node.right)
	}
	if k < node.Key {
		node.left = b.put(k, v, node.left)
	}

	return node

}

func (b *BST) Height() int {
	return b.height(b.root, 0)
}

func (b *BST) HeightMap() (HeightMap, error) {
	if len(b.bstMeta.heightMap) == 0 {
		return nil, errors.New("please run the Height function to populate the latest height map first")
	}
	return b.bstMeta.heightMap, nil
}

func (b *BST) height(curNode *NodeBST, curHeight int) int {
	if curNode == nil {
		return curHeight
	}

	lh := b.height(curNode.left, curHeight+1)
	rh := b.height(curNode.right, curHeight+1)
	// treeH is determined at the bottom of the tree and then travels up the recursive stack
	treeH := max(lh, rh)
	nodeH := treeH - curHeight
	b.heightMap[curNode.Key] = nodeH
	return treeH
}

func (b *BST) CheckBalanced() (bool, error) {
	h := b.Height()
	if h == 0 {
		return true, nil
	}

	heightMap, err := b.HeightMap()
	if err != nil {
		return false, err
	}
	return b.checkBalanced(heightMap, b.root), nil

}
func (b *BST) checkBalanced(hm HeightMap, curNode *NodeBST) bool {
	var lh int
	var rh int

	if curNode == nil {
		return true
	}

	if curNode.left != nil {
		lh = hm[curNode.left.Key]
	}
	if curNode.right != nil {
		rh = hm[curNode.right.Key]
	}

	delta := math.Abs(float64(lh) - float64(rh))
	return delta < 2 && b.checkBalanced(hm, curNode.left) && b.checkBalanced(hm, curNode.right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
