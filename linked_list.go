package data_structures_go

import (
    "fmt"
)

type LLNode struct {
    value   int
    next    *LLNode
}

type LinkedList struct {
    head  *LLNode
    Length  int     // Length (count of node) of linked list
}

func CreateFromSlice(slice []int) *LinkedList {
    ll := new(LinkedList)
    for _, v := range slice {
        ll.AppendTail(v)
    }
    return ll
}

// Add node to tail of linked list
func (ll *LinkedList) AppendTail(value int) {
    if ll.head == nil {
        ll.head = new(LLNode)
        ll.head.value = value
        ll.Length = 1
    } else {
        nn := ll.head
        for ; nn.next != nil; nn = nn.next {}
        nn.next = new(LLNode)
        nn.next.value = value
        ll.Length++
    }
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

// Starting from specified head node returns node at position index
func get_node_at_index(head *LLNode, index int) (*LLNode) {
    if index == 0 || head == nil {
        return head
    }
    return get_node_at_index(head.next, index - 1)
}

// Print out each value in linked list
func print_list(head *LLNode) {
    for ; head != nil; head = head.next {
        fmt.Printf("v: %d, ", head.value)
    }
    fmt.Printf("\n")
}

