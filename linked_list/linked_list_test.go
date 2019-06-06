package linked_list

import (
    "testing"
)

func TestAppendTail(t *testing.T) {
    testval := [...]int{100, 200, 300, 400, 500, 600, 700, 800, 900}
    ll := CreateFromSlice(testval[:])
    if ll.Length != len(testval) {
        t.Errorf("linked list should have %d nodes", len(testval))
    }
    for ii, vv := range *ll.ToArray() {
        if vv != testval[ii] {
            t.Errorf("linked list's node #%d should have value of %d", ii, testval[ii])
        }
    }
}

func TestQuickSort(t *testing.T) {
    expected := [...]int{-300, 200, 400, 500, 800, 880, 900, 950, 1000}
    ll := CreateFromSlice([]int{1000, 200, -300, 400, 500, 950, 800, 880, 900})
    ll.QuickSort()
    for ii, vv := range *ll.ToArray() {
        if vv != expected[ii] {
            t.Errorf("sorted linked list's node #%d should have value of %d", ii, expected[ii])
        }
    }
}

func TestQuickSortSingleItem(t *testing.T) {
    expected := [...]int{1000}
    ll := CreateFromSlice([]int{1000})
    ll.QuickSort()
    for ii, vv := range *ll.ToArray() {
        if vv != expected[ii] {
            t.Errorf("sorted linked list's node #%d should have value of %d", ii, expected[ii])
        }
    }
}

func TestQuickSortEqualValues(t *testing.T) {
    expected := [...]int{10, 10, 1000, 1200, 1200, 1220, 1500, 1500}
    ll := CreateFromSlice([]int{1200, 1220, 1200, 1000, 1500, 1500, 10, 10})
    ll.QuickSort()
    for ii, vv := range *ll.ToArray() {
        if vv != expected[ii] {
            t.Errorf("sorted linked list's node #%d should have value of %d", ii, expected[ii])
        }
    }
}

func TestQuickSortEqualValues2(t *testing.T) {
    expected := [...]int{1200, 1200, 1200}
    ll := CreateFromSlice([]int{1200, 1200, 1200})
    ll.QuickSort()
    for ii, vv := range *ll.ToArray() {
        if vv != expected[ii] {
            t.Errorf("sorted linked list's node #%d should have value of %d", ii, expected[ii])
        }
    }
}
