# GoREST-Tutorial
## Section 3: Concepts
### What is JSON?
JSON stands for JavaScript Object Notation, It's a simple and light-weight data encoding standard. It's generally easier to read, code and decode than XML (SOAP) which is probably why JSON has grow in popularity.

#### JSON Data Types
There are a fairly limited number of data types defined in JSON:
* string - any alphanumeric value encased in double quotes - i.e. "Johnson"
* boolean - true or false
* null
* number - can be integer or float
* array - data encased by []
* object - data encased by {}

#### JSON Structure and Syntax
The syntax rules are pretty simple:
* data field names are encased in double quotes followed by a colon and then the data.
  * "last_name": "Johnson"
  * "record_count": 50
  * "is_retired": false
* Arrays are comma-separated lists of the same data type encased in braces []:
  * "nick_names": ["Big Dave","Bad Dave"]
  * "prime_numbers": [1,2,3,5,7,11,13]
* Objects are groupings of any data type (or other objects) encased in brackets {}
  * "contact": {"last_name": "Johnson", "nick_names":["Big Dave","Bad Dave"], "age":50, "retired": false }
  * { "contact_count": 2, "contacts": [{...},{...} etc...] }
* JSON provides no means of adding comments in the data

Here's a pretty print example of JSON:

```JSON
{
	"contact_count ": 2,
	"contacts": [{
      "last_name": "Johnson",
	    "nick_names": ["Big Dave", "Bad Dave"],
      "age": 50,
      "retired": false
    },{
      "last_name": "Jones",
      "nick_names": ["Billy", "Bill"],
      "age": 58,
      "retired": true
    }]
}
```

For reference, the native Go package used to process JSON: [https://golang.org/pkg/encoding/json/] We'll use this package later in this section to send and receive json data in our Go program.

#### JSON Schemas
JSON schemas are similar in concept to XML schemas, but generally less complex than their XML brethren. A JSON schema is itself JSON, but it's data about the data (metadata) that expresses constraints about JSON documents.

Probably the best way to learn how to use JSON schemas is to play around with an [editor](https://jsonschema.net/#/editor) to see what is possible.

Below is an abbreviated list of the typical constraints used in JSON schemas:
* Check that a field is of the correct data type.
* Describes whether or not a value is required.
* Restricts a valid value to a list of possible ones (enumeration) - i.e. fruit must be one of banana, apple or peach.
* Sets minimum or maximum length or values for a field - i.e. disallow negative numbers, zip must be 5 in length etc.
* Sets a regular expression pattern match to validate input for things like url, ip, phone numbers etc.
* Enforces uniqueness within a list of values.

Imagine if we had to write code to validate all of the above for every field in the JSON received for processing. That's a lot of coding. Fortunately, there are a growing list of schema validation packages for Go on GitHub. The one of primary interest in this tutorial is [gojsonschema](https://github.com/xeipuuv/gojsonschema). It has good community support, easy to use and the one we'll employ for this tutorial.

We'll start using JSON schemas in section 4 but we only touch on the topic in this section because we're focused on basic concepts first and will get into more advanced items later on. Nevertheless, just for the sake of awareness, let's look at a simple example of a JSON schema:

```json
 1 {
 2     "$schema": "http://json-schema.org/draft-04/schema#",
 3     "additionalProperties": false,
 4     "definitions": {},
 5     "id": "http://example.com/example.json",
 6     "properties": {
 7         "age": {
 8             "id": "/properties/age",
 9             "type": "integer"
10         },
11         "last_name": {
12             "id": "/properties/last_name",
13             "type": "string"
14         },
15         "nick_names": {
16             "id": "/properties/nick_names",
17             "items": {
18                 "id": "/properties/nick_names/items",
19                 "type": "string"
20             },
21             "type": "array"
22         },
23         "retired": {
24             "id": "/properties/retired",
25             "type": "boolean"
26         }
27     },
28     "required": [
29         "nick_names",
30         "last_name",
31         "retired",
32         "age"
33     ],
34     "type": "object"
35 }
```

The above schema is relatively simple. It will verify that all data types match and that all fields contain a value (required). This capability is very valuable to the REST programmer because such check are absolutely necessary to ensure that your program receives valid input.  The process of manually coding checks like this would be very tedious and repetitive indeed.

#### Some useful JSON Tutorials, Tools and References
Here is a short list of useful JSON tools and tutorials:
* A simple tutorial / cheat sheet on forming valid JSON: [http://json.org/]
* Is the syntax invalid for your JSON? Fix it here: [http://jsonlint.com]
* A nifty tool that converts valid JSON to a Go data structure (very useful): [http://mholt.github.io/json-to-go/]
* JSON schemas are indispensable! This site will help you create your schemas: [https://jsonschema.net/#/editor]
* JSON linter can help you test your JSON data and schema: [https://jsonschemalint.com]
* The Go package used to process JSON: [https://golang.org/pkg/encoding/json/]
* A growing list of JSON schema validation packages for Go: [https://github.com/search?l=Go&q=go+json+schema+validation]

### What is CGI?
We're going to fire-up the way-back machine here to rediscover something important that happened in the 1970s that is integral to this tutorial. It's called CGI, or Common Gateway Interface. It was first codified as a standard in 1977 and is widely used 40 years later. The original RFC is [here](https://tools.ietf.org/html/rfc3875) and worth a read. In a nutshell, CGI is a standard protocol to convey information to a server-side program (web server) and receive a standardized reply. We use CGI ont the web every day and probably aren't aware of it.

By the way, Go has excellent support for CGI and is one of the easiest ways to get up and going with a CGI/REST program.

#### URL Basics
CGI evolved/improved in the 1980s with the use of URL standards (Uniform Resource Locater). We're all aware of these at least superficially but grasping CGI concepts requires a deeper understanding of the anatomy of a URL:

`scheme:[//[user[:password]@]host[:port]][/path][?query][#fragment]`

Example:

```
ldap://bozo:GoIZwe!rD@localhost:8080/v2/contacts?nick_name=Big+Dave&nick_name=Bad%20Dave&last_name=Johnson#page1
```

In the above example:
* Note that there are several special characters. These are referred to as reserved characters. This is part of the [URL encoding](https://tools.ietf.org/html/rfc3875#section-2.3) standard. The complete list of reserved characters is  = ; / ? : @ & + $ , [ ] % and #.
* Note that protocol does not have to be http/https.
* User name and password can be passed on the URL. Not generally a good practice, but it's supported.
* The path is /v2/contacts and represents our running Go program.
* Everything after the ? is called a query string. This is data that is passed to our Go program.
* Name/Value pairs in the query string are separated by &.
* Name and value in each pair are seperated by =.
* A "fragment" is appended using a #. If the above path pointed to an HTML file, the # would represent an anchor in the page.
* Spaces are represented by +. We can also represent a space using an "escape sequence" that is the hex number for the ascii char we want, preceded by a % - i.e. %20 for a space.
* Note that the field nick_name shows up twice. This is legitimate and will be interpreted as an array in Go.

The go package that handles URL encoding is the [net/url package](https://golang.org/pkg/net/url/#URL). It's worth studying in detail.

**Section 3 - Exercise 1: Parse the example URL**
The solution for exercise 1 is included below (see also exercise1.go in this repo). We can cut paste the code into [https://play.golang.org/](https://play.golang.org/p/l0vg0TmMhN) or just run it with `go run exercise1.go`.

```go
 1 package main
 2
 3 import (
 4         "fmt"
 5         "net/url"
 6 )
 7
 8 func main() {
 9         qstring := "ldap://bozo:GoIZwe!rD@localhost:8080/v2/contacts?first_name=Big+Dave&nick_name=Bad%20Dave&nick_name=Johnson#page1"
10         urlObj, err := url.Parse(qstring)
11         if err != nil {
12                 panic(err) // kinda extreme to panic, but I'm feeling insecure right now
13         }
14         fmt.Printf("Protocol is: %s\n", urlObj.Scheme)
15         fmt.Printf("Host and port are: %s\n", urlObj.Host)
16         fmt.Printf("User name: %s\n", urlObj.User.Username())
17         pw, set := urlObj.User.Password()
18         fmt.Printf("User password is set: %t, and the password is: %s\n", set, pw)
19         fmt.Printf("Path is: %s\n", urlObj.Path)
20         fmt.Printf("RawQuery is: %s\n", urlObj.RawQuery)
21         qvalues := urlObj.Query()
22         for key, arr := range qvalues {
23                 fmt.Printf("\tKey: %s\n", key)
24                 for i := 0; i < len(arr); i++ {
25                         fmt.Printf("\t\tValue[%d]: %s\n", i, arr[i])
26                 }
27         }
28         fmt.Printf("Fragment is: %s\n", urlObj.Fragment)
29 }
```
The output of this program looks like this:
```
Protocol is: ldap
Host and port are: localhost:8080
User name: bozo
User password is set: true, and the password is: GoIZwe!rD
Path is: /v2/contacts
RawQuery is: first_name=Big+Dave&nick_name=Bad%20Dave&nick_name=Johnson
        Key: first_name
                Value[0]: Big Dave
        Key: nick_name
                Value[0]: Bad Dave
                Value[1]: Johnson
Fragment is: page1
```

There are some noteworthy items in the above example:
* The url.Parse() method (line 10) returns a [URL struct object](https://golang.org/pkg/net/url/#URL). Inside that struct is an object of type [Userinfo](https://golang.org/pkg/net/url/#Userinfo) which is [operated on](https://golang.org/src/net/url/url.go?s=9790:9872#L336) by several methods (notably URL.User.Password() and URL.User.Username()).
* Note that the method [URL.Query()](https://golang.org/pkg/net/url/#URL.Query) (line 21) returns a [map of values](https://golang.org/pkg/net/url/#Values) `map[string][]string`. In the URL example above, the field nick_name appears twice, thus map["nick_name"] will return a []string array with two elements.
* Because a field name can appear more than once, we use a nested for loop. The outer loop (line 22) ranges through the map and the inner loop (line 24) iterates through the []string array.

#### Headers
There is more than one way to pass data to our Go program using the CGI protocol. In addition to the query string, we can use headers.  Headers can be passed to the CGI program by the client and also returned by the program to the client. Headers became part of the HTTP protocol in the 90's and is now an integral part of CGI programming.

There are a number of [built-in meta-variables](https://tools.ietf.org/html/rfc3875#section-4.1) as part of the HTTP/CGI protocol (many of these are inserted by the host web service as headers). We can also pass data through custom headers.

In Go, the [net/http](https://golang.org/pkg/net/http/) package handles http requests and responses. Headers are available in either the [Request object](https://golang.org/pkg/net/http/#Request) or the [Response object](https://golang.org/pkg/net/http/#Response) and take on the [same structure](https://golang.org/pkg/net/http/#Header) as query strings - i.e. `map[string][]string`.

**Exercise 2: Read headers from a response object.**

The solution to exercise 2 is included below (see also exercise2.go in this repo). We can run it with `go run exercise2.go`

```go
 1 package main
 2
 3 import (
 4         "fmt"
 5         "net/http"
 6 )
 7
 8 func main() {
 9         res, err := http.Get("http://www.google.com/robots.txt")
10         if err != nil {
11                 panic(err)
12         }
13         for key, arr := range res.Header {
14                 fmt.Printf("\n%s: ", key)
15                 for i := 0; i < len(arr); i++ {
16                         fmt.Printf("%s\t", arr[i])
17                 }
18         }
19         fmt.Printf("\n----\n\nContent-Type: %+v\n", res.Header.Get("content-type"))
20 }
```
The output from this program might look like this:
```
Vary: Accept-Encoding
Content-Type: text/plain
Last-Modified: Mon, 06 Feb 2017 22:00:00 GMT
X-Content-Type-Options: nosniff
X-Xss-Protection: 1; mode=block
Accept-Ranges: bytes
Date: Sat, 06 May 2017 14:29:11 GMT
Expires: Sat, 06 May 2017 14:29:11 GMT
Cache-Control: private, max-age=0
Server: sffe
----

Content-Type: text/plain
```

In the example above:
* The call to the [http.Get()](https://golang.org/pkg/net/http/#Get) method (line 9) returns a response object which contains, among other things, all the headers returned by a particular host.
* The headers need to be iterated through (line 13) in a similar fashion as query parameters in exercise 1.
* It's possible to retrieve a specific header by calling `res.Header.Get("content-type")` (line 19) for example.
* The CGI standard specifies that header names [are not case sensitive](https://tools.ietf.org/html/rfc3875#section-6.3) but their values are.

**Exercise 3: Place header values into a request object.**

The solution to exercise 3 is below (see exercise3.go in this repo). We can run it with `go run exercise3.go`.

In this example, we take advantage of [https://httpbin.org/](https://httpbin.org/) to test requests and responses. Although the site [https://play.golang.org/](https://play.golang.org/) is useful for running code snippets, there are a few things the site won't allow (TCP socket connections for one).

We're also introducing in this exercise another source of data described in the CGI protocol; the body.

```go
 1 package main
 2
 3 import (
 4         "fmt"
 5         "io/ioutil"
 6         "net/http"
 7 )
 8
 9 func main() {
10         client := &http.Client{}
11         req, err := http.NewRequest("GET", "https://httpbin.org/headers", nil)
12         req.Header.Add("foo", "bar")
13         res, err := client.Do(req)
14         if err != nil {
15                 panic(err)
16         }
17         // Close the connection even if we error out
18         defer res.Body.Close()
19
20         // Let's display what we got back in the body
21         body, err := ioutil.ReadAll(res.Body)
22         if err != nil {
23                 panic(err)
24         }
25         fmt.Println("Response Body:", string(body))
26 }
```

The response should look something like this:

```
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
```

In the above example:
* We used the call `req.Header.Add("foo", "bar")` (line 12) to insert a custom header in the request to httpbin.org (called foo with value bar).
* The httpbin.org site replied with a list of headers it received. Note the custom header (called Foo) echoed back in the response.
* Note that the body is in JSON format as indicated by the value of header `Content-Type` (application/json).
* Note the use of `ioutil.ReadAll()` method (line 21) to read the body. We'll get into more detail on this next.

#### The Body and MIME-Types
As we learned in exercise 3, we can receive data back in the CGI message body.

Ok, so we can read the body, but how do we know how to process it? The way to interpret the body content depends primarily on the header named `Content-Type`. In exercise 3, the `Content-Type` is `application/json`. This is referred to in the CGI standard as a [MIME-Type](https://tools.ietf.org/html/rfc2046) (Multipurpose Internet Mail Extensions) and was codified as a standard in the late 1990s. As its name implies, the standard was originally intended for eMail, but it wasn't long after that MIME-Type definitions became widely used in Internet protocols as well. We can find a pretty comprehensive list of MIME-Types [here](http://www.freeformatter.com/mime-types-list.html).

In Go, the body is contained in the request object and is of type [io.Reader](https://golang.org/pkg/io/#Reader). We read the body as a byte stream using the [io/ioutil](https://golang.org/pkg/io/ioutil/) package (using the [ioutil.ReadAll()](https://golang.org/pkg/io/ioutil/#ReadAll) method). It pays to study these packages carefully to understand how they work.

This tutorial is mainly concerned with JSON as the data exchange format for our REST programming but there are a few other MIME-Types frequently used in CGI programming that are worth visiting first.

##### MIME-Type x-www-form-urlencoded

The MIME-Type `x-www-form-urlencoded` showed uo in the 2000s and is often used by web browsers to 'POST' form data to the CGI program. The term POST is a HTTP method (more on that later).

Let's say we are hosting a HTML page with a web form that looks something like this:

```html
1 <form enctype="x-www-form-urlencoded" action="https://httpbin.org/post" method="post">
2     <input type="text" name="nick_name" value="Big Dave">
3     <input type="text" name="nick_name" value="Bad Dave">
4     <input type="text" name="last_name" value="Johnson">
5     <input type="submit" value="Save" />
6 </form>
```

In the above example:
* Note the encoding type (enctype) target (action) and HTTP method (post) on line 1.
* When the user presses the "Save" button (line 5) the browser will encode the form data in x-www-form-urlencoded format and submit (post) it to httpbin.org for processing.

The data 'POSTing' to the CGI program might look like this:

```
POST /post HTTP/1.1
Host: httpbin.org
Content-Type: application/x-www-form-urlencoded

nick_name=Big+Dave&nick_name=Bad+Dave&last_name=Johnson
```

As we can see in the above example, the `Content-Type` header is key to understanding what is contained in the body and how to interpret it. The site httpbin.org must interpret the data being sent and process it into a return. This is the same thing that our own CGI program would do if it were receiving this data.

Fortunately, Go  makes interpreting this type of data very easy with the [Request.ParseForm()](https://golang.org/pkg/net/http/#Request.ParseForm) method. This method parses the form data and places it in the [Request.Form](https://golang.org/pkg/net/http/#Request) data type which is  ... drum roll please ... type `map[string][]string` (surprise).

**Exercise 4: Write a service to receive application/x-www-form-urlencoded.**

It's time to write our first web service. In this exercise we'll be introduced to the [http.ListenAndServe()](https://golang.org/pkg/net/http/#ListenAndServe) method (which stands up a web service) and the [http.HandleFunc()](https://golang.org/pkg/net/http/#HandleFunc) (which handles requests). We'll also use the [JSON encoder](https://golang.org/pkg/encoding/json/) to send back json data to the requester.

The solution to exercise 4 is listed below (see also exercise4.go in this repo). We can issue the command `go run exercise4.go &` to run the program and put it in the background (note the & at the end of the command):

```go
 1 package main
 2
 3 import (
 4         "encoding/json"
 5         "net/http"
 6 )
 7
 8 // This is a handler function to receive POSTed form data
 9 func postContact(w http.ResponseWriter, r *http.Request) {
10         if r.Method == "POST" && r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
11                 r.ParseForm()
12                 w.Header().Set("Content-Type", "application/json")
13                 json.NewEncoder(w).Encode(r.Form)
14                 w.WriteHeader(200)
15         }
16 }
17
18 func main() {
19         http.HandleFunc("/v2/contacts", postContact)
20         err := http.ListenAndServe("127.0.0.1:8080", nil)
21         if err != nil {
22                 panic(err.Error)
23         }
24 }
```

To test the running program, issue a curl command:

```bash
curl -v 127.0.0.1:8080/v2/contacts \
  -d "nick_name=Big Dave" \
  -d "nick_name=Bad Dave" \
  -d "last_name=Johnson"
```

The response coming back from curl might look like this:

```
* Trying 127.0.0.1...
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
```

Issue the command `pkill exercise4` to stop the program.

In the above code example:
* We defined in main() a path that matches a handler function (line 19). This is part of Go's multiplex capabilities. This makes it possible to route different paths to different handlers.
* The function http.HandleFunc() (line 19) requires that the handler function (line 9) accept [two parameters](https://golang.org/pkg/net/http/#Handler): A [ResponseWriter](https://golang.org/pkg/net/http/#ResponseWriter) and a pointer to the Request object (which we are already familiar with).
* We also defined in the main() function a http server that listens for requests (line 20).
* A handler function  `postContact()` (line 9) is defined so that we can read the request and write a response. The ResponseWriter object is used to write back a response to the requester (lines 12-14). The Request object contains the data sent to the CGI program by the client (line 9).
* Note that we are able to test the HTTP method and content-type (line 10) to determine if we received the expected CGI encoding from the client.
* The handler invoked the [Request.ParseForm()](https://golang.org/pkg/net/http/#Request.ParseForm) method (line 11) to parse the data into a native data type `map[string][]string`. Note that Request.ParseForm() must be called before reading the Request.Body.
* The return MIME-Type was set to `application/json` (line 12). Note that the `Content-Type` header must be written before the body or return status.
* The form data was converted to json (using the [json.NewEncoder()](https://golang.org/pkg/encoding/json/) method) and written back as a response (line 13).  
* A status code of 200 (OK) was written to the response (line 14). We can find a list of HTTP status codes and their meaning [here](https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html). The HTTP status code is part of the CGI standard and is used by the client to quickly discovery whether or not the request was successful or not and why.

##### MIME-Type multipart/form-data

The MIME-Type `multipart/form-data` showed up in the late 1990s and is another frequently used encoding to POST data to a server-side CGI program. A web browser can code form data in this format before POSTing it to the server.

Borrowing from exercise 4, let's say we are hosting a HTML page with a web form that is slightly different than before:

```html
1 <form enctype="multipart/form-data" action="https://httpbin.org/post" method="post">
2     <input type="text" name="nick_name" value="Big Dave">
3     <input type="text" name="nick_name" value="Bad Dave">
4     <input type="text" name="last_name" value="Johnson">
5     <input type="submit" value="Save" />
6 </form>
```

In the above example, note that the only thing that changed is the encoding type (enctype). The data submitted to the web server might look similar to the following:

```
POST /post HTTP/1.1
Host: httpbin.org
Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW

------WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="nick_name"

Big Dave
------WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="nick_name"

Bad Dave
------WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="last_name"

Johnson
------WebKitFormBoundary7MA4YWxkTrZu0gW--
```

Ok, this looks a lot more complex to process than the `x-www-form-urlencoded` encoding in exercise 4. Fortunately, Go comes to the rescue again and makes this relatively easy by providing the [Request.ParseMultipartForm()](https://golang.org/pkg/net/http/#Request.ParseMultipartForm) method which places the data in [Request.MultipartForm](https://golang.org/pkg/net/http/#Request) field in the Request object.

**Exercise 5: Write a program to receive multipart/form-data.**

The solution to exercise 5 is listed below (see also exercise5.go in this repo) and we can run it with `go run exercise5.go &`.

```go
1 package main
2
3 import (
4         "encoding/json"
5         "io"
6         "net/http"
7         "strings"
8 )
9
10 // This is a handler function to receive POSTed form data
11 func postContact(w http.ResponseWriter, r *http.Request) {
12         if r.Method == "POST" && strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
13                 r.ParseMultipartForm(0)
14                 w.Header().Set("Content-Type", "application/json")
15                 json.NewEncoder(w).Encode(r.MultipartForm)
16                 file, _, err := r.FormFile("sampledata")
17                 defer file.Close()
18                 if err == nil {
19                         w.Write([]byte("---sample data---\n"))
20                         io.Copy(w, file)
21                 }
22         }
23 }
24
25 func main() {
26         http.HandleFunc("/v2/contacts", postContact)
27         err := http.ListenAndServe("127.0.0.1:8080", nil)
28         if err != nil {
29                 panic(err.Error)
30         }
31 }
```
To test the program, issue a curl command:

```bash
$ curl -v 127.0.0.1:8080/v2/contacts \
  -F "nick_name=Big Dave" \
  -F "nick_name=Bad Dave" \
  -F "last_name=Johnson" \
  -F sampledata=@./exercise5.dat
```
The output might look like the following:

```
> POST /v2/contacts HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.47.0
> Accept: */*
> Content-Length: 1088
> Expect: 100-continue
> Content-Type: multipart/form-data; boundary=------------------------b8d179fba566a850
>
< HTTP/1.1 100 Continue
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Thu, 04 May 2017 18:32:07 GMT
< Content-Length: 259
<
{
  "Value": {
    "last_name": [
      "Johnson"
    ],
    "nick_name": [
      "Big Dave",
      "Bad Dave"
    ]
  },
  "File": {
    "exercise": [
      {
        "Filename": "exercise5.go",
        "Header": {
          "Content-Disposition": [
            "form-data; name=\"exercise\"; filename=\"exercise5.go\""
          ],
          "Content-Type": [
            "application/octet-stream"
          ]
        }
      }
    ]
  }
}
---file---
cGFja2FnZSBtYWluCgppbXBvcnQgKAogICAgICAgICJlbmNvZGluZy9qc29uIgogICAgICAgICJp
byIKICAgICAgICAibmV0L2h0dHAiCiAgICAgICAgInN0cmluZ3MiCikKCi8vIFRoaXMgaXMgYSBo
YW5kbGVyIGZ1bmN0aW9uIHRvIHJlY2VpdmUgUE9TVGVkIGZvcm0gZGF0YQpmdW5jIHBvc3RDb250
YWN0KHcgaHR0cC5SZXNwb25zZVdyaXRlciwgciAqaHR0cC5SZXF1ZXN0KSB7CiAgICAgICAgaWYg
ci5NZXRob2QgPT0gIlBPU1QiICYmIHN0cmluZ3MuQ29udGFpbnMoci5IZWFkZXIuR2V0KCJDb250
ZW50LVR5cGUiKSwgIm11bHRpcGFydC9mb3JtLWRhdGEiKSB7CiAgICAgICAgICAgICAgICByLlBh
cnNlTXVsdGlwYXJ0Rm9ybSgwKQogICAgICAgICAgICAgICAgdy5IZWFkZXIoKS5TZXQoIkNvbnRl
bnQtVHlwZSIsICJhcHBsaWNhdGlvbi9qc29uIikKICAgICAgICAgICAgICAgIGpzb24uTmV3RW5j
b2Rlcih3KS5FbmNvZGUoci5NdWx0aXBhcnRGb3JtKQogICAgICAgICAgICAgICAgZmlsZSwgZXJy
IDo9IHIuTXVsdGlwYXJ0Rm9ybS5GaWxlWyJleGVyY2lzZSJdWzBdLk9wZW4oKQogICAgICAgICAg
ICAgICAgaWYgZXJyID09IG5pbCB7CiAgICAgICAgICAgICAgICAgICAgICAgIHcuV3JpdGUoW11i
eXRlKCItLS1maWxlLS0tXG4iKSkKICAgICAgICAgICAgICAgICAgICAgICAgaW8uQ29weSh3LCBm
aWxlKQogICAgICAgICAgICAgICAgfQogICAgICAgIH0KfQoKZnVuYyBtYWluKCkgewogICAgICAg
IGh0dHAuSGFuZGxlRnVuYygiL3YyL2NvbnRhY3RzIiwgcG9zdENvbnRhY3QpCiAgICAgICAgZXJy
IDo9IGh0dHAuTGlzdGVuQW5kU2VydmUoIjEyNy4wLjAuMTo4MDgwIiwgbmlsKQogICAgICAgIGlm
IGVyciAhPSBuaWwgewogICAgICAgICAgICAgICAgcGFuaWMoZXJyLkVycm9yKQogICAgICAgIH0K
fQo=
```
Be sure to issue a `pkill exercise5` command to stop the program.

In the above example, there are a lot of interesting things going on:
* In the `postContact()` handler (line 12), we altered the test for Content-Type to use the `strings.Contains()` meathod instead of an exact match. The reason for this is that the header name in this case is actually `Content-Type: multipart/form-data; boundary=------------------------b8d179fba566a850`, so we have to search for a partial match.
* We use the `Request.ParseMultipartForm()` method (line 13) instead of the `Request.ParseForm()` method of exercise 4. There are different methods for different MIME-Types.
* The data ends up in `Request.MultipartForm` (line 13) instead of `Request.Form` of exercise 4. There are different data structure for different MIME-Types.
* Note that a file was included in the curl request. This is the main reason to use `multipart/form-data` in the first place. The data and the file are different parts of the request (thus `multipart`).
* The uploaded file can be read by using the [Request.FormFile](https://golang.org/pkg/net/http/#Request.FormFile) data element in the Request object (line 16).
* For demonstration purposes, we just output the raw file as part of the response (line 20).
* There are many gems in the Request object that are worth studying.

##### MIME-Type application/json

There are other MIME-Types, but for the purposes of this tutorial, we will cover just one more of major significance in REST programming; which is (you guessed it) `application/json`.  JSON has it's beginnings in the early 2000's but hasn't been standardized until early this decade. the JSON schema standard is relatively recent (starting in 2013), so our journey with the way-back machine has returned to our current decade.

**Exercise 6: Write a program to receive application/json.**

We've already created examples where we output JSON data from our CGI program. In this exercise, we will also receive JSON data and convert it to a native Go structure for processing. In many ways, it's one of the easiest MIME-Types to work with (we saved the best for the last).

The solution to exercise 6 is listed below (see also exercise6.go in this repo) and we can run it with `go run exercise6.go &`.

```go
1 package main
2
3 import (
4         "crypto/rand"
5         "encoding/json"
6         "fmt"
7         "io/ioutil"
8         "net/http"
9 )
10
11 // structure to hold unmarshaled json data
12 type Contact struct {
13         RecordKey string   `json:"record_key"`
14         LastName  string   `json:"last_name"`
15         NickName  []string `json:"nick_name"`
16 }
17
18 // This generates a new record key
19 func uuid() (uuid string) {
20         b := make([]byte, 16)
21         rand.Read(b)
22         uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
23         return
24 }
25
26 // This is a handler function to receive POSTed form data
27 func postContact(w http.ResponseWriter, r *http.Request) {
28         if r.Method != "POST" || r.Header.Get("Content-Type") != "application/json" {
29                 w.WriteHeader(http.StatusBadRequest)
30                 return
31         }
32         var contact Contact
33         body, _ := ioutil.ReadAll(r.Body)
34         err := json.Unmarshal(body, &contact)
35         if err != nil {
36                 w.WriteHeader(http.StatusBadRequest)
37                 return
38         }
39         contact.RecordKey = uuid()
40         w.Header().Set("Content-Type", "application/json")
41         w.WriteHeader(http.StatusCreated)
42         json.NewEncoder(w).Encode(contact)
43 }
44
45 func main() {
46         http.HandleFunc("/v2/contacts", postContact)
47         err := http.ListenAndServe("127.0.0.1:8080", nil)
48         if err != nil {
49                 panic(err.Error)
50         }
51 }
```

We can test the program by issuing a curl command:

```bash
curl -sv -XPOST 127.0.0.1:8080/v2/contacts \
  -H "Content-Type: application/json" \
  --data-binary '{"last_name":"Johnson","nick_name":["Big Dave","Bad Dave"]}'
```

The response will look similar to the following:

```
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> POST /v2/contacts HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.47.0
> Accept: */*
> Content-Type: application/json
> Content-Length: 59
>
* upload completely sent off: 59 out of 59 bytes
< HTTP/1.1 201 Created
< Content-Type: application/json
< Date: Fri, 05 May 2017 01:10:49 GMT
< Content-Length: 112
<
{
  "record_key": "804621B0-CECA-F3D6-1660-757091435D76",
  "last_name": "Johnson",
  "nick_name": [
    "Big Dave",
    "Bad Dave"
  ]
}
```

Issue a command `pkill exercise6` to stop the program.

Some noteworthy items in the above example:
* We added a small utility function (line 19) that generates unique record keys.
* Note that we started using the [constants](https://golang.org/pkg/net/http/#pkg-constants) in the http package for the status return codes (lines 29, 26 and 41).
* Note that we introduce a new data structure (line 12) to receive the json data. There's a good tutorial on the marshal/unmarshal of json [here](https://blog.golang.org/json-and-go). There's also a good JSON to Go converter [here](https://mholt.github.io/json-to-go/) that will generate some Go code to speed things along.
* We instantiate the `contact` data structure (line 32) and `unmarshal` the JSON data from the body into the contact  (lines 33-34). Note that the data being passed has two fields and the structure has three.
* We then assign a new record ID to the contact (line 39) and convert (Marshal) the amalgamation back to JSON data in the response (line 42).

Now that we've covered numerous examples of how CGI works in Go, it's time to move on to the REST API standard for implementing CGI programs.

### What is REST?

REST (Representational State Transfer) essentially defines the basic framework of nearly any REST-compliant CGI API. REST is more of a set of principles than it is a codified standard. It has its origins in the 1990s and published as a doctoral thesis in 2000 by [Roy Fielding](https://en.wikipedia.org/wiki/Roy_Fielding). The intent - in a nutshell - is to give every web resource a textual representation (in this case a URI) that can be operated upon with a standard set of verbs (GET, PUT, POST, DELETE etc). In this way, we move away from arbitrary resource representations commonly found with SOAP/WSDL and thereby form a more generic and uniform means to interact with data over HTTP.

Unfortunately, REST API implementations vary with some API using verbs as endpoints whereas the growing consensus is to use nouns to describe an endpoint. Some API use the GET method and query strings to alter data, whereas CRUD (create, read, update and delete) operations are considered a best practice via HTTP methods. There is no consensus on whether to return XML or JSON, although JSON is growing in popularity (either data format is technically REST-compliant and capable of representing data in a standardized way). Lastly, there is no documentation standard for APIs, although Swagger and RAML are two popular tools for this purpose.

#### General REST Guidelines

Below are the general guidelines we will follow in this tutorial:
* A collection of resources should be described by a plural noun - i.e. /contacts
* Give endpoint representations for every object in a collection - i.e. /contacts/david.x.johnson
* Endpoint resources should be stateless - i.e. the identifier is all that's needed to identify the resource
* Heuristic versioning should be used to compartmentalize API versions - i.e. /api/v2/contacts
* Operations on resources should be CRUD-like using the appropriate HTTP method- i.e. GET, PUT, POST, DELETE, PATCH etc.
* The use of query parameters should be limited to GET operations, although there may be justifiable exceptions.
* Each REST API call should return a HTTP status code and some content (if appropriate).
* Content is preferably returned in the message body and must match the MIME-Type in the header "Content-Type".
* JSON is the preferred data format for the purposes of this tutorial.
* We'll be using RAML to sketch out and document our API in this tutorial.

#### A Closer Look at HTTP Methods and Status Codes

In REST programming, the HTTP Method is the action verb in a REST call. The client includes the method in the call to the server along with a resource identifier to operate on. The server responds with a status code and perhaps some data (only if necessary). We've already observed this type of communication in previous examples.

POST request from client to create a new contact resource using version 2 of the api:

```
> POST /v2/contacts HTTP/1.1
> Content-Type: application/json
....
```

The server reply contains a status code and perhaps some data as well:

```
< HTTP/1.1 201 Created
< Content-Type: application/json
....
```

The use of HTTP methods can be an area of some confusion, particularly with PUT and POST. This tutorial interprets the most commonly used HTTP methods in the follow way:

[Go Constant Name](https://golang.org/pkg/net/http/#pkg-constants) | [HTTP Method](https://www.w3.org/Protocols/rfc2616/rfc2616-sec9.html) | Interpretation | Typical Return Status
---|---|---|---
MethodGet | GET | Retrieves data without side-effects. In other words, it does not alter the data. Analogous to read.| 200 Ok, 404 Not Found
MethodPost | POST | Uploads a representation of a new resource. Analogous to add new.| 200 Ok, 201 Created, 409 Conflict
MethodPut | PUT | Uploads a replacement of an existing resource or creation of a new one if none exists that matches the data. Analogous to create or replace.| 200 Ok, 201 Created, 204 No Content
MethodPatch | PATCH | RFC 5789 Replaces a portion or property of a resource. Analogous to a partial update.| 200 Ok, 204 No Content, 404 Not Found
MethodDelete  | DELETE | Removes the resource (self explanatory). | 200 Ok, 204 No Content

The following are some of the more common status codes we'll be working with in this tutorial:

[Go Constant Name](https://golang.org/pkg/net/http/#pkg-constants)|[HTTP Status](https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html)|Interpretation|Typical methods
---|---|---|---
StatusOK | 200 | Indicates success and typically contains return data. |GET,POST,PUT
StatusCreated| 201 | Indicates success in creating a new resource but typically does not include return data.| POST, PUT
StatusNoContent | 204 | Indicates success in updating a resource but typically does not include return data.|PUT,PATCH,DELETE
StatusBadRequest| 400 | Typically returned when data payload sent was malformed for the indicated method or Content-Type.| any
StatusUnauthorized| 401 | Unauthorized. Usually an indication that requester credentials could not be authenticated.| any
StatusForbidden | 403 | Forbidden. Usually a case where the requester is authenticated but not allowed to perform the operation.| any
StatusNotFound | 404 | Not found. The resource URI is invalid. | GET, DELETE, PATCH
StatusMethodNotAllowed | 405 | Method not allowed for the indicated resource URI.| any
StatusConflict | 409 | Conflict. Can be used to indicate that a resource can't be created because it already exists.| POST
StatusInternalServerError| 500 | Internal error. This is typically a critical error.| any

Note: The confusion over PUT vs. POST is usually related to the fact that PUT is a create or replace operation; meaning that it will replace an existing resource with a new one or create that new one if the resource doesn't exist (and therefore is potentially destructive). POST - on the other hand - will only create the resource if it doesn't already exist (and therefore not destructive).

### What's next?

Ok, so we learned about JSON and JSON schemas, we had some fun with simple Go programs to process CGI requests and we now have a basic framework to develop a REST program. It's time to code, right? Not so fast, cowboy ... design tools are next ([section 4](../section4)).
