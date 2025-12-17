# Advent of Code 2025 Solutions

Here are my solutions to the programming challenge problems in [Advent of Code 2025](https://adventofcode.com/2025).

I am only posting the main code files in whatever programming language that I may use. And I don't expect to complete all of the challenges based on previous experience.

## Notes

### Day 1

It was a challenge to evaluate the negative rotations correctly.

### Day 2

I am happy with my solution. Later, when I read the [Reddit Thread](https://www.reddit.com/r/adventofcode/) I discovered some interesting other ways to solve it.

I know basic **Regex** but completely overlooked the idea of using Regex. But creating the Regex seems like needing to solve another puzzle and minimizing the coding aspect.

Another interesting method (but seems very inefficient) is to generate repeating **strings** from a slice of the main string to match the whole string.

My solution involved shifting chunks of digits (as integers) to the right and comparing them starting with a cut through the middle. I think that this may be the most efficient way to do it algorithmically.

### Day 3

To solve this I scanned the digits from after the previous largest digit up to the position from the end of the start of the remaining number of digits. Always looking to capture the largest digit and its position.

To make this code more efficient I could have pre-processed the bank data into arrays of integers rather than dealing with the strings.

### Day 4

I refactored my code for part 2 to create a 2D integer slice to store the input data. This allowed for the use of constants which I called EMPTY, ROLL_TO_GO, and ROLL to indicate the state of each position.

My solution involved a function to count and mark the rolls that could be removed, and a function to remove the rolls. Then call them in a loop until no more rolls could be removed.

### Day 5

Part 2 involved sorting the ranges in ascending order of the starting ID numbers. Then the ranges were merged to combine overlapping ranges. Then sum up the total of the IDs.

### Day 6

This had challenges involving string manipulation and converting to numbers.

### Day 7

To solve this I added a weight to the beams that gets distributed at each split to the 2 new beams or added to existing beams in the same position. Then add up the weights at the end.

### Day 8

For this, I created a Pair type and assigned an ID to each pair of vectors. This allowed for easy debugging where the IDs were in order of the distance between the vectors. The circuits were then comprised of lists of IDs where each ID may only be in one of the circuits as the pairs are added.

### Day 9

For part 2 I wrote code to sort the possible rectangles in descending order of area. Then test points around the perimiter of the rectangle to see if they are 
inside the path of the vertices or on the path.

[This article](https://alienryderflex.com/polygon/) had a good algorithm for testing if a point is inside a polygon.

The edges of the path are either vertical or horizontal, so it is straightforward to test points that are on an edge. But points that are not on an edge need
testing to see if they are inside the outer path.

Evaluating all of the rectangles seemed to take a long time so I skipped testing some of the largest rectangles. For example, you could start from a percentage of the maximum rectangle area.

### Day 10

I developed a solution that works for Part 1 and the test data (part1 and 2) but is too inefficient to solve Part 2 with my input data. There are too many combinations of button presses to consider.

My approach to optimise the solution is to find the maximum number of presses for each button to individually hit one of the joltages. Then loop over the ranges of presses of each button. If any of the buttons cause the joltage to be exceeded then the current combination of button presses is skipped.

