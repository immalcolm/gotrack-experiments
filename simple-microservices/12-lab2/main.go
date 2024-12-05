package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var (
	apis map[int]string
	c    chan map[int]interface{}
	wg   sync.WaitGroup
)

func fetchData(API int) {
	defer wg.Done()

	url := apis[API]
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := io.ReadAll(resp.Body); err == nil {
			var result map[string]interface{}
			json.Unmarshal(body, &result)

			var re = make(map[int]interface{})

			switch API {
			case 1:
				if result["success"] == true {
					re[API] = result["rates"].(map[string]interface{})["USD"]
				} else {
					re[API] = result["error"].(map[string]interface{})["info"]
				}
				c <- re
				fmt.Println("Result for API 1 stored")

			case 2:
				if result["main"] != nil {
					re[API] = result["main"].(map[string]interface{})["temp"]
				} else {
					re[API] = (result["message"])
				}
				c <- re
				fmt.Println("Result for API 2 stored")
			case 3:
				if result["status"] == "ok" {

					/*
						articles := result["articles"].([]interface{})
						articleList := make([]map[string]interface{}, len(articles))
						for i, article := range articles {
							articleMap := article.(map[string]interface{})
							articleList[i] = map[string]interface{}{
								"source":      articleMap["source"].(map[string]interface{})["name"],
								"title":       articleMap["title"],
								"description": articleMap["description"],
							}
						}
						re[API] = articleList

						// Format and print the articles
						for _, article := range articleList {
							fmt.Printf("Source: %s\nTitle: %s\nDescription: %s\n\n",
								article["source"], article["title"], article["description"])
						}
					*/

					/*
						re[API] = map[string]interface{}{
							"source":      result["articles"].([]interface{})[0].(map[string]interface{})["source"].(map[string]interface{})["name"],
							"title":       result["articles"].([]interface{})[0].(map[string]interface{})["title"],
							"description": result["articles"].([]interface{})[0].(map[string]interface{})["description"],
						}
					*/
					articles := result["articles"].([]interface{})
					var formattedArticles string
					for _, article := range articles {
						articleMap := article.(map[string]interface{})
						formattedArticles += fmt.Sprintf("Source: %s\nTitle: %s\nDescription: %s\n\n",
							articleMap["source"].(map[string]interface{})["name"],
							articleMap["title"],
							articleMap["description"])
					}
					re[API] = formattedArticles
					// store the result into the channel
					c <- re
					fmt.Println("Result for API 3 stored")
				} else {
					re[API] = result["error"]
					fmt.Println("err")
					fmt.Println(result["error"])
				}
				c <- re
				fmt.Println("Result for API 3 stored")
			} //end switch
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}
func main() {

	apis = make(map[int]string)
	c = make(chan map[int]interface{})

	apis[1] = "http://data.fixer.io/api/latest?access_key=31063bbaa0fbd7f73fc9dae8a4dcaf53"
	apis[2] = "http://api.openweathermap.org/data/2.5/weather?q=SINGAPORE&appid=6ddb1c58d20eb843ae3bab549760c3bf"
	apis[3] = "https://newsapi.org/v2/top-headlines?country=us&category=business&apiKey=68abed74b6aa468aa1c2bcdd04ccaca9"

	wg.Add(3)

	go fetchData(1)
	go fetchData(2)
	go fetchData(3)

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(c) // Close the channel after all goroutines have finished
	}()

	// we expect 3 results in the channel
	/*
		for i := 0; i < 3; i++ {
			fmt.Println(<-c)
		}*/
	// Read and print results from the channel
	for result := range c {
		fmt.Println(result)
	}

	fmt.Println("Done!")
	//fmt.Scanln()
}
