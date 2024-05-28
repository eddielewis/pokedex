package pokeapi

import (
	"encoding/json"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	url := baseURL + "/location-area?offset=0&limit=20"

	if pageURL != nil {
		url = *pageURL
	}

	data, ok := c.cache.Get(url)
	if !ok {
		var err error
		data, err = c.MakeRequest(url)
		if err != nil {
			return LocationAreasResp{}, err
		}
	}

	locationAreasResp := LocationAreasResp{}
	err := json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(url, data)

	return locationAreasResp, nil
}
