package main

import (
	"fmt"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("IMG_20190427_133910.jpg")
	text, err := client.Text()
	fmt.Println("ERR",err)
	fmt.Println(text)
	fmt.Println("ASU")
	// Hello, World!
}
