# Radix Trie

This is mainly based on the [Wikipedia article](https://en.wikipedia.org/wiki/Radix_tree).

Note that a __radix tree__ is a space-optimized [trie](https://en.wikipedia.org/wiki/Trie).

More particularly, this seems to be a __Patricia tree__ - which I believe is the binary form.

![Patricia_trie](https://upload.wikimedia.org/wikipedia/commons/a/ae/Patricia_trie.svg)

## Motivation

I've been doing a lot of high-level programming lately (RESTful APIS, Scala/Akka) so
something a little more low-level sounded like a nice change of pace. And it has also
allowed me to investigate parts of __Go__ that I do not normally run into.

Also, it's a nice opportunity to use
[Table-driven testing](https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go).

## Prerequisites

1. A recent version of Go

2. `make` installed

[Or can simply type `go test -v`.]

## Expected Usage

First four inserts:

```
   1          2             3                4
   |          |             |                |
   r          r             r                r
   |          |             |               / \
 omane      oman            om             om  ubens
             /\            / \            / \
            e  us         an  ulus       an  ulus
                         / \            / \
                        e  us          e   us
```

## To Run

Type the following command:

    make

## Table-driven Tests

In concept, these sounded like a great idea. In practice, I'd have to say my feelings are mixed.

Like all tests, they get unwieldy very quickly. On the other hand, this framework allowed me to
code up the __Find__ tests - and mock up some data to test them with - whereas previously I would
have had to complete the __Insert__ code (and tests) in order to test the matching __Find__ code.
Or maybe do them both in lock-step, step-wise.

As it is, being able to do them both at the same time has been a great help, at least conceptually.

Looking at examples of this testing style, it seemed like it would be possible to grasp - at a high
level - what the tests *were*. But my experience - in practice - is that the framework gets very
cluttered quite quickly, so that this high-level overview is not really possible.

I still think this approach is great - but I'd reserve it for lower-level components, rather than
for system-wise testing.

## To Do

- [ ] Investigate applications of Patricia tries
- [ ] Find out the idiom for stacking __Insert__ and __Find__ tests
- [ ] Investigate whether byte-based __and__ rune-based options are viable
- [ ] Find more examples of tries in use (specifically Rune-based / CJKV {Chinese, Japanese, Korean, Vietnamese})
- [ ] Find out whether the usual practice is to sort trie entries
- [ ] Tests and code for 'retrieve all entries' functionality

## Credits

This follows on from work four of us did (as a group) at the local Golang Meetup.

In particular, Dan coded up the original __Table-driven Tests__, which proved very instructive.

My original fork of this project may be seen here:

    https://github.com/mramshaw/golangvan/tree/master/radix-trie

As I have departed from the original framework (both __design__ and __implementation__) I created this repo.

In particular, I hope opted for a __Rune__-based approach, rather than a __Byte__-based approach. I have opted
for lazy sturctures, which is cheaper in terms of memory requirements but possibly costly in performance terms.
Also, the original effort was to be as efficient, whereas I am not greatly concerned with efficiency here. My
goal was an __MVP__ (minimum viable product; meaning a proof-of-concpet, demo or spike) for learning purposes.
