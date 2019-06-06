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
        ll.head = &LLNode{value, nil}
        ll.tail = ll.head
        ll.Length = 1
    } else {
        ll.tail.next = &LLNode{value, nil}
        ll.tail = ll.tail.next
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

// Print out each value in linked list
func print_list(head *LLNode) {
    for ; head != nil; head = head.next {
        fmt.Printf("v: %d, ", head.value)
    }
    fmt.Printf("\n")
}

