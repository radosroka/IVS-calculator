package gui

import (
	"calculator"
	gtk "github.com/mattn/go-gtk/gtk"
	"strconv"
	"strings"
)

var (
	Display       *gtk.Entry // for Displaying values
	inputMode     = false
	lastOperation = "+"
	lastButton    = "."
	// mapping operator character to function pointers
	operationMap = map[string]func(float64, float64) (float64, error){
		"+": calculator.Plus,
		"-": calculator.Minus,
		"x": calculator.Multiply,
		"/": calculator.Divide,
		"!": calculator.Factorial,
		"^": calculator.Power,
		"√": calculator.NRoot,
		"%": calculator.Mod,
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
		Display.SetText("Error: " + err.Error())
		c.ClearAll()
		lastOperation = "+"
	} else {
		Display.SetText(strconv.FormatFloat(result, 'g', 6, 64))
	}

}

// Click action for operation buttons, returns a handler function
func OperatorButtonClicked(b *gtk.Button, c *calculator.SimpleCalc) func() {
	return func() {
		val, err := strconv.ParseFloat(Display.GetText(), 32)
		if err != nil {
			Display.SetText("Error: Invalid input")
			c.ClearAll()
			lastOperation = "+"
		}
		if lastOperation != "!" {
			executeOperation(c, val)
		}
		inputMode = false
		if strings.Compare(b.GetLabel(), "=") == 0 {
			c.ClearAll()
			lastOperation = "+"
		} else {
			lastOperation = b.GetLabel()
			c.OperationSlot = operationMap[lastOperation]
		}
	}
}

// Click action for input buttons, returns a handler function
func InputButtonClicked(b *gtk.Button, c *calculator.SimpleCalc) func() {
	return func() {
		if lastButton != "." || strings.Compare(b.GetLabel(), ".") != 0 {
			lastButton = b.GetLabel()
			if inputMode {
				Display.SetText(Display.GetText() + b.GetLabel())
			} else {
				Display.SetText(b.GetLabel())
				inputMode = true
			}
		}
	}
}

// Click action for special buttons, returns a handler function
func SpecialButtonClicked(b *gtk.Button, c *calculator.SimpleCalc) func() {
	return func() {
		if strings.Compare(b.GetLabel(), "AC") == 0 {
			c.ClearAll()
			lastOperation = "+"
			Display.SetText("0")
			inputMode = false
		} else if strings.Compare(b.GetLabel(), "+/-") == 0 {
			content := Display.GetText()
			if string([]rune(content)[0]) == "-" {
				content = content[1:len(content)]
			} else {
				content = "-" + content
			}
			Display.SetText(content)
		} else if strings.Compare(b.GetLabel(), "!") == 0 {
			val, err := strconv.ParseFloat(Display.GetText(), 32)
			if err != nil {
				Display.SetText("Error: Invalid input")
				c.ClearAll()
				lastOperation = "+"
			}
			if lastOperation != "!" {
				executeOperation(c, val)
			}

			inputMode = false
			lastOperation = b.GetLabel()
			c.OperationSlot = operationMap[lastOperation]
			executeOperation(c, 0)
			c.ClearAll()
			lastOperation = "+"
		}
	}
}
