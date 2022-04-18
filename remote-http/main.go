package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TodoArray []struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("hello world")
	url := "https://jsonplaceholder.typicode.com/todos/"
	resp, err := http.Get(url)

	if err != nil {
		println(err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("error fetching data")
		return
	}

	var todos TodoArray
	err = json.NewDecoder(resp.Body).Decode(&todos)
	if err != nil {
		println(err)
		return
	}

	for _, todo := range todos {
		println(todo.Title)
	}
}
