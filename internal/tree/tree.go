package tree

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type Node struct {
	Name           string
	Id             string
	StartTimestamp pcommon.Timestamp
	EndTimestamp   pcommon.Timestamp
	Children       []*Node
}

type Forest struct {
	Roots []*Node
}

func NewNode(span ptrace.Span) *Node {
	return &Node{
		Name:           span.Name(),
		Id:             span.SpanID().String(),
		StartTimestamp: span.StartTimestamp(),
		EndTimestamp:   span.EndTimestamp(),
	}
}

func NewNodeWithId(id string) *Node {
	return &Node{
		Id: id,
	}
}

func (n *Node) Update(span ptrace.Span) {
	n.Name = span.Name()
	n.StartTimestamp = span.StartTimestamp()
	n.EndTimestamp = span.EndTimestamp()
}

func NewForest() *Forest {
	return &Forest{
		Roots: nil,
	}
}

func (f *Forest) AddSpans(spans ptrace.SpanSlice) {
	lookup := make(map[string]*Node)

	for i := 0; i < spans.Len(); i++ {
		span := spans.At(i)
		spanId := span.SpanID().String()
		currNode, currExist := lookup[spanId]
		if currExist {
			currNode.Update(span)
		} else {
			currNode = NewNode(span)
			lookup[spanId] = currNode
		}

		parentId := span.ParentSpanID().String()
		if span.ParentSpanID().IsEmpty() {
			f.Roots = append(f.Roots, currNode)
		} else {
			parentNode, parentExist := lookup[parentId]
			if !parentExist {
				parentNode = NewNodeWithId(parentId)
				lookup[parentId] = parentNode
			}
			parentNode.Children = append(parentNode.Children, currNode)
		}
	}
}
