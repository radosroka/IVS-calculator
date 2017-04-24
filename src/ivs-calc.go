/* IVS-calculator
 * Copyright (C) 2017	Radovan Sroka <xsroka00@stud.fit.vutbr.cz>
 * 						Tomáš Sýkora <xsykor25@stud.fit.vutbr.cz>
 *						Michal Cyprian <xcypri01@stud.fit.vutbr.cz>
 *						Jan Mochnak <xmochn00@stud.fit.vutbr.cz>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"calculator"
	gdk "github.com/radosroka/go-gtk/gdk"
	gtk "github.com/radosroka/go-gtk/gtk"
	"gui"
	"os"
	"strings"
)

var (
	nums      = "789/!456x%123-^0.=+√"
	operators = "/!x%-^+=\u221a√"
)

func main() {
	gtk.Init(&os.Args)
	gui.Display = gtk.NewEntry()
	gui.Display.SetSizeRequest(300, 50)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Calculator")
	window.SetBorderWidth(10)
	window.SetDefaultSize(300, 400)
	window.ModifyBG(gtk.STATE_NORMAL, gdk.NewColor("grey"))
	window.Connect("destroy", gui.Quit, nil)

	calc := calculator.NewCalc()

	// Vertical box containing all components
	vbox := gtk.NewVBox(false, 1)

	// Calculator gui.Display as a vertical box
	gui.Display.SetCanFocus(false)
	gui.Display.SetText("0")
	gui.Display.SetAlignment(1.0) // align text to the right
	vbox.Add(gui.Display)

	// Reset button
	additionalBox := gtk.NewHBox(false, 5)
	additionalBox.SetSizeRequest(40, 40)
	resetButton := gtk.NewButtonWithLabel("AC")
	resetButton.SetSizeRequest(30, 30)
	resetButton.Clicked(gui.SpecialButtonClicked(resetButton, calc), nil)
	vbox.Add(resetButton)
	plusMinusButton := gtk.NewButtonWithLabel("+/-")
	plusMinusButton.SetSizeRequest(30, 30)
	plusMinusButton.Clicked(gui.SpecialButtonClicked(plusMinusButton, calc), nil)
	vbox.Add(plusMinusButton)

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
			if strings.Compare(b.GetLabel(), "!") == 0 {
				b.Clicked(gui.SpecialButtonClicked(b, calc), nil)
			} else if strings.Index(operators, string(nums[i*5+j])) != -1 {
				b.Clicked(gui.OperatorButtonClicked(b, calc), nil)
			} else {
				b.Clicked(gui.InputButtonClicked(b, calc), nil)
			}
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
