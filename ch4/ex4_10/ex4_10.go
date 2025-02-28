package ex410

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Suuu775/gopl/ch4/github"
)

func Issues() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s",
			item.Number, item.User.Login, item.Title)
		duration := time.Since(item.CreatedAt)
		category := "over a year"
		if duration < 30*24*time.Hour {
			category = "less than a month"
		} else if duration < 365*24*time.Hour {
			category = "less than a year"
		}
		fmt.Printf(" %s", category)
	}
}
