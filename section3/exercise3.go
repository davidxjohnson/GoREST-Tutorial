package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://httpbin.org/headers", nil)
	req.Header.Add("foo", "bar")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	// Close the connection even if we error out
	defer res.Body.Close()

	// Let's display what we got back in the body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response Body:", string(body))
}
