package flight

type OperationName string

const (
	MinPricesByDestination = "minPricesByDestination"
)

type Operation struct {
	Name      string
	Variables map[string]any
}

type GraphQL struct {
	OperationName OperationName
	Query         string
	Variables     map[string]any
}

func Query() {
}
