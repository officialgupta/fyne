// Package fyne describes the objects and components available to any Fyne app.
// These can all be created, manipulated and tested without renering (for speed).
// Your main package should use the app package to create an application with
// a default driver that will render your UI.
//
// A simple application may look like this:
//
//   package main
//
//   import "github.com/fyne-io/fyne/app"
//   import "github.com/fyne-io/fyne/widget"
//
//   func main() {
//   	a := app.New()
//
//   	w := a.NewWindow("Hello")
//   	w.SetContent(widget.NewVBox(
//   		widget.NewLabel("Hello Fyne!"),
//   		widget.NewButton("Quit", func() {
//   			app.Quit()
//   		})))
//
//   	w.Show()
//   }
package fyne
