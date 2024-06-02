package tgmd

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type doubleSpace struct{}

var DoubleSpace = &doubleSpace{}

var KindDoubleSpace = ast.NewNodeKind("DoubleSpace")

// Extend ...
func (e *doubleSpace) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(NewDoubleSpaceParser(), 500),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewRenderer(), 500),
	))
}

// DoubleSpaceAST abstract semantic tree for "doubleSpace".
type DoubleSpaceAST struct {
	ast.BaseInline
}

// Dump implements Node.Dump.
func (n *DoubleSpaceAST) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

// Kind implements Node.Kind.
func (n *DoubleSpaceAST) Kind() ast.NodeKind {
	return KindHidden
}

// NewDoubleSpace initialize DoubleSpaceAST object.
func NewDoubleSpace() *DoubleSpaceAST {
	return &DoubleSpaceAST{}
}

type doubleSpaceDelimiterProcessor struct{}

// IsDelimiter check incoming byte with object delimiter.
func (p *doubleSpaceDelimiterProcessor) IsDelimiter(b byte) bool {
	return b == SpaceChar.Byte()
}

// CanOpenCloser ...
func (p *doubleSpaceDelimiterProcessor) CanOpenCloser(opener, closer *parser.Delimiter) bool {
	return opener.Char == closer.Char
}

// OnMatch ...
func (p *doubleSpaceDelimiterProcessor) OnMatch(_ int) ast.Node {
	return NewDoubleSpace()
}

var defaultDoubleSpaceDelimiterProcessor = &doubleSpaceDelimiterProcessor{}

type doubleSpaceParser struct{}

var defaultDoubleSpaceParser = &doubleSpaceParser{}

// NewDoubleSpaceParser initialize parser.InlineParser.
func NewDoubleSpaceParser() parser.InlineParser {
	return defaultDoubleSpaceParser
}

// Trigger char for parser.
func (s *doubleSpaceParser) Trigger() []byte {
	return []byte{PipeChar.Byte()}
}

// Parse source.
func (s *doubleSpaceParser) Parse(_ ast.Node, block text.Reader, pc parser.Context) ast.Node {
	before := block.PrecendingCharacter()
	line, segment := block.PeekLine()
	node := parser.ScanDelimiter(line, before, 2, defaultDoubleSpaceDelimiterProcessor)
	if node == nil {
		return nil
	}
	node.Segment = segment.WithStop(segment.Start + node.OriginalLength)
	block.Advance(node.OriginalLength)
	pc.PushDelimiter(node)
	return node
}

// CloseBlock ...
func (s *doubleSpaceParser) CloseBlock(_ ast.Node, _ parser.Context) {
	// nothing to do
}
