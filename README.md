# Fuoco: A Cellular Automata Forest Fire Simulator

A cellular autometa forest fire simulator written in Go. Some useful features:
- Zero external dependencies
- Accounts for elevation, wind, fuel, and moisture
- Easily swappable state transition functions

## How it works


Let `F_k(state, i, j, parameter)` represent a state transition probability function for 
some parameter `k` in cell `(i,j)`, such as transitioning from a cell in a `Ready` state to a `Burning` 
state. Then the probability of a state transition is

```
P(A -> B) = F_1(state, i, j, parameter) * F_2(state, i, j, parameter)...F_k(state, i, j, parameter)
```

Basically, this multiplies the probabilities of a transition for each parameter together. This is highly
unphysical, and yet it produces reasonable results.