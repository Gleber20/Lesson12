package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	mw := io.MultiWriter(file, os.Stdout)
	log.SetOutput(mw)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	scanner := bufio.NewScanner(os.Stdin)
	log.Println("Программа работает, вводите текст. Для выхода из программы введите exit")

	for scanner.Scan() {
		text := scanner.Text()

		if strings.ToLower(text) == "exit" {
			log.Println("Завершение работы программы")
			break
		}
		log.Println(text)
	}
}
