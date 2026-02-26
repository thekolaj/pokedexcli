package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	return handleLocations(cfg.nextLocationsURL, cfg)
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	return handleLocations(cfg.prevLocationsURL, cfg)
}

func handleLocations(url *string, cfg *config) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(url)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
