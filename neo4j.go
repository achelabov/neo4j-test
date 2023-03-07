package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/fatih/structs"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
)

type User struct {
	Name string `json:"name"`
	Lo   int64  `json:"lo"`
	Go   int64  `json:"go"`
}

func handleCreatePartnerRecord(record *db.Record) (*User, error) {
	rawUserNode, found := record.Get("p")
	if !found {
		return nil, fmt.Errorf("could not find column")
	}
	userNode := rawUserNode.(neo4j.Node)
	l, err := neo4j.GetProperty[int64](userNode, "lo")
	if err != nil {
		return nil, err
	}
	g, err := neo4j.GetProperty[int64](userNode, "go")
	if err != nil {
		return nil, err
	}
	name, err := neo4j.GetProperty[string](userNode, "name")
	if err != nil {
		return nil, err
	}

	return &User{Name: name, Lo: l, Go: g}, nil
}

func createUser(ctx context.Context, user *User) neo4j.ManagedTransactionWorkT[*User] {
	return func(tx neo4j.ManagedTransaction) (*User, error) {
		records, err := tx.Run(ctx, "CREATE (p:Partner {name: $name, lo: $lo, go: $go}) RETURN p", map[string]any{
			"name": user.Name,
			"lo":   user.Lo,
			"go":   user.Go,
		})
		if err != nil {
			return nil, err
		}

		record, err := records.Single(ctx)
		if err != nil {
			return nil, err
		}

		return handleCreatePartnerRecord(record)
	}
}

func createUsers(ctx context.Context, usersCount int) neo4j.ManagedTransactionWorkT[[]*User] {
	return func(tx neo4j.ManagedTransaction) ([]*User, error) {
		result, err := tx.Run(ctx, `
		UNWIND RANGE (1,$count) as idx
		CREATE (p:Partner {name: "user" + idx, lo: 100, go: 100})
		RETURN p`, map[string]any{
			"count": usersCount,
		})
		if err != nil {
			return nil, err
		}

		users := make([]*User, usersCount)
		cntr := 0
		for result.Next(ctx) {
			record := result.Record()
			user, err := handleCreatePartnerRecord(record)
			if err != nil {
				return nil, err
			}
			users[cntr] = user
			log.Println(user, "created")
			cntr++
		}

		return users, nil
	}
}

func mapUsers(users []*User) []map[string]interface{} {
	var result = make([]map[string]interface{}, len(users))

	for index, item := range users {
		result[index] = structs.Map(*item)
	}

	return result
}

func createPartner(ctx context.Context, driver neo4j.DriverWithContext, user *User) (*User, error) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	return neo4j.ExecuteWrite(ctx, session, createUser(ctx, user))
}

func createPartners(ctx context.Context, driver neo4j.DriverWithContext, count int) ([]*User, error) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	users := make([]*User, count)

	var err error
	for count > 0 {
		users[count-1], err = neo4j.ExecuteWrite(ctx, session, createUser(ctx, &User{
			Name: "user" + strconv.Itoa(count),
			Lo:   100,
			Go:   100,
		}))
		if err != nil {
			return nil, err
		}
		log.Println(users[count-1], "created")
		count--
	}

	return users, nil
}

func createPartnersUnwind(ctx context.Context, driver neo4j.DriverWithContext, count int) ([]*User, error) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	return neo4j.ExecuteWrite(ctx, session, createUsers(ctx, count))
}

const (
	success = true
)

func createPartnersRelation(ctx context.Context, driver neo4j.DriverWithContext, fromUser, toUser string) error {
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
		MATCH (p1:Partner {name: $u1})
		MATCH (p2:Partner {name: $u2})
		CREATE (p1)-[:HAS_PARTNER]->(p2)`,
			map[string]any{
				"u1": fromUser,
				"u2": toUser,
			})
		if err != nil {
			return nil, err
		}

		return success, nil
	})
	if err != nil {
		return err
	}

	log.Print("created ", fromUser, "-has_partner->", toUser, " relation")
	return nil
}

func createBinaryTreeRelations(ctx context.Context, driver neo4j.DriverWithContext, usersCount int) error {
	for i := 1; i <= usersCount/2; i++ {
		from := "user" + strconv.Itoa(i)

		toLeft := "user" + strconv.Itoa(i*2)
		if err := createPartnersRelation(ctx, driver, from, toLeft); err != nil {
			return err
		}

		toRight := "user" + strconv.Itoa(i*2+1)
		if err := createPartnersRelation(ctx, driver, from, toRight); err != nil {
			return err
		}

	}

	return nil
}

func createBinnaryTreeRelationsLeftUnwind(ctx context.Context, driver neo4j.DriverWithContext, usersCount int) error {
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
		UNWIND RANGE (1,$count) as idx
		MATCH (p1:Partner {name: 'user' + idx})
		MATCH (p2:Partner {name: 'user' + idx * 2})
		CREATE (p1)-[:HAS_PARTNER]->(p2)`,
			map[string]any{
				"count": usersCount / 2,
			})
		if err != nil {
			return nil, err
		}

		return success, nil
	})
	if err != nil {
		return err
	}

	return nil
}

func createBinnaryTreeRelationsRightUnwind(ctx context.Context, driver neo4j.DriverWithContext, usersCount int) error {
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
		UNWIND RANGE (1,$count) as idx
		MATCH (p1:Partner {name: 'user' + idx})
		MATCH (p2:Partner {name: 'user' + idx * 2 + 1})
		CREATE (p1)-[:HAS_PARTNER]->(p2)`,
			map[string]any{
				"count": usersCount / 2,
			})
		if err != nil {
			return nil, err
		}

		return success, nil
	})
	if err != nil {
		return err
	}

	return nil
}

func createBinnaryTreeRelationsUnwind(ctx context.Context, driver neo4j.DriverWithContext, usersCount int) error {
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	if err := createBinnaryTreeRelationsLeftUnwind(ctx, driver, 100); err != nil {
		return err
	}
	if err := createBinnaryTreeRelationsRightUnwind(ctx, driver, 100); err != nil {
		return err
	}

	return nil
}

func createBinaryTree(ctx context.Context, driver neo4j.DriverWithContext, verticesCount int) error {
	_, err := createPartnersUnwind(ctx, driver, verticesCount)
	if err != nil {
		return err
	}

	if err := createBinaryTreeRelations(ctx, driver, verticesCount); err != nil {
		return err
	}

	return nil
}

func getUsers(ctx context.Context, headVertex string, minDepth, maxDepth int) neo4j.ManagedTransactionWorkT[[]*User] {
	return func(tx neo4j.ManagedTransaction) ([]*User, error) {
		result, err := tx.Run(ctx,
			fmt.Sprintf("MATCH (:Partner {name: $name})-[:HAS_PARTNER*%d..%d]->(p:Partner) RETURN p",
				minDepth, maxDepth),
			map[string]any{
				"name": headVertex,
				"min":  minDepth,
				"max":  maxDepth,
			})
		if err != nil {
			return nil, err
		}

		users := make([]*User, 0)
		for result.Next(ctx) {
			record := result.Record()
			user, err := handleCreatePartnerRecord(record)
			if err != nil {
				return nil, err
			}
			users = append(users, user)
		}

		return users, nil
	}
}

func getPartners(ctx context.Context, driver neo4j.DriverWithContext,
	headVertex string, minDepth, maxDepth int) ([]*User, error) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	return neo4j.ExecuteRead(ctx, session, getUsers(ctx, headVertex, minDepth, maxDepth))
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50000*time.Second)
	defer cancel()

	driver, err := neo4j.NewDriverWithContext("neo4j://localhost:7687", neo4j.BasicAuth("neo4j", "password", ""))
	if err != nil {
		log.Fatal(err)
	}
	defer driver.Close(ctx)
	/*
		partners, err := getPartners(ctx, driver, "user1", 0, 16)
		if err != nil {
			log.Fatal(err)
		}
		for _, p := range partners {
			log.Println(p, "received")
		}
		log.Println("users count:", len(partners))
	*/
	/*
		partner, err := createPartner(ctx, driver, &User{Name: "user1", Lo: 100, Go: 100})
		if err != nil {
			log.Fatal(err)
		}
	*/
	startTime := time.Now().UnixMilli()
	_, err = createPartnersUnwind(ctx, driver, 101)
	if err != nil {
		log.Fatal(err)
	}

	if err := createBinnaryTreeRelationsUnwind(ctx, driver, 100); err != nil {
		log.Fatal(err)
	}
	log.Println("execution time:", time.Now().UnixMilli()-startTime)

	/*
		if err := createBinaryTreeRelations(ctx, driver, 50000); err != nil {
			log.Fatal(err)
		}
	*/
	/*
		if err := createBinaryTree(ctx, driver, 10001); err != nil {
			log.Fatal(err)
		}
	*/
	/*
		_, err = createPartnersUnwind(ctx, driver, 101)
		if err != nil {
			log.Fatal(err)
		}
	*/
}
