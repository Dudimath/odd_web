package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", oddHandler)
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func oddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("static/index.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		numStr := r.FormValue("number")
		num, err := strconv.Atoi(numStr)
		if err != nil {
			http.Error(w, "Invalid Input", http.StatusBadRequest)
			return
		}
		result := fmt.Sprintf("%d is not an odd number", num)
		if isOdd(num) {
			result = fmt.Sprintf("%d is an odd number", num)
		}
		tmpl := template.Must(template.ParseFiles("static/index.html"))
		tmpl.Execute(w, struct{ Result string }{Result: result})
	}
}

func isOdd(num int) bool {
	if num%2 == 0 {
		return false
	}
	return true
}
