package pol_type

import "testing"
import "pol_type"

type boolPair struct {
	bool1 bool
	bool2 bool
}

//Negative number constants don't exist
var TestsBool = []boolPair{
	{true,false},
	{true,true},
	{false,false},
	{false,true},
}

func TestBoolEquals(t *testing.T) {
	for _, pair := range TestsBool{
		real_answer := pol_type.NewBool(pair.bool1 == pair.bool2)
		lhs := pol_type.NewBool(pair.bool1)
		rhs := pol_type.NewBool(pair.bool2)
		my_answer := lhs.Equals(rhs)
		if  real_answer != my_answer {
			t.Error(
				"For", pair.bool1, "==", pair.bool2,
				"expected", real_answer,
				"got", my_answer,
			)
		}
	}
}

func TestBoolNotEquals(t *testing.T) {
	for _, pair := range TestsBool{
		real_answer := pol_type.NewBool(pair.bool1 != pair.bool2)
		lhs := pol_type.NewBool(pair.bool1)
		rhs := pol_type.NewBool(pair.bool2)
		my_answer := lhs.NotEquals(rhs)
		if  real_answer != my_answer {
			t.Error(
				"For", pair.bool1, "==", pair.bool2,
				"expected", real_answer,
				"got", my_answer,
			)
		}
	}
}
