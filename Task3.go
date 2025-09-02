package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	data, err := os.ReadFile("users.json")
	if err != nil {
		fmt.Println("Ошибка при чтении файла JSON", err)
		return
	}
	var user []User
	err = json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println(err)
		return
	}
	f := excelize.NewFile()
	sheet := "Sheet1"

	f.SetCellValue(sheet, "A1", "Name")
	f.SetCellValue(sheet, "A1", "Age")

	for i, user := range user {
		row := i + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), user.Name)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), user.Age)
	}
	if err := f.SaveAs("report.xlsx"); err != nil {
		fmt.Println("Произошла ошибка при сохранении файла Excel", err)
		return
	}
	fmt.Println("Создан файл report.xlsx, проверьте результат")
}
