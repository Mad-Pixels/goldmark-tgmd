package tgmd

import (
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	ext "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/renderer"
	textm "github.com/yuin/goldmark/text"
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

func TGMD(c *config) goldmark.Markdown {
	return goldmark.New(
		goldmark.WithRenderer(
			renderer.NewRenderer(
				renderer.WithNodeRenderers(util.Prioritized(NewRenderer(c), 1000)),
			),
		),
		goldmark.WithExtensions(Strikethroughs),
		goldmark.WithExtensions(Hidden),
	)
}

func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindText, r.renderText)

	reg.Register(ast.KindFencedCodeBlock, r.code)
	reg.Register(ast.KindCodeBlock, r.code)
	reg.Register(ast.KindBlockquote, r.blockquote)
	reg.Register(ast.KindHeading, r.heading)
	reg.Register(ast.KindListItem, r.listItem)
	reg.Register(ast.KindEmphasis, r.emphasis)
	reg.Register(ast.KindLink, r.renderLink)
	reg.Register(ast.KindList, r.list)
	reg.Register(ast.KindCodeSpan, r.codeSpan)

	// re-define.
	reg.Register(ext.KindStrikethrough, r.strikethrough)
	reg.Register(KindHidden, r.hidden)
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

func (r *Renderer) heading(w util.BufWriter, _ []byte, node ast.Node, entering bool) (
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

func (r *Renderer) renderLink(w util.BufWriter, _ []byte, node ast.Node, entering bool) (
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

func (r *Renderer) emphasis(w util.BufWriter, _ []byte, node ast.Node, entering bool) (
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

func (r *Renderer) list(w util.BufWriter, _ []byte, node ast.Node, entering bool) (
	ast.WalkStatus, error,
) {
	n := node.(*ast.List)
	if !entering {
		if n.Parent().Kind().String() == ast.KindDocument.String() {
			writeNewLine(w)
		}
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) listItem(w util.BufWriter, _ []byte, node ast.Node, entering bool) (
	ast.WalkStatus, error,
) {
	n := node.(*ast.ListItem)
	if entering {
		writeNewLine(w)
		if n.Parent().Parent().Kind().String() == ast.KindDocument.String() {
			writeRowBytes(w, []byte{Space.Byte(), Space.Byte()})
			writeRune(w, r.Config.listBullets[0])
		} else {
			if n.Parent().Parent().Parent().Parent() != nil {
				if n.Parent().Parent().Parent().Parent().Kind().String() == ast.KindListItem.String() {
					writeRowBytes(w, []byte{Space.Byte(), Space.Byte(), Space.Byte(), Space.Byte(), Space.Byte(), Space.Byte()})
					writeRune(w, r.Config.listBullets[2])
				} else {
					writeRowBytes(w, []byte{Space.Byte(), Space.Byte(), Space.Byte(), Space.Byte()})
					writeRune(w, r.Config.listBullets[1])
				}
			}
		}
		writeRowBytes(w, []byte{Space.Byte()})
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) blockquote(w util.BufWriter, _ []byte, node ast.Node, entering bool) (
	ast.WalkStatus, error,
) {
	writeNewLine(w)
	n := node.(*ast.Blockquote)
	if entering {
		writeRowBytes(w, []byte{GreaterThan.Byte(), Space.Byte()})
	} else {
		if n.Parent().Kind().String() == ast.KindDocument.String() {
			writeNewLine(w)
		}
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) strikethrough(w util.BufWriter, _ []byte, node ast.Node, entering bool) (
	ast.WalkStatus, error,
) {
	writeWrapperArr(w.Write(Strikethrough.Bytes()))
	return ast.WalkContinue, nil
}

func (r *Renderer) hidden(w util.BufWriter, _ []byte, node ast.Node, entering bool) (
	ast.WalkStatus, error,
) {
	writeWrapperArr(w.Write(Hiddend.Bytes()))
	return ast.WalkContinue, nil
}

func (r *Renderer) codeSpan(w util.BufWriter, _ []byte, node ast.Node, entering bool) (
	ast.WalkStatus, error,
) {
	writeWrapperArr(w.Write(Form.Bytes()))
	return ast.WalkContinue, nil
}

func (r *Renderer) code(w util.BufWriter, source []byte, node ast.Node, entering bool) (
	ast.WalkStatus, error,
) {
	n := node.(interface {
		Lines() *textm.Segments
	})
	var content string
	l := n.Lines().Len()
	for i := 0; i < l; i++ {
		line := n.Lines().At(i)
		content += string(line.Value(source))
	}
	content = strings.ReplaceAll(content, "\t", "    ")

	nn := node.(*ast.FencedCodeBlock)
	if entering {
		writeNewLine(w)
		writeWrapperArr(w.Write(Code.Bytes()))
		writeWrapperArr(w.Write(nn.Language(source)))
	} else {
		w.WriteByte('\n')
		w.WriteString(content)
		writeWrapperArr(w.Write(Code.Bytes()))
	}
	return ast.WalkContinue, nil
}
