package main

import (
	"fmt"
	"net/http"
)

const (
	nodesPath = "/nodes"
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
	http.ListenAndServe(":8080", nil)
}

// http

type httpHandler struct {
	index *graphIndex
}

func (h *httpHandler) nodes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nodes: %+v", h.index.List())
}

// storage

type graphStore struct {
}

// indexing

func createIndex(graphStore *graphStore) *graphIndex {
	return &graphIndex{}
}

type graphIndex struct {
}

type nodeInfo struct {
	message string
}

func (i *graphIndex) List() []string {
	return []string{"node1", "node2"}
}

func (i *graphIndex) Get(vname string) (*nodeInfo, error) {
	return &nodeInfo{
		message: "some node",
	}, nil
}
