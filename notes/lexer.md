# Lexical Analysis
Source Code ---> Tokens ---> Abstract Syntax Tree
             |
      Lexical Analysis


For example,
```
let x = 5 + 5;
```
 |
 V
```
[
  LET,
  IDENTIFIER("x"),
  EQUAL_SIGN,
  INTEGER(5),
  PLUS_SIGN,
  INTEGER(5),
  SEMICOLON
]
```




## Our Language
```
let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
}

let result = add(five, ten);
```



