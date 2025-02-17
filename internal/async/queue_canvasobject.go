// Code generated by go run gen.go; DO NOT EDIT.

package async

import (
	"sync"
	"sync/atomic"

	"github.com/Anwert/fyne/v2"
)

var itemCanvasObjectPool = sync.Pool{
	New: func() interface{} { return &itemCanvasObject{next: nil, v: nil} },
}

// In puts the given value at the tail of the queue.
func (q *CanvasObjectQueue) In(v fyne.CanvasObject) {
	i := itemCanvasObjectPool.Get().(*itemCanvasObject)
	i.next = nil
	i.v = v

	var last, lastnext *itemCanvasObject
	for {
		last = loadCanvasObjectItem(&q.tail)
		lastnext = loadCanvasObjectItem(&last.next)
		if loadCanvasObjectItem(&q.tail) == last {
			if lastnext == nil {
				if casCanvasObjectItem(&last.next, lastnext, i) {
					casCanvasObjectItem(&q.tail, last, i)
					atomic.AddUint64(&q.len, 1)
					return
				}
			} else {
				casCanvasObjectItem(&q.tail, last, lastnext)
			}
		}
	}
}

// Out removes and returns the value at the head of the queue.
// It returns nil if the queue is empty.
func (q *CanvasObjectQueue) Out() fyne.CanvasObject {
	var first, last, firstnext *itemCanvasObject
	for {
		first = loadCanvasObjectItem(&q.head)
		last = loadCanvasObjectItem(&q.tail)
		firstnext = loadCanvasObjectItem(&first.next)
		if first == loadCanvasObjectItem(&q.head) {
			if first == last {
				if firstnext == nil {
					return nil
				}
				casCanvasObjectItem(&q.tail, last, firstnext)
			} else {
				v := firstnext.v
				if casCanvasObjectItem(&q.head, first, firstnext) {
					atomic.AddUint64(&q.len, ^uint64(0))
					itemCanvasObjectPool.Put(first)
					return v
				}
			}
		}
	}
}

// Len returns the length of the queue.
func (q *CanvasObjectQueue) Len() uint64 {
	return atomic.LoadUint64(&q.len)
}
