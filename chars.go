package tgmd

// SpecialRune define custom rune object.
type SpecialRune rune

// Rune from SpecialRune.
func (sr SpecialRune) Rune() rune {
	return rune(sr)
}

// SpecialChar define custom byte object.
type SpecialChar byte

// Byte from SpecialChar.
func (sc SpecialChar) Byte() byte {
	return byte(sc)
}

// Escaped return SpecialChar as escaped byte char.
func (sc SpecialChar) Escaped() []byte {
	return append([]byte{SlashChar.Byte()}, sc.Byte())
}

// SpecialTag define Markdown formatting characters.
type SpecialTag [3]SpecialChar

// Bytes from SpecialTags.
func (st SpecialTag) Bytes() []byte {
	return []byte{st[0].Byte(), st[1].Byte(), st[2].Byte()}
}

// define characters.
const (
	UnderscoreChar   SpecialChar = '_'
	AsteriskChar     SpecialChar = '*'
	OpenBracketChar  SpecialChar = '['
	CloseBracketChar SpecialChar = ']'
	OpenParenChar    SpecialChar = '('
	CloseParenChar   SpecialChar = ')'
	OpenBraceChar    SpecialChar = '{'
	CloseBraceChar   SpecialChar = '}'
	HashChar         SpecialChar = '#'
	PlusChar         SpecialChar = '+'
	MinusChar        SpecialChar = '-'
	EqualChar        SpecialChar = '='
	DotChar          SpecialChar = '.'
	TildeChar        SpecialChar = '~'
	PipeChar         SpecialChar = '|'
	ExclamationChar  SpecialChar = '!'
	GreaterThanChar  SpecialChar = '>'
	LessThanChar     SpecialChar = '<'
	BackqouteChar    SpecialChar = '`'
	SpaceChar        SpecialChar = ' '
	NewLineChar      SpecialChar = '\n'
	SlashChar        SpecialChar = '\\'
	TabChar          SpecialChar = '\t'
)

// define symbols.
const (
	CircleSymbol   SpecialRune = '•'
	TriangleSymbol SpecialRune = '‣'
	SquareSymbol   SpecialRune = '▪'
)

// define Telegram Markdown formatting tags.
var (
	BoldTg          SpecialTag = [3]SpecialChar{AsteriskChar, AsteriskChar, AsteriskChar}
	StrikethroughTg SpecialTag = [3]SpecialChar{TildeChar, TildeChar, TildeChar}
	UnderlineTg     SpecialTag = [3]SpecialChar{UnderscoreChar, UnderscoreChar}
	HiddenTg        SpecialTag = [3]SpecialChar{PipeChar, PipeChar}
	ItalicsTg       SpecialTag = [3]SpecialChar{UnderscoreChar}
	CodeTg          SpecialTag = [3]SpecialChar{BackqouteChar, BackqouteChar, BackqouteChar}
	SpanTg          SpecialTag = [3]SpecialChar{BackqouteChar}
)

// define escape map.
var escape = map[byte][]byte{
	UnderscoreChar.Byte():   UnderscoreChar.Escaped(),
	AsteriskChar.Byte():     AsteriskChar.Escaped(),
	OpenBracketChar.Byte():  OpenBracketChar.Escaped(),
	CloseBracketChar.Byte(): CloseBracketChar.Escaped(),
	OpenParenChar.Byte():    OpenParenChar.Escaped(),
	CloseParenChar.Byte():   CloseParenChar.Escaped(),
	OpenBraceChar.Byte():    OpenBraceChar.Escaped(),
	CloseBraceChar.Byte():   CloseBraceChar.Escaped(),
	HashChar.Byte():         HashChar.Escaped(),
	PlusChar.Byte():         PlusChar.Escaped(),
	MinusChar.Byte():        MinusChar.Escaped(),
	EqualChar.Byte():        EqualChar.Escaped(),
	DotChar.Byte():          DotChar.Escaped(),
	ExclamationChar.Byte():  ExclamationChar.Escaped(),
	GreaterThanChar.Byte():  GreaterThanChar.Escaped(),
	LessThanChar.Byte():     LessThanChar.Escaped(),
	TildeChar.Byte():        TildeChar.Escaped(),
	PipeChar.Byte():         PipeChar.Escaped(),
	BackqouteChar.Byte():    BackqouteChar.Escaped(),
}
