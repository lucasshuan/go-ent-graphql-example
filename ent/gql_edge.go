// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (u *User) Indication(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedIndication(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.IndicationOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryIndication().All(ctx)
	}
	return result, err
}

func (u *User) Indicated(ctx context.Context) (*User, error) {
	result, err := u.Edges.IndicatedOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryIndicated().Only(ctx)
	}
	return result, MaskNotFound(err)
}