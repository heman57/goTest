package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/html"
)

func main() {

	db, err := sql.Open("mysql", "user:root@tcp(127.0.0.1:3306)/db")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10)
	isNewURL := true
	rankAdd, err := db.Prepare("INSERT INTO ranks(id, Person, Occupation, Rating) VALUES(?,?,?,?)")

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	for isNewURL {
		resp, err := client.Get("https://rejstrik.penize.cz/#personRating")
		if err != nil {
			fmt.Println("ERROR: Failed to crawl")
			return
		}
		b := resp.Body
		defer b.Close() // close Body when the function returns
		node, _ := html.Parse(b)
		document := goquery.NewDocumentFromNode(node)
		document.Find("div").Each(func(i int, s *goquery.Selection) {
			class, _ := s.Attr("id")
			if strings.HasPrefix(class, "rating") {
				fmt.Println("ID: ", class)
				var rank string
				row := db.QueryRow(`SELECT id FROM ranks where id like ? limit 1;`, class)
				switch err := row.Scan(&rank); err {
				case sql.ErrNoRows:
					fmt.Println("New record:", class)
					person := s.Find("a").Text()
					fmt.Println("Person: ", person)
					occupation := s.Find("p").Text()
					occupation = occupation[len(person)+3 : len(occupation)]
					fmt.Println("Occupation: ", occupation)
					ratingText := s.Find("form").Text()
					ratingText = ratingText[0 : strings.Count(ratingText, "")-3]
					fmt.Println("Rating: ", strings.Replace(ratingText, "+", "", -1))
					rankAdd.Exec(class, person, occupation, strings.Replace(ratingText, "+", "", -1))

				case nil:
					fmt.Println("Is already in:", class)

				default:
					panic(err)
				}

			}
		})
	}

	//END

}
