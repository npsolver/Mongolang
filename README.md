# Mongolang
A tool to generate Golang code from Mongo Queries.



## Implementation
The project consists of a compiler to convert a Mongo query to Golang code. The different parts of the compiler follow the well known steps like Scanning, Parsing, etc.

The motivation behind this implementation is that the process is a lot similar to converting C++ code to Assembly language. The only difference is that we are compiling the code for using it in another language, rather than to a lower level language for easier compilation (like using an assembler).

More information about the compiler can be found [here](compiler).


