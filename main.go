package main

import (
	"fmt"
	"log"

	"github.com/jgpawletko/rsbe-client-go/rsbe"
)


func main() {
	c := new(rsbe.Config)
	c.BaseURL = "http://localhost:3000"
	c.User = "foo"
	c.Password = "bar"

	rsbe.ConfigureClient(c)

	partners, err := rsbe.PartnerList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("Partners")
	fmt.Printf("%s\n", partners)
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("Partner List Entry")
	fmt.Printf("%s\n", partners[0].ToString())
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("Partner Show Entry")
	partner, err := rsbe.PartnerGet(partners[0].ID)
	fmt.Printf("%s\n", partner.ToString())
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("Collections")
	collections, err := rsbe.PartnerCollectionsList(partner.ID)
	fmt.Printf("%s\n", collections)
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("Collection List Entry")
	fmt.Printf("%s\n", collections[0].ToString())
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("Collection Show Entry")
	collection, err := rsbe.CollectionGet(collections[0].ID)
	fmt.Printf("%s\n", collection.ToString())
	fmt.Println("------------------------------------------------------------------------------")
}
