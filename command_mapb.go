package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

func commandMapb(c *Config) error {
	var res *http.Response
	var err error
	if c.Previous != nil {
		res, err = http.Get(*c.Previous)
		if err != nil {
		return err
		}
	} else {
		fmt.Println("you're on the first page")
		return nil
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	
	if err != nil {
		return err
	}

	page := Page{}

	err = json.Unmarshal(body, &page)
	if err != nil {
		return err
	}

	c.Next = page.Next
	
	c.Previous = page.Previous

	for _, result := range page.Results {
		fmt.Println(result.Name)
	}

	return nil
}