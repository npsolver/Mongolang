# Parser
In the parsing step, we use the scanned tokens and try put in into a context-free grammar. To explain this, let us consider the scanned tokens we got from the scanning step,
```
ID db
DOT .
ID users
DOT .
ID find
FIRSTBRACKETLEFT (
SECONDBRACKETLEFT {
ID age
COLON :
INT 18
SECONDBRACKETRIGHT }
FIRSTBRACKETRIGHT )
```
We parse this into a parsetree,
```
START
|---BOF BOF
|---COMMAND
|   |---DBNAME
|   |   |---ID db
|   |---DOT .
|   |---COLLECTIONNAME
|   |   |---ID users
|   |---DOT .
|   |---COLLECTIONCOMMAND
|       |---ID find
|       |---FIRSTBRACKETLEFT (
|       |---OPTIONS
|       |   |---OPTION
|       |       |---SET
|       |           |---SECONDBRACKETLEFT {
|       |           |---FIELDS
|       |           |   |---FIELD
|       |           |       |---FIELDNAME
|       |           |       |   |---ID age
|       |           |       |---COLON :
|       |           |       |---VALUE
|       |           |           |---INT 18
|       |           |---SECONDBRACKETRIGHT }
|       |---FIRSTBRACKETRIGHT )
|---EOF EOF
```
where the `START` symbol is the root of the tree. Every element of the tree is a symbol. The leaves of the tree are precisely the scanned terminating symbols we got earlier and have a value assigned to them. Whereas all other symbols are non-terminating ones. 

## Context-Free Grammar
A context-free grammar refers to a set of rules used to generate the above tree. Each rule defines a way to reduce a given set of symbols in a parsing step. 

For example, the rule below:
```
FIELD -> FIELDNAME COLON VALUE
```
refers to a field passed in a query consisting of a field name, colon, and a value. Here the `COLON` is a terminating symbol while the other two are non-terminating. You can find the whole list of rules for our parser in [parser.cfg](https://github.com/npsolver/Mongolang/blob/master/core/parser/parser.cfg).

When it comes to generating the tree itself using the rules, we use an extended form of DFA. You can find more information about it [here](https://github.com/npsolver/Mongolang/tree/master/core/parser/edfa).