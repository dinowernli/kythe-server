package main

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	nodesPath       = "/nodes"
	nodePathPattern = "/node/"
)

// TODO(dino): Split this yolo-file up into packages.

func main() {
	fmt.Println("A server for Kythe... under construction")

	// Load the graph store.
	graphStore := &graphStore{}

	// Construct an index from it.
	index := createIndex(graphStore)

	// Create a handler and serve.
	handler := &httpHandler{index: index}
	http.HandleFunc(nodesPath, handler.nodes)
	http.HandleFunc(nodePathPattern, handler.node)
	http.ListenAndServe(":8080", nil)
}

// http

type httpHandler struct {
	index *graphIndex
}

func (h *httpHandler) nodes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world. Nodes: \n")
	names := h.index.List()
	for _, name := range names {
		fmt.Fprintf(w, "<a href=\"%s%s\">%s</a>\n", nodePathPattern, name, name)
	}
}

func (h *httpHandler) node(w http.ResponseWriter, r *http.Request) {
	nodeName := strings.TrimPrefix(r.URL.Path, nodePathPattern)
	fmt.Fprintf(w, "Hello world. Node name: %s", nodeName)
}

// storage

type graphStore struct {
}

// indexing

func createIndex(graphStore *graphStore) *graphIndex {
	return &graphIndex{
		nodes: []*nodeInfo{
			&nodeInfo{name: "name1", message: "m1"},
			&nodeInfo{name: "name2", message: "m2"},
		},
	}
}

type graphIndex struct {
	nodes []*nodeInfo
}

type nodeInfo struct {
	name    string
	message string
}

// List returns the list of node names available in this index.
func (i *graphIndex) List() []string {
	result := []string{}
	for _, node := range i.nodes {
		result = append(result, node.name)
	}
	return result
}

// Get return the available information about the supplied node.
func (i *graphIndex) Get(vname string) (*nodeInfo, error) {
	return &nodeInfo{
		message: "some node",
	}, nil
}
