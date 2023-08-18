package main

import (
	"awesomeProject1/file"
	_ "awesomeProject1/file"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	some := file.NewCache("hichat.txt")

	err := some.Update()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/cache", some.HandleRequest)
	http.HandleFunc("/", textHandler)
	http.HandleFunc("/hichat", mHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func textHandler(w http.ResponseWriter, r *http.Request) {

	text := "hi chat test"
	file, err := os.Create("hichat.txt")
	if err != nil {
		fmt.Printf("ошибка создать файл[%s]", err.Error())
		return
	}
	writeString, err := file.WriteString(text)
	if err != nil {
		fmt.Printf("ошибка создать файл[%s]", err.Error())

	}
	println(file.Name())
	println(writeString)

	openfile, err := os.Open(file.Name())
	if err != nil {
		fmt.Printf("ошибка чтения файла[%s]", err.Error())
		return
	}
	bufer := make([]byte, writeString)
	for {
		_, err := openfile.Read(bufer)
		if err != io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("ошибка чтения файла[%s]", err.Error())
			break
		}

	}
	fmt.Fprintf(w, string(bufer))
}

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

	fmt.Fprintf(w, strconv.Itoa(number1*number2))

}

func yHandler(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query().Get("y")
	fmt.Fprintf(w, "x= %s", queryParam)
}
