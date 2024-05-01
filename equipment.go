package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	CategoryAlpinbriller   = 6859
	CategoryAlpinhjelm     = 6877
	CategoryAlpinsko       = 8313
	CategoryAlpinskiVoksen = 8315
	CategoryAlpinskiBarn   = 8314
	CategoryAlpinski       = 12859
	CategoryAlpinstaver    = 6922
)

const (
	StatePublished = "published"
)

const (
	StatusBorrowed = "borrowed"
)

func (c Client) GetEquipment(result map[int][]string, category int, state string, status string) error {
	variants, err := c.getVariants(category, StatePublished)
	if err != nil {
		return err
	}

	for _, variant := range variants.Data {
		instances := variant.Relationships.Instances.Data

		for _, instance := range instances {
			if instance.Attributes.Status != status {
				continue
			}

			equipment, ok := result[instance.Attributes.LastLoanID]
			if ok {
				result[instance.Attributes.LastLoanID] = append(equipment, instance.Attributes.Identifier)
			} else {
				result[instance.Attributes.LastLoanID] = make([]string, 1)
				result[instance.Attributes.LastLoanID][0] = instance.Attributes.Identifier
			}
		}
	}

	return nil
}

func (c Client) getVariants(category int, state string) (*ListVariantResponse, error) {
	client := http.Client{}

	endpoint := fmt.Sprintf("%s/equipment/%d/variants?limit=0&filters[state]=%s", c.host, category, state)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", c.token)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data ListVariantResponse
	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
