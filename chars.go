package tgmd

// SpecialRune ...
type SpecialRune rune

// Rune ...
func (sr SpecialRune) Rune() rune {
	return rune(sr)
}

// SpecialChar ...
type SpecialChar byte

// Byte ...
func (sc SpecialChar) Byte() byte {
	return byte(sc)
}

// Escaped ...
func (sc SpecialChar) Escaped() []byte {
	return append([]byte("\\"), sc.Byte())
}

// SpecialTag ...
type SpecialTag [3]SpecialChar

// Bytes ...
func (st SpecialTag) Bytes() []byte {
	return []byte{st[0].Byte(), st[1].Byte(), st[2].Byte()}
}

const (
	Underscore   SpecialChar = '_'
	Asterisk     SpecialChar = '*'
	OpenBracket  SpecialChar = '['
	CloseBracket SpecialChar = ']'
	OpenParen    SpecialChar = '('
	CloseParen   SpecialChar = ')'
	OpenBrace    SpecialChar = '{'
	CloseBrace   SpecialChar = '}'
	Hash         SpecialChar = '#'
	Plus         SpecialChar = '+'
	Minus        SpecialChar = '-'
	Equal        SpecialChar = '='
	Dot          SpecialChar = '.'
	Tilde        SpecialChar = '~'
	Pipe         SpecialChar = '|'
	Exclamation  SpecialChar = '!'
	GreaterThan  SpecialChar = '>'
	LessThan     SpecialChar = '<'
	Backqoute    SpecialChar = '`'
	Space        SpecialChar = ' '
	NewLine      SpecialChar = '\n'
)

const (
	SymbolCircle   SpecialRune = '•'
	SymbolTriangle SpecialRune = '‣'
	SymbolSquare   SpecialRune = '▪'
)

var (
	Bold          SpecialTag = [3]SpecialChar{Asterisk, Asterisk, Asterisk}
	Strikethrough SpecialTag = [3]SpecialChar{Tilde, Tilde, Tilde}
	Underline     SpecialTag = [3]SpecialChar{Underscore, Underscore}
	Mono          SpecialTag = [3]SpecialChar{Backqoute, Backqoute}
	Hidden        SpecialTag = [3]SpecialChar{Pipe, Pipe}
	Italics       SpecialTag = [3]SpecialChar{Underscore}
)

var escape = map[byte][]byte{
	Underscore.Byte():   Underscore.Escaped(),
	Asterisk.Byte():     Asterisk.Escaped(),
	OpenBracket.Byte():  OpenBracket.Escaped(),
	CloseBracket.Byte(): CloseBracket.Escaped(),
	OpenParen.Byte():    OpenParen.Escaped(),
	CloseParen.Byte():   CloseParen.Escaped(),
	OpenBrace.Byte():    OpenBrace.Escaped(),
	CloseBrace.Byte():   CloseBrace.Escaped(),
	Hash.Byte():         Hash.Escaped(),
	Plus.Byte():         Plus.Escaped(),
	Minus.Byte():        Minus.Escaped(),
	Equal.Byte():        Equal.Escaped(),
	Dot.Byte():          Dot.Escaped(),
	Exclamation.Byte():  Exclamation.Escaped(),
	GreaterThan.Byte():  GreaterThan.Escaped(),
	LessThan.Byte():     LessThan.Escaped(),
	NewLine.Byte():      NewLine.Escaped(),
	//Tilde.Byte():        Tilde.Escaped(),
	Backqoute.Byte(): Backqoute.Escaped(),
}
