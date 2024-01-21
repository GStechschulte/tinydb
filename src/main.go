package main

import (
	"tinydb/src/memtable"
)

func main() {
	tree := memtable.NewTree()

	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(25)
	tree.Insert(25)
	tree.Insert(5)

	tree.PrintTree()
}
