# Simple Fetch API Web Service


## Sample APIs
https://fixer.io
Using `json.NewDecoder` and `json.Unmarshal` are both common methods for decoding JSON data in Go. Each method has its own pros and cons, depending on the specific use case.

**Using `json.NewDecoder`**
### Pros:
**Streaming:**
- `json.NewDecoder` can handle large JSON payloads efficiently by streaming the data directly from the io.Reader without loading the entire content into memory.
- This is useful for processing large JSON responses or files.

**Flexibility:**
- It allows for incremental decoding, which means you can decode parts of the JSON data as they are read from the stream.
- This is useful for processing JSON data in chunks or when dealing with continuous data streams.

**Error Handling:**
- Errors can be detected and handled immediately as the data is being read and decoded.

### Cons:

**Complexity:**
It can be more complex to use when you need to handle nested structures or multiple JSON objects in a single stream.
Requires careful management of the io.Reader to ensure proper resource cleanup.



### Using `json.Unmarshal`
### Pros:
**Simplicity:**
- `json.Unmarshal` is straightforward and easy to use for decoding JSON data from a byte slice.
- It is ideal for simple use cases where the entire JSON payload can be read into memory at once.

**Readability:**
- The code is often more readable and concise, making it easier to understand and maintain.

**Common Use Case:**
It is commonly used for decoding JSON data from HTTP responses, files, or other sources where the entire content can be read into memory.

### Cons:
**Memory Usage:**
- `json.Unmarshal` requires the entire JSON payload to be read into memory as a byte slice before decoding.
- This can be inefficient and problematic for very large JSON payloads, leading to high memory consumption.

**Error Handling:**
- Errors are detected only after the entire JSON payload has been read and unmarshaled, which may delay error detection.
