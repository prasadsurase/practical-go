package main

import (
	"fmt"
	"sync"
	"time"
)

type Payment struct {
	Sender   string
	Receiver string
	Amount   float64
	once     sync.Once
}

func (p *Payment) process(t time.Time) {
	ts := t.Format(time.RFC3339)
	fmt.Printf("[%s] %s -> %.2f -> %s\n", ts, p.Sender, p.Amount, p.Receiver)
}

func (p *Payment) Process() {
	t := time.Now()
	p.once.Do(func() { p.process(t) })
}

func main() {

	p1 := Payment{
		Sender:   "Tom",
		Receiver: "Jerry",
		Amount:   123.45,
	}

	p1.Process()
	p1.Process()
	p1.Process()
}
