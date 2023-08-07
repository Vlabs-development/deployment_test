// Copyright 2022-present Vlabs Development Kft
//
// All rights reserved under a proprietary license.
//
// Code generated by entc, DO NOT EDIT.

package ent

// CreateCarInput represents a mutation input for creating cars.
type CreateCarInput struct {
	IsSold *bool
	Name   string
	Price  int
}

// Mutate applies the CreateCarInput on the CarMutation builder.
func (i *CreateCarInput) Mutate(m *CarMutation) {
	if v := i.IsSold; v != nil {
		m.SetIsSold(*v)
	}
	m.SetName(i.Name)
	m.SetPrice(i.Price)
}

// SetInput applies the change-set in the CreateCarInput on the CarCreate builder.
func (c *CarCreate) SetInput(i CreateCarInput) *CarCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateCarInput represents a mutation input for updating cars.
type UpdateCarInput struct {
	IsSold *bool
	Name   *string
	Price  *int
}

// Mutate applies the UpdateCarInput on the CarMutation builder.
func (i *UpdateCarInput) Mutate(m *CarMutation) {
	if v := i.IsSold; v != nil {
		m.SetIsSold(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Price; v != nil {
		m.SetPrice(*v)
	}
}

// SetInput applies the change-set in the UpdateCarInput on the CarUpdate builder.
func (c *CarUpdate) SetInput(i UpdateCarInput) *CarUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateCarInput on the CarUpdateOne builder.
func (c *CarUpdateOne) SetInput(i UpdateCarInput) *CarUpdateOne {
	i.Mutate(c.Mutation())
	return c
}