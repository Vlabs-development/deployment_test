// Copyright 2022-present Vlabs Development Kft
//
// All rights reserved under a proprietary license.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"dealership/ent/cars"
	"dealership/ent/dealership"
	"dealership/ent/predicate"
	"errors"
	"fmt"
)

// CarsWhereInput represents a where input for filtering Cars queries.
type CarsWhereInput struct {
	Predicates []predicate.Cars  `json:"-"`
	Not        *CarsWhereInput   `json:"not,omitempty"`
	Or         []*CarsWhereInput `json:"or,omitempty"`
	And        []*CarsWhereInput `json:"and,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *CarsWhereInput) AddPredicates(predicates ...predicate.Cars) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the CarsWhereInput filter on the CarsQuery builder.
func (i *CarsWhereInput) Filter(q *CarsQuery) (*CarsQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyCarsWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyCarsWhereInput is returned in case the CarsWhereInput is empty.
var ErrEmptyCarsWhereInput = errors.New("ent: empty predicate CarsWhereInput")

// P returns a predicate for filtering carsslice.
// An error is returned if the input is empty or invalid.
func (i *CarsWhereInput) P() (predicate.Cars, error) {
	var predicates []predicate.Cars
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, cars.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Cars, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, cars.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Cars, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, cars.And(and...))
	}
	predicates = append(predicates, i.Predicates...)

	switch len(predicates) {
	case 0:
		return nil, ErrEmptyCarsWhereInput
	case 1:
		return predicates[0], nil
	default:
		return cars.And(predicates...), nil
	}
}

// DealershipWhereInput represents a where input for filtering Dealership queries.
type DealershipWhereInput struct {
	Predicates []predicate.Dealership  `json:"-"`
	Not        *DealershipWhereInput   `json:"not,omitempty"`
	Or         []*DealershipWhereInput `json:"or,omitempty"`
	And        []*DealershipWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "city" field predicates.
	City             *string  `json:"city,omitempty"`
	CityNEQ          *string  `json:"cityNEQ,omitempty"`
	CityIn           []string `json:"cityIn,omitempty"`
	CityNotIn        []string `json:"cityNotIn,omitempty"`
	CityGT           *string  `json:"cityGT,omitempty"`
	CityGTE          *string  `json:"cityGTE,omitempty"`
	CityLT           *string  `json:"cityLT,omitempty"`
	CityLTE          *string  `json:"cityLTE,omitempty"`
	CityContains     *string  `json:"cityContains,omitempty"`
	CityHasPrefix    *string  `json:"cityHasPrefix,omitempty"`
	CityHasSuffix    *string  `json:"cityHasSuffix,omitempty"`
	CityEqualFold    *string  `json:"cityEqualFold,omitempty"`
	CityContainsFold *string  `json:"cityContainsFold,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *DealershipWhereInput) AddPredicates(predicates ...predicate.Dealership) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the DealershipWhereInput filter on the DealershipQuery builder.
func (i *DealershipWhereInput) Filter(q *DealershipQuery) (*DealershipQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyDealershipWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyDealershipWhereInput is returned in case the DealershipWhereInput is empty.
var ErrEmptyDealershipWhereInput = errors.New("ent: empty predicate DealershipWhereInput")

// P returns a predicate for filtering dealerships.
// An error is returned if the input is empty or invalid.
func (i *DealershipWhereInput) P() (predicate.Dealership, error) {
	var predicates []predicate.Dealership
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, dealership.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Dealership, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, dealership.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Dealership, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, dealership.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, dealership.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, dealership.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, dealership.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, dealership.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, dealership.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, dealership.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, dealership.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, dealership.IDLTE(*i.IDLTE))
	}
	if i.City != nil {
		predicates = append(predicates, dealership.CityEQ(*i.City))
	}
	if i.CityNEQ != nil {
		predicates = append(predicates, dealership.CityNEQ(*i.CityNEQ))
	}
	if len(i.CityIn) > 0 {
		predicates = append(predicates, dealership.CityIn(i.CityIn...))
	}
	if len(i.CityNotIn) > 0 {
		predicates = append(predicates, dealership.CityNotIn(i.CityNotIn...))
	}
	if i.CityGT != nil {
		predicates = append(predicates, dealership.CityGT(*i.CityGT))
	}
	if i.CityGTE != nil {
		predicates = append(predicates, dealership.CityGTE(*i.CityGTE))
	}
	if i.CityLT != nil {
		predicates = append(predicates, dealership.CityLT(*i.CityLT))
	}
	if i.CityLTE != nil {
		predicates = append(predicates, dealership.CityLTE(*i.CityLTE))
	}
	if i.CityContains != nil {
		predicates = append(predicates, dealership.CityContains(*i.CityContains))
	}
	if i.CityHasPrefix != nil {
		predicates = append(predicates, dealership.CityHasPrefix(*i.CityHasPrefix))
	}
	if i.CityHasSuffix != nil {
		predicates = append(predicates, dealership.CityHasSuffix(*i.CityHasSuffix))
	}
	if i.CityEqualFold != nil {
		predicates = append(predicates, dealership.CityEqualFold(*i.CityEqualFold))
	}
	if i.CityContainsFold != nil {
		predicates = append(predicates, dealership.CityContainsFold(*i.CityContainsFold))
	}
	if i.Name != nil {
		predicates = append(predicates, dealership.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, dealership.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, dealership.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, dealership.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, dealership.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, dealership.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, dealership.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, dealership.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, dealership.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, dealership.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, dealership.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, dealership.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, dealership.NameContainsFold(*i.NameContainsFold))
	}

	switch len(predicates) {
	case 0:
		return nil, ErrEmptyDealershipWhereInput
	case 1:
		return predicates[0], nil
	default:
		return dealership.And(predicates...), nil
	}
}
