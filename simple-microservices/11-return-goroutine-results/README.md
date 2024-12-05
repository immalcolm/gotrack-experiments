# Using WaitGroups
To ensure that the application ends when all the goroutines have completed their tasks, you can use a `sync.WaitGroup`. A `WaitGroup` waits for a collection of goroutines to finish executing. You can add a counter to the `WaitGroup` for each goroutine you start, and then decrement the counter when each goroutine completes.

Here's how you can modify your code to use a `WaitGroup` to ensure the application ends when all the goroutines are done:

1. Import the sync package.
2. Create a WaitGroup variable.
3. Add a counter to the WaitGroup for each goroutine.
4. Decrement the counter when each goroutine completes.
5. Wait for all goroutines to finish before exiting the main function.

## Channels
Channels in Go are a way for different parts of a program to communicate with each other. Think of channels as pipes or conduits through which you can send and receive messages between different parts of your program, specifically between different goroutines (which are lightweight threads managed by Go).

### Key Points:

1. **Communication:**
Channels allow one part of your program to send information to another part. This is useful when you have multiple tasks running at the same time (concurrently) and they need to share data.

2. **Synchronization:**
Channels help synchronize the execution of different tasks. When one task sends a message through a channel, it can wait until another task receives that message, ensuring that both tasks are coordinated.

3. **Type-Safe:**
Channels are type-safe, meaning you can only send and receive specific types of data through a channel. For example, if you create a channel for integers, you can only send and receive integers through that channel.

### How Channels Work:
1. Creating a Channel:
- You create a channel using the make function. For example, `ch := make(chan int)` creates a channel that can carry integer values.

2. Sending Data:
- You send data to a channel using the `<-` operator. For example, `ch <- 42` sends the integer `42` to the channel ch.

3. Receiving Data:
- You receive data from a channel using the `<-` operator. For example, `value := <-ch` receives a value from the channel ch and stores it in the variable value.