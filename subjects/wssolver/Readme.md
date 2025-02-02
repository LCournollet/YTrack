# Word Scrabble Solver
## Objectives

Word Scramble Solver consists on receiving a Word Scramble field and finding all the words contained in it.

The given field will always be of size [10]x[10] and will always contain runes.

You are given a list of words to work with , all the words to find will be contained in this list.

All the possible words have to be found.


The Words to find can only be in line , Column or Diagonal like so :

![](SolverVector.png)



## Instructions

Write a function that reads the Word Scramble field and prints the words found.

The Program must draw the letter table like so :

![](ExampleTable.png)


Here is an example of the program running :

![](ExampleSolverMM.gif)

You are only allowed the packages listed bellow : 

- "buffio"

- "log"

- "os"

- "github.com/01-edu/z01"
## Ressources

Here is an example of letter table to work with and test your code : 
```
var field = [10][10]rune{
	{'c', 'o', 't', 'd', 't', 'r', 's', 'n', 'e', 'c'},
	{'r', 'e', 'e', 'o', 't', 'e', 'o', 'h', 'u', 'c'},
	{'ê', 'u', 'h', 'h', 't', 'r', 'l', 'a', 'o', 'a'},
	{'p', 'f', 'w', 'a', 'e', 'r', 'e', 'a', 'a', 'f'},
	{'e', 's', 't', 'p', 'r', 't', 'a', 'e', 'r', 'é'},
	{'s', 'a', 'p', 'y', 'u', 'i', 'o', 'e', 's', 'd'},
	{'p', 'a', 'e', 'o', 'i', 't', 'c', 'd', 't', 'e'},
	{'n', 'n', 'c', 'n', 'c', 'w', 'b', 'o', 'b', 'n'},
	{'f', 'o', 'u', 'r', 'c', 'h', 'e', 't', 't', 'e'},
	{'b', 'a', 'i', 'e', 's', 't', 'h', 'n', 'w', 's'},
}
```

Here is the list of Words :

[words.txt](https://github.com/Lyon-Ynov-Campus/YTrack/blob/master/subjects/wssolver/words.txt)

### Docs:
[Readfile.doc](https://golangdocs.com/golang-read-file-line-by-line)

