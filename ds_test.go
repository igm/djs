package disjointset_test

import (
	"fmt"
	ds "github.com/igm/disjointsets"
	"testing"
)

func TestUnionInt(t *testing.T) {
	var data ds.UnionInt = make([]int, 100)
	ds.Init(data)
	for i, v := range data {
		if i != v {
			t.Errorf("Union initialization problem\n")
		}
	}
}

func TestRankUnionInt(t *testing.T) {
	ru := ds.NewRankUnion(1000)
	for i := 0; i < 10; i++ {
		if i != ds.Find(ru, i) {
			t.Errorf("Union initialization problem\n")
		}
	}

	ds.Union(ru, 1, 2)
	ds.Union(ru, 1, 3)
	ds.Union(ru, 8, 9)
	if !ds.Connected(ru, 1, 2) {
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
	ds.Init(set)

	for _, node := range set.nodes {
		if node != node.parent {
			t.Errorf("Union initialization problem\n")
		}
	}
	if ds.Union(set, set.nodes[0], set.nodes[2]); !ds.Connected(set, set.nodes[0], set.nodes[2]) {
		t.Errorf("Sets should be connected\n")
	}
	if set.nodes[0].rank != 1 || set.nodes[2].rank != 0 {
		t.Errorf("Wrong rank update\n")
	}

	if ds.Union(set, set.nodes[0], set.nodes[3]); set.nodes[0].rank != 1 || set.nodes[3].rank != 0 {
		t.Errorf("Wrong rank update\n")
	}
	if ds.Union(set, set.nodes[4], set.nodes[5]); set.nodes[4].rank != 1 || set.nodes[5].rank != 0 {
		t.Errorf("Wrong rank update\n")
	}
	if ds.Union(set, set.nodes[0], set.nodes[4]); set.nodes[0].rank != 2 || set.nodes[4].rank != 1 {
		t.Errorf("Wrong rank update\n")
	}
}

func ExampleInit() {
	var data ds.UnionInt = make([]int, 20)

	ds.Init(data)
	ds.Union(data, 10, 11)
	ds.Union(data, 1, 2)

	fmt.Println(ds.Connected(data, 10, 11))
	fmt.Println(data)
	// Output: true
	//[0 2 2 3 4 5 6 7 8 9 11 11 12 13 14 15 16 17 18 19]
}
