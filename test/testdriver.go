package test

import "github.com/fyne-io/fyne"

type testDriver struct {
}

func (d *testDriver) CanvasForObject(fyne.CanvasObject) fyne.Canvas {
	// cheating as we only have a single test window
	return windows[0].Canvas()
}

func (d *testDriver) CreateWindow(string) fyne.Window {
	return NewTestWindow(nil)
}

func (d *testDriver) AllWindows() []fyne.Window {
	return windows
}

func (d *testDriver) RenderedTextSize(text string, size int, style fyne.TextStyle) fyne.Size {
	return fyne.NewSize(len(text)*size, size)
}

func (d *testDriver) Run() {
	// no-op
}

func (d *testDriver) Quit() {
	// no-op
}

// NewTestDriver sets up and registers a new dummy driver for test purpose
func NewTestDriver() fyne.Driver {
	driver := new(testDriver)
	// make a single dummy window for rendering tests
	NewTestWindow(nil)

	return driver
}
