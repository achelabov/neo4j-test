package main

import (
	"context"
	"testing"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var driver, _ = neo4j.NewDriverWithContext("neo4j://localhost:7687", neo4j.BasicAuth("neo4j", "password", ""))

func BenchmarkCreatePartners(b *testing.B) {
	createBinaryTree(context.Background(), driver, 1001)
}
