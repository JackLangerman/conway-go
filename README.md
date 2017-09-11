# conway-go


## Intro to Go

### Task

Implement Conway's Game of life

### The rules of life say:
A living cell with two or three neighboring living cells survives into the next generation. A living cell with fewer than two living neighbors dies of loneliness and a living cell with more than three living neighbors will die from overcrowding.
Each empty/dead cell that has exactly three living neighbors--no more, no fewer-- comes to life in the next generation.

### Input:
Get the initial state from a file, named "life.txt" in the current working directory.
Each row will have dashes for dead cells and asterisks for live ones.
To be nice, I will make all lines the same length.
Output to standard output (i.e. the screen):
Display the original state of the world, as read in from the file and 10 additional generations.
In each generation, each row will have spaces and asterisks. (spaces look nicer than dashes)
Spaces represent empty/dead cells
Asterisks represent occupied/living cells

### Suggestion: 
For your internal representation, create an extra row on the top and bottom as well as an extra column to the right and the left of the actual data. Mark the extra cells as dead. This will simplify your code, as you won't need special checks for "boundary" cells.

Sample input and output for the program are attached

### Good programming style: 
By now you know good programming style!
Naming
Use of functions
Length of lines and functions
Good commenting.

### TODO
* goify comments -- Make comments full sentances [https://golang.org/doc/effective_go.html#commentary](https://golang.org/doc/effective_go.html#commentary)
* run go fmtt
