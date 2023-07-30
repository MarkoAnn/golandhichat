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
	//str1 := "x"
	//str2 := "y"
	number1, err := strconv.Atoi(str1)
	number2, err1 := strconv.Atoi(str2)
	if err != nil {
		//panic(err)
		fmt.Fprintf(w, "error400")
	}
	if err1 != nil {
		//	panic(err1)
		fmt.Fprintf(w, "error400")
	}

	fmt.Fprintf(w, "result = ", number2*number1)
	if number2*number1 == 0 {
		fmt.Fprintf(w, "error 400")
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
