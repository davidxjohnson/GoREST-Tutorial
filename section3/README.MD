# GoREST-Tutorial
### Section 3: Concepts
#### What is JSON
Content TBD
#### What is CGI?
CGI stands for Common Gateway Interface. It's a standard protocol to convey information to a server-side program (web server) and receive a standardized reply. It was first standardized in 1977 and is still widely used today. You can read the RFC [here](https://tools.ietf.org/html/rfc3875) if you like.

The Go language has good support for this protocol and is one of the easiest ways to get up and going with a REST program (more on that later in this tutorial).
#### URL Basics
Understanding CGI begins with an understanding of the anatomy of a URL:  

`scheme:[//[user[:password]@]host[:port]][/path][?query][#fragment]`

`ldap://bozo:GoIZwe!rD@localhost:8080/v2/contacts?nick_name=Big+Dave&nick_name=Bad%20Dave&last_name=Johnson#page1`

In the above example:
* Note in the above URL that there are several special characters. These are referred to as reserved characters. This is part of the [URL encoding](https://tools.ietf.org/html/rfc3875#section-2.3) standard. The complete list of reserved characters is  = ; / ? : @ & + $ , [ ] % and #.
* Note that protocol does not have to be http/https.
* User name and password can be passed on the URL. Not generally a good practice, but it's supported.
* The /path is /v2/contacts and represents your running Go program.
* Everything after the ? is called a query string. This is data that is passed to your Go program.
* Name/Value pairs in the query string are separated by &.
* Name and value in each pair are seperated by =.
* A "fragment" is appended using a #. If the above path pointed to an HTML file, the # would represent an anchor in the page.
* Spaces are represented by +. You can also represent a space using an "escape sequence" that is the hex number for the ascii char you want, preceded by a % - i.e. %20 for a space.
* Note that the field nick_name shows up twice. This is legitimate and will be interpreted as an array in Go.

The go package that handles URL encoding is the [net/url package](https://golang.org/pkg/net/url/#URL). It's worth studying in detail.

**Section 3 - Exercise 1: Parse the example URL**
The solution for exercise 1 is included below (and included in this repo). You can cut paste the code into [https://play.golang.org/](https://play.golang.org/p/l0vg0TmMhN) or invoke it with `go run exercise1.go`.

```
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
}```

There are some noteworthy items in the above example:
* The url.Parse method returns a [URL struct object](https://golang.org/pkg/net/url/#URL). Inside that struct is an object of type [Userinfo](https://golang.org/pkg/net/url/#Userinfo) which is [operated on](https://golang.org/src/net/url/url.go?s=9790:9872#L336) by several methods (notably User.Password() and User.Username()).
* Note that the method [url.RawQuery](https://golang.org/pkg/net/url/#URL.Query) returns a [map of values](https://golang.org/pkg/net/url/#Values) that are type map[string][]string. In the URL example above, the field nick_name appears twice, thus map["nick_name"] will return a []string array with two elements.
* Because a field name can appear more than once, we use a nested for loop. The outer loop ranges through the map and the inner loop iterates through the array.

#### Headers
There is more than one way to pass data to our Go program using the CGI protocol. In addition to the  query string, we can use headers. Headers can be passed to the CGI program by the client and also returned by the program to the client.

There are a number of [built-in meta-variables](https://tools.ietf.org/html/rfc3875#section-4.1) as part of the CGI protocol (many of these are inserted by the host web service as headers). You can also pass data through custom headers.

In Go, the net/http package handles http requests and responses. Headers are available in either the [request object](https://golang.org/pkg/net/http/#Request) or the [response object](https://golang.org/pkg/net/http/#Response) and take on the [same structure](https://golang.org/pkg/net/http/#Header) as query strings - i.e. map[string][]string.

**Section 3 - Exercise 2: Read headers from a response object.**
The solution to exercise 2 is included below (and as a file in this repo). You can invoke it with `go run exercise2.go`

````
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
````
In the example above:
* The call to http.Get returns a response object which contains, among other things, all the headers returned by a particular host.
* The headers need to be iterated through in a similar fashion as query parameters in exercise 1.
* It's possible to retrieve a specific header by calling `res.Header.Get("content-type")` for example.
* The CGI standard specifies that reference to the headers [is not case sensitive](https://tools.ietf.org/html/rfc3875#section-6.3).

**Section 3 - Exercise 3: Place header values into a request object.**
The solution to exercise 3 is below (and included as a file in this repo). You can invoke it with `go run exercise3.go`.

In this example, we take advantage of a [nifty web site](https://httpbin.org/) that we can use to test requests and responses. The site [https://play.golang.org/](https://play.golang.org/) is useful for running code snippets, but there are a few things they won't allow (TCP socket connections for one).

We're also introducing another source of data described in the CGI protocol; the body.

````
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

        // let's read the body of the response object
        body, err := ioutil.ReadAll(res.Body)
        if err != nil {
                panic(err)
        }
        // Let's display what we got back in the Content-Type header and the body
        fmt.Printf("Content-Type: %s\n", res.Header.Get("content-type"))
        fmt.Println(string(body))
}
````

The response should look something like this:
````
Content-Type: application/json
{
  "headers": {
    "Accept-Encoding": "gzip",
    "Connection": "close",
    "Foo": "bar",
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/1.1"
  }
}
````

In the above example:
* We used the call `req.Header.Add("foo", "bar")` to insert a custom header in the request to httpbin.org (called foo with value bar).
* The httpbin.org site replied with a list of headers it received. Note the custom header (called Foo) echoed back in the response.
* Note that the body is in JSON format as indicated by the header `Content-Type`.

#### The Body and Mime-Types
In exercise 3 we received data back in the message body.

In Go, the body is contained in the request object and is of type [io.Reader](https://golang.org/pkg/io/#Reader). We read the body as a byte stream using the [io/ioutil](https://golang.org/pkg/io/ioutil/) package (using the [ioutil.ReadAll](https://golang.org/pkg/io/ioutil/#ReadAll) method). It pays to study these packages carefully to understand how they work.

The way to interpret the body content depends on the header named `Content-Type` ... in this example, the value is `application/json`. This is referred to in the CGI standard as a [Mime-Type](https://tools.ietf.org/html/rfc2046). You can find a pretty comprehensive list of Mime-Types [here](http://www.freeformatter.com/mime-types-list.html).

This tutorial is mainly concerned with JSON as the data exchange format for our REST programming (more on that later in this tutorial) but there's another very commonly used Mime-Type called `x-www-form-urlencoded`.

The Mime-Type `x-www-form-urlencoded` is often what is used by web browsers to 'POST' form data to the CGI program. The term POST is a HTTP method (more on that later).

Let's say you have a web form that looks like this:
````
<form enctype="multipart/form-data" action="https://httpbin.org/post" method="post">
    <input type="text" name="nick_name" value="Big Dave">
    <input type="text" name="nick_name" value="Bad Dave">
    <input type="text" name="last_name" value="Johnson">
    <input type="submit" value="Save" />
</form>
````

The data 'POSTing' to the CGI program might look like this:

````
POST /post HTTP/1.1
Host: httpbin.org
Content-Type: application/x-www-form-urlencoded

nick_name=Big+Dave&nick_name=Bad+Dave&last_name=Johnson
````
As you can see in the above example, the `Content-Type` header is key to understanding what exactly is contained in the body and how to interpret it.

Fortunately, Go  makes interpreting this type of data very easy with the [Request.ParseForm](https://golang.org/pkg/net/http/#Request.ParseForm) method. This method parses the form data and places it in the [Request.Form](https://golang.org/pkg/net/http/#Request) data type which is  ... drum roll please ... type `map[string][]string` (surprise).

**Section 3 - Exercise 4: Write a service to receive form data.**

It's time to write your first web service. In this exercise you'll be introduced to http.ListenAndServe (which stands up a web service) and http.HandleFunc (which handles requests). We'll also use the JSON encoder to send back json data to the requester.

The solution to exercise 4 is listed below (and also included as a file in this repo). You can use `go run exercise4.go &` to run the program and put it in the background:

````
package main

import (
        "encoding/json"
        "net/http"
)

// This is a handler function to receive POSTed form data
func postContact(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" && r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
                r.ParseForm()
                w.Header().Set("Content-Type", "application/json")
                json.NewEncoder(w).Encode(r.Form)
                w.WriteHeader(200)
        }
}

func main() {
        http.HandleFunc("/v2/contacts", postContact)
        err := http.ListenAndServe("127.0.0.1:8080", nil)
        if err != nil {
                panic(err.Error)
        }
}
````
To test the running program, issue a curl command:
````
curl -v 127.0.0.1:8080/v2/contacts -d "nick_name=Big Dave" -d "nick_name=Bad Dave" -d "last_name=Johnson"
````
The response coming back from curl might look like this:

````
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> POST /v2/contacts HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.47.0
> Accept: */*
> Content-Length: 55
> Content-Type: application/x-www-form-urlencoded
>
* upload completely sent off: 55 out of 55 bytes
2017/05/04 00:18:21 http: multiple response.WriteHeader calls
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Thu, 04 May 2017 05:18:21 GMT
< Content-Length: 62
<
{"last_name":["Johnson"],"nick_name":["Big Dave","Bad Dave"]}
````
Issue a command `pkill exercise4` to stop the program.

In the above example:
* Defined in main() a path that matches a handler function. This is part of Go's multiplex capabilities. This makes it possible to route different paths to different handlers.
* The function http.HandleFunc requires that the handler function accept [two parameters](https://golang.org/pkg/net/http/#Handler): A [ResponseWriter](https://golang.org/pkg/net/http/#ResponseWriter) and a pointer to the Request object (which we are already familiar with). The ResponseWriter is used to write back a response to the requester. The Request object contains the data sent to the CGI program by the client.
* Defined in the main() function a http server that listens for requests.
* A handler function  `postContact` is defined that can read the request and write a response.
* In the handler function, we are able to test the HTTP method and content-type to determine if we received the expected CGI encoding from the client.
* The handler invoked the [Request.ParseForm](https://golang.org/pkg/net/http/#Request.ParseForm) method to parse the data into a native data type map[string][]string.
* The return Mime-Type was set to `application/json`.
* The form data was converted to json (using the [json.NewEncoder](https://golang.org/pkg/encoding/json/) method) and written back as a response.  It pays to carefully study this package ot understand how it works.
* A status code of 200 (OK) was written to the response. You can find a list of HTTP status codes and their meaning [here](https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html).
* Note that ParseForm() must be called before reading the Request.Body.
* Note that the `Content-Type` header must be written before the body or return status.
