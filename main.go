package main

import "fyne.io/fyne/v2/app"
import "fyne.io/fyne/v2/container"
import "fyne.io/fyne/v2/widget"

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Пример приложения на Golang")

	hello := widget.NewLabel("Привет, мир!")
	content := container.NewVBox(hello)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
