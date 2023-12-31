//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {

	ex, err := entgql.NewExtension(
		entgql.WithWhereInputs(true),
		entgql.WithSchemaPath("../presenters/graph/ent.graphqls"),
	)

	err = entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureModifier,
			gen.FeatureNamedEdges,
			gen.FeatureEntQL,
		},
		Header: `
		// Copyright 2022-present Vlabs Development Kft
		//
		// All rights reserved under a proprietary license.
		//
		// Code generated by entc, DO NOT EDIT.
		`,
		Templates: entgql.AllTemplates,
	}, entc.Extensions(ex))

	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}

	ex, err = entgql.NewExtension(
		entgql.WithWhereInputs(true),
		entgql.WithSchemaPath("../presenters/graph_admin/ent.graphqls"),
	)

	err = entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureModifier,
			gen.FeatureNamedEdges,
			gen.FeatureEntQL,
		},
		Header: `
		// Copyright 2022-present Vlabs Development Kft
		//
		// All rights reserved under a proprietary license.
		//
		// Code generated by entc, DO NOT EDIT.
		`,
		Templates: entgql.AllTemplates,
	}, entc.Extensions(ex))

	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
