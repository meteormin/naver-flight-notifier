package flight

import "github.com/goccy/go-json"

type Operation struct {
	Name      string
	Variables map[string]any
}

type GraphQL[T any] struct {
	operationName string
	query         string
	variables     T
}

func (g *GraphQL[T]) OperationName() string {
	return g.operationName
}

func (g *GraphQL[T]) Query() string {
	return g.query
}

func (g *GraphQL[T]) WithVariables(variables T) *GraphQL[T] {
	g.variables = variables
	return g
}

func (g *GraphQL[T]) Variables() T {
	return g.variables
}

func (g *GraphQL[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"operationName": g.operationName,
		"query":         g.query,
		"variables":     g.variables,
	})
}

func NewMinPricesByDestination() *GraphQL[VariablesMinPricesByDestination] {
	return &GraphQL[VariablesMinPricesByDestination]{operationName: OperationMinPricesByDestination, query: QueryMinPricesByDestination}
}
