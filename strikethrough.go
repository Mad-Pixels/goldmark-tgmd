package tgmd

import (
	"github.com/yuin/goldmark"
	ext "github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type strikethrough struct{}

var Strikethroughs = &strikethrough{}

// Extend ...
func (e *strikethrough) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(ext.NewStrikethroughParser(), 500),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewRenderer(nil), 500),
	))
}
