// Code generated by ent, DO NOT EDIT.

package ent

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Name          string
	IndicationIDs []int
	IndicatedID   *int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	m.SetName(i.Name)
	if v := i.IndicationIDs; len(v) > 0 {
		m.AddIndicationIDs(v...)
	}
	if v := i.IndicatedID; v != nil {
		m.SetIndicatedID(*v)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	Name                *string
	ClearIndication     bool
	AddIndicationIDs    []int
	RemoveIndicationIDs []int
	ClearIndicated      bool
	IndicatedID         *int
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearIndication {
		m.ClearIndication()
	}
	if v := i.AddIndicationIDs; len(v) > 0 {
		m.AddIndicationIDs(v...)
	}
	if v := i.RemoveIndicationIDs; len(v) > 0 {
		m.RemoveIndicationIDs(v...)
	}
	if i.ClearIndicated {
		m.ClearIndicated()
	}
	if v := i.IndicatedID; v != nil {
		m.SetIndicatedID(*v)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdate builder.
func (c *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdateOne builder.
func (c *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}