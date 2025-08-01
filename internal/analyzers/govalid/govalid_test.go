package govalid_test

import (
	"flag"
	"testing"

	"github.com/gostaticanalysis/codegen/codegentest"

	"github.com/sivchari/govalid/internal/analyzers/govalid"
	"github.com/sivchari/govalid/internal/analyzers/markers"
	"github.com/sivchari/govalid/internal/analyzers/registry"
)

var update bool

func init() {
	flag.BoolVar(&update, "update", false, "update golden files")
}

func Test(t *testing.T) {
	registry := registry.NewRegistry(
		registry.AddAnalyzers(markers.Initializer()),
		registry.AddGenerators(govalid.Initializer()),
	)

	if err := registry.Init(nil); err != nil {
		t.Fatalf("failed to initialize analyzers: %v", err)
	}

	govalid, err := registry.Generator(govalid.Name)
	if err != nil {
		t.Fatalf("failed to get govalid generator: %v", err)
	}

	markers := []string{
		"required",
		"lt",
		"gt",
		"maxlength",
		"maxitems",
		"minitems",
		"minlength",
		"gte",
		"lte",
		"enum",
		"email",
		"uuid",
		"url",
		"cel",
		"length",
	}

	for _, tc := range markers {
		t.Run(tc, func(t *testing.T) {
			results := codegentest.Run(t, codegentest.TestData(), govalid, tc)
			codegentest.Golden(t, results, update)
		})
	}
}
