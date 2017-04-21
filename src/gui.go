package main

import (
	"./calculator"
	"./mathlib"
	gdk "github.com/mattn/go-gtk/gdk"
	gtk "github.com/mattn/go-gtk/gtk"
	"os"
	"strconv"
	"strings"
)

var (
	display   *gtk.Entry // for displaying values
	inputMode = false
	nums      = "789/!456x%123-^0.=+√"
	operators = "/!x%-^+=\u221a√"
	// mapping operator character to function pointers
	oparation_map = map[string]func(float64, float64) (float64, error){
		"+": mathlib.Plus,
		"-": mathlib.Minus,
		"x": mathlib.Multiply,
		"/": mathlib.Divide,
		"!": mathlib.Factorial,
		"^": mathlib.Power,
		"√": mathlib.NRoot,
		"%": mathlib.Mod,
	}
)

// End the program
func Quit() {
	gtk.MainQuit()
}

// Helper function, execute current operation and show result
func executeOperation(c *calculator.SimpleCalc, val float64) {
	c.Execute(val)
	result, err := c.GetResult()
	if err != nil {
		display.SetText(err.Error())
	} else {
		display.SetText(strconv.FormatFloat(result, 'g', 6, 64))
	}

}

// on button click action, returns a handler function
func ButtonClicked(b *gtk.Button, c *calculator.SimpleCalc) func() {
	return func() {
		if strings.Index(operators, b.GetLabel()) != -1 {
			val, _ := strconv.ParseFloat(display.GetText(), 32)
			executeOperation(c, val)
			inputMode = false
			if strings.Compare(b.GetLabel(), "=") == 0 {
				c.ClearAll()
			} else {
				c.OperationSlot = oparation_map[b.GetLabel()]
				if strings.Compare(b.GetLabel(), "!") == 0 {
					executeOperation(c, 0)
				}
			}
		} else if strings.Compare(b.GetLabel(), "AC") == 0 {
			c.ClearAll()
			display.SetText("0")
			inputMode = false
		} else {
			if inputMode {
				display.SetText(display.GetText() + b.GetLabel())
			} else {
				display.SetText(b.GetLabel())
				inputMode = true
			}
		}
	}
}

func main() {
	gtk.Init(&os.Args)
	display = gtk.NewEntry()
	display.SetSizeRequest(300, 50)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Calculator")
	window.SetBorderWidth(10)
	window.SetDefaultSize(200, 200)
	window.ModifyBG(gtk.STATE_NORMAL, gdk.NewColor("grey"))
	window.Connect("destroy", Quit, nil)

	calc := calculator.NewCalc()

	// Vertical box containing all components
	vbox := gtk.NewVBox(false, 1)

	// Calculator display as a vertical box
	display.SetCanFocus(false)
	display.SetText("0")
	display.SetAlignment(1.0) // align text to the right
	vbox.Add(display)

	// Reset button
	additionalBox := gtk.NewHBox(false, 5)
	additionalBox.SetSizeRequest(40, 40)
	resetButton := gtk.NewButtonWithLabel("AC")
	resetButton.SetSizeRequest(30, 30)
	resetButton.Clicked(ButtonClicked(resetButton, calc), nil)
	vbox.Add(resetButton)

	// Vertical box containing all buttons
	buttons := gtk.NewVBox(false, 5)
	var b *gtk.Button

	for i := 0; i < 4; i++ {
		hbox := gtk.NewHBox(false, 5) // a horizontal box for each 4 buttons
		for j := 0; j < 5; j++ {
			if i*5+j != 19 {
				// this is an ugly hack as I didn't know how to use unicode from nums[] in this algorithm
				b = gtk.NewButtonWithLabel(string(nums[i*5+j]))
				b.SetSizeRequest(35, 35)
			} else {
				b = gtk.NewButtonWithLabel(string("\u221a"))
				b.SetSizeRequest(35, 35)
			}
			b.Clicked(ButtonClicked(b, calc), nil)
			hbox.Add(b)
		}
		buttons.Add(hbox) // add horizonatal box to the vertical buttons box
	}

	vbox.Add(buttons)

	window.Add(vbox)
	window.SetSizeRequest(250, 250)
	window.ShowAll()

	gtk.Main()
}
