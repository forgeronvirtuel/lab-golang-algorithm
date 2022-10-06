package linkedlist

import (
	"bytes"
	"math/rand"
	"testing"
)

const dataSize = 2_000_000

func generateBytes() []byte {
	token := make([]byte, 4)
	rand.Read(token)
	return token
}

func TestBytesLinkedListItem_Add(t *testing.T) {
	var ll *LIFOBytesLinkedList

	ll = ll.Add([]byte("A"))
	if bytes.Compare(ll.last.msg, []byte("A")) != 0 {
		t.Errorf("Expected message to be A, got %s", string(ll.last.msg))
	}

	ll = ll.Add([]byte("B"))
	if bytes.Compare(ll.last.msg, []byte("B")) != 0 {
		t.Errorf("Expected message to be B, got %s", string(ll.last.msg))
	}
}

func TestBytesLinkedListItem_Get(t *testing.T) {
	var ll *LIFOBytesLinkedList

	msg1 := generateBytes()
	msg2 := generateBytes()

	lli := &bytesLinkedListItem{
		msg: msg2,
		ll: &bytesLinkedListItem{
			msg: msg1,
			ll:  nil,
		},
	}

	ll = &LIFOBytesLinkedList{
		last:   lli,
		length: 2,
	}

	if bytes.Compare(ll.Get(), msg2) != 0 {
		t.Fatalf("Not the same message")
	}
	if ll.Length() != 1 {
		t.Fatalf("not the same length: got %d, expected %d", ll.Length(), 1)
	}
}

func BenchmarkLinkedListItem_Add(b *testing.B) {
	var ll *LIFOBytesLinkedList

	data := make([][]byte, dataSize)
	for i := 0; i < dataSize; i++ {
		data[i] = generateBytes()
	}

	ll = ll.Add(data[0])

	b.ResetTimer()
	b.StartTimer()
	for i := 1; i < dataSize; i++ {
		ll = ll.Add(data[i])
	}
	b.StopTimer()
}

func BenchmarkLinkedListItem_Get(b *testing.B) {
	var ll *LIFOBytesLinkedList

	for i := 0; i < dataSize; i++ {
		ll = ll.Add(generateBytes())
	}

	b.ResetTimer()
	b.StartTimer()
	for i := 1; i < dataSize; i++ {
		ll.Get()
	}
	b.StopTimer()
}

func TestBytesLinkedListItem_Length(t *testing.T) {
	var ll *LIFOBytesLinkedList
	if ll.Length() != 0 {
		t.Errorf("length %d, expected 0", ll.Length())
	}

	ll = ll.Add(generateBytes())
	if ll.Length() != 1 {
		t.Errorf("Got length %d, expected 1", ll.Length())
	}

	ll = ll.Add(generateBytes())
	if ll.Length() != 2 {
		t.Errorf("Got length %d, expected 2", ll.Length())
	}
}

func TestQueue(t *testing.T) {
	var queue *Queue

	if msg, q := queue.Consume(); msg != nil || q != nil {
		t.Fatalf("msg = %v\nq = %v", msg, q)
	}
	if queue.Length() != 0 {
		t.Fatalf("not equal to zero")
	}

	msgs := [][]byte{
		generateBytes(),
		generateBytes(),
		generateBytes(),
	}

	queue = queue.
		Produce(msgs[0]).
		Produce(msgs[1]).
		Produce(msgs[2])

	if queue.Length() != 3 {
		t.Fatalf("Should be 3 items, got %d", queue.Length())
	}

	if msg, _ := queue.Consume(); bytes.Compare(msg, msgs[0]) != 0 {
		t.Fatalf("#1 Not equal")
	}
	if queue.length != 2 {
		t.Fatalf("Should be 2")
	}
	if msg, _ := queue.Consume(); bytes.Compare(msg, msgs[1]) != 0 {
		t.Fatalf("#2 Not equal")
	}
	if queue.length != 1 {
		t.Fatalf("Should be 1")
	}
	if msg, _ := queue.Consume(); bytes.Compare(msg, msgs[2]) != 0 {
		t.Fatalf("#3 Not equal")
	}
	if queue.length != 0 {
		t.Fatalf("Should be 0")
	}
	if msg, q := queue.Consume(); msg != nil || q != nil {
		t.Fatalf("msg = %v\nq = %v", msg, q)
	}
	if queue.length != 0 {
		t.Fatalf("Should be 0")
	}
}

func BenchmarkQueue_Produce(b *testing.B) {
	var data = make([][]byte, dataSize)
	for i := 0; i < dataSize; i++ {
		data[i] = generateBytes()
	}
	var queue *Queue
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < dataSize; i++ {
		queue = queue.Produce(data[i])
	}
	b.StopTimer()
}

func BenchmarkQueue_Consume(b *testing.B) {
	var data = make([][]byte, dataSize)
	for i := 0; i < dataSize; i++ {
		data[i] = generateBytes()
	}
	var queue *Queue
	for i := 0; i < dataSize; i++ {
		queue = queue.Produce(data[i])
	}

	b.ResetTimer()
	b.StartTimer()
	for queue.Length() != 0 {
		queue.Consume()
	}
	b.StopTimer()
}
