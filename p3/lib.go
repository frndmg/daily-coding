package p3

import (
	"encoding/json"
)

// Node type
type Node struct {
	Val   interface{} `tag:"val"`
	Left  *Node       `tag:"left"`
	Right *Node       `tag:"right"`
}

// Serialize function
func Serialize(n *Node) string {
	b, _ := json.Marshal(n)
	return string(b)
}

// Deserialize function
func Deserialize(s string) *Node {
	var n Node
	json.Unmarshal([]byte(s), &n)
	return &n
}
