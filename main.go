package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(`input.txt`)
	if err != nil {
		fmt.Println("Ошибка открытия текстового файл", err)
		return
	}
	defer file.Close()

	var textBuilder strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		textBuilder.WriteString(scanner.Text())
		textBuilder.WriteString(" ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}
	text := strings.ToLower(strings.TrimSpace(textBuilder.String()))
	words := strings.Fields(text)

	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}

	outfile, err := os.Create(`output.csv`)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outfile.Close()

	writer := csv.NewWriter(outfile)
	defer writer.Flush()
	writer.Write([]string{"Слово", "Частота"})

	for word, count := range counts {
		writer.Write([]string{word, strconv.Itoa(count)})
	}
	fmt.Println("Результат записан в файл output.csv")
}
