package node

import (
	"encoding/json"
	"io"
)

type Node struct {
	stdin io.Reader
}

type Message struct {
	Src  string
	Dest string
	Body json.RawMessage
}

func NewNode(stdin io.Reader) *Node {
	return &Node{stdin: stdin}
}
