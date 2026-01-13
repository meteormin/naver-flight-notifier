package flight

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/goccy/go-json"
)

type NaverFlight struct {
	URL string
}

func (f *NaverFlight) MinPricesByDestination(v VariablesMinPricesByDestination) (*MinPricesByDestination, error) {
	g := NewMinPricesByDestination()

	body, err := json.Marshal(g.WithVariables(v))
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(f.URL, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(errBody))
	}

	var result MinPricesByDestination
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func NewNaverFlight(url string) *NaverFlight {
	return &NaverFlight{
		URL: url,
	}
}
