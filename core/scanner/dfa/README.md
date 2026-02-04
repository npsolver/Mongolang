# Deterministic finite automation

A DFA is a finite state machine that accepts or rejects a string. This can be used in our scanner to check whether a string can be considered a symbol (through an accept) or not (through a reject).

This is done through state transitions. We define states(nodes) and transitions(edges) between them hence making a DFA tree. Each transition represents a character and a transversal through the DFA tree from the start node describes a specific string. Some states are marked as accepting which means that the string up to that point represents a specific symbol. 

A DFA is defined in a file with the following structure:

```
.STATES
start
id !
int !
.TRANSITIONS
start a-z A-Z _ id
id a-z A-Z _ id
start 1-9 int
int 0-9 int
.END
```

It contains 2 sections: STATES and TRANSITIONS. States with an ! at the end of them are accepting states and represent a terminating symbol with the same name in capital letters. Transitions are defined similarly. `start 1-9 int` represents transitions from start to int using integers from 1 to 9.

This DFA above can be used to parse IDs (which have letters and underscores) and integers. 

The algorithm used to match the given strings to symbols is called Maximal Munch Algorithm. Basically it means on each iteration parse the longest possible symbol. For example if the string was 12345, instead of getting 5 int symbols (as above) each representing the values from 1 to 5, we would just get 1 symbol with the value 12345.