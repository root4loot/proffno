package main

import (
	"fmt"

	"github.com/root4loot/proffno"
)

func main() {
	// Get subsidiaries of a company and its sub-subsidiaries owned by more than 50%
	subs, _ := proffno.GetSubsidiaries("DnB Bank ASA", 2, 50)

	for _, sub := range subs {
		fmt.Println(sub.Name)
	}
}
