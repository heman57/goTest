package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type jsonPOSTInput struct {
	Email  string `json:"email"`
	Status string `json:"status"`
	Client string `json:"clientID"`
}

type jsonPUTInput struct {
	Email  []string `json:"emails"`
	Status string   `json:"status"`
	Client string   `json:"clientID"`
}

func getEmail(db *sql.DB, email string, client string) (isIn bool, status string, created time.Time) {
	row := db.QueryRow(`SELECT email, status, created FROM blacklist WHERE email like ? and client like ?;`, email, client)
	switch err := row.Scan(&email, &status, &created); err {
	case sql.ErrNoRows:
		isIn = false
	case nil:
		isIn = true
	default:
		isIn = false
		panic(err)
	}
	return
}

func getEmails(db *sql.DB, client string) (emails []string, statusArray []string, createdArray []time.Time) {
	var email string
	var status string
	var created time.Time
	rows, err := db.Query(`SELECT email, status, created FROM blacklist WHERE client = ?;`, client)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&email, &status, &created)
		if err != nil {
			panic(err)
		}
		emails = append(emails, email)
		statusArray = append(statusArray, status)
		createdArray = append(createdArray, created)
		fmt.Println(email, status, created)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return
}

func resolveGet(requestURL []string, client string, db *sql.DB, w http.ResponseWriter) {
	if len(requestURL) > 2 {
		// email address is specified
		email := requestURL[2]
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if re.MatchString(email) {
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
		} else {
			w.Write([]byte(`{"status": "not in blacklist", "email": "` + email + `"}`))
		}
	} else {
		// email address is not specified all emails in response
		emails, status, created := getEmails(db, client)
		outJSON := `{"emails":[`
		for index := range emails {
			outJSON += `{"status": "` + status[index] + `", "email": "` + emails[index] + `", "created": "` + created[index].String() + `"},`
		}
		if outJSON[len(outJSON)-1:] == "," {
			outJSON = strings.TrimSuffix(outJSON, ",")
		}
		outJSON += `]}`
		w.Write([]byte(outJSON))
	}
}

func resolvePost(body []byte, postInput jsonPOSTInput, db *sql.DB, w http.ResponseWriter) {
	blckIn, _ := db.Prepare("INSERT INTO blacklist(email,status,client,created) VALUES(?,?,?,?)")
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
}

func resolvePut(body []byte, postInput jsonPUTInput, db *sql.DB, w http.ResponseWriter) {
	blckIn, _ := db.Prepare("INSERT INTO blacklist(email,status,client,created) VALUES(?,?,?,?)")
	if len(body) > 3 {
		er := json.Unmarshal(body, &postInput)
		if er != nil {
			log.Fatal(er)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": {"code": "400", "status": "Bad request", "message": "Failed to add email to black list, data input is not correct."}}`))
			return
		}
		outJSON := `{"items": [`
		email := postInput.Email
		client := postInput.Client
		for index := range email {
			isIn, _, _ := getEmail(db, email[index], client)
			if isIn {
				outJSON += `{ "email": "` + email[index] + `", "status": "Email was already added to blacklist"},`
			} else {
				time := time.Now()
				status := postInput.Status
				client := postInput.Client
				blckIn.Exec(email[index], status, client, time)
				outJSON += `{"email": "` + email[index] + `","status": "Email was added to blacklist"},`
			}
		}
		if outJSON[len(outJSON)-1:] == "," {
			outJSON = strings.TrimSuffix(outJSON, ",")
		}
		outJSON += `]}`
		w.Write([]byte(outJSON))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": {"code": "400", "status": "Bad request", "message": "Failed to add email to black list, data input is not correct."}}`))
		return
	}
}

func resolveDelete(requestURL []string, client string, db *sql.DB, w http.ResponseWriter) {
	delblck, _ := db.Prepare(`DELETE FROM blacklist  WHERE email like ? and client like ?;`)
	email := requestURL[2]
	isIn, _, _ := getEmail(db, email, client)
	if isIn {
		delblck.Exec(email, client)
		w.Write([]byte(`{"status": "deleted", "email": "` + email + `"}`))
	} else {
		w.Write([]byte(`{"status": "Email is not in blacklist", "email": "` + email + `"}`))
	}
}

func main() {
	db, err := sql.Open("mysql", "user@blacklist557:Admmin147@tcp(blacklist557.mysql.database.azure.com)/blacklist?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10)
	var postInput jsonPOSTInput
	var putInput jsonPUTInput
	//main function
	APIRoot := func(w http.ResponseWriter, req *http.Request) {
		var client string
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": {"code": "500", "status": "Error", "message": "Request failed."}}`))
			return
		} else {

			requestURL := strings.Split(req.URL.Path, "/")
			if requestURL[1] == "blacklist" {
				if len(req.Header["Clientid"]) > 0 {
					client = req.Header["Clientid"][0]
				} else {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(`{"error": {"code": "400", "status": "Bad request", "message": "Client ID missing, data input is not correct."}}`))
					return
				}
				switch req.Method {
				case "GET":
					resolveGet(requestURL, client, db, w)
				case "POST":
					resolvePost(body, postInput, db, w)
				case "PUT":
					resolvePut(body, putInput, db, w)
				case "DELETE":
					resolveDelete(requestURL, client, db, w)

				default:
					w.WriteHeader(http.StatusMethodNotAllowed)
					w.Write([]byte(`{"error": {"code": "405", "status": "Forbiden", "message": "Method not allowed!"}}`))
				}
			} else {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte(`{"error": {"code": "403", "status": "Forbiden", "message": "Access denied!"}}`))
			}
		}
	}
	http.HandleFunc("/", APIRoot)
	http.ListenAndServe(":80", nil)
}
