# What is a parser exactly?
- A parser
    - is a software component that takes inputt data and builds a data structure
        - some kind of parse tree
        - abstract syntax tree
        - or other hierarchical structure
    - giving a strutural representation of the input
    - checking for corect syntax in the process
    - is often preceded by a separate lexical analyser, which creates tokens from the sequence of input characters

A parser turns its input into a data structure that represents the input.


In most interpreters and compilers the data structure used for the internal representation of the source code is called a “syntax tree” or an “abstract syntax tree” (AST for short). The “abstract” is based on the fact that certain details visible in the source code are omitted in the AST. Semicolons, newlines, whitespace, comments, braces, bracket and parentheses – depending on the language and the parser these details are not represented in the AST, but merely guide the parser when constructing it.

# Why not a parser generator?
like yacc, bison, ANTLR. 


# 2 strategies for parser.
- top-down
    - recursive descent parsing (recommended for newcomers)
    - early parsing
    - predictive parsing
- bottom-up


# What we create is "Recursive Descent Parser". 
- Top down operator precedence. 
- called "Pratt Parser"

# Our tradeoff
- not fast
- no proof of correctness
- no recovery process


# References
- https://zenn.dev/pandaman64/books/pratt-parsing/viewer/introduction#pratt%E3%83%91%E3%83%BC%E3%82%B5%E3%81%AE%E6%A7%8B%E9%80%A0

- https://scrapbox.io/mrsekut-p/Nim%E3%81%A7%E8%87%AA%E4%BD%9C%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%97%E3%83%AA%E3%82%BF%E2%91%A1_Pratt_Parser%E3%82%92%E5%AE%9F%E8%A3%85%E3%81%97%E3%81%9F

> 記号同士の優先順位の強さを見て、周りの数字などがどちらに結びつくかを決定します。
なるほど。このメンタルモデルの獲得ができていなかった。
