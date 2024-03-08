package tgmd

import "github.com/yuin/goldmark/util"

type config struct {
	headings [6]Element
}

// NewConfig for generate Tg styles.
func NewConfig() *config {
	return &config{
		headings: [6]Element{
			{
				Style:  Bold,
				Prefix: "# ",
			},
			{
				Style:  Bold,
				Prefix: "",
			},
			{
				Style:  Italics,
				Prefix: "# ",
			},
			{
				Style:  Italics,
				Prefix: "",
			},
			{
				Style:  Italics,
				Prefix: "~",
			},
			{
				Style:  Italics,
				Prefix: "",
			},
		},
	}
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
	writeNewLine(w)
}
