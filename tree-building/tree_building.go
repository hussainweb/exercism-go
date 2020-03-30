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
func Build(records []Record) (root *Node, err error) {
	if len(records) == 0 {
		root = nil
		return
	}

	// Let's validate first.
	err = validateRecords(records)
	if err != nil {
		return nil, err
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make(map[int]*Node)
	for _, r := range records {
		nodes[r.ID] = &Node{ID: r.ID}
		if r.ID == 0 {
			continue
		}

		parentNode := nodes[r.Parent]
		currentNode := nodes[r.ID]
		parentNode.Children = append(parentNode.Children, currentNode)
	}

	return nodes[0], nil
}

func validateRecords(records []Record) error {
	nodes := make(map[int]Node)
	largestID := -1
	for _, r := range records {
		if _, ok := nodes[r.ID]; ok {
			return errors.New("Duplicate record")
		}
		if r.ID == 0 && r.Parent != 0 {
			return errors.New("Invalid root")
		}
		if r.ID != 0 && r.Parent >= r.ID {
			return errors.New("Parent cannot be larger than or same as ID")
		}

		nodes[r.ID] = Node{ID: r.ID}
		if r.ID > largestID {
			largestID = r.ID
		}
	}

	if _, ok := nodes[0]; !ok {
		return errors.New("No root node")
	}

	if largestID+1 != len(nodes) {
		return errors.New("The IDs are not sequential")
	}

	// All good, no errors
	return nil
}
