package main

import (
	
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/html"
)

func getHref(t html.Token) (ok bool, href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}
	return
}

func crawl(url string, ch chan string, chFinished chan bool) {

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	//client.Get(url)
	resp, err := client.Get(url)

	defer func() {
		chFinished <- true
	}()
	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return
	}
	b := resp.Body
	defer b.Close() // close Body when the function returns
	z := html.NewTokenizer(b)
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := z.Token()
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}
			ok, url := getHref(t)
			if !ok {
				continue
			}
			hasProto := strings.Index(url, "http") == 0
			if hasProto {
				ch <- url
			}
		}
	}
}

func uniqueCZ(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[strings.TrimSpace(entry)]; !value {
			keys[entry] = true
			czArray := strings.SplitAfter(entry, ".")
			if strings.HasPrefix(czArray[len(czArray)-1], "cz") {
				list = append(list, strings.TrimSpace(entry))
			}
		}
	}
	return list
}

func main() {
	db, err := sql.Open("mysql", "user:root@tcp(127.0.0.1:3306)/db")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10)
	isNewURL := true
	var startURL string
	startURL = "http://www.seznam.cz"
	stall, err := db.Prepare("INSERT INTO allurls(url) VALUES(?)")
	stnew, err := db.Prepare("INSERT INTO newurls(url) VALUES(?)")
	stindx, err := db.Prepare("INSERT INTO indexed(url) VALUES(?)")
	delnew, err := db.Prepare(`DELETE FROM newurls where url like ? limit 1;`)
	//updfnd, err := db.Prepare("UPDATE indexed SET count = ? WHERE url = ?")

	//START
	for isNewURL {
		//for q := 0; q < 3; q++ {
		foundURLs := make(map[string]bool)
		var Urls []string
		chUrls := make(chan string)
		chFinished := make(chan bool)
		var URLNoFolders []string

		go crawl(startURL, chUrls, chFinished)
		fmt.Println("!!! Now proceed", startURL)
		for c := 0; c < 1; {
			select {
			case url := <-chUrls:
				foundURLs[url] = true
			case <-chFinished:
				c++
			}
		}

		for url := range foundURLs {
			fmt.Println("Founded: ", url)

			QueryOutURL := strings.SplitAfter(url, "&")
			SplitedURL := strings.SplitAfter(QueryOutURL[0], "//")

			if len(SplitedURL) > 1 {
				URLNoFolders = strings.Split(SplitedURL[1], "/")
			} else {
				URLNoFolders = strings.Split(SplitedURL[0], "/")
			}

			URLFinal := strings.SplitAfter(URLNoFolders[0], ".")
			URLString := ""
			if len(URLFinal) > 2 {
				if len(URLFinal[len(URLFinal)-2]) > 0 {
					URLString = URLFinal[len(URLFinal)-2]
				}
			}

			if len(URLFinal) > 1 {
				if len(URLFinal[len(URLFinal)-1]) > 1 {
					URLString = URLString + URLFinal[len(URLFinal)-1][0:2]
				}
			}

			if strings.HasPrefix(URLNoFolders[0], "www") {
				URLString = URLFinal[0] + URLString
			}

			FinalURL := SplitedURL[0] + URLString
			Urls = append(Urls, FinalURL)
		}

		uniqueUrls := uniqueCZ(Urls)
		fmt.Println(uniqueUrls)
		for item := range uniqueUrls {
			row := db.QueryRow(`SELECT url FROM allurls WHERE url like ?;`, uniqueUrls[item])

			switch err := row.Scan(&uniqueUrls[item]); err {
			case sql.ErrNoRows:
				fmt.Println("No rows were returned! New inserted", uniqueUrls[item])
				stnew.Exec(uniqueUrls[item])
				stall.Exec(uniqueUrls[item])
			case nil:
				fmt.Println("Already founded", uniqueUrls[item])

			default:
				panic(err)
			}

		}
		stindx.Exec(startURL)
	

		var newurl string
		row := db.QueryRow(`SELECT url FROM newurls limit 1;`)
		switch err := row.Scan(&newurl); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
			//isNewURL = false
		case nil:
			fmt.Println("Ready to go:", newurl)
			startURL = newurl
			delnew.Exec(startURL)
		default:
			panic(err)
		}

		//startURL = uniqueUrls[0]

		close(chUrls)

		defer db.Close()
	}

	//END

}
