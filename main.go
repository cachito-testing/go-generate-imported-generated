package main

import "github.com/cachito-testing/go-generate-imported-generated/foobar"

//go:generate go run internal/generate/generatefoobar.go

func main() {
	foobar.Output()
}
