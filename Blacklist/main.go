package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func getEmail(db *sql.DB, email string, client string) (isIn bool, status string, created time.Time) {
	row := db.QueryRow(`SELECT email, status, created FROM blacklist WHERE email like ? and client like ?;`, email, client)
	switch err := row.Scan(&email, &status, &created); err {
	case sql.ErrNoRows:
		isIn = false
	case nil:
		isIn = true
	default:
		panic(err)
	}
	return
}

func main() {
	db, err := sql.Open("mysql", "user:root@tcp(127.0.0.1:3306)/blacklist?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10)

	blckIn, err := db.Prepare("INSERT INTO blacklist(email,status,client,created) VALUES(?,?,?,?)")
	delblck, err := db.Prepare(`DELETE FROM blacklist where email like ? limit 1;`)
	//updfnd, err := db.Prepare("UPDATE indexed SET count = ? WHERE url = ?")

	type jsonInput struct {
		Email  string `json:"email"`
		Status string `json:"status"`
		Client string `json:"clientID"`
	}
	var postInput jsonInput

	APIRoot := func(w http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body)

		if strings.Split(req.URL.Path, "/")[1] == "blacklist" {
			switch req.Method {
			case "GET":
				email := strings.Split(req.URL.Path, "/")[2]
				client := "1"
				isIn, status, created := getEmail(db, email, client)
				if isIn {
					w.Write([]byte(`{"status": "` + status + `", "email": "` + email + `", "created": "` + created.String() + `"}`))
				} else {
					email := "@" + strings.Split(email, "@")[1]
					isIn, status, created := getEmail(db, email, client)
					if isIn {
						w.Write([]byte(`{"status": "` + status + `", "email": "` + email + `", "created": "` + created.String() + `"}`))
					} else {
						w.Write([]byte(`{"status": "not in blacklist", "email": "` + email + `"}`))
					}
				}
			case "POST":
				if len(body) > 3 {
					er := json.Unmarshal(body, &postInput)
					if er != nil {
						log.Fatal(er)
						w.WriteHeader(http.StatusBadRequest)
						w.Write([]byte(`{"error": {"code": "400", "status": "Bad request", "message": "Failed to add email to black list, data input is not correct."}}`))
						return
					}
					email := postInput.Email
					client := postInput.Client
					isIn, _, created := getEmail(db, email, client)
					if isIn {
						w.Write([]byte(`{"status": "Email was already added to blacklist", "email": "` + email + `", "created": "` + created.String() + `"}`))
					} else {
						time := time.Now()
						status := postInput.Status
						client := postInput.Client
						blckIn.Exec(email, status, client, time)
						w.Write([]byte(`{"status": "Email was added to blacklist", "email": "` + email + `"}`))
					}
				} else {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(`{"error": {"code": "400", "status": "Bad request", "message": "Failed to add email to black list, data input is not correct."}}`))
					return
				}
			case "PUT":
				fmt.Println("Not implemented.")
			case "DELETE":
				email := strings.Split(req.URL.Path, "/")[2]
				client := "1"
				isIn, _, _ := getEmail(db, email, client)
				if isIn {
					delblck.Exec(email, client)
					w.Write([]byte(`{"status": "deleted", "email": "` + email + `"}`))

				} else {
					w.Write([]byte(`{"status": "Email is not in blacklist", "email": "` + email + `"}`))
				}
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte(`{"error": {"code": "405", "status": "Forbiden", "message": "Method not allowed!"}}`))
			}
		} else {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"error": {"code": "403", "status": "Forbiden", "message": "Method not allowed!"}}`))
		}
	}
	http.HandleFunc("/", APIRoot)
	http.ListenAndServe(":80", nil)
}
