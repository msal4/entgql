// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (a *AssignmentQuery) CollectFields(ctx context.Context, satisfies ...string) *AssignmentQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		a = a.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return a
}

func (a *AssignmentQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *AssignmentQuery {
	return a
}
