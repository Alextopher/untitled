# Untitled

This is a small toy programming language I'm writing that targets brainfuck. WIP. Written in Go.

## Lexer

I followed rob pike's video on writing a lexer in Go https://www.youtube.com/watch?v=HxaD_trXwRE

## Grammar

```
program -> function | function program
function -> "func" identifier "(" parameters ")" type block
parameters -> parameter | parameter "," parameters | ""
parameter -> identifier type
type -> "byte" | "void"
block -> "{" statements "}"
statements -> statement | statement statements | ""
statement -> declaration | assignment | expression ";" | bf | ";"
declaration -> type identifier ";" | type identifier "=" expression ";"
assignment -> identifier ASSIGNOP expression ";"
expression -> simple_expression | simple_expression RELOP simple_expression
simple_expression -> term | simple_expression ADDOP term
term -> factor | term MULOP factor
factor -> identifier | identifier (expression_list) | number | "(" expression ")"
expression_list -> expression | expression "," expression_list | ""
bf -> "``" bf_body "``"
bf_body -> bf_command | bf_command bf_body
bf_command -> "+" | "-" | "." | "," | "[" | "]" | ">" | "<" | ">" | "(" identifier ")"
```

## List of overloadable operators

```
Add(L, R) : L + R
Sub(L, R) : L - R
Mul(L, R) : L * R
AddEq(L, R) : L += R
SubEq(L, R) : L -= R
MulEq(L, R) : L *= R

Add(L, &R) : L + &R
Sub(L, &R) : L - &R
Mul(L, &R) : L * &R
AddEq(L, &R) : L += &R
SubEq(L, &R) : L -= &R
MulEq(L, &R) : L *= &R

Increment(L) : L++
Decrement(L) : L--
```