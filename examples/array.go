package main

import (
	"fmt"
	ds "github.com/igm/djs"
)

func main() {
	var data ds.UnionInt = make([]int, 20)

	ds.Init(data)
	ds.Union(data, 10, 11)
	ds.Union(data, 1, 2)

	fmt.Println(ds.Connected(data, 10, 11))
	fmt.Println(data)
}
