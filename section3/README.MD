# GoREST-Tutorial
## Section 3: Concepts
### What is JSON?
Content TBD
### What is CGI?
CGI stands for Common Gateway Interface. It's a standard protocol to convey information to a server-side program (web server) and receive a standardized reply. It was first codified as a standard in 1977 and is still widely used today. The original RFC [here](https://tools.ietf.org/html/rfc3875) has the details.

The Go language has good support for this protocol and is one of the easiest ways to get up and going with a REST program.

### URL Basics
Grasping CGI concept requires an understanding of the anatomy of a URL:  

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

```
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

There are some noteworthy items in the above example:
* The url.Parse method (line 10) returns a [URL struct object](https://golang.org/pkg/net/url/#URL). Inside that struct is an object of type [Userinfo](https://golang.org/pkg/net/url/#Userinfo) which is [operated on](https://golang.org/src/net/url/url.go?s=9790:9872#L336) by several methods (notably User.Password() and User.Username()).
* Note that the method [url.Query](https://golang.org/pkg/net/url/#URL.Query) (line 21) returns a [map of values](https://golang.org/pkg/net/url/#Values) `map[string][]string`. In the URL example above, the field nick_name appears twice, thus map["nick_name"] will return a []string array with two elements.
* Because a field name can appear more than once, we use a nested for loop. The outer loop (line 22) ranges through the map and the inner loop (line 24) iterates through the []string array.

### Headers
There is more than one way to pass data to our Go program using the CGI protocol. In addition to the  query string, we can use headers. Headers can be passed to the CGI program by the client and also returned by the program to the client.

There are a number of [built-in meta-variables](https://tools.ietf.org/html/rfc3875#section-4.1) as part of the CGI protocol (many of these are inserted by the host web service as headers). We can also pass data through custom headers.

In Go, the [net/http](https://golang.org/pkg/net/http/) package handles http requests and responses. Headers are available in either the [request object](https://golang.org/pkg/net/http/#Request) or the [response object](https://golang.org/pkg/net/http/#Response) and take on the [same structure](https://golang.org/pkg/net/http/#Header) as query strings - i.e. `map[string][]string`.

**Exercise 2: Read headers from a response object.**

The solution to exercise 2 is included below (see also exercise2.go in this repo). We can run it with `go run exercise2.go`

```
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

In the example above:
* The call to [http.Get](https://golang.org/pkg/net/http/#Get) (line 9) returns a response object which contains, among other things, all the headers returned by a particular host.
* The headers need to be iterated through (line 13) in a similar fashion as query parameters in exercise 1.
* It's possible to retrieve a specific header by calling `res.Header.Get("content-type")` (line 19) for example.
* The CGI standard specifies that reference to the headers [is not case sensitive](https://tools.ietf.org/html/rfc3875#section-6.3) but their values are.

**Exercise 3: Place header values into a request object.**

The solution to exercise 3 is below (see exercise3.go in this repo). We can run it with `go run exercise3.go`.

In this example, we take advantage of [https://httpbin.org/](https://httpbin.org/) to test requests and responses. Although the site [https://play.golang.org/](https://play.golang.org/) is useful for running code snippets, there are a few things the site won't allow (TCP socket connections for one).

We're also introducing in this exercise another source of data described in the CGI protocol; the body.

```
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
* Note the use of `ioutil.ReadAll` (line 21) to read the body. We'll get into more detail on this next.

### The Body and MIME-Types
In exercise 3 we received data back in the CGI message body.

In Go, the body is contained in the request object and is of type [io.Reader](https://golang.org/pkg/io/#Reader). We read the body as a byte stream using the [io/ioutil](https://golang.org/pkg/io/ioutil/) package (using the [ioutil.ReadAll](https://golang.org/pkg/io/ioutil/#ReadAll) method). It pays to study these packages carefully to understand how they work.

Ok, so we can read the body, but how do we know how to process it? The way to interpret the body content depends on the header named `Content-Type`. In this example, the value is `application/json`. This is referred to in the CGI standard as a [MIME-Type](https://tools.ietf.org/html/rfc2046). We can find a pretty comprehensive list of MIME-Types [here](http://www.freeformatter.com/mime-types-list.html).

This tutorial is mainly concerned with JSON as the data exchange format for our REST programming but there are a few other MIME-Types frequently used in CGI programming.

#### MIME-Type x-www-form-urlencoded

The MIME-Type `x-www-form-urlencoded` is often what is used by web browsers to 'POST' form data to the CGI program. The term POST is a HTTP method (more on that later).

Let's say we are hosting a HTML page with a web form that looks like this:

```
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

As we can see in the above example, the `Content-Type` header is key to understanding what exactly is contained in the body and how to interpret it. The site httpbin.org must interpret the data being sent and process it into a return. This is the same thing that our own CGI program would do if it were receiving this data.

Fortunately, Go  makes interpreting this type of data very easy with the [Request.ParseForm](https://golang.org/pkg/net/http/#Request.ParseForm) method. This method parses the form data and places it in the [Request.Form](https://golang.org/pkg/net/http/#Request) data type which is  ... drum roll please ... type `map[string][]string` (surprise).

**Exercise 4: Write a service to receive form data.**

It's time to write our first web service. In this exercise we'll be introduced to http.ListenAndServe (which stands up a web service) and http.HandleFunc (which handles requests). We'll also use the JSON encoder to send back json data to the requester.

The solution to exercise 4 is listed below (see also exercise4.go in this repo). We can use `go run exercise4.go &` to run the program and put it in the background (note the & at the end of the command):

```
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

```
curl -v 127.0.0.1:8080/v2/contacts -d "nick_name=Big Dave" -d "nick_name=Bad Dave" -d "last_name=Johnson"
```

The response coming back from curl might look like this:

```
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
```

Issue a command `pkill exercise4` to stop the program.

In the above code example:
* We defined in main() a path that matches a handler function (line 19). This is part of Go's multiplex capabilities. This makes it possible to route different paths to different handlers.
* The function http.HandleFunc (line 19) requires that the handler function (line 9) accept [two parameters](https://golang.org/pkg/net/http/#Handler): A [ResponseWriter](https://golang.org/pkg/net/http/#ResponseWriter) and a pointer to the Request object (which we are already familiar with).
* We also defined in the main() function a http server that listens for requests (line 20).
* A handler function  `postContact` (line 9) is defined so that we can read the request and write a response. The ResponseWriter is used to write back a response to the requester (lines 12-14). The Request object contains the data sent to the CGI program by the client (line 9).
* Note that we are able to test the HTTP method and content-type (line 10) to determine if we received the expected CGI encoding from the client.
* The handler invoked the [Request.ParseForm](https://golang.org/pkg/net/http/#Request.ParseForm) method (line 11) to parse the data into a native data type `map[string][]string`. Note that ParseForm() must be called before reading the Request.Body.
* The return MIME-Type was set to `application/json` (line 12). Note that the `Content-Type` header must be written before the body or return status.
* The form data was converted to json (using the [json.NewEncoder](https://golang.org/pkg/encoding/json/) method) and written back as a response (line 13).  
* A status code of 200 (OK) was written to the response (line 14). We can find a list of HTTP status codes and their meaning [here](https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html). The HTTP status code is part of the CGI standard and is used by the client to quickly discovery whether or not the request was successful or not and why.

#### MIME-Type multipart/form-data

The MIME-Type `multipart/form-data` is another frequently used method to POST data to a server-side CGI program. A web browser can code form data in this format before POSTing it to the server.

Borrowing from exercise 4, let's say we are hosting a HTML page with a web form that slightly different than before:

```
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

Ok, this looks a lot more complex to process than the `x-www-form-urlencoded` encoding in exercise 4. Fortunately, Go comes to the rescue again and makes this relatively easy by providing the [ParseMultipartForm](https://golang.org/pkg/net/http/#Request.ParseMultipartForm) method which places the data in [MultipartForm](https://golang.org/pkg/net/http/#Request) field in the Request object.

**Exercise 5: Write a program to receive multipart/form-data.**

The solution to exercise 5 is listed below (see also exercise5.go in this repo) and we can run it with `go run exercise5.go &`.

```
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

```
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
} [477 bytes data]
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Thu, 04 May 2017 18:32:07 GMT
< Content-Length: 259
<
{ [259 bytes data]
100  1347  100   259  100  1088   5638  23687 --:--:-- --:--:-- --:--:-- 24727
* Connection #0 to host 127.0.0.1 left intact
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
* In the `postContact` handler (line 12), we altered the test for Content-Type to use `Contains` instead of an exact match. The reason for this is that the header name in this case is actually `Content-Type: multipart/form-data; boundary=------------------------b8d179fba566a850`, so we have to search for a partial match.
* We use the `ParseMultipartForm` method (line 13) instead of the `ParseForm` method of exercise 4. There are different methods for different MIME-Types.
* The data ends up in `Request.MultipartForm` (line 13) instead of `Request.Form` of exercise 4. There are different data structure for different MIME-Types.
* Note that a file was included in the curl request. This is the main reason to use `multipart/form-data` in the first place. The data and the file are different parts of the request (thus `multipart`).
* The uploaded file can be read by using the [FormFile](https://golang.org/pkg/net/http/#Request.FormFile) data element in the Request object (line 16).
* For demonstration purposes, we just output the raw file as part of the response (line 20).
* There are many gems in the Request object that are worth studying.

#### MIME-Type application/json

There are other MIME-Types, but for the purposes of this tutorial, we will cover just one more of major significance in REST programming; which is (you guessed it) `application/json`.

**Exercise 6: Write a program to receive application/json.**

We've already created examples where we output JSON data from our CGI program. In this exercise, we will also receive JSON data and convert it to a native Go structure for processing. In many ways, it's one of the easiest MIME-Types to work with (we saved the best for the last).

The solution to exercise 6 is listed below (see also exercise6.go in this repo) and we can run it with `go run exercise6.go &`.


```
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

```
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

Some noteworthy items the above example:
* We added a small utility function (line 19) that generates unique record keys.
* Note that we started using the [constants](https://golang.org/pkg/net/http/#pkg-constants) in the http package for the status return codes (lines 29,26 and 41).
* Note that we introduce a new data structure (line 12) to receive the json data. There's a good tutorial on the marshal/unmarshal of json [here](https://blog.golang.org/json-and-go). There's also a good JSON to Go converter [here](https://mholt.github.io/json-to-go/) that will generate some Go code to speed things along.
* We instantiate the `contact` data structure (line 32) and `unmarshal` the JSON data from the body into the contact  (lines 33-34). Note that the data being passed has two fields and the structure has three.
* We then assign a new record ID to the contact (line 39) and convert (Marshal) the amalgamation back to JSON data in the response (line 42).

Now that we've covered how CGI works in Go, it's time to move on to the REST API standard for implementing CGI programs.

### What is REST?

Content TBD
