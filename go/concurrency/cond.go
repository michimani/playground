package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	Name string
	Q    []struct{}
	C    *sync.Cond
}

func (q *Queue) len() int {
	return len(q.Q)
}

func (q *Queue) add() {
	q.Q = append(q.Q, struct{}{})
	fmt.Printf("Added to %s. (len: %d)\n", q.Name, q.len())
}

func NewQueue(name string, size int, cond *sync.Cond) Queue {
	return Queue{
		Name: name,
		Q:    make([]struct{}, 0, size),
		C:    cond,
	}
}

func SignalToQueue() {
	c := sync.NewCond(&sync.Mutex{})
	queue := NewQueue("Q1", 10, nil)

	remove := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue.Q = queue.Q[1:] // 先頭の要素を取り除く
		fmt.Printf("Remove from %s.\n", queue.Name)
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for queue.len() == 2 {
			c.Wait()
		}
		queue.add()
		go remove(1 * time.Second)
		c.L.Unlock()
	}

	fmt.Printf("Remaining item count in %s is %d.\n", queue.Name, queue.len())
}

func Broadcast() {
	c := sync.NewCond(&sync.Mutex{})
	queues := []Queue{
		NewQueue("Q1", 10, c),
		NewQueue("Q2", 10, c),
		NewQueue("Q3", 10, c),
	}

	remove := func(q *Queue, delay time.Duration) {
		time.Sleep(delay)
		q.C.L.Lock()
		q.Q = q.Q[1:] // 先頭の要素を取り除く
		fmt.Printf("Remove from %s. (len: %d)\n", q.Name, q.len())
		q.C.L.Unlock()
		q.C.Broadcast()
		// q.C.Signal()
	}

	var wg sync.WaitGroup
	for i, q := range queues {
		q := q
		wg.Add(1)
		go func(q *Queue, i int) {
			for i := 0; i < 10; i++ {
				q.C.L.Lock()
				for q.len() == 2 {
					q.C.Wait()
					if q.len() == 2 {
						fmt.Printf("%s 「Whats?」\n", q.Name)
					}
				}
				q.add()
				d := time.Duration(200 * i)
				go remove(q, d*time.Millisecond)
				q.C.L.Unlock()
			}
			wg.Done()
		}(&q, (i+1)*3)
	}

	wg.Wait()
}
