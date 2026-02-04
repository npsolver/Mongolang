# Extended Deterministic Finite Automation
For our parsing step, we again use a DFA but extend it a little bit. Now, each state has a list of `item`'s. Essentially, consider the rule,
```
FIELD -> FIELDNAME COLON VALUE
```
Each state might contain a modified version of the rule, which is called an item, like the following
```
FIELD -> FIELDNAME * COLON VALUE
```
where the `*` symbol is called a bookmark. If a state contains the above item, and then we transition to another state using a `COLON` symbol, the new state will contain the item
```
FIELD -> FIELDNAME COLON * VALUE
```
In addition, we add new rules to the new state by recursion. For example since the above item is already in the new state and we have that the bookmark is on the left of `VALUE`, we will also add in the following to the state:
```
VALUE -> * LIST 
VALUE -> * SET 
VALUE -> * ID
VALUE -> * INT
```
since `VALUE` is a non-terminating symbol and the above rules have `VALUE` on the left hand side of it. Then for example if we transition again to another state using `ID`, then the state will have the item,
```
VALUE -> ID *
```
and since we cannot go any further and the bookmark is at the end of the item, this state will be an accepting state.

We used an LR(0) bottom-up parsing algorithm for our parser. You can watch an illustration of the extended DFA and the algorithm itself [here](https://student.cs.uwaterloo.ca/~cs241/videos/index.php?video=m6-v3-Parsing-Using-LR0-Algorithm).
