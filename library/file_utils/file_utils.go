package file_utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	bufferSize = 32 * 1024
)

// ForEachLineDo performs the given action for each line in the file.
func ForEachLineDo(filename string, action func(line string)) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		action(scanner.Text())
	}

	return scanner.Err()
}

// ForEachLineDoIterator performs a given action for each line in the file.
// The action has both a line index and the text string.
func ForEachLineDoIterator(filename string, action func(index int, line string)) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	lineNumber := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		action(lineNumber, scanner.Text())
		lineNumber++
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return nil
}

// LoadAllLines reads all the lines of a file all at once.
func LoadAllLines(filename string) (lines []string, err error) {
	var total int
	total, err = CountLines(filename)
	if err != nil {
		return nil, err
	}

	lines = make([]string, 0, total)
	err = ForEachLineDoIterator(filename, func(index int, line string) {
		if index%1_000_000 == 0 && index > 0 {
			log.Println(fmt.Sprintf("%d million lines read from %s", index/1_000_000, filename))
		}

		lines = append(lines, line)
	})

	if err != nil {
		return nil, err
	}

	return lines, nil
}

// CountLines counts the number of lines in a file.
func CountLines(filename string) (lines int, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return -1, err
	}

	defer file.Close()

	buffer := make([]byte, bufferSize)
	lines = 0
	newline := []byte{'\n'}
	newlineTerminator := false
	firstLoop := true
	for {
		var count int
		count, err = file.Read(buffer)
		lines += bytes.Count(buffer[:count], newline)
		if count != 0 {
			newlineTerminator = buffer[count-1] == '\n'
		}

		if err == io.EOF {
			if !newlineTerminator && !firstLoop {
				lines++
			}

			return lines, nil
		}
		if err != nil {
			return -1, err
		}

		firstLoop = false
	}
}

// WriteLines writes all the given lines to a file.
func WriteLines(filename string, lines []string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	for index, line := range lines {
		if index%1_000_000 == 0 && index > 0 {
			log.Println(fmt.Sprintf("%d million lines written to %s", index/1_000_000, filename))
		}

		_, err = fmt.Fprintln(file, line)
		if err != nil {
			return err
		}
	}

	return nil
}
