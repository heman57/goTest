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
	//	rankAdd, err := db.Prepare("INSERT INTO ranks(id, Person, Occupation, Rating) VALUES(?,?,?,?)")

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	domainAdd, err := db.Prepare("INSERT INTO newdomains (domain, status, clength, description, title, heading, header, indexed) VALUES (?,?,?,?,?,?,?,?)")
	domainAddFail, err := db.Prepare("INSERT INTO newdomains (domain, status, indexed) VALUES(?,?,?)")
	delnew, err := db.Prepare(`DELETE FROM newdomains where domain like ? limit 1;`)

	for isNewURL {

		var newurl string
		row := db.QueryRow(`SELECT domain FROM newdomains where indexed = 'f' limit 1;`)
		switch err := row.Scan(&newurl); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
			isNewURL = false
		case nil:
			fmt.Println("Ready to go:", newurl)

			delnew.Exec(newurl)
		default:
			panic(err)
		}

		resp, err := client.Get(newurl)
		if err != nil {
			fmt.Println("ERROR: Failed to crawl, not exist")
			domainAddFail.Exec(newurl, "Read failiure", "t")
			//return
		} else {
			s := resp.Status        //string
			h := resp.Header        //http.header
			b := resp.Body          // io.read
			l := resp.ContentLength //int
			desc := ""

			defer b.Close() // close Body when the function returns
			node, _ := html.Parse(b)
			document := goquery.NewDocumentFromNode(node)

			fmt.Println(node, "------------------")
			if node != nil {
				document.Find("h1").Text()
				fmt.Println("Status: ", s)
				fmt.Println(document.Find("title").Text())
				fmt.Println("H1: ", document.Find("h1").Text())
				fmt.Println("Header: ", h)
				fmt.Println("Clength: ", l)
				document.Find("meta").Each(func(i int, s *goquery.Selection) {
					class, _ := s.Attr("name")
					cntnt, _ := s.Attr("content")
					if class == "Description" {
						fmt.Println(class, " > ", cntnt)
						desc = cntnt
					}
				})
				var hdr string
				for key, value := range h {
					hdr = hdr + "[" + key + " = " + strings.Join(value, ", ") + "]"
				}
				domainAdd.Exec(newurl, s, l, desc, document.Find("title").Text(), document.Find("h1").Text(), hdr, "t")
			}
			fmt.Println("-------------------------------------------------")
		}
	}

	//END

}
