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

What does "`O(1)` additional memory" mean? A stack (of any implementation,
linked list or variable-sized array) will require `O(n)` additional memory,
where `n` is the number of elements in the stack.

## Interview Analysis

"[Hard]" is not quite right. "[Impossible]" is more like it.


## Around the web

This is remarkably widespread.

* [Homework in 300-level algorithms class](http://jeffe.cs.illinois.edu/teaching/algorithms/hwex/s04/hw4.pdf)
* [Exam problem #3](https://www.cs.duke.edu/education/assets_documents/exam_532_algo_fall16.pdf)
* [Hardest interview questions ever](https://www.scien.cx/2021/09/09/the-hardest-coding-interview-questions-ever/)
* [Stack Exchange Question](https://stackoverflow.com/questions/53577545/two-stacks-with-a-deque-whats-the-purpose-of-implementing-it?rq=1)
* [Another Stack Exchange Question](https://stackoverflow.com/questions/624704/design-a-stack-that-can-also-dequeue-in-o1-amortized-time)
* [Solution that doesn't realy fit](https://dev.to/sharansharma94/google-interview-problem-solution-ae2)
* [Probably incorrect big-O analysis](http://trsong.github.io/python/java/2021/02/02/DailyQuestionsFeb.html#apr-21-2021-easy-special-stack)

In light of the fact that *nobody* knows if there's an O(1) algorithm
that Exam Problem #3 describes, I wonder if some Duke CS students aren't due
a letter grade increment, or maybe a refund, or at least an apology?

Apparently the source of this is [Web Problem 34](https://algs4.cs.princeton.edu/13stacks/):
---
Queue with a constant number of stacks.
Implement a queue with a constant number of stacks
so that each queue operations takes a constant (worst-case) number of stack operations.
Warning: Very high degree of difficulty. 
---

Apparently it used to say 3 stacks.
