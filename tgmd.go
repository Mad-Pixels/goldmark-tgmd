package tgmd

import (
	"fmt"
	"strings"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type Option interface {
	SetOption(*Config)
}

type Config struct{}

func (r Config) SetOption(c *Config) { *c = r }

type Renderer struct {
	Config Config
}

func NewRenderer(opts ...Option) renderer.NodeRenderer {
	r := &Renderer{
		Config: Config{},
	}
	for _, opt := range opts {
		opt.SetOption(&r.Config)
	}
	return r
}

func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindHeading, r.renderHeading)

	//
	reg.Register(ast.KindText, r.renderText)
	reg.Register(ast.KindLink, r.renderLink)
	reg.Register(ast.KindEmphasis, r.emphasis)
}

func (r *Renderer) emphasis(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Emphasis)
	if n.Level == 2 {
		_, _ = w.WriteString(textBold)
	}
	if n.Level == 1 {
		_, _ = w.WriteString(textItalics)
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderLink(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Link)
	if entering {
		_, _ = w.WriteString("[")
	} else {
		_, _ = w.WriteString(fmt.Sprintf("](%s)", string(n.Destination)))
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderHeading(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString(textBold)
	} else {
		_, _ = w.WriteString(textBold + "\n")
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderText(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	n := node.(*ast.Text)
	segment := n.Segment.Value(source)

	var escapedText strings.Builder
	for _, runeValue := range segment {
		switch runeValue {
		case '_', '*', '[', ']', '(', ')', '#', '+', '-', '=', '{', '}', '.', '!', '>', '<':
			escapedText.WriteString("\\")
		}
		escapedText.WriteByte(runeValue)
	}

	_, _ = w.WriteString(escapedText.String())
	return ast.WalkContinue, nil
}
