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

import "testing"

var optests = []struct {
	a      float64
	b      float64
	op     func(float64, float64) (float64, error)
	result float64
}{
	{0, 8, Plus, 8},
	{0, 8, Minus, -8},
	{2, 8, Multiply, 16},
	{4, 2, Divide, 2},
	{4, 2, Factorial, 24},
	{2, 2, Power, 4},
	{16, 4, NRoot, 2},
	{2, 2, Mod, 0},
}

func TestOperations(t *testing.T) {

	for _, tt := range optests {
		if res, err := tt.op(tt.a, tt.b); res != tt.result {
			t.Errorf("Expected %v, but got %v", tt.result, res)
		} else if err != nil {
			t.Error(err)
		}
	}
}

func TestDivisionByZero(t *testing.T) {
	if _, err := Divide(10, 0); err == nil {
		t.Error("Division by zero should be not possible")
	}
}

func TestFactorial(t *testing.T) {
	if _, err := Factorial(1.5, 8); err.Error() != "Factorial of non integer or negative value can't be calculated" {
		t.Error("Expected error when trying to factor non-interger or negative value")
	}

	if _, err := Factorial(-2, 8); err.Error() != "Factorial of non integer or negative value can't be calculated" {
		t.Error("Expected error when trying to factor non-interger or negative value")
	}

	if _, err := Factorial(26, 8); err.Error() != "Integer overflow" {
		t.Error("Expected integer overflow")
	}
}
