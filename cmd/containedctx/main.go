package main

import (
	"containedctx"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(containedctx.Analyzer) }
