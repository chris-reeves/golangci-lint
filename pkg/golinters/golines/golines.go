package golines

import (
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/pkg/config"
	"github.com/golangci/golangci-lint/pkg/goanalysis"
	"github.com/golangci/golangci-lint/pkg/goformatters"
	golinesbase "github.com/golangci/golangci-lint/pkg/goformatters/golines"
	"github.com/golangci/golangci-lint/pkg/golinters/internal"
)

const linterName = "golines"

func New(settings *config.GoLinesSettings) *goanalysis.Linter {
	a := goformatters.NewAnalyzer(
		internal.LinterLogger.Child(linterName),
		"Checks if code is formatted, and fixes long lines",
		golinesbase.New(settings),
	)

	return goanalysis.NewLinter(
		a.Name,
		a.Doc,
		[]*analysis.Analyzer{a},
		nil,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
