package main

import "fmt"

type itemType int

const (
	itemError   itemType = iota // error occurred, val is the error message
	itemWarning                 // warning occurred, val is the warning message
	itemEOF                     // end of file

	itemSemiColon  // ;
	itemIdentifier // identifier
	itemNumber     // number (could be hex or decimal)
	itemChar       // char (single character)

	itemWhiteSpace // whitespace
	itemNewLine    // \n or ;
	itemLeftBrace  // {
	itemRightBrace // }
	itemComma      // ,
	itemDot        // .
	itemCopy       // &

	// Keywords
	itemIf       // "if"
	itemWhile    // "while"
	itemFunction // "func"
	itemReturn   // "return"
	itemBreak    // "break"
	itemContinue // "continue"
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

func (t itemType) String() string {
	switch t {
	case itemError:
		return fmt.Sprintf("ERROR")
	case itemWarning:
		return fmt.Sprintf("WARNING")
	case itemEOF:
		return "EOF"
	case itemSemiColon:
		return "SEMICOLON"
	case itemIdentifier:
		return fmt.Sprintf("IDENTIFIER")
	case itemNumber:
		return fmt.Sprintf("NUMBER")
	case itemChar:
		return fmt.Sprintf("CHAR")
	case itemWhiteSpace:
		return fmt.Sprintf("WHITE SPACE")
	case itemNewLine:
		return fmt.Sprintf("NEW LINE")
	case itemLeftBrace:
		return fmt.Sprintf("LEFT BRACE")
	case itemRightBrace:
		return fmt.Sprintf("RIGHT BRACE")
	case itemComma:
		return fmt.Sprintf("COMMA")
	case itemDot:
		return fmt.Sprintf("DOT")
	case itemCopy:
		return fmt.Sprintf("COPY")
	case itemIf:
		return fmt.Sprintf("IF")
	case itemWhile:
		return fmt.Sprintf("WHILE")
	case itemFunction:
		return fmt.Sprintf("FUNCTION")
	case itemReturn:
		return fmt.Sprintf("RETURN")
	case itemBreak:
		return fmt.Sprintf("BREAK")
	case itemContinue:
		return fmt.Sprintf("CONTINUE")
	case itemTyp:
		return fmt.Sprintf("TYPE")
	case itemBF:
		return fmt.Sprintf("BRAINFUCK")
	case itemInc:
		return fmt.Sprintf("INCREMENT")
	case itemDec:
		return fmt.Sprintf("DECREMENT")
	case itemOut:
		return fmt.Sprintf("OUTPUT")
	case itemIn:
		return fmt.Sprintf("INPUT")
	case itemLeftBracket:
		return fmt.Sprintf("LEFT BRACKET")
	case itemRightBracket:
		return fmt.Sprintf("RIGHT BRACKET")
	case itemLeft:
		return fmt.Sprintf("LEFT")
	case itemRight:
		return fmt.Sprintf("RIGHT")
	case itemPlus:
		return fmt.Sprintf("PLUS")
	case itemMinus:
		return fmt.Sprintf("MINUS")
	case itemAssign:
		return fmt.Sprintf("ASSIGN")
	case itemMult:
		return fmt.Sprintf("MULT")
	case itemPlusEqual:
		return fmt.Sprintf("PLUS EQUAL")
	case itemMinusEqual:
		return fmt.Sprintf("MINUS EQUAL")
	case itemMultEqual:
		return fmt.Sprintf("MULT EQUAL")
	case itemPlusPlus:
		return fmt.Sprintf("PLUS PLUS")
	case itemMinusMinus:
		return fmt.Sprintf("MINUS MINUS")
	case itemLeftParen:
		return fmt.Sprintf("LEFT PAREN")
	case itemRightParen:
		return fmt.Sprintf("RIGHT PAREN")
	case itemAnd:
		return fmt.Sprintf("AND")
	case itemOr:
		return fmt.Sprintf("OR")
	case itemNot:
		return fmt.Sprintf("NOT")
	case itemEqual:
		return fmt.Sprintf("EQUAL")
	case itemNotEqual:
		return fmt.Sprintf("NOT EQUAL")
	case itemGreater:
		return fmt.Sprintf("GREATER")
	case itemLess:
		return fmt.Sprintf("LESS")
	case itemGreaterEqual:
		return fmt.Sprintf("GREATER EQUAL")
	case itemLessEqual:
		return fmt.Sprintf("LESS EQUAL")
	default:
		return fmt.Sprintf("UNKNOWN")
	}
}

func (i item) String() string {
	switch i.typ {
	case itemError:
		return fmt.Sprintf("error: %s", i.val)
	case itemWarning:
		return fmt.Sprintf("warning: %s", i.val)
	case itemEOF:
		return "EOF"
	case itemSemiColon:
		return ";"
	case itemIdentifier:
		return fmt.Sprintf("[id: %s]", i.val)
	case itemNumber:
		return fmt.Sprintf("[num: %s]", i.val)
	case itemChar:
		return fmt.Sprintf("[char: %s]", i.val)
	case itemWhiteSpace:
		return fmt.Sprintf("%s", i.val)
	case itemNewLine:
		return "\n"
	case itemLeftBrace:
		return "{"
	case itemRightBrace:
		return "}"
	case itemComma:
		return "[,]"
	case itemDot:
		return "[.]"
	case itemCopy:
		return "[&]"
	case itemIf:
		return "[if]"
	case itemWhile:
		return "[while]"
	case itemFunction:
		return "[func]"
	case itemReturn:
		return "[return]"
	case itemBreak:
		return "[break]"
	case itemContinue:
		return "[continue]"
	case itemTyp:
		return "[type]"
	case itemBF:
		return "[```]"
	case itemInc:
		return "[+]"
	case itemDec:
		return "[-]"
	case itemOut:
		return "[.]"
	case itemIn:
		return "[,]"
	case itemLeftBracket:
		return "[open]"
	case itemRightBracket:
		return "[close]"
	case itemLeft:
		return "[<]"
	case itemRight:
		return "[>]"
	case itemPlus:
		return "[+]"
	case itemMinus:
		return "[-]"
	case itemAssign:
		return "[=]"
	case itemMult:
		return "[*]"
	case itemPlusEqual:
		return "[+=]"
	case itemMinusEqual:
		return "[-=]"
	case itemMultEqual:
		return "[*=]"
	case itemPlusPlus:
		return "[++]"
	case itemMinusMinus:
		return "[--]"
	case itemLeftParen:
		return "[(]"
	case itemRightParen:
		return "[)]"
	case itemAnd:
		return "[&&]"
	case itemOr:
		return "[||]"
	case itemNot:
		return "[!]"
	case itemEqual:
		return "[==]"
	case itemNotEqual:
		return "[!=]"
	case itemGreater:
		return "[>]"
	case itemLess:
		return "[<]"
	case itemGreaterEqual:
		return "[>=]"
	case itemLessEqual:
		return "[<=]"
	default:
		return fmt.Sprintf("[%d, %s]", i.typ, i.val)
	}
}
