package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/01-edu/z01"
)

var field = [10][10]rune{
	{'m', 'o', 'u', 't', 'o', 'n', 'a', 'i', 't', 'f'},
	{'n', 'a', 't', 'i', 'v', 'i', 't', 'é', 'u', 'o'},
	{'y', 'w', 'v', 'f', 'o', 'p', 't', 'i', 's', 'i'},
	{'b', 'c', 'r', 'è', 'c', 'h', 'e', 'e', 'b', 'n'},
	{'m', 'i', 't', 'e', 'a', 'a', 'g', 'i', 'e', 's'},
	{'y', 'è', 'b', 'e', 'y', 'a', 'm', 'l', 'r', 'p'},
	{'d', 'r', 'r', 'l', 'm', 'c', 'h', 'i', 'e', 'n'},
	{'e', 't', 't', 'e', 'e', 'b', 'é', 'b', 'é', 'e'},
	{'a', 't', 'l', 'o', 'a', 'a', 'r', 'i', 'o', 'o'},
	{'m', 's', 'r', 'e', 'l', 'i', 'g', 'i', 'o', 'n'},
}

func Contains(x string, y string) bool {
	StrToCheck := []string{}
	WordToCheck := []string{}
	Verif := ""
	for _, i := range x {
		StrToCheck = append(StrToCheck, string(i))
	}
	for _, j := range y {
		WordToCheck = append(WordToCheck, string(j))
	}
	for k := 0; k < len(StrToCheck)-1; k++ {
		if StrToCheck[k] == WordToCheck[0] {
			for i := 0; i < len(WordToCheck); i++ {
				if k+i >= len(StrToCheck) {
					Verif = ""
					break
				}
				if WordToCheck[i] == StrToCheck[k+i] {
					Verif += StrToCheck[k+i]
				} else {
					Verif = ""
				}
			}
			if Verif == y {
				Verif = ""
				return true
			}
		}
	}
	return false
}

func readWords() []string {
	// open file
	f, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	tabWords := []string{}

	for scanner.Scan() {
		// do something with a line
		// fmt.Printf("%s\n", scanner.Text())
		tabWords = append(tabWords, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return tabWords
}

func draw() {
	for _, row := range field {
		for _, i := range row {
			z01.PrintRune(i)
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')

	}
	z01.PrintRune('\n')
}

func testColumn(w []string, x [10][10]rune) []string {
	foundWords := []string{}
	for k := 0; k < 10; k++ {
		testColumn := ""
		for i := 0; i < 10; i++ {
			testColumn += string(x[i][k])
		}
		for j := 0; j < len(w); j++ {
			if Contains(testColumn, w[j]) {
				foundWords = append(foundWords, w[j])
			}
		}
	}
	return foundWords
}

func testLine(w []string, x [10][10]rune) []string {
	foundWords := []string{}

	for k := 0; k < 10; k++ {
		testline := ""
		for i := 0; i < 10; i++ {
			testline += string(x[k][i])
		}
		for j := 0; j < len(w); j++ {
			if Contains(testline, w[j]) {
				foundWords = append(foundWords, w[j])
			}
		}
	}
	return foundWords
}

func testDiagDown(w []string, x [10][10]rune) []string {
	foundWords := []string{}
	var b int
	for i := 0; i < 10; i++ {
		testDiag := ""
		b = 0
		for a := i; a < 10; a++ {
			testDiag += string(x[a][b])
			b++
		}
		for j := 0; j < len(w); j++ {
			if Contains(testDiag, w[j]) {
				foundWords = append(foundWords, w[j])
			}
		}
	}
	for i := 0; i < 10; i++ {
		testDiag := ""
		b = 0
		for a := i; a < 10; a++ {
			testDiag += string(x[b][a])
			b++
		}
		for j := 0; j < len(w); j++ {
			if Contains(testDiag, w[j]) {
				foundWords = append(foundWords, w[j])
			}
		}
	}
	return foundWords
}

func testDiagUp(w []string, x [10][10]rune) []string {
	foundWords := []string{}
	var b int
	for i := 9; i > 0; i-- {
		testDiag := ""
		b = 0
		for a := i; a >= 0; a-- {
			testDiag += string(x[a][b])
			b++
		}
		for j := 0; j < len(w); j++ {
			if Contains(testDiag, w[j]) {
				foundWords = append(foundWords, w[j])
			}
		}
	}
	for i := 1; i < 10; i++ {
		testDiag := ""
		b = 9
		for a := i; a < 10; a++ {
			testDiag += string(x[b][a])
			b--
		}
		for j := 0; j < len(w); j++ {
			if Contains(testDiag, w[j]) {
				foundWords = append(foundWords, w[j])
			}
		}
	}
	return foundWords
}

func main() {
	draw()
	words := readWords()
	fmt.Println("Mots trouvés =")
	fmt.Print(testLine(words, field))
	fmt.Print(testColumn(words, field))
	fmt.Print(testDiagDown(words, field))
	fmt.Println(testDiagUp(words, field))
}
