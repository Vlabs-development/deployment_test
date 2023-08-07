// Copyright 2022-present Vlabs Development Kft
//
// All rights reserved under a proprietary license.
//
// Code generated by entc, DO NOT EDIT.

package ent

// CreateDealershipInput represents a mutation input for creating dealerships.
type CreateDealershipInput struct {
	City   string
	Name   string
	CarIDs []int
}

// Mutate applies the CreateDealershipInput on the DealershipMutation builder.
func (i *CreateDealershipInput) Mutate(m *DealershipMutation) {
	m.SetCity(i.City)
	m.SetName(i.Name)
	if v := i.CarIDs; len(v) > 0 {
		m.AddCarIDs(v...)
	}
}

// SetInput applies the change-set in the CreateDealershipInput on the DealershipCreate builder.
func (c *DealershipCreate) SetInput(i CreateDealershipInput) *DealershipCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateDealershipInput represents a mutation input for updating dealerships.
type UpdateDealershipInput struct {
	City         *string
	Name         *string
	ClearCars    bool
	AddCarIDs    []int
	RemoveCarIDs []int
}

// Mutate applies the UpdateDealershipInput on the DealershipMutation builder.
func (i *UpdateDealershipInput) Mutate(m *DealershipMutation) {
	if v := i.City; v != nil {
		m.SetCity(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearCars {
		m.ClearCars()
	}
	if v := i.AddCarIDs; len(v) > 0 {
		m.AddCarIDs(v...)
	}
	if v := i.RemoveCarIDs; len(v) > 0 {
		m.RemoveCarIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateDealershipInput on the DealershipUpdate builder.
func (c *DealershipUpdate) SetInput(i UpdateDealershipInput) *DealershipUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateDealershipInput on the DealershipUpdateOne builder.
func (c *DealershipUpdateOne) SetInput(i UpdateDealershipInput) *DealershipUpdateOne {
	i.Mutate(c.Mutation())
	return c
}