package data_structures_go

type LLNode struct {
    value   int
    next    *LLNode
}

type LinkedList struct {
    head  *LLNode
    size  int
}

func (ll *LinkedList) AppendTail(value int) {
    if ll.head == nil {
        ll.head = new(LLNode)
        ll.head.value = value
        ll.size = 1
    } else {
        nn := ll.head
        for ; nn.next != nil; nn = nn.next {}
        nn.next = new(LLNode)
        nn.next.value = value
        ll.size++
    }
}

func (ll *LinkedList) ToArray() *[]int {
    if ll.size == 0 {
        return nil
    }
    a := make([]int, ll.size)
    var ii int = 0
    for nn := ll.head; nn != nil; nn = nn.next {
        a[ii] = nn.value
        ii++
    }
    return &a
}
