# Scanner

The process of scanning takes in a MongDB query and outputs a sequence of tokens. In this case, a token is a terminating symbol, which is defined as,
```
type Symbol struct {
	name          string
	val           string
	isTerminating bool
}
```
The scanning process returns a list of symbols whose `isTerminating` field is set to `true`. The purpose of this is to divide the query up into separate chunks.

Since it's easier to have a single struct instead of two separate ones, the terminating and non-terminating structs were combined to the above and named `Symbol`.

One can understand the reasoning behind the terminating field if they go through the [Parser](https://github.com/npsolver/Mongolang/tree/master/core/parser) section.

## Example
For the query below:
```
db.users.find({ age: 18 })
```
we get the following list of tokens where each line represent a token's `name` and `val` (the `isTerminating` field is `true` for all of them):
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

## Process
The process of scanning involves using a Deterministic Finite Automata (DFA). The configuration file containing the states and rules of the DFA that we use for our scanner can be found in [scanner.dfa](https://github.com/npsolver/Mongolang/blob/master/core/scanner/scanner.dfa). 

For more information about how the DFA works, see [DFA](https://github.com/npsolver/Mongolang/tree/master/core/scanner/dfa). 
