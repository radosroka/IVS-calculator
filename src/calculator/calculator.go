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


package calculator

import (
	"fmt"
)

// Defines operations of calculator
type Calculator interface {
	ClearAll()
	Execute(float64)
	Show()
	GetResult()
}

// Data structure to store result of calculation
type result struct {
	value float64
	err   error
}

// Stores structure containing current result and slot for operation to be performed
type SimpleCalc struct {
	result        result
	OperationSlot func(float64, float64) (float64, error)
}

// SimpleCalc constructor
func NewCalc() *SimpleCalc {
	c := new(SimpleCalc)
	c.result = result{
		value: 0,
		err:   nil,
	}
	c.OperationSlot = Plus
	return c
}

// Reset calculator to initial state
func (c *SimpleCalc) ClearAll() {
	c.result.value = 0
	c.result.err = nil
	c.OperationSlot = Plus
}

// Executes selected Operation
func (c *SimpleCalc) Execute(operand float64) {
	c.result.value, c.result.err = c.OperationSlot(c.result.value, operand)
}

// Shows result
func (c *SimpleCalc) Show() {
	fmt.Println(c.result)
}

// Returns current result and error code
func (c *SimpleCalc) GetResult() (float64, error) {
	return c.result.value, c.result.err
}
