# Exercise WaitGroup
Using Scanln() to ensure that main will only exit after the goroutines have finished execution is not the best way of doing it.

You have learnt how to make sure of WaitGroup when studying concurrency in Go Advanced. Now improve on the previous example to incorporate the use of WaitGroup.

Write a Go program to display the list of news sources. Include the following fields:
- Name of the news source
- Title of the news
- Description of the news

**Given the following URL:**
https://newsapi.org/v2/top-headlines?country=us&category=business&apiKey=<api_key>
https://newsapi.org/v2/top-headlines?country=us&category=business&apiKey=68abed74b6aa468aa1c2bcdd04ccaca9

0d45b317eb00401881b5ba3ed4091dcf