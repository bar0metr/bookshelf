package main;

import (
"net/http"
"fmt"
	"database/sql"

	_ "github.com/lib/pq"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}
func requestHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, worlddd")
}


func main(){
	db, err := sql.Open("postgres", "user=postgres dbname=tefst sslmode=verify-full")
	if err != nil {
			http.HandleFunc("/", requestHandler2)
	}

	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
fmt.Printf("%v\n", rows)
//	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8181", nil)

	
}
