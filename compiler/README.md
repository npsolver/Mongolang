## Compiler

The compiler consists of the following parts:

### Scanner
Scanning a piece of code just means converting it into a sequence of tokens that are specific to that language. 

The reason this is necessary to be done first is because a single piece of code might be interpreted in more than one way. And usually, we only care about what a single part of the code represents (like a Mongo collection), instead of what value that piece of code actually holds (like the name of a collection). Tokenizing the code makes it easier for later steps where we check that the code actually represents a query in the parsing phase.

For example, a query like,
```
db.getCollection("customers").find({})
```
may be converted into tokens like
```
DBNAME DOT DBCOMMAND LPAREN QUOTE WORD QUOTE RPAREN DOT COLCOMMAND LPAREN LBRACE RBRACE RPAREN
```
and the token structure has a field called `value` which stores the actual value the token represents. Like the `DBCOMMAND` token above represents `getCollection`.

More information about the Scanner can be found [here](scanner).


