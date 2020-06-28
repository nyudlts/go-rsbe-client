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
	fmt.Println("SEs List Entry")
	ses, err := rsbe.CollectionSEsList(collection.ID)
	fmt.Printf("%s\n", ses[0].ToString())
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("SE Show Entry")
	se, err := rsbe.SEGet(ses[0].ID)
	fmt.Printf("%s\n", se.ToString())
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("SE FMD List Entry")
	fmds, err := rsbe.SEFMDsList(se.ID)
	fmt.Printf("%v\n", fmds)
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Printf("%s\n", fmds[0].ToString())
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("FMD Show Entry")
	fmd, err := rsbe.FMDGet(fmds[0].ID)
	fmt.Printf("%s\n", fmd.ToString())
	fmt.Println("------------------------------------------------------------------------------")
	pe := rsbe.PartnerEntry{
		Code:    "bork",
		Name:    "Bork Partner",
		RelPath: "bork",
	}
	fmt.Println("PartnerCreate: Before")
	fmt.Printf("%v", pe)
	err = rsbe.PartnerCreate(&pe)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PartnerCreate: After")
	fmt.Printf("%v", pe)
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("PartnerUpdate BEFORE:")
	partner, err = rsbe.PartnerGet(pe.ID)
	fmt.Printf("%s\n", partner.ToString())
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("PartnerUpdate:")
	partner.Code = "goofy"
	partner.Name = "GOOFY Partner"
	err = rsbe.PartnerUpdate(&partner)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PartnerUpdate: AFTER")
	fmt.Printf("%v", partner)
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("PartnerGet:")
	partner, err = rsbe.PartnerGet(partner.ID)
	fmt.Printf("%s\n", partner.ToString())
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("PartnerDelete:")
	err = rsbe.PartnerDelete(partner.ID)
	fmt.Printf("%v", partner)
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("PartnerGet:")
	partner, err = rsbe.PartnerGet(partner.ID)
	fmt.Printf("%s\n", partner.ToString())
	fmt.Println("------------------------------------------------------------------------------")
}
