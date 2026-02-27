package pokeapi

import (
	"encoding/json"
)

// GetLocation -
func (c *Client) GetLocation(identifier string) (Location, error) {
	url := baseURL + "/location-area/" + identifier
	dat, err := c.Get(url)
	if err != nil {
		return Location{}, err
	}

	resp := Location{}
	err = json.Unmarshal(dat, &resp)
	if err != nil {
		return Location{}, err
	}

	return resp, nil
}
