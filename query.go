package esgenerics

import "context"

// QueryHandler is an interface for creating queries.
//
//go:generate mockery --name=QueryHandler --inpackage --case=snake
type QueryHandler interface {
	BuildQuery(ctx context.Context) QueryMap
}
