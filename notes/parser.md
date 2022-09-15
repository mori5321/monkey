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


