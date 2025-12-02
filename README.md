# Advent of Code 2025 Solutions

Here are my solutions to the programming challenge problems in [Advent of Code 2025](https://adventofcode.com/2025).

I am only posting the main code files in whatever programming language that I may use. And I don't expect to complete all of the challenges based on previous experience.

## Notes

### Day 1

I tried hard to solve part 2 with modulo arithmetic etc. but ended up iterating over each click. I couldn't get the negative rotations to evaluate correctly.

### Day 2

I am happy with my solution. Later, when I read the [Reddit Thread](https://www.reddit.com/r/adventofcode/) I discovered some interesting other ways to solve it.

I know basic **Regex** but completely overlooked the idea of using Regex. But creating the Regex seems like needing to solve another puzzle and minimizing the coding aspect.

Another interesting method (but seems very inefficient) is to generate repeating **strings** from a slice of the main string to match the whole string.

My solution involved shifting chunks of digits (as integers) to the right and comparing them starting with a cut through the middle. I think that this may be the most efficient way to do it algorithmically.

## Programming languages used

This year I should maybe try to code the solutions in a variety of different programming languages to re-kindle my knowledge of them and to learn some new ones.

Languages used so far:

* Go
