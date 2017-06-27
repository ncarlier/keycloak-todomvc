package client

import (
	"log"
	"encoding/json"
	"todo/api"
)

func (c *Client) List() {
	resp, err := c.Get("todos", nil)

	if err != nil {
		log.Fatalf("Unable to contact remote server : %s", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Bad response from server : %s (%d)", resp.Status, resp.StatusCode)
		return
	}

	todosResponse := new(api.TodosResponse)
	json.NewDecoder(resp.Body).Decode(todosResponse)

	log.Printf("You have %d todo(s) :", todosResponse.Page.TotalElements)

	for _, todo := range todosResponse.Embedded.Todos {
		if(todo.Completed) {
			log.Printf("[x] %s", todo.Title)
		} else {
			log.Printf("[ ] %s", todo.Title)
		}
	}
}
