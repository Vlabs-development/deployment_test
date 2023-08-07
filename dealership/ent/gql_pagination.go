// Copyright 2022-present Vlabs Development Kft
//
// All rights reserved under a proprietary license.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"dealership/ent/cars"
	"dealership/ent/dealership"
	"errors"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[int]
	PageInfo       = entgql.PageInfo[int]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// CarsEdge is the edge representation of Cars.
type CarsEdge struct {
	Node   *Cars  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// CarsConnection is the connection containing edges to Cars.
type CarsConnection struct {
	Edges      []*CarsEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *CarsConnection) build(nodes []*Cars, pager *carsPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Cars
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Cars {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Cars {
			return nodes[i]
		}
	}
	c.Edges = make([]*CarsEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &CarsEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// CarsPaginateOption enables pagination customization.
type CarsPaginateOption func(*carsPager) error

// WithCarsOrder configures pagination ordering.
func WithCarsOrder(order *CarsOrder) CarsPaginateOption {
	if order == nil {
		order = DefaultCarsOrder
	}
	o := *order
	return func(pager *carsPager) error {
		if err := order.Direction.Validate(); err != nil {
			return err
		}
		if order.Field == nil {
			order.Field = DefaultCarsOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithCarsFilter configures pagination filter.
func WithCarsFilter(filter func(*CarsQuery) (*CarsQuery, error)) CarsPaginateOption {
	return func(pager *carsPager) error {
		if filter == nil {
			return errors.New("CarsQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type carsPager struct {
	reverse bool
	order   *CarsOrder
	filter  func(*CarsQuery) (*CarsQuery, error)
}

func newCarsPager(opts []CarsPaginateOption, reverse bool) (*carsPager, error) {
	pager := &carsPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultCarsOrder
	}
	return pager, nil
}

func (p *carsPager) applyFilter(query *CarsQuery) (*CarsQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *carsPager) toCursor(c *Cars) Cursor {
	return p.order.Field.toCursor(c)
}

func (p *carsPager) applyCursors(query *CarsQuery, after, before *Cursor) (*CarsQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultCarsOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *carsPager) applyOrder(query *CarsQuery) *CarsQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultCarsOrder.Field {
		query = query.Order(DefaultCarsOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *carsPager) orderExpr(query *CarsQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultCarsOrder.Field {
			b.Comma().Ident(DefaultCarsOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Cars.
func (c *CarsQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...CarsPaginateOption,
) (*CarsConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newCarsPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if c, err = pager.applyFilter(c); err != nil {
		return nil, err
	}
	conn := &CarsConnection{Edges: []*CarsEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = c.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if c, err = pager.applyCursors(c, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		c.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := c.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	c = pager.applyOrder(c)
	nodes, err := c.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// CarsOrderField defines the ordering field of Cars.
type CarsOrderField struct {
	// Value extracts the ordering value from the given Cars.
	Value    func(*Cars) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) cars.OrderOption
	toCursor func(*Cars) Cursor
}

// CarsOrder defines the ordering of Cars.
type CarsOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *CarsOrderField `json:"field"`
}

// DefaultCarsOrder is the default ordering of Cars.
var DefaultCarsOrder = &CarsOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &CarsOrderField{
		Value: func(c *Cars) (ent.Value, error) {
			return c.ID, nil
		},
		column: cars.FieldID,
		toTerm: cars.ByID,
		toCursor: func(c *Cars) Cursor {
			return Cursor{ID: c.ID}
		},
	},
}

// ToEdge converts Cars into CarsEdge.
func (c *Cars) ToEdge(order *CarsOrder) *CarsEdge {
	if order == nil {
		order = DefaultCarsOrder
	}
	return &CarsEdge{
		Node:   c,
		Cursor: order.Field.toCursor(c),
	}
}

// DealershipEdge is the edge representation of Dealership.
type DealershipEdge struct {
	Node   *Dealership `json:"node"`
	Cursor Cursor      `json:"cursor"`
}

// DealershipConnection is the connection containing edges to Dealership.
type DealershipConnection struct {
	Edges      []*DealershipEdge `json:"edges"`
	PageInfo   PageInfo          `json:"pageInfo"`
	TotalCount int               `json:"totalCount"`
}

func (c *DealershipConnection) build(nodes []*Dealership, pager *dealershipPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Dealership
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Dealership {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Dealership {
			return nodes[i]
		}
	}
	c.Edges = make([]*DealershipEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &DealershipEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// DealershipPaginateOption enables pagination customization.
type DealershipPaginateOption func(*dealershipPager) error

// WithDealershipOrder configures pagination ordering.
func WithDealershipOrder(order *DealershipOrder) DealershipPaginateOption {
	if order == nil {
		order = DefaultDealershipOrder
	}
	o := *order
	return func(pager *dealershipPager) error {
		if err := order.Direction.Validate(); err != nil {
			return err
		}
		if order.Field == nil {
			order.Field = DefaultDealershipOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithDealershipFilter configures pagination filter.
func WithDealershipFilter(filter func(*DealershipQuery) (*DealershipQuery, error)) DealershipPaginateOption {
	return func(pager *dealershipPager) error {
		if filter == nil {
			return errors.New("DealershipQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type dealershipPager struct {
	reverse bool
	order   *DealershipOrder
	filter  func(*DealershipQuery) (*DealershipQuery, error)
}

func newDealershipPager(opts []DealershipPaginateOption, reverse bool) (*dealershipPager, error) {
	pager := &dealershipPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultDealershipOrder
	}
	return pager, nil
}

func (p *dealershipPager) applyFilter(query *DealershipQuery) (*DealershipQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *dealershipPager) toCursor(d *Dealership) Cursor {
	return p.order.Field.toCursor(d)
}

func (p *dealershipPager) applyCursors(query *DealershipQuery, after, before *Cursor) (*DealershipQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultDealershipOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *dealershipPager) applyOrder(query *DealershipQuery) *DealershipQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultDealershipOrder.Field {
		query = query.Order(DefaultDealershipOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *dealershipPager) orderExpr(query *DealershipQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultDealershipOrder.Field {
			b.Comma().Ident(DefaultDealershipOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Dealership.
func (d *DealershipQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...DealershipPaginateOption,
) (*DealershipConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newDealershipPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if d, err = pager.applyFilter(d); err != nil {
		return nil, err
	}
	conn := &DealershipConnection{Edges: []*DealershipEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = d.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if d, err = pager.applyCursors(d, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		d.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := d.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	d = pager.applyOrder(d)
	nodes, err := d.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// DealershipOrderField defines the ordering field of Dealership.
type DealershipOrderField struct {
	// Value extracts the ordering value from the given Dealership.
	Value    func(*Dealership) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) dealership.OrderOption
	toCursor func(*Dealership) Cursor
}

// DealershipOrder defines the ordering of Dealership.
type DealershipOrder struct {
	Direction OrderDirection        `json:"direction"`
	Field     *DealershipOrderField `json:"field"`
}

// DefaultDealershipOrder is the default ordering of Dealership.
var DefaultDealershipOrder = &DealershipOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &DealershipOrderField{
		Value: func(d *Dealership) (ent.Value, error) {
			return d.ID, nil
		},
		column: dealership.FieldID,
		toTerm: dealership.ByID,
		toCursor: func(d *Dealership) Cursor {
			return Cursor{ID: d.ID}
		},
	},
}

// ToEdge converts Dealership into DealershipEdge.
func (d *Dealership) ToEdge(order *DealershipOrder) *DealershipEdge {
	if order == nil {
		order = DefaultDealershipOrder
	}
	return &DealershipEdge{
		Node:   d,
		Cursor: order.Field.toCursor(d),
	}
}