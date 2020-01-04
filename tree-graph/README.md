## Tree

A binary search tree is a binary tree in which every node fits a specific ordering property: all left descendents <= n < all right descendents 

The definition of a binary search tree can vary slightly with respect to equality. 

Note that this inequity must be true for all of a node's descendents, not just its immediate children. 

**Binary Tree vs Binary Search Tree**

**Balanced vs Unbalanced**

**Complete Binary Trees**
Every level of the tree is fully filled, except for perhaps the last level. To that extend that
the last level is filled, it is filled left to right. 

**Full Binary Trees**
A full binary tree is a binary tree in which every node has either zero or two children. That is, no nodes have only one child. 

**Perfect Binary Trees**
A perfect binary tree is one that is both full and complete.

### Binary Tree Traversal 

````text
In-order: left branch -> current node -> right branch
We perform on a binary search tree, it visits the nodes in ascending order (hence the name "in-order")
Pre-order: current node -> left branch -> right branch 
Visiting the current node before its child nodes, hence the name "pre-order"
Post-order: left branch -> right branch -> current node
Visiting the current node after its child nodes, hence the name "post-order"
````

### Traversal Trick
````text
To quickly generate a traversal:
- Trace a path counterclockwise
- As you pass a node on the proper side, process it.
  - pre-order: left side
  - in-order: bottom
  - post-order: right side
````