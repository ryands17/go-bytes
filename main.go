package main

import (
	"fmt"
	"log"
	"slices"
	"time"

	"github.com/ryands17/go-bytes/cmd/brands"
	"github.com/ryands17/go-bytes/cmd/builders"
	"github.com/ryands17/go-bytes/cmd/features"
	"github.com/ryands17/go-bytes/cmd/iterators"
	"github.com/ryands17/go-bytes/cmd/routines"
	"github.com/ryands17/go-bytes/cmd/structures"
	"github.com/ryands17/go-bytes/cmd/utils"
	"github.com/ryands17/go-bytes/cmd/utils/bitmasks"
)

func main() {
	// using sets
	s := structures.NewSet[int]()

	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(1)

	fmt.Println("Size of set", s.Size())

	// branded types
	user := brands.GeneralUser{
		ID:       "1",
		Name:     "User 1",
		UserType: brands.Admin,
	}

	adminUser, err := brands.IsAdmin(user)
	if err != nil {
		log.Fatalf("User is not an admin! Exiting.")
	}
	utils.PrintJSON(adminUser)

	// copying struct fields via reflection
	var user2 brands.GeneralUser
	utils.CopyStructFields(&user2, &user)
	fmt.Println("Copied user")
	utils.PrintJSON(user2)

	// goroutines fetch example
	routines.FetchAllUsers(5)

	// bitmasks example
	permissions := bitmasks.READ
	fmt.Println("Has read access:", permissions.Has(bitmasks.READ))

	permissions.Set(bitmasks.WRITE)
	fmt.Println("Has read access:", permissions.Has(bitmasks.READ))
	fmt.Println("Has write access:", permissions.Has(bitmasks.WRITE))

	permissions.Toggle(bitmasks.READ)
	fmt.Println("Has read access:", permissions.Has(bitmasks.READ))

	// convert value to pointer
	userPtr := utils.PointerTo(user)
	utils.PrintJSON(*userPtr)

	// marshalling struct to JSON manually
	jsonBytes, _ := utils.MarshalStruct(user)
	fmt.Println("Marshalled JSON:", string(jsonBytes))

	// iterators with sequences
	primeNumbers := iterators.Primes(iterators.NonZeroIntegers(100))
	utils.PrintJSON(slices.Collect(primeNumbers))

	// create a database client (the options pattern)
	db1 := builders.NewDBClient(
		builders.WithUrl("localhost:5432"),
		builders.WithConnections(10),
		builders.WithTimeout(10*time.Second),
	)
	db1.Connect()

	// create a database client (the fluent pattern)
	db2 := builders.NewDBClientFluent().
		WithURL("localhost:5432").
		WithConnections(10).
		WithTimeout(10 * time.Second).
		Build()
	db2.Connect()

	// in-memory cache with TTL
	cache := utils.NewCache(3*time.Second, nil)

	if err = cache.Set("key1", "value1", nil); err != nil {
		log.Fatalf("Failed to set cache: %v", err)
	}

	if value, found := cache.Get("key1"); found {
		fmt.Println("Cache hit for key1: ", value)
	}

	// expire key1 after 3 seconds
	time.Sleep(3 * time.Second)
	if _, found := cache.Get("key1"); !found {
		fmt.Println("key1 has expired from cache")
	}

	// build flags example
	// to get premium features, run: go run -tags=premium main.go or go build -tags=premium
	fmt.Printf("Features available: %+v\n", features.AvailableFeatures())
}
