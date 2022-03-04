package main

type itemType int

const (
	itemError   itemType = iota // error occurred, val is the error message
	itemWarning                 // warning occurred, val is the warning message
	itemEOF                     // end of file

	itemSemiColon  // ;
	itemIdentifier // identifier
	itemNumber     // number (could be hex or decimal)
	itemChar       // char (single character)

	itemLeftBrace  // {
	itemRightBrace // }
	itemComma      // ,
	itemDot        // .
	itemCopy       // &

	// Keywords
	itemIf       // "if"
	itemWhile    // "while"
	itemFunction // "function"
	itemReturn   // "return"
	itemBreak    // "break"
	itemContinue // "continue"
	itemConst    // "const"
	itemTyp      // "type"

	// Brainfuck operators
	itemBF           // surround brainfuck code with ```
	itemInc          // increment +
	itemDec          // decrement -
	itemOut          // output .
	itemIn           // input ,
	itemLeftBracket  // left bracket [
	itemRightBracket // right bracket ]
	itemLeft         // left move <
	itemRight        // right move >

	// Operators
	itemPlus       // +
	itemMinus      // -
	itemAssign     // =
	itemMult       // *
	itemPlusEqual  // +=
	itemMinusEqual // -=
	itemMultEqual  // *=
	itemPlusPlus   // ++
	itemMinusMinus // --
	itemLeftParen  // (
	itemRightParen // )

	// Logic operators
	itemAnd          // &&
	itemOr           // ||
	itemNot          // !
	itemEqual        // ==
	itemNotEqual     // !=
	itemGreater      // >
	itemLess         // <
	itemGreaterEqual // >=
	itemLessEqual    // <=
)

type item struct {
	typ itemType
	val string
}

func (i item) String() string {
	switch i.typ {
	case itemError:
		return "error: " + i.val
	case itemEOF:
		return "EOF"
	}

	// long strings will be truncated
	if len(i.val) > 10 {
		return "[" + i.val[:10] + "..." + "]"
	}

	return "[" + i.val + "]"
}
