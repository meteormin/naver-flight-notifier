package flight

type NaverFlight struct {
	URL string
}

func (f *NaverFlight) Query() {
}

func NewNaverFlight(url string) *NaverFlight {
	return &NaverFlight{
		URL: url,
	}
}
