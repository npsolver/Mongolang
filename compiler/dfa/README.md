## Deterministic Finite Automata (DFA)

A DFA consists of a list of states and ways to transfer, called transitions, from one state to another. One usually starts from the start state. Then using the transition rules, we hop from one state to another. There are a few states marked as goal states, reaching which, one is said to have reached the goal.

One can use DFAs for recognition problems. Say we want to tokenize a piece of code. Each goal state of the DFA can represent a specific token. Then when we scan a sequence of characters in the code, we can move from one state to another and identify a token when we reach a goal state.

### Implementation

The state of a DFA can represent anything. So we should make it possible for the user to do so by using generic types. So we define states as,
```

```
