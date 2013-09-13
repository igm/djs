// In computing, a disjoint-set data structure (or union find data structure) is a data structure that keeps track of a set of elements partitioned into a number of disjoint (nonoverlapping) subsets. A union-find algorithm is an algorithm that performs two useful operations on such a data structure:
//
//  Find  - determine which subset a particular element is in. This can be used for determining if two elements are in the same subset.
//  Union - join two subsets into a single subset.
package djs

// A type, typically a collection. This is the minimum contract any structure needs to conform in order
// to be used as underlaying data structure for union-find algorithm.
// Path compression is the only iprovement applied in this case.
type Interface interface {
	// Sets parent p to element c
	SetParent(c, p interface{})
	// Returns back a parent of element
	GetParent(c interface{}) interface{}
}

// Adds iteration to optionally perform initial state setup
type InitInterface interface {
	Interface
	Each(func(a interface{}))
}

// Additional methods to enable "union by rank".
type RankInterface interface {
	Interface
	GetRank(a interface{}) int
	SetRank(a interface{}, rank int)
}

func Init(u InitInterface) {
	u.Each(func(a interface{}) {
		u.SetParent(a, a)
		if g, ok := u.(RankInterface); ok {
			g.SetRank(a, 0)
		}
	})
}

func root(u Interface, a interface{}) interface{} {
	for ; a != u.GetParent(a); a = u.GetParent(a) {
		u.SetParent(a, u.GetParent(u.GetParent(a)))
	}
	return a
}

//  Determine which subset a particular element is in. This can be used for determining if two elements are in the same subset.
func Find(u Interface, a interface{}) interface{} {
	return root(u, a)
}

// Convenient method that compares result of Find for both elements.
func Connected(u Interface, a, b interface{}) bool {
	return root(u, a) == root(u, b)
}

// Join two subsets into a single subset.
func Union(u Interface, a, b interface{}) {
	aa := Find(u, a)
	bb := Find(u, b)

	if g, ok := u.(RankInterface); ok {
		ra := g.GetRank(aa)
		rb := g.GetRank(bb)
		if ra > rb {
			u.SetParent(bb, aa)
		} else if ra < rb {
			u.SetParent(aa, bb)
		} else {
			g.SetRank(aa, ra+1)
			u.SetParent(bb, aa)
		}
	} else {
		u.SetParent(aa, bb)
	}
}

// Convenience types for common cases

// UnionInt attaches the methods of Interface to []int
type UnionInt []int

func (m UnionInt) SetParent(c, p interface{})          { m[c.(int)] = p.(int) }
func (m UnionInt) GetParent(c interface{}) interface{} { return m[c.(int)] }

func (m UnionInt) Each(fn func(a interface{})) {
	for a := range m {
		fn(a)
	}
}

type rankUnion struct {
	sets  []int
	ranks []int
}

func (m *rankUnion) SetParent(c, p interface{})          { m.sets[c.(int)] = p.(int) }
func (m *rankUnion) GetParent(c interface{}) interface{} { return m.sets[c.(int)] }
func (m *rankUnion) SetRank(a interface{}, rank int)     { m.ranks[a.(int)] = rank }
func (m *rankUnion) GetRank(a interface{}) int           { return m.ranks[a.(int)] }
func (m *rankUnion) Each(fn func(a interface{})) {
	for a := range m.sets {
		fn(a)
	}
}

func NewRankUnion(size int) *rankUnion {
	ret := &rankUnion{sets: make([]int, size), ranks: make([]int, size)}
	Init(ret)
	return ret
}
