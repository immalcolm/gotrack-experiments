package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	connectToAPI()
}

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func connectToAPI() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts", nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var result []Post
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}
	for _, post := range result {
		fmt.Printf("UserID: %d, ID: %d, Title: %s, Body: %s\n",
			post.UserID, post.ID, post.Title, post.Body)
	}

	//var result []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)

}
