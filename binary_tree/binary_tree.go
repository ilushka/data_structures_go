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

func postorder(root *BTNode, f func(int)) {
    if root == nil {
        return
    }
    postorder(root.left, f)
    postorder(root.right, f)
    f(root.value)
}

func inorder(root *BTNode, f func(int)) {
    if root == nil {
        return
    }
    inorder(root.left, f)
    f(root.value)
    inorder(root.right, f)
}

func (bt *BinaryTree) Insert(value int) {
    if bt.Length == 0 {
        bt.root = &BTNode{value, nil, nil}
    } else {
        insert(bt.root, value)
    }
    bt.Length++
}

// Perform preorder triversal over binary tree while executing provided
// function for each node
func (bt *BinaryTree) Preorder(f func(int)) {
    preorder(bt.root, f)
}

// Perform postorder triversal over binary tree while executing provided
// function for each node
func (bt *BinaryTree) Postorder(f func(int)) {
    postorder(bt.root, f)
}

// Perform inorder triversal over binary tree while executing provided
// function for each node
func (bt *BinaryTree) Inorder(f func(int)) {
    inorder(bt.root, f)
}
