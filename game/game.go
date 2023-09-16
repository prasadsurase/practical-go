package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Item struct {
	X int
	Y int
}

type Player struct {
	Name string
	Item
	Keys []Key
}

type Mover interface {
	Move(int, int)
}

// func NewItem(x, y int) Item {
// func NewItem(x, y int) *Item {
// func NewItem(x, y int) (Item, error) {
func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}
	// The Go compiler does "escape analysis" and will allocation i on the heap
	return &i, nil
}

// zero vs missing value
const (
	maxX = 1000
	maxY = 600
)

func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

func (p *Player) FoundKey(k Key) error {
	if k < Jade || k >= invalidKey {
		return fmt.Errorf("invalid key: %v", k)
	}
	if !containsKey(p.Keys, k) {
		p.Keys = append(p.Keys, k)
	}
	return nil
}

func containsKey(keys []Key, k Key) bool {
	for _, k2 := range keys {
		if k2 == k {
			return true
		}
	}
	return false
}

type Key byte

// Go's version of enum
const (
	Jade Key = iota + 1
	Copper
	Crystal
	invalidKey //internal. not exported.
)

// implement Stringer interface
func (k Key) String() string {
	switch k {
	case Jade:
		return "jade"
	case Copper:
		return "copper"
	case Crystal:
		return "crystal"
	}
	return fmt.Sprintf("<Key %d", k)
}

func main() {
	i1 := Item{100, 100}

	p1 := Player{
		Name: "Prasad",
		Item: Item{500, 300},
	}

	ms := []Mover{&i1, &p1}

	for _, i := range ms {
		fmt.Printf("Mover: %#v\n", i)
	}
	k := Jade
	fmt.Printf("%d %#v\n", k, k)
	fmt.Println(k.String())

	json.NewEncoder(os.Stdout).Encode("hello world")
	json.NewEncoder(os.Stdout).Encode(time.Now())

	p1.FoundKey(Jade)
	fmt.Println(p1.Keys)
	p1.FoundKey(Copper)
	fmt.Println(p1.Keys)
	p1.FoundKey(Crystal)
	fmt.Println(p1.Keys)

	p1.FoundKey(Jade)
	fmt.Println(p1.Keys)
}
