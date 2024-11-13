package tgmd

import "github.com/yuin/goldmark/util"

var Config = &config{
	headings: [6]Element{
		{
			Style:  BoldTg,
			Prefix: "",
		},
		{
			Style:  BoldTg,
			Prefix: "",
		},
		{
			Style:  BoldTg,
			Prefix: "",
		},
		{
			Style:  ItalicsTg,
			Prefix: "",
		},
		{
			Style:  ItalicsTg,
			Prefix: "",
		},
		{
			Style:  ItalicsTg,
			Prefix: "",
		},
	},
	listBullets: [3]rune{
		CircleSymbol.Rune(),
		SquareSymbol.Rune(),
		TriangleSymbol.Rune(),
	},
}

type config struct {
	headings    [6]Element
	listBullets [3]rune
}

// UpdateHeading1 change default H1 style.
func (c *config) UpdateHeading1(e Element) {
	c.headings[0] = e
}

// UpdateHeading2 change default H2 style.
func (c *config) UpdateHeading2(e Element) {
	c.headings[1] = e
}

// UpdateHeading3 change default H3 style.
func (c *config) UpdateHeading3(e Element) {
	c.headings[2] = e
}

// UpdateHeading4 change default H4 style.
func (c *config) UpdateHeading4(e Element) {
	c.headings[3] = e
}

// UpdateHeading5 change default H5 style.
func (c *config) UpdateHeading5(e Element) {
	c.headings[4] = e
}

// UpdateHeading6 change default H6 style.
func (c *config) UpdateHeading6(e Element) {
	c.headings[5] = e
}

// UpdatePrimaryListBullet change default primary bullet.
func (c *config) UpdatePrimaryListBullet(r rune) {
	c.listBullets[0] = r
}

// UpdateSecondaryListBullet change default primary bullet.
func (c *config) UpdateSecondaryListBullet(r rune) {
	c.listBullets[1] = r
}

// UpdateAdditionalListBullet change default primary bullet.
func (c *config) UpdateAdditionalListBullet(r rune) {
	c.listBullets[2] = r
}

// Element styles object.
type Element struct {
	Style   SpecialTag
	Prefix  string
	Postfix string
}

func (e Element) writeStart(w util.BufWriter) {
	writeSpecialTagStart(w, e.Style, StringToBytes(e.Prefix))
}

func (e Element) writeEnd(w util.BufWriter) {
	writeSpecialTagEnd(w, e.Style, StringToBytes(e.Postfix))
}
