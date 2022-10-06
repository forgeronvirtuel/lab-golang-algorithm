package linkedlist

import "fmt"

type LinkedListError struct {
	msg       string
	operation string
}

func (err *LinkedListError) Error() string {
	return fmt.Sprintf("%s: %s", err.operation, err.msg)
}

type bytesLinkedListItem struct {
	msg []byte
	ll  *bytesLinkedListItem
}

func (ll *bytesLinkedListItem) lifoAdd(msg []byte) *bytesLinkedListItem {
	if ll == nil {
		return &bytesLinkedListItem{
			msg: msg,
			ll:  nil,
		}
	}

	return &bytesLinkedListItem{msg: msg, ll: ll}
}

type Queue struct {
	first  *bytesLinkedListItem
	last   *bytesLinkedListItem
	length int
}

func (q *Queue) Produce(msg []byte) *Queue {
	if q == nil {
		blk := &bytesLinkedListItem{
			msg: msg,
			ll:  nil,
		}
		return &Queue{
			first:  blk,
			last:   blk,
			length: 1,
		}
	}

	blk := &bytesLinkedListItem{
		msg: msg,
		ll:  nil,
	}
	q.last.ll = blk
	q.last = blk
	q.length++
	return q
}

func (q *Queue) Consume() ([]byte, *Queue) {
	if q == nil || q.first == nil {
		return nil, nil
	}
	blk := q.first
	q.first = blk.ll
	q.length--
	return blk.msg, q
}

func (q *Queue) Length() int {
	if q == nil {
		return 0
	}
	return q.length
}

type LIFOBytesLinkedList struct {
	last   *bytesLinkedListItem
	length int
}

func (ll *LIFOBytesLinkedList) Add(msg []byte) *LIFOBytesLinkedList {
	if ll == nil {
		return &LIFOBytesLinkedList{
			last: &bytesLinkedListItem{
				msg: msg,
				ll:  nil,
			},
			length: 1,
		}
	}

	ll.last = ll.last.lifoAdd(msg)
	ll.length += 1
	return ll
}

func (ll *LIFOBytesLinkedList) Get() []byte {
	msg := ll.last.msg
	ll.last = ll.last.ll
	ll.length -= 1
	return msg
}

func (ll *LIFOBytesLinkedList) Length() int {
	if ll == nil {
		return 0
	}
	return ll.length
}
