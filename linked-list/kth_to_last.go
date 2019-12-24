package list

// Implement an algorithm to find the kth to last element of a singly linked list.
// Hints: #8, #25, #41, #67, #126

// For this solution, we have defined k that passing in k = 1 would return the last element
// k = 2 would return the second to the last element, and so on.

// Solution #1: if linked list size is known
// If the size of the linked list is known, the kth to last element is the (length - k)th element.
// We can just iterate through the linked list to find this element.

// Solution #2: recursive
