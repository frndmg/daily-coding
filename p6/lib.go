package p6

import (
	"fmt"
	"unsafe"
)

// ErrIndexOutOfRange type
type ErrIndexOutOfRange struct {
	idx int
}

func (e ErrIndexOutOfRange) Error() string {
	return fmt.Sprintf("xorlinkedlist: index out of range at %v", e.idx)
}

// Direction type
type Direction int

const (
	// Left direction
	Left Direction = -1
	// Right direction
	Right Direction = 1
)

// XORLinkedList type
type XORLinkedList struct {
	cursor *XORCursor
}

// XORCursor type
type XORCursor struct {
	idx  int
	curr *XORNode
	prev *XORNode
	next *XORNode
}

// XORNode type
type XORNode struct {
	v    interface{}
	both uintptr
}

// Add function
func (l *XORLinkedList) Add(v interface{}) {
	if l.cursor == nil {
		l.cursor = &XORCursor{0, &XORNode{v, 0}, nil, nil}
		return
	}

	l.cursor.Add(v)
}

// Add function
func (c *XORCursor) Add(v interface{}) {
	for c.next != nil {
		c.Move(Right)
	}

	next := &XORNode{v, uintptr(unsafe.Pointer(c.curr))}
	c.curr.both = uintptr(unsafe.Pointer(
		uintptr(unsafe.Pointer(c.prev)) ^ uintptr(unsafe.Pointer(next))))
	c.next = next
	c.Move(Right)
}

// Move function
func (c *XORCursor) Move(d Direction) bool {
	nPtr := uintptr(unsafe.Pointer(c.curr))

	switch d {
	case Right:
		if c.next == nil {
			return false
		}
		next := (*XORNode)(unsafe.Pointer(c.next.both ^ nPtr))
		c.prev = c.curr
		c.curr = c.next
		c.next = next
	case Left:
		if c.prev == nil {
			return false
		}
		prev := (*XORNode)(unsafe.Pointer(c.prev.both ^ nPtr))
		c.next = c.curr
		c.curr = c.prev
		c.prev = prev
	}

	c.idx += int(d)
	return true
}

// Get function
func (l *XORLinkedList) Get(idx int) (interface{}, error) {

	d := Right
	if l.cursor.idx > idx {
		d = Left
	}

	for n := l.cursor.curr; n != nil; n = l.cursor.curr {
		if l.cursor.idx == idx {
			return n.v, nil
		}

		if !l.cursor.Move(d) {
			break
		}
	}

	return nil, ErrIndexOutOfRange{idx}
}
