package main

import (
	"os"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	fifth "github.com/h8gi/fifth/lib"
)

func main() {
	i := fifth.NewInterpreter(os.Stdin)

	myApp := app.New()
	myWindow := myApp.NewWindow("ca-explorer")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter forth command...")

	content := widget.NewVBox(input, widget.NewButton("Run", func() {
		i.SetString(input.Text)
		err := i.Run()
		if err == fifth.QuitError {
			myApp.Quit()
		}
		if err != nil {
			//
		}
	}))
	myWindow.SetContent(content)

	myWindow.ShowAndRun()
}
