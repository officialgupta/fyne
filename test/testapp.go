// Package test provides utility drivers for running UI tests without rendering
package test

import (
	"image/color"

	"github.com/fyne-io/fyne"
)

// ensure we have a dummy app loaded and ready to test
func init() {
	NewApp()
}

type testApp struct {
	driver   *testDriver
	settings fyne.Settings
}

func (a *testApp) Icon() fyne.Resource {
	return nil
}

func (a *testApp) SetIcon(fyne.Resource) {
	// no-op
}

func (a *testApp) NewWindow(title string) fyne.Window {
	return NewTestWindow(nil)
}

func (a *testApp) OpenURL(url string) {
	// no-op
}

func (a *testApp) Run() {
	// no-op
}

func (a *testApp) Quit() {
	// no-op
}

func (a *testApp) applyThemeTo(content fyne.CanvasObject, canvas fyne.Canvas) {
	if themed, ok := content.(fyne.ThemedObject); ok {
		themed.ApplyTheme()
		canvas.Refresh(content)
	}
	if wid, ok := content.(fyne.Widget); ok {
		// we cannot use the renderer cache as that is in the widget package (import loop)
		render := wid.CreateRenderer()
		render.ApplyTheme()
		canvas.Refresh(content)

		for _, o := range render.Objects() {
			a.applyThemeTo(o, canvas)
		}
	}
	if c, ok := content.(*fyne.Container); ok {
		for _, o := range c.Objects {
			a.applyThemeTo(o, canvas)
		}
	}
}

func (a *testApp) applyTheme() {
	for _, window := range a.driver.AllWindows() {
		content := window.Content()

		a.applyThemeTo(content, window.Canvas())
	}
}

func (a *testApp) Driver() fyne.Driver {
	return a.driver
}

func (a *testApp) Settings() fyne.Settings {
	return a.settings
}

// NewApp returns a new dummy app used for testing..
// It loads a test driver which creates a virtual window in memory for testing.
func NewApp() fyne.App {
	test := &testApp{driver: NewTestDriver().(*testDriver), settings: &testSettings{}}
	test.Settings().SetTheme(&dummyTheme{})
	fyne.SetCurrentApp(test)

	listener := make(chan fyne.Settings)
	test.Settings().AddChangeListener(listener)
	go func() {
		for {
			_ = <-listener
			test.applyTheme()
		}
	}()

	return test
}

type dummyTheme struct {
}

func (dummyTheme) BackgroundColor() color.Color {
	return color.White
}

func (dummyTheme) ButtonColor() color.Color {
	return color.Black
}

func (dummyTheme) TextColor() color.Color {
	return color.Black
}

func (dummyTheme) PrimaryColor() color.Color {
	return color.Black
}

func (dummyTheme) FocusColor() color.Color {
	return color.Black
}

func (dummyTheme) TextSize() int {
	return 1
}

func (dummyTheme) TextFont() fyne.Resource {
	return nil
}

func (dummyTheme) TextBoldFont() fyne.Resource {
	return nil
}

func (dummyTheme) TextItalicFont() fyne.Resource {
	return nil
}

func (dummyTheme) TextBoldItalicFont() fyne.Resource {
	return nil
}

func (dummyTheme) TextMonospaceFont() fyne.Resource {
	return nil
}

func (dummyTheme) Padding() int {
	return 1
}

func (dummyTheme) IconInlineSize() int {
	return 1
}

type testSettings struct {
	theme fyne.Theme

	changeListeners []chan fyne.Settings
}

func (s *testSettings) AddChangeListener(listener chan fyne.Settings) {
	s.changeListeners = append(s.changeListeners, listener)
}

func (s *testSettings) SetTheme(theme fyne.Theme) {
	s.theme = theme

	s.apply()
}

func (s *testSettings) Theme() fyne.Theme {
	return s.theme
}

func (s *testSettings) apply() {
	for _, listener := range s.changeListeners {
		go func(listener chan fyne.Settings) {
			listener <- s
		}(listener)
	}
}
