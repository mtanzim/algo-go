package bst

import (
	"reflect"
	"testing"
)

func TestNewNode(t *testing.T) {
	type args struct {
		k int
		v interface{}
		l *NodeBST
		r *NodeBST
	}
	tests := []struct {
		name string
		args args
		want *NodeBST
	}{
		{
			name: "init",
			args: args{k: 1, v: "hello", l: nil, r: nil},
			want: &NodeBST{BSTContent: &BSTContent{Key: 1, Value: "hello"}, left: nil, right: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.k, tt.args.v, tt.args.l, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBST(t *testing.T) {

	type args struct {
		node *NodeBST
	}

	node := NewNode(1, "jello", nil, nil)

	tests := []struct {
		name string
		args args
		want *BST
	}{
		{
			name: "init bst",
			args: args{node},
			want: &BST{root: node, bstMeta: &bstMeta{make(map[int]int)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBST(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func arrangeTree() (*BST, []BSTContent) {
	treeEntries := []BSTContent{
		{Key: 1, Value: "apples"},
		{Key: 2, Value: "oranges"},
		{Key: 0, Value: "pears"},
		{Key: -9, Value: "dogs"},
		{Key: 3, Value: "cats"},
	}
	bst := TreeFromNodes(treeEntries)
	return bst, treeEntries
}

func TestBST_Put(t *testing.T) {
	bst, _ := arrangeTree()

	wantLeft := NewNode(0, "pears", NewNode(-9, "dogs", nil, nil), nil)
	wantRight := NewNode(2, "oranges", nil, NewNode(3, "cats", nil, nil))
	wantNode := NewNode(1, "apples", wantLeft, wantRight)
	wantBst := NewBST(wantNode)

	t.Run("simple insert", func(t *testing.T) {
		if !reflect.DeepEqual(wantBst, bst) {
			t.Errorf("BST.Insert() = %v, want %v", bst, wantBst)
		}
	})
}

func TestBST_String(t *testing.T) {
	bst, _ := arrangeTree()
	got := bst.String()
	want :=
		`key: -9, value: dogs
key: 0, value: pears
key: 1, value: apples
key: 2, value: oranges
key: 3, value: cats
`
	if got != want {
		t.Errorf("BST.String() = {%v}, want {%v}", got, want)
	}

	t.Run("stringer", func(t *testing.T) {
	})
}

func TestBSTPostOrder_String(t *testing.T) {
	bst, _ := arrangeTree()
	got := bst.PostOrderString()
	want :=
		`|x||x||-9||x||0||x||x||x||3||2||1|`
	if got != want {
		t.Errorf("BST.String() = {%v}, want {%v}", got, want)
	}

	t.Run("stringer", func(t *testing.T) {
	})
}

func TestBST_Get(t *testing.T) {

	bst, treeEntries := arrangeTree()

	for _, entry := range treeEntries {
		t.Run(entry.String(), func(t *testing.T) {
			got, _ := bst.Get(entry.Key)
			want := entry.Value
			if got != want {
				t.Errorf("BST.Put() = %v, want %v", got, want)
			}
		})
	}

	t.Run("invalid key", func(t *testing.T) {
		got, err := bst.Get(-908)
		if got != nil {
			t.Errorf("BST.Put() = %v, want %v", got, nil)
		}
		if err == nil {
			t.Errorf("BST.Put() error = %v, want %v", err, nil)
		}
	})

}

func TestBST_Inorder(t *testing.T) {
	node := NewNode(1, "apples", nil, nil)
	bst := NewBST(node)
	bst.Put(2, "oranges")
	bst.Put(0, "pears")
	bst.Put(-9, "dogs")
	bst.Put(3, "cats")
	got := bst.Inorder()
	want := []*BSTContent{
		&BSTContent{-9, "dogs"},
		&BSTContent{0, "pears"},
		&BSTContent{1, "apples"},
		&BSTContent{2, "oranges"},
		&BSTContent{3, "cats"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("BST.Inorder() = %v, want %v", got, want)
	}

}

func TestBST_CheckBalanced(t *testing.T) {

	type tc struct {
		name        string
		treeEntries []BSTContent
		want        bool
	}

	testCases := []tc{
		tc{
			"linear",
			[]BSTContent{
				{Key: 1, Value: "apples"},
				{Key: 2, Value: "oranges"},
				{Key: 3, Value: "pears"},
				{Key: 4, Value: "dogs"},
			},
			false},

		tc{
			"balanced",
			[]BSTContent{
				{Key: 17, Value: "apples"},
				{Key: 11, Value: "oranges"},
				{Key: 10, Value: "pears"},
				{Key: 19, Value: "dogs"},
				{Key: 31, Value: "doggo"},
			},
			true},

		tc{"simple",
			[]BSTContent{
				{Key: 1, Value: "apples"},
				{Key: 2, Value: "oranges"},
			},
			true},

		tc{"single key",
			[]BSTContent{
				{Key: 1, Value: "apples"},
			},
			true},

		tc{"empty",
			[]BSTContent{},
			true},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			bst := TreeFromNodes(test.treeEntries)
			got, _ := bst.CheckBalanced()

			if got != test.want {
				t.Errorf("BST.HeightMap() = %v, want %v", got, test.want)
			}
		})
	}

}

func TestBST_Height(t *testing.T) {
	treeEntries := []BSTContent{
		{Key: 1, Value: "apples"},
		{Key: 2, Value: "oranges"},
		{Key: -9, Value: "pears"},
		{Key: -4, Value: "dogs"},
		{Key: -88, Value: "cats"},
		{Key: 3, Value: "cats"},
		{Key: 65, Value: "cats"},
		{Key: 67, Value: "cats"},
		{Key: -3, Value: "cats"},
		{Key: -4, Value: "cats"},
		{Key: -5, Value: "cats"},
		{Key: -6, Value: "cats"},
		{Key: 55, Value: "cats"},
		{Key: -56, Value: "cats"},
		{Key: 78, Value: "cats"},
		{Key: 55, Value: "cats"},
		{Key: 21, Value: "cats"},
	}
	t1 := func(t *testing.T) {

		bst := TreeFromNodes(treeEntries)
		want := 6
		got := bst.Height()

		if got != want {
			t.Errorf("BST.Height() = %v, want %v", got, want)
		}

	}

	t2 := func(t *testing.T) {
		bst, _ := arrangeTree()
		want := 3
		got := bst.Height()

		if got != want {
			t.Errorf("BST.Height() = %v, want %v", got, want)
		}
	}

	t3 := func(t *testing.T) {
		bst := TreeFromNodes(treeEntries)
		want := map[int]int{-88: 2, -56: 1, -9: 4, -6: 1, -5: 2, -4: 3, -3: 1, 1: 6, 2: 5, 3: 4, 21: 1, 55: 2, 65: 3, 67: 2, 78: 1}
		bst.Height()
		got, _ := bst.HeightMap()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("BST.HeightMap() = %v, want %v", got, want)
		}
	}

	testMap := map[string]func(t *testing.T){
		"1":          t1,
		"2":          t2,
		"height map": t3,
	}

	for k, v := range testMap {
		t.Run(k, v)
	}

}

func TestBST_Validate(t *testing.T) {
	treeEntries := []BSTContent{
		{Key: 1, Value: "apples"},
		{Key: 2, Value: "oranges"},
		{Key: -9, Value: "pears"},
		{Key: -4, Value: "dogs"},
		{Key: -88, Value: "cats"},
		{Key: -6, Value: "cats"},
		{Key: 55, Value: "cats"},
		{Key: -56, Value: "cats"},
		{Key: 78, Value: "cats"},
		{Key: 55, Value: "cats"},
		{Key: 21, Value: "cats"},
	}

	bst := TreeFromNodes(treeEntries)
	want := true
	if got := bst.Validate(); got != want {
		t.Errorf("BST.Validate() = %v, want %v", got, want)
	}
}
