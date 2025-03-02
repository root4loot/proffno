package main

import (
	"fmt"

	"github.com/root4loot/proffno"
)

func main() {
	// Get subsidiaries with ownership percentage greater than 50%
	subs, _ := proffno.GetSubsidiaries("DnB Bank ASA", 50)

	for _, sub := range subs {
		fmt.Println(sub.Name)
	}
}
