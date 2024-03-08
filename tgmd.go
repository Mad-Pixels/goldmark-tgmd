package tgmd

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type Renderer struct {
	Config *config
}

func NewRenderer(c *config) renderer.NodeRenderer {
	return &Renderer{
		Config: c,
	}
}

func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindText, r.renderText)

	reg.Register(ast.KindHeading, r.renderHeading)
	reg.Register(ast.KindEmphasis, r.emphasis)
	reg.Register(ast.KindLink, r.renderLink)
}

func (r *Renderer) renderText(w util.BufWriter, source []byte, node ast.Node, entering bool) (
	ast.WalkStatus, error,
) {
	if !entering {
		return ast.WalkContinue, nil
	}
	n := node.(*ast.Text)
	render(w, n.Segment.Value(source))
	return ast.WalkContinue, nil
}

func (r *Renderer) renderHeading(w util.BufWriter, _ []byte, node ast.Node, entering bool) (
	ast.WalkStatus, error,
) {
	n := node.(*ast.Heading)
	if entering {
		r.Config.headings[n.Level].writeStart(w)
	} else {
		r.Config.headings[n.Level].writeEnd(w)
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderLink(w util.BufWriter, source []byte, node ast.Node, entering bool) (
	ast.WalkStatus, error,
) {
	n := node.(*ast.Link)
	if entering {
		writeRowBytes(w, []byte{OpenBracket.Byte()})
	} else {
		writeRowBytes(w, []byte{CloseBracket.Byte(), OpenParen.Byte()})
		writeRowBytes(w, n.Destination)
		writeRowBytes(w, []byte{CloseParen.Byte()})
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) emphasis(w util.BufWriter, source []byte, node ast.Node, entering bool) (
	ast.WalkStatus, error,
) {
	n := node.(*ast.Emphasis)
	if n.Level == 2 {
		writeRowBytes(w, Bold.Bytes())
	}
	if n.Level == 1 {
		writeRowBytes(w, Italics.Bytes())
	}
	return ast.WalkContinue, nil
}
