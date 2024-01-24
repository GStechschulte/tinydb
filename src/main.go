package main

import (
	"time"
	"tinydb/src/memtable"
)

func main() {
	tree := memtable.NewTree()

	t1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 2, 0, 0, 0, 1, time.UTC)
	t3 := time.Date(2024, 1, 1, 0, 0, 0, 2, time.UTC)

	tree.Insert(t1, 1.0)
	tree.Insert(t2, 2.0)
	tree.Insert(t3, 3.0)

	// Old test cases for value of type int
	// tree.Insert(13)
	// tree.Insert(8)
	// tree.Insert(17)
	// tree.Insert(1)
	// tree.Insert(11)
	// tree.Insert(15)
	// tree.Insert(25)
	// tree.Insert(6)
	// tree.Insert(22)
	// tree.Insert(27)

	tree.PrintTree()
}
