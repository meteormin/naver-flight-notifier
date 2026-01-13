package flight

// MinPricesByDestination
const (
	OperationMinPricesByDestination = "minPricesByDestination"
	QueryMinPricesByDestination     = `query minPricesByDestination($departureLocationCode: String, $arrivalLocationCode: String, $continentIds: [Int!], $countryCodes: [String!], $departureDate: String, $returnDate: String, $departureMonths: [String!], $duration: [Int!], $isDomestic: Boolean, $isMappable: Boolean, $price: [Int!], $size: Int, $stops: Int, $themeIds: [Int!], $timeCategories: [DepartureTimeCategory!], $tripDays: [Int!], $tripType: String) {
  minPricesByDestination(
    departureLocationCode: $departureLocationCode
    arrivalLocationCode: $arrivalLocationCode
    continentIds: $continentIds
    countryCodes: $countryCodes
    departureDate: $departureDate
    returnDate: $returnDate
    departureMonths: $departureMonths
    duration: $duration
    isDomestic: $isDomestic
    isMappable: $isMappable
    price: $price
    size: $size
    stops: $stops
    themeIds: $themeIds
    timeCategories: $timeCategories
    tripDays: $tripDays
    tripType: $tripType
  ) {
    origin {
      airportCode
      airportName
      cityName
      isDomestic
      location {
        latitude
        longitude
        __typename
      }
      popularity
      type
      __typename
    }
    destination {
      airportCode
      airportName
      cityImage {
        imageUrl
        __typename
      }
      cityName
      event
      isDomestic
      location {
        latitude
        longitude
        __typename
      }
      popularity
      type
      __typename
    }
    departureDate
    airlineCodes
    continentId
    countryCode
    duration
    isDomestic
    minPrice
    returnDate
    stops
    themeIds
    timeCategory
    tripDays
    tripType
    __typename
  }
}`
)

type VariablesMinPricesByDestination struct {
	DepartureLocationCode string   `json:"departureLocationCode,omitempty"`
	ArrivalLocationCode   string   `json:"arrivalLocationCode,omitempty"`
	ContinentIds          []int    `json:"continentIds,omitempty"`
	CountryCodes          []string `json:"countryCodes,omitempty"`
	DepartureDate         string   `json:"departureDate,omitempty"`
	ReturnDate            string   `json:"returnDate,omitempty"`
	DepartureMonths       []string `json:"departureMonths,omitempty"`
	Duration              []int    `json:"duration,omitempty"`
	IsDomestic            bool     `json:"isDomestic,omitempty"`
	IsMappable            bool     `json:"isMappable,omitempty"`
	Price                 []int    `json:"price,omitempty"`
	Size                  int      `json:"size,omitempty"`
	Stops                 int      `json:"stops,omitempty"`
	ThemeIds              []int    `json:"themeIds,omitempty"`
	TimeCategories        []string `json:"timeCategories,omitempty"`
	TripDays              []int    `json:"tripDays,omitempty"`
	TripType              string   `json:"tripType,omitempty"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Typename  string  `json:"__typename"`
}

type CityImage struct {
	ImageUrl string `json:"imageUrl"`
	Typename string `json:"__typename"`
}

type Destination struct {
	AirportCode string   `json:"airportCode"`
	AirportName string   `json:"airportName"`
	CityName    string   `json:"cityName"`
	IsDomestic  bool     `json:"isDomestic"`
	Location    Location `json:"location"`
	Popularity  int      `json:"popularity"`
	Type        string   `json:"type"`
	Typename    string   `json:"__typename"`
}

type Origin struct {
	Destination
	CityImage CityImage `json:"cityImage"`
}

type MinPricesByDestination struct {
	Origin        Origin      `json:"origin"`
	Destination   Destination `json:"destination"`
	DepartureDate string      `json:"departureDate"`
	AirlineCodes  []string    `json:"airlineCodes"`
	ContinentId   int         `json:"continentId"`
	CountryCode   string      `json:"countryCode"`
	Duration      int         `json:"duration"`
	IsDomestic    bool        `json:"isDomestic"`
	MinPrice      int         `json:"minPrice"`
	ReturnDate    string      `json:"returnDate"`
	Stops         int         `json:"stops"`
	ThemeIds      []int       `json:"themeIds"`
	TimeCategory  string      `json:"timeCategory"`
	TripDays      []int       `json:"tripDays"`
	TripType      string      `json:"tripType"`
	Typename      string      `json:"__typename"`
}
