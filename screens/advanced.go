package screens

import (
	"fmt"

	"fyne.io/fyne"
	//"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func scaleString(c fyne.Canvas) string {
	return fmt.Sprintf("%0.2f", c.Scale())
}

func prependTo(g *widget.Group, s string) {
	g.Prepend(widget.NewLabel(s))
}

// AdvancedScreen loads a panel that shows details and settings that are a bit
// more detailed than normally needed.
func AdvancedScreen(win fyne.Window) fyne.CanvasObject {

	label := widget.NewLabel("Just type...")
	generic := widget.NewGroupWithScroller("Generic")
	desk := widget.NewGroupWithScroller("Desktop")

	return widget.NewHBox(
		widget.NewVBox(
			widget.NewButton("Custom Theme", func() {
				//fyne.CurrentApp().Settings().SetTheme(newCustomTheme())
			}),
			widget.NewButton("Fullscreen", func() {
				win.SetFullScreen(!win.FullScreen())
			}),
		),

		fyne.NewContainerWithLayout(layout.NewBorderLayout(label, nil, nil, nil),
			label,
			fyne.NewContainerWithLayout(layout.NewGridLayout(2),
				generic, desk,
			),
		))
}
