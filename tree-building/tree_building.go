package tree

import (
	"errors"
	"sort"
)

// Record represents a record in a tree
type Record struct {
	ID     int
	Parent int
}

// Node represents a node in a tree
type Node struct {
	ID       int
	Children []*Node
}

// Build constructs a tree from a set of records.
func Build(records []Record) (*Node, error) {
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make(map[int]*Node)
	for i, r := range records {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, errors.New("not in sequence or has bad parent")
		}

		nodes[r.ID] = &Node{ID: r.ID}

		if r.ID != 0 {
			parentNode := nodes[r.Parent]
			parentNode.Children = append(parentNode.Children, nodes[r.ID])
		}
	}

	return nodes[0], nil
}
