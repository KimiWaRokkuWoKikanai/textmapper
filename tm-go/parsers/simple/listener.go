// generated by Textmapper; DO NOT EDIT

package simple

import (
	"fmt"
)

type NodeType int

type Listener func(t NodeType, offset, endoffset int)

const (
	NoType NodeType = iota
	NodeTypeMax
)

var nodeTypeStr = [...]string{
	"NONE",
}

func (t NodeType) String() string {
	if t >= 0 && int(t) < len(nodeTypeStr) {
		return nodeTypeStr[t]
	}
	return fmt.Sprintf("node(%d)", t)
}
