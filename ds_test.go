package djs_test

import (
	"fmt"
	"github.com/igm/djs"
	"testing"
)

func TestUnionInt(t *testing.T) {
	var data djs.UnionInt = make([]int, 100)
	djs.Init(data)
	for i, v := range data {
		if i != v {
			t.Errorf("Union initialization problem\n")
		}
	}
}

func TestRankUnionInt(t *testing.T) {
	ru := djs.NewRankUnion(1000)
	for i := 0; i < 10; i++ {
		if i != djs.Find(ru, i) {
			t.Errorf("Union initialization problem\n")
		}
	}

	djs.Union(ru, 1, 2)
	djs.Union(ru, 1, 3)
	djs.Union(ru, 8, 9)
	if !djs.Connected(ru, 1, 2) {
		t.Errorf("Sets should be connected\n")
	}

}

type (
	Node struct {
		parent *Node
		rank   int
	}
	Sets struct{ nodes []*Node }
)

func (s *Sets) GetParent(p interface{}) interface{} { return p.(*Node).parent }
func (s *Sets) SetParent(c, p interface{})          { c.(*Node).parent = p.(*Node) }
func (s *Sets) Each(fn func(interface{})) {
	for k, v := range s.nodes {
		if v == nil {
			v = &Node{}
			s.nodes[k] = v
		}
		fn(v)
	}
}

// optional RankInterface
func (s *Sets) SetRank(n interface{}, rank int) { n.(*Node).rank = rank }
func (s *Sets) GetRank(n interface{}) int       { return n.(*Node).rank }

func TestCustomStruct(t *testing.T) {
	set := &Sets{nodes: make([]*Node, 20)}
	djs.Init(set)

	for _, node := range set.nodes {
		if node != node.parent {
			t.Errorf("Union initialization problem\n")
		}
	}
	if djs.Union(set, set.nodes[0], set.nodes[2]); !djs.Connected(set, set.nodes[0], set.nodes[2]) {
		t.Errorf("Sets should be connected\n")
	}
	if set.nodes[0].rank != 1 || set.nodes[2].rank != 0 {
		t.Errorf("Wrong rank update\n")
	}

	if djs.Union(set, set.nodes[0], set.nodes[3]); set.nodes[0].rank != 1 || set.nodes[3].rank != 0 {
		t.Errorf("Wrong rank update\n")
	}
	if djs.Union(set, set.nodes[4], set.nodes[5]); set.nodes[4].rank != 1 || set.nodes[5].rank != 0 {
		t.Errorf("Wrong rank update\n")
	}
	if djs.Union(set, set.nodes[0], set.nodes[4]); set.nodes[0].rank != 2 || set.nodes[4].rank != 1 {
		t.Errorf("Wrong rank update\n")
	}
}

func ExampleInit() {
	var data djs.UnionInt = make([]int, 20)

	djs.Init(data)
	djs.Union(data, 10, 11)
	djs.Union(data, 1, 2)

	fmt.Println(djs.Connected(data, 10, 11))
	fmt.Println(data)
	// Output: true
	//[0 2 2 3 4 5 6 7 8 9 11 11 12 13 14 15 16 17 18 19]
}

func ExampleNewRankUnion() {
	// create structure with nodes adresses from 0 to 9 (10 in length)
	uf := djs.NewRankUnion(10)
	// join 1 and 2
	djs.Union(uf, 1, 2)
	// join 3 and 4
	djs.Union(uf, 3, 4)
	fmt.Println(djs.Find(uf, 2))
	fmt.Println(djs.Find(uf, 1))
	// Output: 1
	// 1
}
