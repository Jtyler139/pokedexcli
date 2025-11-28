package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

func commandMap(c *Config) error {
	var res *http.Response
	var err error
	if c.Next != "" {
		res, err = http.Get(c.Next)
		if err != nil {
		return err
		}
	} else {
		res, err = http.Get("https://pokeapi.co/api/v2/location-area")
		if err != nil {
		return err
		}
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
