package main

import (
	"bufio"
	"os"
	"regexp"
	"testing"
)

func TestNoCount(t *testing.T) {
	f_in, _ := os.Open("input/empty.txt")
	defer f_in.Close()
	scanner := bufio.NewScanner(f_in)
	err := ReadInput(scanner)
	if err == nil {
		t.Errorf("[%s] estimated but no error happens", ErrNoTestsCount)
	} else if err.Error() != ErrNoTestsCount {
		t.Errorf("[%s] estimated but [%s] happens", ErrNoTestsCount, err.Error())
	}
}

func TestWrongCount(t *testing.T) {
	f_in, _ := os.Open("input/wrong_tests_count.txt")
	defer f_in.Close()
	scanner := bufio.NewScanner(f_in)
	err := ReadInput(scanner)
	if err == nil {
		t.Errorf("[%s] estimated but no error happens", ErrWrongTestsCount)
	} else if err.Error() != ErrWrongTestsCount {
		t.Errorf("[%s] estimated but [%s] happens", ErrWrongTestsCount, err.Error())
	}
}

func TestNoDimensions(t *testing.T) {
	f_in, _ := os.Open("input/no dimensions.txt")
	defer f_in.Close()
	scanner := bufio.NewScanner(f_in)
	err := ReadInput(scanner)
	if err == nil {
		t.Errorf("[%s] estimated but no error happens", ErrNoDimension)
	} else {
		matched, _ := regexp.MatchString(ErrNoDimension, err.Error())
		if !matched {
			t.Errorf("[%s] estimated but [%s] happens", ErrNoDimension, err.Error())
		}
	}
}

func TestWrongDimensions(t *testing.T) {
	f_in, _ := os.Open("input/wrong dimensions.txt")
	defer f_in.Close()
	scanner := bufio.NewScanner(f_in)
	err := ReadInput(scanner)
	if err == nil {
		t.Errorf("[%s] estimated but no error happens", ErrWrongDimensions)
	} else {
		matched, _ := regexp.MatchString(ErrWrongDimensions, err.Error())
		if !matched {
			t.Errorf("[%s] estimated but [%s] happens", ErrWrongDimensions, err.Error())
		}
	}
}

func TestWrongLines(t *testing.T) {
	f_in, _ := os.Open("input/invalid lines.txt")
	defer f_in.Close()
	scanner := bufio.NewScanner(f_in)
	err := ReadInput(scanner)
	if err == nil {
		t.Errorf("[%s] estimated but no error happens", ErrWrongLinesCount)
	} else {
		matched, _ := regexp.MatchString(ErrWrongLinesCount, err.Error())
		if !matched {
			t.Errorf("[%s] estimated but [%s] happens", ErrWrongLinesCount, err.Error())
		}
	}
}

func TestWrongLen(t *testing.T) {
	f_in, _ := os.Open("input/invalid len.txt")
	defer f_in.Close()
	scanner := bufio.NewScanner(f_in)
	err := ReadInput(scanner)
	if err == nil {
		t.Errorf("[%s] estimated but no error happens", ErrWrongLenCount)
	} else {
		matched, _ := regexp.MatchString(ErrWrongLenCount, err.Error())
		if !matched {
			t.Errorf("[%s] estimated but [%s] happens", ErrWrongLenCount, err.Error())
		}
	}
}

func TestLessLines(t *testing.T) {
	f_in, _ := os.Open("input/less lines.txt")
	defer f_in.Close()
	scanner := bufio.NewScanner(f_in)
	err := ReadInput(scanner)
	if err == nil {
		t.Errorf("[%s] estimated but no error happens", ErrLessLines)
	} else {
		matched, _ := regexp.MatchString(ErrLessLines, err.Error())
		if !matched {
			t.Errorf("[%s] estimated but [%s] happens", ErrLessLines, err.Error())
		}
	}
}

func TestLine(t *testing.T) {
	f_in, _ := os.Open("input/invalid line.txt")
	defer f_in.Close()
	scanner := bufio.NewScanner(f_in)
	err := ReadInput(scanner)
	if err == nil {
		t.Errorf("[%s] estimated but no error happens", ErrInvalidPixels)
	} else {
		matched, _ := regexp.MatchString(ErrInvalidPixels, err.Error())
		if !matched {
			t.Errorf("[%s] estimated but [%s] happens", ErrInvalidPixels, err.Error())
		}
	}
}

func TestMissedEmptyLine(t *testing.T) {
	f_in, _ := os.Open("input/missed empty line.txt")
	defer f_in.Close()
	scanner := bufio.NewScanner(f_in)
	err := ReadInput(scanner)
	if err == nil {
		t.Errorf("[%s] estimated but no error happens", ErrMissedEmptyLine)
	} else {
		matched, _ := regexp.MatchString(ErrMissedEmptyLine, err.Error())
		if !matched {
			t.Errorf("[%s] estimated but [%s] happens", ErrMissedEmptyLine, err.Error())
		}
	}
}

func TestOk(t *testing.T) {
	f_in, _ := os.Open("input/2.txt")
	defer f_in.Close()

	scanner := bufio.NewScanner(f_in)
	err := ReadInput(scanner)
	if err == nil {
		Calculate()
		f_out, _ := os.Create("output/2.txt")
		writer := bufio.NewWriter(f_out)
		defer f_out.Close()

		if WriteResults(writer) == nil {
			f_compare, _ := os.Open("input/2 result.txt")
			defer f_compare.Close()

			_, _ = f_out.Seek(0, 0)
			reader := bufio.NewScanner(f_out)
			comparer := bufio.NewScanner(f_compare)
			result := true
			var s1 string
			var s2 string
			for reader.Scan() {
				if comparer.Scan() {
					s1 = reader.Text()
					s2 = comparer.Text()
					if s1 != s2 {
						result = false
						break
					}
				} else {
					result = false
					break
				}
			}
			if !result {
				t.Errorf("Results not matched!")
			}
		}
	}
}

func TestNOk(t *testing.T) {
	f_in, _ := os.Open("input/2.txt")
	defer f_in.Close()

	scanner := bufio.NewScanner(f_in)
	err := ReadInput(scanner)
	if err == nil {
		Calculate()
		f_out, _ := os.Create("output/2.txt")
		writer := bufio.NewWriter(f_out)
		defer f_out.Close()

		if WriteResults(writer) == nil {
			f_compare, _ := os.Open("input/1.txt")
			defer f_compare.Close()

			_, _ = f_out.Seek(0, 0)
			reader := bufio.NewScanner(f_out)
			comparer := bufio.NewScanner(f_compare)
			result := true
			var s1 string
			var s2 string
			for reader.Scan() {
				if comparer.Scan() {
					s1 = reader.Text()
					s2 = comparer.Text()
					if s1 != s2 {
						result = false
						break
					}
				} else {
					result = false
					break
				}
			}
			if result {
				t.Errorf("Results matched but shouldn't!")
			}
		}
	}
}
