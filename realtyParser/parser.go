package realtyParser

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

type Parser struct {
	collector *colly.Collector
	sources   []Source
}

func NewParser() *Parser {
	collector := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))
	return &Parser{
		collector: collector,
		sources:   sources,
	}
}
