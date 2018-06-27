package trafficlightco

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// ratingsRequest is the parameter used for the request on the API
type ratingsRequest struct {
	CompanyID string `json:"company"`
}

// Rating is the response given by the API
// color is either green, yellow, red, black
type Rating struct {
	Color     string    `json:"color"`
	CompanyID string    `json:"company"`
	Created   time.Time `json:"created"`
}

// GetRating will return the rating of a company
func (c *Client) GetRating(cID string) (*Rating, error) {
	if cID == "" {
		return nil, fmt.Errorf("GetRating - empty cID")
	}
	rr := ratingsRequest{CompanyID: cID}
	b, errM := json.Marshal(rr)
	if errM != nil {
		return nil, fmt.Errorf("GetRating - Marshal: %v", errM)
	}

	req, errNR := http.NewRequest(
		http.MethodPost,
		c.apiURL.String()+"/v1.0/ratings",
		bytes.NewReader(b),
	)
	if errNR != nil {
		return nil, fmt.Errorf("GetRating - NewRequest: %v", errNR)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.apiKey)

	resp, errDo := c.httpClient.Do(req)
	if errDo != nil {
		return nil, fmt.Errorf("GetRating - Do: %v", errDo)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatalf("GetRating - Do: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf(
			"GetRating - got status %d for url %s and data %s",
			resp.StatusCode,
			req.URL.String(),
			string(b),
		)
	}
	if resp.Header.Get("content-type") != "application/json" {
		return nil, fmt.Errorf("GetRating - content-type not application/json")
	}

	decoder := json.NewDecoder(resp.Body)
	var rat Rating
	errD := decoder.Decode(&rat)
	if errD != nil {
		return nil, fmt.Errorf("GetRating - Decode: %v", errD)
	}

	return &rat, nil
}
