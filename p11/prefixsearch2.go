package p11

// Node type
type Node struct {
	val      rune
	leaf     bool
	depth    int
	parent   *Node
	children map[rune]*Node
}

// Trie type
type Trie struct {
	root *Node
}

// NewTrie function
func NewTrie() *Trie {
	return &Trie{
		root: &Node{children: make(map[rune]*Node)},
	}
}

const leafVal rune = 0x0

// Add function
func (t *Trie) Add(w string) {
	rs := []rune(w)
	i, n := findNode(t.root, rs)

	for ; i < len(w); i++ {
		r := rs[i]
		next := &Node{
			val:      r,
			depth:    n.depth + 1,
			parent:   n,
			children: make(map[rune]*Node),
		}
		n.children[r] = next
		n = next
	}

	n.children[leafVal] = &Node{
		val:    leafVal,
		leaf:   true,
		depth:  n.depth + 1,
		parent: n,
	}
}

// SearchPrefix function
func (t *Trie) SearchPrefix(w string) []string {
	i, n := findNode(t.root, []rune(w))
	if i == len(w) {
		return collect(n)
	}

	return []string{}
}

func findNode(root *Node, rs []rune) (int, *Node) {
	n := root
	i := 0

	for ; i < len(rs) && n != nil; i++ {
		r := rs[i]

		if c, ok := n.children[r]; ok {
			n = c
		} else {
			break
		}
	}

	return i, n
}

func collect(node *Node) []string {
	out := make([]string, 0)

	ns := []*Node{node}
	for len(ns) > 0 {
		n := ns[0]
		ns = ns[1:]

		for _, c := range n.children {
			ns = append(ns, c)
		}

		if n.leaf {
			word := ""
			for p := n.parent; p.depth > 0; p = p.parent {
				word = string(p.val) + word
			}
			out = append(out, word)
		}
	}

	return out
}

// PrefixSearch2 function
func PrefixSearch2(s string, a []string) []string {
	t := NewTrie()
	for _, w := range a {
		t.Add(w)
	}
	return t.SearchPrefix(s)
}
