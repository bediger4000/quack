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

### My Attempt

I did this in Go, because that's what I'm using as my primary language right now.

I wrote a Go interface for a quack:

```go
type Quack interface {
    Push(interface{})
    Pop() interface{}
    Pull() interface{}
    Print(io.Writer)
}
```
I wrote a [2-stack quack](quack/stack.go) implementation of that interface,
and a [doubly-linked list quack](quack/list.go) to see if my
imperative programming language intuition about this problem works.

The [program](q1.go) "interprets" pushes, pops and pulls on one or the other
implementation of a quack:

```sh
$ go build q1.go
$ ./q1 push 1 push 2 push 3 pull push 4 quack
pull 1
Push stack: 4 -> 
Pull stack: 2 -> 3 -> 
$ ./q1 -list push 1 push 2 push 3 pull push 4 quack
pull 1
4 -> 3 -> 2 -> 

```

By default, the program uses the 2-stack implementation.
Using a `-list` flag, you can get it to use a doubly-linked list implementation.

Alas, I'm not smart enough to see how a 3-stack implementation might make sense,
so I didn't do one.
I'm also not smart enough to do the 6-stack,
O(1) implementation without a lot of hints.

## Interview Analysis

The Chris Okasaki [paper](https://www.cambridge.org/core/journals/journal-of-functional-programming/article/simple-and-efficient-purely-functional-queues-and-deques/7B3036772616B39E87BF7FBD119015AB)
that talks about O(1) lazy-evaluation implementations of this problem says this
about the 2-stack implementation:

---

Each operation requires only O(1) amortized time,
but any particular remove might require O(n) time.

---

I have no idea what the Google interviewers are after.


## Around the web

This is remarkably widespread.

* [Homework in 300-level algorithms class](http://jeffe.cs.illinois.edu/teaching/algorithms/hwex/s04/hw4.pdf)
* [Exam problem #3](https://www.cs.duke.edu/education/assets_documents/exam_532_algo_fall16.pdf)
* [Hardest interview questions ever](https://www.scien.cx/2021/09/09/the-hardest-coding-interview-questions-ever/)
* [Stack Exchange Question](https://stackoverflow.com/questions/53577545/two-stacks-with-a-deque-whats-the-purpose-of-implementing-it?rq=1)
* [Another Stack Exchange Question](https://stackoverflow.com/questions/624704/design-a-stack-that-can-also-dequeue-in-o1-amortized-time)
* [Solution that doesn't really fit](https://dev.to/sharansharma94/google-interview-problem-solution-ae2)
* [Informal big-O analysis](http://trsong.github.io/python/java/2021/02/02/DailyQuestionsFeb.html#apr-21-2021-easy-special-stack)

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
