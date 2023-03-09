package main

import (
	"context"
	"testing"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var driver, _ = neo4j.NewDriverWithContext("neo4j://localhost:7687", neo4j.BasicAuth("neo4j", "password", ""))
var ctx = context.Background()

// 10k tree
func BenchmarkGetPartnersFrom1To5Lvl10kVerticesTree(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getPartners(ctx, driver, "user1", 1, 5)
	}
}

func BenchmarkGetPartnersFrom5To10Lvl10kVerticesTree(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getPartners(ctx, driver, "user1", 5, 10)
	}
}
func BenchmarkGetPartnersFrom10To14Lvl10kVerticesTree(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getPartners(ctx, driver, "user1", 10, 14)
	}
}

func BenchmarkGetPartnersFrom1To14Lvl10kVerticesTree(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getPartners(ctx, driver, "user1", 1, 14)
	}
}

// 100k tree
func BenchmarkGetPartnersFrom16To16Lvl100kVerticesTree(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getPartners(ctx, driver, "user1", 16, 16)
	}
}

func BenchmarkGetPartnersFrom13To16Lvl100kVerticesTree(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getPartners(ctx, driver, "user1", 13, 16)
	}
}

func BenchmarkGetPartnersFrom10To13Lvl100kVerticesTree(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getPartners(ctx, driver, "user1", 10, 13)
	}
}

func BenchmarkGetPartnersFrom7To10Lvl100kVerticesTree(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getPartners(ctx, driver, "user1", 7, 10)
	}
}

func BenchmarkGetPartnersFrom1To16Lvl100kVerticesTree(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getPartners(ctx, driver, "user1", 1, 16)
	}
}

// bonus
func BenchmarkGetMainBonusBronze(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := GetMainBonus(ctx, driver, "user1", 2, 2)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkGetMainBonusBronzePro(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := GetMainBonus(ctx, driver, "user1", 2, 3)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkGetMainBonusSilver(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := GetMainBonus(ctx, driver, "user1", 2, 4)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkGetMainBonusSilverPro(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := GetMainBonus(ctx, driver, "user1", 2, 5)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkGetMainBonusGold(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := GetMainBonus(ctx, driver, "user1", 2, 6)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkGetMainBonusGoldPro(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := GetMainBonus(ctx, driver, "user1", 2, 7)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkGetMainBonusPlatinum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := GetMainBonus(ctx, driver, "user1", 2, 8)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkGetMainBonusPlatinimPro(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := GetMainBonus(ctx, driver, "user1", 2, 9)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkGetMainBonusDiamond(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := GetMainBonus(ctx, driver, "user1", 2, 10)
		if err != nil {
			b.Fail()
		}
	}
}
