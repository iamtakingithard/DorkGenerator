package main

import (
	"bufio"
	"log"
	"os"
)

func readWords() []string {
	file, err := os.Open("keywords.txt")

	if err != nil {
		log.Fatal("no words")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	defer func() {
		_ = file.Close()
	}()

	return words

}

func readFormats() []string {
	file, err := os.Open("formats.txt")

	if err != nil {
		log.Fatal("no formats.txt")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var formats []string

	for scanner.Scan() {
		formats = append(formats, scanner.Text())
	}

	defer func() {
		_ = file.Close()
	}()

	return formats

}

func readTypes() []string {
	file, err := os.Open("types.txt")

	if err != nil {
		log.Fatal("no types")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var types []string

	for scanner.Scan() {
		types = append(types, scanner.Text())
	}

	defer func() {
		_ = file.Close()
	}()

	return types

}

func writeFile(filename string, content string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	if _, err := f.WriteString(content); err != nil {
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}
}

func main() {
	words := readWords()
	formats := readFormats()
	types := readTypes()
	for _, word := range words {
		for _, format := range formats {
			for _, type_ := range types {
				writeFile("dorks.txt", word+format+type_+"\n")
			}
		}
	}
}
