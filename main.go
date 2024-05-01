package main

import (
	"flag"
	"log"
)

func main() {
	auth := flag.String("auth", "", "Authorization header.")

	flag.Parse()

	client := NewClient(*auth)

	categories := []int{
		// CategoryAlpinbriller,
		// CategoryAlpinhjelm,
		CategoryAlpinsko,
		CategoryAlpinskiVoksen,
		CategoryAlpinskiBarn,
		CategoryAlpinski,
		// CategoryAlpinstaver,
	}

	loans := make(map[int][]string)

	for _, category := range categories {
		err := client.GetLoansWithEquipment(loans, category, StatePublished, StatusBorrowed)
		if err != nil {
			log.Println("[ERROR]", "GetEquipment", category, err)
		}
	}

	for loanID, equipment := range loans {
		log.Printf("%d: %v\n", loanID, equipment)
	}
	log.Println(len(loans))
}
