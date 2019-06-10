package linked_list

import (
    "fmt"
)

type LLNode struct {
    value   int
    next    *LLNode
}

type LinkedList struct {
    head    *LLNode
    tail    *LLNode
    Length  int     // Length (count of node) of linked list
}

// Create linked list from elements in a slice
func CreateFromSlice(slice []int) *LinkedList {
    ll := new(LinkedList)
    for _, v := range slice {
        ll.Append(v)
    }
    return ll
}

// Return node at specified index and, potentially, its previous node
func get(prev *LLNode, curr *LLNode, index int) (*LLNode, *LLNode) {
    if curr == nil || index < 0 {
        return nil, nil
    }
    if index == 0 {
        return prev, curr
    }
    return get(curr, curr.next, index - 1)
}

// Set first element of linked list
func init_ll(ll *LinkedList, value int) {
    ll.head = &LLNode{value, nil}
    ll.tail = ll.head
    ll.Length++
}

// Print out each value in linked list
func print_list(head *LLNode) {
    for ; head != nil; head = head.next {
        fmt.Printf("v: %d, ", head.value)
    }
    fmt.Printf("\n")
}

// Return value at node at specified index
func (ll *LinkedList) Get(index int) (int, error) {
    if _, c := get(nil, ll.head, index); c == nil {
        return 0, fmt.Errorf("Node index %d does not exist", index)
    } else {
        return c.value, nil
    }
}

// Add node to tail of linked list
func (ll *LinkedList) Append(value int) {
    if ll.Length == 0 {
        ll.head = &LLNode{value, nil}
        ll.tail = ll.head
    } else {
        ll.tail.next = &LLNode{value, nil}
        ll.tail = ll.tail.next
    }
    ll.Length++
}

// Add node at specified index
func (ll *LinkedList) AppendAtIndex(value int, index int) (error) {
    if index > ll.Length || index < 0 {
        return fmt.Errorf("Cannot insert node at index %d in list of size %d", index, ll.Length)
    }
    if index == ll.Length {
        ll.Append(value)
    } else if index == 0 {
        ll.head = &LLNode{value, ll.head}
        ll.Length++
    } else {
        p, c := get(nil, ll.head, index)
        p.next = &LLNode{value, c}
        ll.Length++
    }
    return nil
}

// Returns array of values in linked list
func (ll *LinkedList) ToArray() (*[]int) {
    if ll.Length == 0 {
        return nil
    }
    a := make([]int, ll.Length)
    var ii int = 0
    for nn := ll.head; nn != nil; nn = nn.next {
        a[ii] = nn.value
        ii++
    }
    return &a
}


