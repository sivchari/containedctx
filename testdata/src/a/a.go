package a

import "context"

type ok struct {
	i int
	s string
}

type ng struct {
	ctx context.Context // want "found a struct that contains a Context field"
}

type empty struct{}
