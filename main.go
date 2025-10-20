package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Викторина")
	myWindow.Resize(fyne.NewSize(1920, 1080))
	label := widget.NewLabel("Викторина на 4 вопросоа. " +
		"Отвечайте на все вопросы русским языком начиная с заглавной буквы")
	label1 := widget.NewLabel("1 вопрос - \"Столица России?\"\n")
	entry1 := widget.NewEntry()
	label2 := widget.NewLabel("2 вопрос \"Столица Японии?\"\n")
	entry2 := widget.NewEntry()
	label3 := widget.NewLabel("3 вопрос\"Столица Китая?\"\n")
	entry3 := widget.NewEntry()
	label4 := widget.NewLabel("4 вопрос \"Столица Франции?\"\n")
	entry4 := widget.NewEntry()
	resultTest := widget.NewLabel("")
	MyButton := widget.NewButton("Готов", func() {
		answer1 := entry1.Text
		answer2 := entry2.Text
		answer3 := entry3.Text
		answer4 := entry4.Text
		scores := 0
		if answer1 == "Москва" {
			scores++
		}
		if answer2 == "Токио" {
			scores++
		}
		if answer3 == "Пекин" {
			scores++
		}
		if answer4 == "Париж" {
			scores++
		}
		if scores == 4 {
			resultTest.SetText(fmt.Sprintf("Все правильно: %d", scores))
		} else {
			resultTest.SetText(fmt.Sprintf("Ваши баллы: %d", scores))
		} //Кнопка готов, реализовать
	})
	myWindow.SetContent(container.NewVBox(
		label,
		label1,
		entry1,
		label2,
		entry2,
		label3,
		entry3,
		label4,
		entry4,
		MyButton,
		resultTest, //для готовности
		//для ника
	))
	myWindow.ShowAndRun()
}