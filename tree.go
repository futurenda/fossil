package main

import (
	"path"
	"strings"
)

type node struct {
	Name     string
	Contents []byte
	Children map[string]*node
}

func appendNode(root *node, pathSlice []string, contents []byte) {
	if len(pathSlice) == 1 {
		key := pathSlice[0]
		root.Children[key].Name = key
		root.Children[key].Contents = contents
		return
	}
	dir := pathSlice[0]
	if root.Children[dir] == nil {
		root.Children[dir] = &node{
			Name:     dir,
			Children: make(map[string]*node),
		}
	}
	appendNode(root.Children[dir], pathSlice[1:], contents)
}

func buildTree(kv map[string][]byte) *node {
	root := &node{
		Children: make(map[string]*node),
	}

	for filepath, contents := range kv {
		filepath = path.Clean(filepath)
		appendNode(root, strings.Split(filepath, "/"), contents)
	}

	return root
}
