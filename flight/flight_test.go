package flight

import (
	"testing"

	"github.com/goccy/go-json"
)

const testURL = "https://flight-api.naver.com/graphql"

func TestMinPricesByDestination_Body(t *testing.T) {
	g := NewMinPricesByDestination()
	v := VariablesMinPricesByDestination{
		CountryCodes:          []string{"JP"},
		DepartureMonths:       []string{"202602", "202603", "202604"},
		DepartureLocationCode: "SEL",
		Duration:              []int{1, 24},
		IsMappable:            true,
		Price:                 []int{10000, 3000000},
		TripDays:              []int{1, 15},
	}
	body, err := json.Marshal(g.WithVariables(v))
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(body))
}

func TestMinPricesByDestination(t *testing.T) {
	f := NewNaverFlight(testURL)

	v := VariablesMinPricesByDestination{
		CountryCodes:          []string{"JP"},
		DepartureMonths:       []string{"202602", "202603", "202604"},
		DepartureLocationCode: "SEL",
		Duration:              []int{1, 24},
		IsMappable:            true,
		Price:                 []int{10000, 3000000},
		TripDays:              []int{1, 15},
	}

	res, err := f.MinPricesByDestination(v)
	if err != nil {
		t.Error(err)
	}

	jsonB, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(jsonB))
}
