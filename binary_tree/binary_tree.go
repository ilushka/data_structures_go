package binary_tree


type BTNode struct {
    value   int
    left    *BTNode
    right   *BTNode
}

type BinaryTree struct {
    root    *BTNode
    Length  int
}

func insert(root *BTNode, value int) {
    var node **BTNode
    if value < root.value {
        node = &root.left
    } else {
        node = &root.right
    }
    if *node == nil {
        *node = &BTNode{value, nil, nil}
    } else {
        insert(*node, value)
    }
}

func preorder(root *BTNode, f func(int)) {
    if root == nil {
        return
    }
    f(root.value)
    preorder(root.left, f)
    preorder(root.right, f)
}

func (bt *BinaryTree) Insert(value int) {
    if bt.Length == 0 {
        bt.root = &BTNode{value, nil, nil}
    } else {
        insert(bt.root, value)
    }
    bt.Length++
}

func (bt *BinaryTree) Preorder(f func(int)) {
    preorder(bt.root, f)
}

