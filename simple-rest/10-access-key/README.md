
**1. Extracting Query Parameters:**
The expression `r.URL.Query()` is used to extract the query parameters from the URL of the incoming HTTP request r. 
The Query method returns a `url.Values` type, which is essentially a map of string keys to slices of string values (`map[string][]string`). This map contains all the query parameters present in the URL.

`kv := r.URL.Query()`

For example, if the URL is `http://example.com/search?q=golang&sort=asc`, the kv map will contain:
```go
map[string][]string{
    "q":    {"golang"},
    "sort": {"asc"},
}
```

2. **Iterating Over the Query Parameters:** The `for` loop iterates over the key-value pairs in the kv map. The loop variable k represents the key (the name of the query parameter), and v represents the value (a slice of strings containing the values for that query parameter).

3. **Printing the Key-Value Pairs:** Inside the loop, the `fmt.Println` function is used to print each key-value pair to the console. The `fmt.Println` function takes any number of arguments, formats them using their default string representations, and prints them separated by spaces, followed by a newline.

For the example URL `http://example.com/search?q=golang&sort=asc`, the output will be:
```
q [golang]
sort [asc]
```

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    kv := r.URL.Query()
    for k, v := range kv {
        fmt.Println(k, v) // print out the key/value pair
    }
    fmt.Fprintln(w, "Query parameters printed to console")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

## References
- https://golang.org/pkg/net/http/#ResponseWriter
