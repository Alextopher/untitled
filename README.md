# Untitled

This is a small toy programming language I'm writing that targets brainfuck. WIP. Written in Go.

## Lexer

I followed rob pike's video on writing a lexer in Go https://www.youtube.com/watch?v=HxaD_trXwRE

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