package file

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Cache struct {
	filePath string
	data     []byte
}

func NewCache(filePath string) *Cache {
	c := Cache{
		filePath: filePath,
	}
	return &c
}

func (c *Cache) Update() error {
	file, err := os.Open("hichat.txt")
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer file.Close()

	d := make([]byte, 64)

	for {
		n, err := file.Read(d)
		if err == io.EOF {
			break
		}

		fmt.Print(string(d[:n]))
	}
	fmt.Println(d)
	c.data = d
	return nil
}

func (c *Cache) HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Println(string(c.data))
	w.Write(c.data)

}
