package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Ты передал параметр: %s", queryParam)
}
func mHandler(w http.ResponseWriter, r *http.Request) {
	str2 := r.URL.Query().Get("x")
	str1 := r.URL.Query().Get("y")
	number1, err := strconv.Atoi(str1)
	number2, err1 := strconv.Atoi(str2)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "status code:400")
		return
	}
	if err1 != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "status code:400")
		return
	}
	if err1 == nil {
		fmt.Fprintf(w, strconv.Itoa(number1*number2))
	}

}

func yHandler(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query().Get("y")
	fmt.Fprintf(w, "x= %s", queryParam)
}

func main() {
	http.HandleFunc("/", mHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
