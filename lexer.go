package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type stateFn func(*lexer) stateFn

type lexer struct {
	input string
	start int
	pos   int
	width int
	state stateFn
	items chan item
}

const eof = -1
const (
	symbols    = "+-*/%<>=!&|^`~[]{}.,;:?()'\""
	digits     = "0123456789"
	whitespace = " \t\r\n"
)

/* -- Detector functions for character groups -- */

// Detects if a rune is a symbol
func isSymbol(r rune) bool {
	return strings.IndexRune(symbols, r) >= 0
}

// Detects if a rune is a digit
func isDigit(r rune) bool {
	return strings.IndexRune(digits, r) >= 0
}

// Detects if a rune is whitespace
func isWhiteSpace(r rune) bool {
	return strings.IndexRune(whitespace, r) >= 0
}

// Detects is a rune can be part of an identifier
func isIdentRune(r rune) bool {
	return !isWhiteSpace(r) && !isSymbol(r)
}

func lex(input string) *lexer {
	l := &lexer{
		input: input,
		items: make(chan item),
	}

	go l.run()

	return l
}

func (l *lexer) run() {
	for l.state = lexStart; l.state != nil; {
		l.state = l.state(l)
	}
	close(l.items)
}

// next returns the next rune in the input or returns EOF
func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	if r == utf8.RuneError {
		l.errorf("invalid UTF-8 encoding, ALL UNTITLED programs are written in UTF-8")
		return eof
	}
	l.pos += l.width
	return r
}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{itemError, fmt.Sprintf(format, args...)}
	return nil
}

func (l *lexer) warnf(format string, args ...interface{}) stateFn {
	l.items <- item{itemWarning, fmt.Sprintf(format, args...)}
	return nil
}

// emit passes an item back to the client
func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

// ignore skips over the pending input before this point
func (l *lexer) ignore() {
	l.start = l.pos
}

// backup steps back one rune. (only can be used once per call of next)
func (l *lexer) backup() {
	l.pos -= l.width
}

// peek returns but does not consume the next rune in the input
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// accept consumes the next rune if it's from the valid set
func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes (including zero runs) from the valid set
func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

func (l *lexer) acceptUntil(runes string) {
	for strings.IndexRune(runes, l.next()) < 0 {
	}
	l.backup()
}

// acceptFunc consumes the next rune it matches the predicate
func (l *lexer) acceptFunc(f func(r rune) bool) bool {
	if f(l.next()) {
		return true
	}
	l.backup()
	return false
}

// acceptRunFunc consumes a run of runes (including zero runs) as long as the predicate matches
func (l *lexer) acceptFuncRun(f func(r rune) bool) {
	for f(l.next()) {
	}
	l.backup()
}

/* -- state functions -- */

// lex is the entry point for lexer and starts the state machine
func lexStart(l *lexer) stateFn {
	r := l.peek()

	if isWhiteSpace(r) {
		return lexWhiteSpace
	} else if isDigit(r) {
		return lexNumber
	} else if isSymbol(r) {
		return lexOp
	} else if r == eof {
		l.emit(itemEOF)
		return nil
	} else {
		return lexText
	}
}

// lexWhitespace consumes all contiguous whitespace and acts as the start state
func lexWhiteSpace(l *lexer) stateFn {
	l.acceptRun(whitespace)
	l.ignore()
	return lexStart
}

// lexNumber consumes a number it's not up to the lexer to validate the number
func lexNumber(l *lexer) stateFn {
	digits := "0123456789"
	// is it a hex number?
	if l.accept("0") && l.accept("xX") {
		digits = "0123456789abcdefABCDEF"
	}
	l.acceptRun(digits)

	// we do not accept floating point numbers
	if l.peek() == '.' {
		// be nice and finish reading the number after the decimal point
		l.acceptRun(digits)
		return l.errorf("floating points not supported: %q", l.input[l.start:l.pos])
	}

	// next character cannot be a letter
	if unicode.IsLetter(l.peek()) {
		return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
	}

	l.emit(itemNumber)
	return lexStart
}

// lexOp parses looking for items from the "Math operators", and "Logical operators" sections
// can enter "brainfuck" mode
func lexOp(l *lexer) stateFn {
	switch l.next() {
	case '+':
		// check += and ++
		if l.accept("=") {
			l.emit(itemPlusEqual)
		} else if l.accept("+") {
			l.emit(itemPlusPlus)
		} else {
			l.emit(itemPlus)
		}
	case '-':
		// check -= and --
		if l.accept("=") {
			l.emit(itemMinusEqual)
		} else if l.accept("-") {
			l.emit(itemMinusMinus)
		} else {
			l.emit(itemMinus)
		}
	case '*':
		// check *=
		if l.accept("=") {
			l.emit(itemMultEqual)
		} else {
			l.emit(itemMult)
		}
	case '=':
		// check ==
		if l.accept("=") {
			l.emit(itemEqual)
		} else {
			l.emit(itemAssign)
		}
	case '!':
		// check !=
		if l.accept("=") {
			l.emit(itemNotEqual)
		} else {
			l.emit(itemNot)
		}
	case '<':
		// check <=
		if l.accept("=") {
			l.emit(itemLessEqual)
		} else {
			l.emit(itemLess)
		}
	case '>':
		// check >=
		if l.accept("=") {
			l.emit(itemGreaterEqual)
		} else {
			l.emit(itemGreater)
		}
	case '&':
		if l.accept("&") {
			l.emit(itemAnd)
		} else {
			l.emit(itemCopy)
		}
	case '|':
		// must be ||
		if l.accept("|") {
			l.emit(itemOr)
		} else {
			l.errorf("invalid operator: %q", l.input[l.start:l.pos])
		}
	case '(':
		l.emit(itemLeftParen)
	case ')':
		l.emit(itemRightParen)
	case '{':
		l.emit(itemLeftBrace)
	case '}':
		l.emit(itemRightBrace)
	case ';':
		l.emit(itemSemiColon)
	case ',':
		l.emit(itemComma)
	case '.':
		l.emit(itemDot)
	case '\'':
		r := l.next()
		if r == eof {
			return l.errorf("unterminated character constant")
		}
		if r == '\'' {
			return l.errorf("empty character constant")
		}
		if l.accept("'") {
			l.emit(itemChar)
		} else {
			l.errorf("invalid character constant: %q", l.input[l.start:l.pos])
		}
	case '`':
		// require two more "`"s to enter brainfuck mode
		if l.accept("`") && l.accept("`") {
			l.emit(itemBF)
			return lexBrainfuck
		} else {
			l.errorf("failed to enter brainfuck mode: %q.\nopen a BF block with ```", l.input[l.start:l.pos])
		}
	default:
		l.errorf("invalid operator: %q", l.input[l.start:l.pos])
	}

	return lexStart
}

// lexText consumes all text until a non-text character is found
func lexText(l *lexer) stateFn {
	l.acceptFuncRun(isIdentRune)

	// check if the text is a keyword
	text := l.input[l.start:l.pos]
	switch text {
	case "if":
		l.emit(itemIf)
	case "while":
		l.emit(itemWhile)
	case "function":
		l.emit(itemFunction)
	case "return":
		l.emit(itemReturn)
	case "break":
		l.emit(itemBreak)
	case "continue":
		l.emit(itemContinue)
	case "const":
		l.emit(itemConst)
	case "type":
		l.emit(itemTyp)
	default:
		// if the text is not a keyword, its and identifier
		l.emit(itemIdentifier)
	}

	return lexStart
}

// lexBrainfuck is for inline brainfuck
func lexBrainfuck(l *lexer) stateFn {
	r := l.next()

	switch r {
	case '+':
		l.emit(itemInc)
	case '-':
		l.emit(itemDec)
	case '>':
		l.emit(itemRight)
	case '<':
		l.emit(itemLeft)
	case '.':
		l.emit(itemOut)
	case ',':
		l.emit(itemIn)
	case '[':
		l.emit(itemLeftBracket)
	case ']':
		l.emit(itemRightBracket)
	case '(':
		return lexBFIndentifier
	case ')':
		return l.errorf("unmatched ')'")
	case '`':
		// require two more "`"s to exit brainfuck mode
		if l.accept("`") && l.accept("`") {
			return lexStart
		} else {
			return l.errorf("failed to exit brainfuck mode: %q.\nclose a BF block with ```", l.input[l.start:l.pos])
		}
	case eof:
		return l.errorf("unmatched \"```\"")
	default:
		// everything else in brainfuck is a comment, do nothing
	}

	return lexBrainfuck
}

// lexBFIndentifier consumes a brainfuck identifier
// \w*[id]+\w*)
func lexBFIndentifier(l *lexer) stateFn {
	// Read whitespace
	l.acceptRun(whitespace)

	// ignore opening parenthesis and whitespace
	l.ignore()

	// Read AT LEAST one character for the identifier
	if !l.acceptFunc(isIdentRune) {
		return l.errorf("invalid identifier %q", l.input[l.start:l.pos])
	}
	l.acceptFuncRun(isIdentRune)

	// Emit the identifier right now
	l.emit(itemIdentifier)

	// Make sure we have a closing parenthesis
	l.acceptRun(whitespace)
	if !l.accept(")") {
		return l.errorf("expected ) to close identifier")
	}

	return lexBrainfuck
}
