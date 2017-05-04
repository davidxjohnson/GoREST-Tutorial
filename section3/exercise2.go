package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		panic(err)
	}
	for key, arr := range res.Header {
		fmt.Printf("\n%s: ", key)
		for i := 0; i < len(arr); i++ {
			fmt.Printf("%s\t", arr[i])
		}
	}
	fmt.Printf("\n----\n\nContent-Type: %+v\n", res.Header.Get("content-type"))
}
