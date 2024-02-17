package main

import (
	"fmt"
	"time"
	"tinydb/src/memtable"
)

func main() {

	path := "/tmp/example_1.sst"
	// defer os.RemoveAll(path)

	tree := memtable.NewTree()

	t1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 2, 0, 0, 0, 1, time.UTC)
	t3 := time.Date(2024, 1, 1, 0, 0, 0, 2, time.UTC)

	tree.Insert(t1, 1.0)
	tree.Insert(t2, 2.0)
	tree.Insert(t3, 52.6)

	fmt.Printf("Total size of tree = %d bytes\n", tree.Size())

	// tree.PrintTree()

	// sorted := memtable.InOrderTraversal(tree)
	// fmt.Println("In order traversal of tree:", sorted)

	tree.Flush(path)
}
