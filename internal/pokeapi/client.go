package pokeapi

import (
	"net/http"
	"time"
	"github.com/jtyler139/pokedexcli/internal/pokecache"
)

type Client struct {
	cache		pokecache.Cache	
	httpClient 	http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		cache:		pokecache.NewCache(timeout),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}