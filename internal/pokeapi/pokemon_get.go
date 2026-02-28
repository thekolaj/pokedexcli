package pokeapi

import (
	"encoding/json"
)

// GetPokemon -
func (c *Client) GetPokemon(identifier string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + identifier
	dat, err := c.Get(url)
	if err != nil {
		return Pokemon{}, err
	}

	resp := Pokemon{}
	err = json.Unmarshal(dat, &resp)
	if err != nil {
		return Pokemon{}, err
	}

	return resp, nil
}
