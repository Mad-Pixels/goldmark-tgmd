package tgmd

import (
	"github.com/yuin/goldmark"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type HiddenAst struct {
	gast.BaseInline
}

// Dump implements Node.Dump.
func (n *HiddenAst) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

var KindHidden = gast.NewNodeKind("Hidden")

// Kind implements Node.Kind.
func (n *HiddenAst) Kind() gast.NodeKind {
	return KindHidden
}

func NewHidden() *HiddenAst {
	return &HiddenAst{}
}

// aaa
// aaa
// aaa
// aaa
// aaa
type hiddenDelimiterProcessor struct {
}

func (p *hiddenDelimiterProcessor) IsDelimiter(b byte) bool {
	return b == '|'
}

func (p *hiddenDelimiterProcessor) CanOpenCloser(opener, closer *parser.Delimiter) bool {
	return opener.Char == closer.Char
}

func (p *hiddenDelimiterProcessor) OnMatch(consumes int) gast.Node {
	return NewHidden()
}

var defaultHiddenDelimiterProcessor = &hiddenDelimiterProcessor{}

type hiddenParser struct {
}

var defaultHiddenParser = &hiddenParser{}

func NewHiddenParser() parser.InlineParser {
	return defaultHiddenParser
}

func (s *hiddenParser) Trigger() []byte {
	return []byte{'|'}
}

func (s *hiddenParser) Parse(parent gast.Node, block text.Reader, pc parser.Context) gast.Node {
	before := block.PrecendingCharacter()
	line, segment := block.PeekLine()
	node := parser.ScanDelimiter(line, before, 2, defaultHiddenDelimiterProcessor)
	if node == nil {
		return nil
	}
	node.Segment = segment.WithStop(segment.Start + node.OriginalLength)
	block.Advance(node.OriginalLength)
	pc.PushDelimiter(node)
	return node
}

func (s *hiddenParser) CloseBlock(parent gast.Node, pc parser.Context) {
	// nothing to do
}

type hidden struct {
}

var Hidden = &hidden{}

func (e *hidden) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(NewHiddenParser(), 500),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewRenderer(nil), 500),
	))
}
