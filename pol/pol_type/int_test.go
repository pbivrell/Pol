package pol_type

import "testing"
import "pol_type"

type intPair struct {
	int1 int
	int2 int
}

//Negative number constants don't exist
var TestsInt = []intPair{
	{1,1},
	{1,-1},
	{0,1},
	{1,0},
	{-1,0},
	{0,-1},
	{-1, -1},
	{8,2},
	{2,8},
}

func TestIntAdd(t *testing.T) {
	for _, pair := range TestsInt{
		real_answer := pol_type.NewIntFromInt(pair.int1 + pair.int2)
		lhs := pol_type.NewIntFromInt(pair.int1)
		rhs := pol_type.NewIntFromInt(pair.int2)
		my_answer := lhs.Add(rhs)
		if  real_answer != my_answer {
			t.Error(
				"For", pair.int1, "+", pair.int2,
				"expected", real_answer,
				"got", my_answer,
			)
		}
	}
}

func TestIntSubtract(t *testing.T) {
	for _, pair := range TestsInt{
		real_answer := pol_type.NewIntFromInt(pair.int1 - pair.int2)
		lhs := pol_type.NewIntFromInt(pair.int1)
		rhs := pol_type.NewIntFromInt(pair.int2)
		my_answer := lhs.Subtract(rhs)
		if  real_answer != my_answer {
			t.Error(
				"For", pair.int1, "-", pair.int2,
				"expected", real_answer,
				"got", my_answer,
			)
		}
	}
}


func TestIntMultiply(t *testing.T) {
	for _, pair := range TestsInt{
		real_answer := pol_type.NewIntFromInt(pair.int1 * pair.int2)
		lhs := pol_type.NewIntFromInt(pair.int1)
		rhs := pol_type.NewIntFromInt(pair.int2)
		my_answer := lhs.Multiply(rhs)
		if  real_answer != my_answer {
			t.Error(
				"For", pair.int1, "*", pair.int2,
				"expected", real_answer,
				"got", my_answer,
			)
		}
	}
}


func TestIntDivide(t *testing.T) {
	for _, pair := range TestsInt{

		//Skip divide by 0 tests as they throw runtime errors in both languages
		if pair.int2 == 0 {
			continue
		}
		real_answer := pol_type.NewIntFromInt(pair.int1 / pair.int2)
		lhs := pol_type.NewIntFromInt(pair.int1)
		rhs := pol_type.NewIntFromInt(pair.int2)
		my_answer := lhs.Divide(rhs)
		if  real_answer != my_answer {
			t.Error(
				"For", pair.int1, "/", pair.int2,
				"expected", real_answer,
				"got", my_answer,
			)
		}
	}
}

func TestIntMod(t *testing.T) {
	for _, pair := range TestsInt{
		//Skip divide by 0 tests as they throw runtime errors in both languages
		if pair.int2 == 0 {
			continue
		}
		real_answer := pol_type.NewIntFromInt(pair.int1 % pair.int2)
		lhs := pol_type.NewIntFromInt(pair.int1)
		rhs := pol_type.NewIntFromInt(pair.int2)
		my_answer := lhs.Mod(rhs)
		if  real_answer != my_answer {
			t.Error(
				"For", pair.int1, "%", pair.int2,
				"expected", real_answer,
				"got", my_answer,
			)
		}
	}
}

func TestIntExponent(t *testing.T) {
	for _, pair := range TestsInt{
		temp := pair.int1
		//Skip raising to a negative power tests as they throw runtime errors in pol
		if pair.int2 < 0 {
			continue
		//Anything to 0th power is 1 in pol
		}else if pair.int2 == 0 {
			temp = 1
		}

		//Manually compute int1 ^ int2 because go doesn't support raising ints to powers
		for i := 0; i < pair.int2-1; i++{
			temp *= pair.int1
		}

		real_answer := pol_type.NewIntFromInt(temp)
		lhs := pol_type.NewIntFromInt(pair.int1)
		rhs := pol_type.NewIntFromInt(pair.int2)
		my_answer := lhs.Exponent(rhs)
		if  real_answer != my_answer {
			t.Error(
				"For", pair.int1, "^", pair.int2,
				"expected", real_answer,
				"got", my_answer,
			)
		}
	}
}


func TestIntEquals(t *testing.T) {
	for _, pair := range TestsInt{
		real_answer := pol_type.NewBool(pair.int1 == pair.int2)
		lhs := pol_type.NewIntFromInt(pair.int1)
		rhs := pol_type.NewIntFromInt(pair.int2)
		my_answer := lhs.Equals(rhs)
		if  real_answer.GetValue() != my_answer.GetValue() {
			t.Error(
				"For", pair.int1, "==", pair.int2,
				"expected", real_answer,
				"got", my_answer,
			)
		}
	}
}

func TestIntNotEquals(t *testing.T) {
	for _, pair := range TestsInt{
		real_answer := pol_type.NewBool(pair.int1 == pair.int2)
		lhs := pol_type.NewIntFromInt(pair.int1)
		rhs := pol_type.NewIntFromInt(pair.int2)
		my_answer := lhs.NotEquals(rhs)
		if  real_answer.GetValue() == my_answer.GetValue() {
			t.Error(
				"For", pair.int1, "!=", pair.int2,
				"expected", real_answer,
				"got", my_answer,
			)
		}
	}
}

func TestIntLess(t *testing.T) {
	for _, pair := range TestsInt{
		real_answer := pol_type.NewBool(pair.int1 < pair.int2)
		lhs := pol_type.NewIntFromInt(pair.int1)
		rhs := pol_type.NewIntFromInt(pair.int2)
		my_answer := lhs.Less(rhs)
		if  real_answer.GetValue() != my_answer.GetValue() {
			t.Error(
				"For", pair.int1, "<", pair.int2,
				"expected", real_answer,
				"got", my_answer,
			)
		}
	}
}

func TestIntGreater(t *testing.T) {
	for _, pair := range TestsInt{
		real_answer := pol_type.NewBool(pair.int1 > pair.int2)
		lhs := pol_type.NewIntFromInt(pair.int1)
		rhs := pol_type.NewIntFromInt(pair.int2)
		my_answer := lhs.Greater(rhs)
		if  real_answer.GetValue() != my_answer.GetValue() {
			t.Error(
				"For", pair.int1, ">", pair.int2,
				"expected", real_answer,
				"got", my_answer,
			)
		}
	}
}

func TestIntLessEquals(t *testing.T) {
	for _, pair := range TestsInt{
		real_answer := pol_type.NewBool(pair.int1 <= pair.int2)
		lhs := pol_type.NewIntFromInt(pair.int1)
		rhs := pol_type.NewIntFromInt(pair.int2)
		my_answer := lhs.LessEquals(rhs)
		if  real_answer.GetValue() != my_answer.GetValue() {
			t.Error(
				"For", pair.int1, "<=", pair.int2,
				"expected", real_answer,
				"got", my_answer,
			)
		}
	}
}

func TestIntGreterEqual(t *testing.T) {
	for _, pair := range TestsInt{
		real_answer := pol_type.NewBool(pair.int1 >= pair.int2)
		lhs := pol_type.NewIntFromInt(pair.int1)
		rhs := pol_type.NewIntFromInt(pair.int2)
		my_answer := lhs.GreaterEquals(rhs)
		if  real_answer.GetValue() != my_answer.GetValue() {
			t.Error(
				"For", pair.int1, ">=", pair.int2,
				"expected", real_answer,
				"got", my_answer,
			)
		}
	}
}
