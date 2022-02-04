package main

import (
	"fmt"
	"os"

	//"io"
	"bufio"
	"errors"
	"strconv"
	"strings"
	"sync"
)

type Bitmap struct {
	lines, length int
	pixels        [][]bool
	distances     [][]int
}

const maxLines int = 182
const maxLength int = 182

const white byte = 49
const black byte = 48

const ErrNoTestsCount = "no tests count provided"
const ErrWrongTestsCount = "tests count not a number"
const ErrNoDimension = "no dimensions provided"
const ErrWrongDimensions = "wrong dimensions"
const ErrWrongLinesCount = "wrong lines number"
const ErrWrongLenCount = "wrong length"
const ErrLessLines = "not enough lines"
const ErrInvalidPixels = "invalid pixels line"
const ErrMissedEmptyLine = "empty line required after data"

// load data from given scanner
func (bmp *Bitmap) Load(reader *bufio.Scanner) error {
	if !reader.Scan() {
		return errors.New(ErrNoDimension)
	}
	str := reader.Text()
	strs := strings.Split(str, " ")
	if len(strs) != 2 {
		return fmt.Errorf(ErrWrongDimensions+": %s", str)
	}

	t, err := strconv.Atoi(strs[0])
	if err != nil || t < 1 || t > maxLines {
		return fmt.Errorf(ErrWrongLinesCount+" %s (should be number between 1 and %d", strs[0], maxLines)
	}
	bmp.lines = t

	t, err = strconv.Atoi(strs[1])
	if err != nil || t < 1 || t > maxLength {
		return fmt.Errorf(ErrWrongLenCount+" %s (should be number between 1 and %d", strs[0], maxLength)
	}
	bmp.length = t

	bmp.pixels = make([][]bool, bmp.lines)
	bmp.distances = make([][]int, bmp.lines)
	t = 0
	var bts []byte
	for t < bmp.lines {
		if !reader.Scan() {
			return fmt.Errorf(ErrLessLines+" (%d required, %d provided)", bmp.length, t)
		}
		str = reader.Text()
		bmp.pixels[t] = make([]bool, bmp.length)
		bmp.distances[t] = make([]int, bmp.length)
		bts = []byte(str)
		if len(bts) != bmp.length {
			return fmt.Errorf(ErrInvalidPixels+": %s", str)
		}
		for index, b := range bts {
			if b == black {
				bmp.pixels[t][index] = false
			} else if b == white {
				bmp.pixels[t][index] = true
			} else {
				return fmt.Errorf(ErrInvalidPixels+": %s", str)
			}
		}
		t++
	}

	if !reader.Scan() {
		return reader.Err()
	}

	str = reader.Text()
	if str != "" {
		return fmt.Errorf(ErrMissedEmptyLine+", got %s", str)
	}
	return nil
}

// save data to given writer
func (bmp *Bitmap) Save(writer *bufio.Writer) error {
	s := "%d"
	for i := 1; i < bmp.length; i++ {
		s = s + " %d"
	}
	s = s + "\n"
	v := make([]interface{}, bmp.length)
	for _, l := range bmp.distances {
		for i, val := range l {
			v[i] = val
		}
		fmt.Fprintf(writer, s, v...)
	}
	fmt.Fprintln(writer, "")
	return nil
}

func (bmp *Bitmap) checkPixel(dX, dY int) bool {
	if dX >= 0 && dY >= 0 && dX < bmp.length && dY < bmp.lines {
		return bmp.pixels[dY][dX]
	}
	return false
}

// calculate distances
func (bmp *Bitmap) Calculate() {
	wg := &sync.WaitGroup{}
	wg.Add(bmp.length * bmp.lines)

	for i := range bmp.pixels {
		for j := range bmp.pixels[i] {
			go func(wgr *sync.WaitGroup, _i, _j int) {
				dX := 0
				dY := 0
				if bmp.pixels[_i][_j] {
					bmp.distances[_i][_j] = 0
				} else {
					distance := 1
					found := false
					for distance < bmp.length+bmp.lines && !found {
						dX = distance
						dY = 0
						for dX >= 0 {
							if bmp.checkPixel(_j+dX, _i+dY) ||
								bmp.checkPixel(_j+dX, _i-dY) ||
								bmp.checkPixel(_j-dX, _i+dY) ||
								bmp.checkPixel(_j-dX, _i-dY) {
								bmp.distances[_i][_j] = distance
								found = true
								break
							}
							dX--
							dY++
						}
						distance++
					}
				}
				wg.Done()
			}(wg, i, j)
		}
	}
	wg.Wait()
}

var bitmaps []Bitmap

// read input data (any scanner)
func ReadInput(reader *bufio.Scanner) error {
	var s string
	if !reader.Scan() {
		return errors.New(ErrNoTestsCount)
	}
	s = reader.Text()
	testsNumber, err := strconv.Atoi(s)
	if err != nil {
		return errors.New(ErrWrongTestsCount)
	}
	bitmaps = make([]Bitmap, testsNumber)
	currentTest := 0
	for currentTest < testsNumber {
		err = bitmaps[currentTest].Load(reader)
		if err != nil {
			return fmt.Errorf("error on test %d: %s", currentTest+1, err.Error())
		}
		currentTest++
	}
	return nil
}

// Calculate distances

func Calculate() {
	for _, bmp := range bitmaps {
		bmp.Calculate()
	}
}

// write results (any writer)
func WriteResults(writer *bufio.Writer) error {
	for _, bmp := range bitmaps {
		err := bmp.Save(writer)
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

func main() {
	//example usage

	// create scanner (file, stdin etc)
	f_in, _ := os.Open("input/1.txt")
	defer f_in.Close()
	scanner := bufio.NewScanner(f_in)
	// load data
	if err := ReadInput(scanner); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		// calculate all bitmaps
		Calculate()

		// create writer (file, stdin etc)
		f_out, _ := os.Create("output/1.txt")
		writer := bufio.NewWriter(f_out)
		defer f_out.Close()
		// save results
		if err := WriteResults(writer); err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}

}
