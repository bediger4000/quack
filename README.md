# Quack Data Structure
## Daily Coding Problem: Problem #1116 [Hard]

This problem was asked by Google.

A quack is a data structure combining properties of both stacks and queues.
It can be viewed as a list of elements written left to right such that three
operations are possible:

* `push(x)`: add a new item x to the left end of the list
* `pop()`: remove and return the item on the left end of the list
* `pull()`: remove the item on the right end of the list.

Implement a quack using three stacks and `O(1)` additional memory,
so that the amortized time for any push, pop, or pull operation is `O(1)`.

## Analysis

What's up with the definition of `pull()`?
`pop()` is to remove *and return*,
but `pull()` only removes?
That makes it hard to test.
I'm going to have `pull()` return something so I can try it out.

Do we care about the type of the "new item x" added to the left end of the list
specified for `push(x)`?

## Interview Analysis
