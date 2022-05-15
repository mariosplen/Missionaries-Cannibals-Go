# Missionaries and Cannibals

Missionaries and Cannibals problem solved using Go during an optional university assignment for the subject of Artificial Intelligence

## The Problem

> In the missionaries and cannibals problem, three missionaries and three cannibals must cross a river using a boat which can carry at most two people, under the constraint that, for both banks, if there are missionaries present on the bank, they cannot be outnumbered by cannibals (if they were, the cannibals would eat the missionaries). The boat cannot cross the river by itself with no people on board. And, in some variations, one of the cannibals has only one arm and cannot row.[[1]](https://en.wikipedia.org/wiki/Missionaries_and_cannibals_problem#cite_note-PressmanSingmaster-1)

<img src="image.png" alt="MC" width= “100%”/>

## Tools

- Go Programming Language
- General search algorithm from the book Artificial Intelligence ISBN: 978-618-5196-44-8

## State Representation

A state can be represented by the (m c b) notation, where m is the number of missionaries on
the left, c is the number of cannibals on the left, and b indicates whether the boat is on the
left bank or right bank. For example, the initial state is (3 3 L) and the goal state is
(0 0 R). https://faculty.cc.gatech.edu/~riedl/classes/2014/cs3600/homeworks/missionaries-soln.pdf

## Usage

```
go run *.go
```

## Output

One of the solutions:

```
(3,3,L)->(3,1,R)->(3,2,L)->(3,0,R)->(3,1,L)->(1,1,R)->(2,2,L)->(0,2,R)->(0,3,L)->(0,1,R)->(1,1,L)->(0,0,R)
```
