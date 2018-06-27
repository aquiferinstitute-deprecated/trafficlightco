package trafficlightco

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// CompaniesResponse is the JSON response from the companies endpoint
type CompaniesResponse struct {
	Results  []Company `json:"results"`
	Previous int64     `json:"previous"`
	Next     int64     `json:"next"`
}

// Company is represented by its name and id
type Company struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// Search will return a list of companies matching the possible criteria
func (c *Client) Search(co, na, iType, iValue string) ([]Company, error) {

	req, errReq := buildSearchRequest(c.apiURL, c.apiKey, co, na, iType, iValue)
	if errReq != nil {
		return nil, fmt.Errorf("Search: %v", errReq)
	}

	resp, errDo := c.httpClient.Do(req)
	if errDo != nil {
		return nil, fmt.Errorf("Search - Do: %v", errDo)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatalf("Search - Do: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"Search - got status %d for %s",
			resp.StatusCode, req.URL.String(),
		)
	}
	if resp.Header.Get("content-type") != "application/json" {
		return nil, fmt.Errorf("Search - content-type not application/json")
	}

	decoder := json.NewDecoder(resp.Body)
	var rez CompaniesResponse
	errD := decoder.Decode(&rez)
	if errD != nil {
		return nil, fmt.Errorf("Search - Decode: %v", errD)
	}

	return rez.Results, nil
}

// buildSearchQuery will build the search request
func buildSearchRequest(
	apiURL fmt.Stringer, apiKey,
	co, na, iType, iValue string,
) (*http.Request, error) {
	req, errNR := http.NewRequest(
		http.MethodGet,
		apiURL.String()+"/v1.0/companies",
		nil,
	)
	if errNR != nil {
		return nil, fmt.Errorf("buildSearchRequest - NewRequest: %v", errNR)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)

	q := req.URL.Query()
	if co != "" {
		q.Add("country", co)
	}
	if na != "" {
		q.Add("name", na)
	}
	if iType != "" && iValue != "" {
		q.Add("identifier_type", iType)
		q.Add("identifier_value", iValue)
	}

	if len(q) == 0 {
		return nil, fmt.Errorf("buildSearchRequest - no search criteria")
	}

	req.URL.RawQuery = q.Encode()

	return req, nil
}
