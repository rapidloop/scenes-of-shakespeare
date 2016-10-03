package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var db *sql.DB

type result struct {
	Work        string
	WorkID      string
	Act         int
	Scene       int
	Description string
	Snippet     template.HTML
}

func main() {
	// get the port to listen on
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// get the detail of the database to connect to
	dburl := os.Getenv("DATABASE_URL")
	if dburl == "" {
		log.Fatal("DATABASE_URL env var not set")
	}

	// connect to the database
	var err error
	if db, err = sql.Open("postgres", dburl); err != nil {
		log.Fatal(err)
	}

	// compile the templates
	tplHome := template.Must(template.New(".").Parse(tplStrHome))
	tplResults := template.Must(template.New(".").Parse(tplStrResults))
	tplScene := template.Must(template.New(".").Parse(tplStrScene))

	// handler to render a scene page
	http.HandleFunc("/scene", func(w http.ResponseWriter, r *http.Request) {
		// get parameters and do some basic validation
		a := r.FormValue("a")
		if v, err := strconv.Atoi(a); err != nil || v < 0 || v > 5 {
			http.Error(w, "not found", 404)
			return
		}
		s := r.FormValue("s")
		if v, err := strconv.Atoi(a); err != nil || v < 0 || v > 15 {
			http.Error(w, "not found", 404)
			return
		}
		wid := r.FormValue("w")
		if len(wid) < 5 || len(wid) > 14 {
			http.Error(w, "not found", 404)
			return
		}
		// fetch the work title
		var title string
		if err := db.QueryRow(sqlGetWork, wid).Scan(&title); err != nil {
			http.Error(w, "not found", 404)
			return
		}
		// fetch the scene description and body text
		var desc, body string
		if err := db.QueryRow(sqlGetScene, wid, a, s).Scan(&desc, &body); err != nil {
			http.Error(w, "not found", 404)
			return
		} else {
			body = strings.Replace(body, "\n\n", "<p>", -1)
			body = strings.Replace(body, "\n", "<br>", -1)
			tplScene.Execute(w, map[string]interface{}{
				"Work":        title,
				"Act":         a,
				"Scene":       s,
				"Description": desc,
				"Body":        template.HTML(body),
			})
		}
	})

	// handler to render home page and query results pages
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.FormValue("q")
		if q == "" {
			tplHome.Execute(w, nil)
			return
		}
		if len(q) > 100 {
			q = q[:100]
		}
		rows, err := db.Query(sqlSearch, q)
		if err != nil {
			http.Error(w, "not found", 404)
			return
		}
		defer rows.Close()
		results := make([]result, 0, 10)
		for rows.Next() {
			var r result
			var snip string
			if err := rows.Scan(&r.Work, &r.WorkID, &r.Act, &r.Scene, &r.Description, &snip); err != nil {
				http.Error(w, "not found", 404)
				return
			}
			r.Snippet = template.HTML(strings.Replace(snip, "\n", "<br>", -1))
			results = append(results, r)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, "not found", 404)
			return
		}
		tplResults.Execute(w, map[string]interface{}{
			"Results": results,
			"Query":   q,
		})
	})

	// start the http server
	http.ListenAndServe(":"+port, nil)
}
