package main

import (
	"bytes"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	fifth "github.com/h8gi/fifth/lib"
)

func main() {
	i := fifth.NewInterpreter()
	stdout := new(bytes.Buffer)
	i.SetWriter(stdout)
	myApp := app.New()
	myWindow := myApp.NewWindow("ca-explorer")

	inputArea := widget.NewMultiLineEntry()
	outputArea := widget.NewLabel("")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{"forth code", inputArea},
		},
		OnSubmit: func() {
			i.SetString(inputArea.Text)
			err := i.Run()
			if err == fifth.QuitError {
				myApp.Quit()
			}
			if err != nil {

			}
			out := stdout.String()
			outputArea.SetText(out)
		},
	}

	myWindow.SetContent(form)

	myWindow.ShowAndRun()
}
