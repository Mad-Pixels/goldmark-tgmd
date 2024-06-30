package tgmd

import (
	"github.com/yuin/goldmark"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var KindDoubleSpace = gast.NewNodeKind("DoubleSpace")

var uncloseCounterKey = parser.NewContextKey()

type unclosedCounter struct {
	Single int
	Double int
}

func (u *unclosedCounter) Reset() {
	u.Single = 0
	u.Double = 0
}

func getUnclosedCounter(pc parser.Context) *unclosedCounter {
	v := pc.Get(uncloseCounterKey)
	if v == nil {
		v = &unclosedCounter{}
		pc.Set(uncloseCounterKey, v)
	}
	return v.(*unclosedCounter)
}

type doubleSpaceParser struct{}

var defaultDoubleSpaceParser = &doubleSpaceParser{}

// NewDoubleSpaceParser ...
func NewDoubleSpaceParser() parser.InlineParser {
	return defaultDoubleSpaceParser
}

func (s *doubleSpaceParser) Trigger() []byte {
	return []byte{SpaceChar.Byte()}
}

func (s *doubleSpaceParser) Parse(_ gast.Node, block text.Reader, pc parser.Context) gast.Node {
	line, _ := block.PeekLine()
	c := line[0]
	if len(line) > 1 {
		if c == ' ' {
			if line[1] == ' ' { // '  '
				node := gast.NewString([]byte{NewLineChar.Byte()})
				node.SetCode(true)
				block.Advance(2)
				return node
			}
			return nil
		}
	}

	return nil
}

func (s *doubleSpaceParser) CloseBlock(_ gast.Node, pc parser.Context) {
	getUnclosedCounter(pc).Reset()
}

type doubleSpace struct{}

// DoubleSpace ...
var DoubleSpace = &doubleSpace{}

func (e *doubleSpace) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(NewDoubleSpaceParser(), 500),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewRenderer(), 500),
	))
}
