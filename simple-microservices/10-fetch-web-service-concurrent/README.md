# Fetching Web services concurrently
Be sure to call the `Scanln()` function before the end of the main() function; otherwise, the application will exit before the goroutines have the chance to finish execution

Alternatively, sync.Waitgroup together with its Wait() function can be used

**Starting Goroutines:** The `go` keyword is used to **start new goroutines**. A goroutine is a lightweight thread managed by the Go runtime. 
By prefixing the `fetchData` function calls with `go`, the program starts three separate goroutines to execute the fetchData function concurrently.

### Breakdown
Here's how the fetchData function works in conjunction with the active selection:

- The fetchData function takes an integer argument (API) and uses it to determine which URL to fetch data from. The URLs are stored in a global variable `apis`, which is a map of integers to strings.
- The function makes an HTTP GET request to the specified URL using `http.Get`.
- If the request is successful, the response body is read using `io.ReadAll`.
- The JSON data from the response body is unmarshaled into a `map[string]interface{}`.
- A `switch` statement is used to handle the response based on the API argument. Different cases handle different APIs and process the JSON data accordingly.