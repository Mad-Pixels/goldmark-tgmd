package tgmd

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var KindHidden = ast.NewNodeKind("Hidden")

// HiddenAST abstract semantic tree for "hidden".
type HiddenAST struct {
	ast.BaseInline
}

// Dump implements Node.Dump.
func (n *HiddenAST) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

// Kind implements Node.Kind.
func (n *HiddenAST) Kind() ast.NodeKind {
	return KindHidden
}

// NewHidden initialize HiddenAST object.
func NewHidden() *HiddenAST {
	return &HiddenAST{}
}

type hiddenDelimiterProcessor struct{}

// IsDelimiter check incoming byte with object delimiter.
func (p *hiddenDelimiterProcessor) IsDelimiter(b byte) bool {
	return b == PipeChar.Byte()
}

// CanOpenCloser ...
func (p *hiddenDelimiterProcessor) CanOpenCloser(opener, closer *parser.Delimiter) bool {
	return opener.Char == closer.Char
}

// OnMatch ...
func (p *hiddenDelimiterProcessor) OnMatch(consumes int) ast.Node {
	return NewHidden()
}

var defaultHiddenDelimiterProcessor = &hiddenDelimiterProcessor{}

type hiddenParser struct{}

var defaultHiddenParser = &hiddenParser{}

// NewHiddenParser initialize parser.InlineParser.
func NewHiddenParser() parser.InlineParser {
	return defaultHiddenParser
}

// Trigger char for parser.
func (s *hiddenParser) Trigger() []byte {
	return []byte{PipeChar.Byte()}
}

// Parse source.
func (s *hiddenParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
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

// CloseBlock ...
func (s *hiddenParser) CloseBlock(parent ast.Node, pc parser.Context) {
	// nothing to do
}

type hidden struct{}

var Hidden = &hidden{}

// Extend ...
func (e *hidden) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(NewHiddenParser(), 500),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewRenderer(), 500),
	))
}
