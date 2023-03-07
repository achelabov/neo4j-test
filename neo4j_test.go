package main

import (
	"context"
	"testing"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var driver, _ = neo4j.NewDriverWithContext("neo4j://localhost:7687", neo4j.BasicAuth("neo4j", "password", ""))

func BenchmarkGetPartnersFrom1To5Lvl(b *testing.B) {
	ctx := context.Background()
	for n := 0; n < b.N; n++ {
		getPartners(ctx, driver, "user1", 1, 5)
	}
}

func BenchmarkGetPartnersFrom5To10Lvl(b *testing.B) {
	ctx := context.Background()
	for n := 0; n < b.N; n++ {
		getPartners(ctx, driver, "user1", 5, 10)
	}
}
func BenchmarkGetPartnersFrom10To14Lvl(b *testing.B) {
	ctx := context.Background()
	for n := 0; n < b.N; n++ {
		getPartners(ctx, driver, "user1", 10, 14)
	}
}

func BenchmarkGetPartnersFrom1To14Lvl(b *testing.B) {
	ctx := context.Background()
	for n := 0; n < b.N; n++ {
		getPartners(ctx, driver, "user1", 1, 14)
	}
}

func BenchmarkCreate10kPartners(b *testing.B) {
	ctx := context.Background()
	createPartnersUnwind(ctx, driver, 10000)
}

func BenchmarkCreate100BinaryTreeRelations(b *testing.B) {
	ctx := context.Background()
	for n := 0; n < b.N; n++ {
		createBinnaryTreeRelationsUnwind(ctx, driver, 100)
	}
}
