package main

import (
	"fmt"
	"log"

	"github.com/ryands17/go-bytes/cmd/brands"
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
}
