package main

// uncomment this with ButtonClicked()
import (
	gtk "github.com/mattn/go-gtk/gtk"
	"os"
	//"strconv"
	//"strings"
)

var (
	display   *gtk.Entry // for displaying values
	inputMode = false
	nums      = "789/!456x^123-e0.=+√"
	operators = "/!x^-e+√="
)

// End the program
func Quit() {
	gtk.MainQuit()
}

// on button click action, returns a handler function
/* TODO: implement these functions in backend:
                   Calculation() - set values to some result***--ˇˇ strucutre, 
				   Reset() - put system to initial state, 
				   GetResult() - returns string with the newest result
				    */
/*func ButtonClicked(b *gtk.Button) func() {
	return func() {
		if strings.Index(operators, b.GetLabel()) != -1 {
			val, _ := strconv.ParseFloat(display.GetText(), 32)
			Calculation(float32(val), b.GetLabel())
			display.SetText(GetResult())
			inputMode = false
		} else {
			if inputMode {
				display.SetText(display.GetText() + b.GetLabel())
			} else {
				display.SetText(b.GetLabel())
				inputMode = true
				// ***result --^^ structure to save actual computations info
				if result.operator == "=" {
					Reset()
				}
			}
		}
	}
}*/

func main() {
	gtk.Init(&os.Args)
	display = gtk.NewEntry()
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Calculator")
	window.Connect("destroy", Quit, nil)

	// Vertical box containing all components
	vbox := gtk.NewVBox(false, 1)

	// Calculator display as a vertical box
	display.SetCanFocus(false)
	display.SetText("0")
	display.SetAlignment(1.0) // align text to the right
	vbox.Add(display)

	// Vertical box containing all buttons
	buttons := gtk.NewVBox(false, 5)
	var b *gtk.Button

	for i := 0; i < 4; i++ {
		hbox := gtk.NewHBox(false, 5) // a horizontal box for each 4 buttons
		for j := 0; j < 5; j++ {
			if i*5+j != 19 {
				// this is an ugly hack as I didn't know how to use unicode from nums[] in this algorithm
				b = gtk.NewButtonWithLabel(string(nums[i*5+j]))
			} else {
				b = gtk.NewButtonWithLabel(string("\u221a"))
			}
			//b.Clicked(ButtonClicked(b), nil)
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
