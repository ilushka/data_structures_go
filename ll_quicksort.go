package data_structures_go

// Quicksort partitioning based on selected pivot. Uses tail node as pivot aka. Lomuto scheme.
func qs_partition(head *LLNode, tail *LLNode) (*LLNode, *LLNode, *LLNode, *LLNode, *LLNode) {
    if head.next == nil || head == nil || tail == nil {
        return nil, nil, nil, nil, nil
    }

    var min_hd, min_tl, max_hd, max_tl *LLNode = nil, nil, nil, nil
    for j := head; j != tail; j = j.next {
        if (j.value < tail.value) {
            // add j to list with elements < pivot
            if min_hd == nil {
                min_hd = j
            } else {
                min_tl.next = j
            }
            min_tl = j
        } else {
            // add j to list with elements >= pivot
            if max_hd == nil {
                max_hd = j
            } else {
                max_tl.next = j
            }
            max_tl = j
        }
    }

    // terminate min and max lists
    tail.next = nil
    if min_tl != nil {
        min_tl.next = nil
    }
    if max_tl != nil {
        max_tl.next = nil
    }
    return min_hd, min_tl, tail, max_hd, max_tl
}

// Recursion portion of Quicksort
func qs_sort(head *LLNode, tail *LLNode) (*LLNode, *LLNode) {
    if head == nil {
        return nil, nil
    } else if head.next == nil {
        return head, head
    }

    min_hd, min_tl, pivot, max_hd, max_tl := qs_partition(head, tail)
    if pivot == nil {
        return nil, nil
    }

    par0_hd, par0_tl := qs_sort(min_hd, min_tl)
    par1_hd, par1_tl := qs_sort(max_hd, max_tl)

    var new_head, new_tail *LLNode = par0_hd, par1_tl
    if par0_hd == nil {
        // start sorted partition with pivot
        new_head = pivot
    } else {
        // point pivot to tail of sorted min partition
        par0_tl.next = pivot
    }
    if par1_tl == nil {
        // set tail of partition to pivot
        new_tail = pivot
        pivot.next = nil
    } else {
        // point pivot to head of sorted max partition
        pivot.next = par1_hd
    }
    return new_head, new_tail
}

// Sort linked list using Quicksort
func (ll *LinkedList) QuickSort() {
    if ll.Length == 0 || ll.Length == 1 {
        return
    }
    tail := get_node_at_index(ll.head, ll.Length - 1)
    new_head, _ := qs_sort(ll.head, tail)
    ll.head = new_head
}

