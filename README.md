# Mongolang
A tool to generate Golang code from Mongo Queries.



## Implementation
The project consists of a compiler to convert a Mongo query to Golang code. The different parts of the compiler follow the well known steps like Scanning, Parsing, etc.

The motivation behind this implementation is that the process is a lot similar to converting C++ code to Assembly language. The only difference is that we are compiling the code for using it in another language, rather than to a lower level language for easier compilation (like using an assembler).

More information about the compiler can be found here





### Step 1: Scanning
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

