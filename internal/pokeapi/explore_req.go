package pokeapi

import "encoding/json"

func (c *Client) ExploreLocationArea(areaName string) (LocationAreaExploreResp, error) {
	url := baseURL + "/location-area/" + areaName

	data, ok := c.cache.Get(url)
	if !ok {
		var err error
		data, err = c.MakeRequest(url)
		if err != nil {
			return LocationAreaExploreResp{}, err
		}
	}

	locationAreaResp := LocationAreaExploreResp{}
	err := json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreaExploreResp{}, err
	}

	c.cache.Add(url, data)

	return locationAreaResp, nil
}
