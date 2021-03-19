// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// +build ignore

package main

import (
	"fmt"
	"log"
	"strings"
	"text/template"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	// A usage for custom templates with external functions.
	// One template is defined in the option below, and the
	// rest are provided with the `Templates` option.
	opts := []entc.Option{
		entc.TemplateFiles("template/stringer.tmpl"),
	}
	err := entc.Generate("./schema", &gen.Config{
		Header: `
			// Copyright 2019-present Facebook Inc. All rights reserved.
			// This source code is licensed under the Apache 2.0 license found
			// in the LICENSE file in the root directory of this source tree.

			// Code generated by entc, DO NOT EDIT.
		`,
		Templates: []*gen.Template{
			gen.MustParse(gen.NewTemplate("static").
				Funcs(template.FuncMap{"title": strings.ToTitle}).
				ParseFiles("template/static.tmpl")),
			gen.MustParse(gen.NewTemplate("debug").
				Funcs(template.FuncMap{"byName": byName}).
				ParseFiles("template/debug.tmpl")),
		},
		Hooks: []gen.Hook{TagFields("json")},
	}, opts...)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

// byName returns a node in the graph by its label/name.
func byName(g *gen.Graph, name string) (*gen.Type, error) {
	for _, n := range g.Nodes {
		if n.Name == name {
			return n, nil
		}
	}
	return nil, fmt.Errorf("node %q was not found in the graph", name)
}

// TagFields tags all fields defined in the schema with the given struct-tag.
func TagFields(name string) gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			for _, node := range g.Nodes {
				for _, field := range node.Fields {
					field.StructTag = fmt.Sprintf("%s:%q", name, field.Name)
				}
			}
			return next.Generate(g)
		})
	}
}
