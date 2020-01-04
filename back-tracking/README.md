# Backtracking

## Arm's length recursion

A poor style where unnecessary tests are performed before performing recursive calls. Typically, the tests try to avoid making a call into what would otherwise be a base case. 

For example: escapeMaze  
````text
- Our code recursively tries to explore up, down, left, and right.
- Some of those directions may lead to walls or off the board. 
    Shouldn't we test before making calls in these directions?
````