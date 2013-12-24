package ll

import "fmt"

type Cell struct {
	Value interface{}
	Next  *Cell
}

type SLL struct {
	Cell **Cell
}

func (s SLL) String() string {
	acc := "("
	cell := *s.Cell
	for {
		acc += fmt.Sprintf("%#v", cell.Value)
		if cell.Next != nil {
			acc += ", "
			cell = cell.Next
		} else {
			acc += ")"
			return acc
		}
	}
}

func (sll *SLL) Cons(value interface{}) {
	if sll.Cell != nil {
		newCell := &Cell{value, *sll.Cell}
		sll.Cell = &newCell
	} else {
		cell := &Cell{value, nil}
		sll.Cell = &cell
	}
}

func (sll *SLL) Count() int {
	acc := 0
	if sll.Cell != nil {
		for next := *sll.Cell; next != nil; {
			acc += 1
			next = next.Next
		}
	}
	return acc
}

func (sll *SLL) Reverse() {
	cell := *sll.Cell
	next := cell.Next
	cell.Next = nil
	for next != nil {
		newNext := next.Next
		next.Next = cell
		cell = next
		next = newNext
	}
	sll.Cell = &cell
}
