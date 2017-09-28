package main

import (
	"fmt"
	"net/url"
)

func main() {
	qstring := "ldap://bozo:GoIZwe!rD@localhost:8080/v2/contacts?first_name=Big+Dave&nick_name=Bad%20Dave&nick_name=Johnson#page1"
	urlObj, err := url.Parse(qstring)
	if err != nil {
		panic(err) // kinda extreme to panic, but I'm feeling insecure right now
	}
	fmt.Printf("Protocol is: %s\n", urlObj.Scheme)
	fmt.Printf("Host and port are: %s\n", urlObj.Host)
	fmt.Printf("User name: %s\n", urlObj.User.Username())
	pw, set := urlObj.User.Password()
	fmt.Printf("User password is set: %t, and the password is: %s\n", set, pw)
	fmt.Printf("Path is: %s\n", urlObj.Path)
	fmt.Printf("RawQuery is: %s\n", urlObj.RawQuery)
	qvalues := urlObj.Query()
	for key, arr := range qvalues {
		fmt.Printf("\tKey: %s\n", key)
		for i := 0; i < len(arr); i++ {
			fmt.Printf("\t\tValue[%d]: %s\n", i, arr[i])
		}
	}
	fmt.Printf("Fragment is: %s\n", urlObj.Fragment)
}
