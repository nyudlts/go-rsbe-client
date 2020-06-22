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

	partners, err := rsbe.PartnerIndex()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", partners)

	id := partners[0].Id
	
	partner, err := rsbe.PartnerShow(id)

	fmt.Printf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%d\n%s\n",
		partner.Id, partner.Code, partner.Name, partner.Created_at,
		partner.Updated_at, partner.PartnersURL, partner.CollectionsURL,
		partner.LockVersion, partner.RelPath)
}
