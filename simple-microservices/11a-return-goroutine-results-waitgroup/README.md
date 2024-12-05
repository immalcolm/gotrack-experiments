
## Explanation:

1. **Import the `sync` Package:**
- The `sync` package is imported to use the `WaitGroup`.

2. **Declare a `WaitGroup` Variable:**
- `var wg sync.WaitGroup` declares a `WaitGroup` variable named `wg`.

3. **Add a Counter to the WaitGroup:**
- `wg.Add(3)` adds a counter to the WaitGroup for each goroutine that will be started.

4. **Decrement the Counter When Each Goroutine Completes:**
- `defer wg.Done()` is added at the beginning of the `fetchData` function to decrement the counter when the goroutine completes.

5. **Wait for All Goroutines to Finish:**
- `wg.Wait()` is called in the `main` function to block until all goroutines have finished executing.

6. **Close the Channel:**
- `close(c)` is called after `wg.Wait()` to close the channel after all goroutines have finished.

7. **Read and Print Results from the Channel:**
- A `for` loop is used to read and print results from the channel until it is closed.

By using a WaitGroup, you can ensure that the main function waits for all goroutines to complete before exiting, allowing the application to end gracefully when all tasks are done.
