# Radix Trie

[![Build status](https://travis-ci.org/mramshaw/radix-trie.svg?branch=master)](https://travis-ci.org/mramshaw/radix-trie)
[![Coverage Status](http://codecov.io/github/mramshaw/radix-trie/coverage.svg?branch=master)](http://codecov.io/github/mramshaw/radix-trie?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mramshaw/radix-trie?style=flat-square)](https://goreportcard.com/report/github.com/mramshaw/radix-trie)
[![GoDoc](https://godoc.org/github.com/mramshaw/radix-trie?status.svg)](https://godoc.org/github.com/mramshaw/radix-trie)
[![GitHub release](https://img.shields.io/github/release/mramshaw/radix-trie.svg?style=flat-square)](https://github.com/mramshaw/radix-trie/releases)

This is mainly based on the [Wikipedia article](https://en.wikipedia.org/wiki/Radix_tree).

Note that a __radix tree__ is a space-optimized [trie](https://en.wikipedia.org/wiki/Trie).

According to the article cited above, tries were first suggested in 1959 by a Frenchman
named René de la Briandais (`tri` is French for `sort`). According to Donald Knuth’s
research in The Art of Computer Programming:

> Trie memory for computer searching was first recommended by René de la Briandais.
> He pointed out that we can save memory space at the expense of running time if we
> use a linked list for each node vector, since most of the entries in the vectors
> tend to be empty.

[Note that this is not the approach being followed here.]

More particularly, this seems to be a __Patricia trie__ - which I believe is the binary form.

![Patricia_trie](https://upload.wikimedia.org/wikipedia/commons/a/ae/Patricia_trie.svg)

UPDATE: I found this example (also from Wikipedia) that shows an interesting edge case:

![Edge_case](https://upload.wikimedia.org/wikipedia/commons/6/63/An_example_of_how_to_find_a_string_in_a_Patricia_trie.png)

It shows that __slow__ can be both a node - and a leaf - at the same time. This is something
which I had not considered (another example might be __real__ and __realistic__). Interestingly,
my understanding was that the root node of the trie (it may - or may not - be true that trie
comes from _retrieval_) is always a single character whereas this diagram shows the root node
as the entire word `slow`.

For retrieval purposes I am inclined to persist with using a single character as a root.
The only uses for a trie that I have been able to find are spell-checking and autocomplete
for search bars and the like. So being able to respond quickly to that first typed character
sounds like what I am after. In the case of __bytes__ this would make it possible to use
the initial byte as an offset index, although I doubt this would be practical for runes.

## Motivation

I've been doing a lot of high-level programming lately (RESTful APIs, Scala/Akka) so
something a little more low-level sounded like a nice change of pace. And it has also
allowed me to investigate parts of __Go__ that I do not normally run into.

More importantly, it's a nice opportunity to use
[Table-driven testing](https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go).

## Prerequisites

1. A recent version of Go

2. `make` installed

[Or can simply type `go test -v`.]

## Expected Usage

First seven inserts:

```
  1 (romane) 2 (romanus) 3 (romulus)   4 (rubens)   5 (ruber)         6 (rubicon)        7 (rubicundus)
  |          |           |             |            |                 |                  |
  r          r           r             r            r                 r                  r
  |          |           |            / \          / \               / \                / \
omane       oman         om          om  ubens    om  \             om  ub             om  ub
            / \         / \         / \          / \   \           / \    \           / \    \
           e   us      an ulus    an  ulus     an  ulus \        an  ulus  *        an  ulus  *
                      /  \       /  \         /  \      ube     /  \      / \      /  \      / \
                     e    us    e    us      e    us   /   \   e    us   e  icon  e    us   e   \
                                                      ns    r           / \                / \   \
                                                                       ns  r              ns  r  ic
                                                                                                /  \
                                                                                              on  undus
```

## To Run

Type the following command:

    $ make

The results should look as follows:

    GOPATH=/go GOOS=linux GOARCH=amd64 gofmt -d -e -s -w *.go
    GOPATH=/go GOOS=linux GOARCH=amd64 golint -set_exit_status *.go
    GOPATH=/go GOOS=linux GOARCH=amd64 go tool vet *.go
    GOPATH=/go GOOS=linux GOARCH=amd64 go test -v
    === RUN   TestInsert
    --- PASS: TestInsert (0.00s)
    === RUN   TestFind
    --- PASS: TestFind (0.00s)
    PASS
    ok

## Table-driven Tests

In concept, these sounded like a great idea. In practice, I'd have to say my feelings are mixed.

Like all tests, they get unwieldy very quickly.

[For an example of how tests can become opaque very quickly, check out the
[following example](https://github.com/mramshaw/RESTful-Recipes/blob/master/src/test/main_test.go).
While I made every effort to code these tests to be as transparent as possible, I do not think
the results - in terms of being able to `grok` what I am actually testing __for__ - are all that
easy to sort out. Doing so requires scrolling through pages of tests in order to form an overview.]

On the other hand, the table-driven test framework allowed me to code up the __Find__ tests - and
mock up some data to test them with - while I was still working on the __Insert__ tests, whereas
previously I would have had to complete the __Insert__ tests (and code) in order to test the matching
__Find__ tests (and code). Or maybe do them both in lock-step, step-wise.

As it is, being able to do them both at the same time has been a great help, at least conceptually.

Looking at examples of this testing style, it seemed like it would be possible to quickly grasp - at
a high level - what the tests ***were***. But my experience - in practice - is that this framework
still gets very cluttered quite quickly, so that this high-level overview is not really possible.

I still think this approach is great - but I'd probably reserve it for lower-level components, rather
than for system testing.

## To Do

- [ ] Investigate applications of Patricia tries
- [x] Refactor tests to avoid some of the duplicated code
- [x] Add code and tests to allow for entries such as "slow", "slower", "slowly"
- [x] Find out the idiom for stacking __Insert__ and __Find__ tests (avoiding mocks)
- [ ] Investigate whether byte-based __and__ rune-based options are viable
- [ ] Find more examples of tries in use - specifically Rune-based CJKV (Chinese, Japanese, Korean, Vietnamese)
- [x] Add example of Chinese Rune-based trie
- [ ] Find out whether the usual practice is to sort trie entries (the Wikipedia example __is__ sorted)
- [ ] Tests and code for 'retrieve all entries' functionality

## Credits

This follows on from work four of us did as a group in a couple of hours at the local Golang Meetup.

In particular, Dan coded up the original __Table-driven Tests__, which proved very instructive.

Blake coded up the original ASCII art, which has been fun to elaborate on.

My original fork of this project may be seen here:

    https://github.com/mramshaw/golangvan/tree/master/radix-trie

As I have departed heavily from the original framework (both __design__ and __implementation__) I created this repo.

In particular, I have followed a __Rune__-based approach, rather than a __Byte__-based approach. I have opted
for lazy structures, which are cheaper in terms of memory requirements but possibly costly in performance terms.
Also, the original goal was to be as efficient as possible, whereas I am not greatly concerned with efficiency
here. My goal is an __MVP__ (minimum viable product; meaning a proof-of-concept, demo or spike) for learning purposes.
