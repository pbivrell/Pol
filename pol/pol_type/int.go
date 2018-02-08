package pol_type

import "strconv"
import "fmt"

func (i Pol_Integer) GetType() Type_identifier {
    return i.id
}

func (i Pol_Integer) String() string {
    ret := ""
    if i.negative {
        ret += "-"
    }
    ret += strconv.FormatUint(i.value,10)
    return ret
}

func (i Pol_Integer) ToString() Pol_String{
    return NewString(i.String())
}

func (i Pol_Integer) ToRational() Pol_Rational{
    return NewRational(i.String())
}

func (i Pol_Integer) Convert(c Pol_Type) Pol_Type{
    if c.GetType() == STRING {
        return NewString(i.String())
    }else if c.GetType() == RATIONAL {
        return NewRational(i.String())
    }

    return NewNothing()
}


type Pol_Integer struct {
    id Type_identifier
    value uint64
    negative bool
}

func (a Pol_Integer) Check(op string, b Pol_Type) {
    if(a.GetType() != b.GetType()){
        printError(fmt.Sprintf("left and right hand side of operator '%s' must be of the same type given: %s and %s",op,a.GetType(),b.GetType()))
    }
}

func NewInt(constant string) Pol_Integer {
    conv,_ := strconv.Atoi(constant)
    isNegative := false
    if conv < 0{
        isNegative = true
        conv = 0 - conv
    }
    ret := Pol_Integer{INT,uint64(conv),isNegative}
    //fmt.Println(ret.String())
    return ret
}

func NewIntFromInt(value int) Pol_Integer {
    isNegative := false
    if value < 0{
        isNegative = true
        value = 0 - value
    }
    return Pol_Integer{INT,uint64(value),isNegative}
}

func (i Pol_Integer) makeIntoInt() int {
    if i.negative {
        return int(i.value) * -1
    }
    return int(i.value)
}

func (i Pol_Integer) Hash() uint {
    return uint(i.value)
}

func unGeneric(pt Pol_Type) Pol_Integer {
    retVal, isInt:= pt.(Pol_Integer)
    if !isInt{
        printError(fmt.Sprintf("%s is not an Int",pt.String()))
    }
    return retVal
}

func (i1 Pol_Integer) Add(i2 Pol_Type) Pol_Type {
    i1.Check("+",i2)
    i3 := unGeneric(i2)
    a := i1.makeIntoInt()
    b := i3.makeIntoInt()
    return NewIntFromInt(a + b)
}

func (i1 Pol_Integer) Subtract(i2 Pol_Type) Pol_Type {
    i1.Check("-",i2)
    i3 := unGeneric(i2)
    i3.negative = !i3.negative
    return i1.Add(i3)
}

func (i1 Pol_Integer) Multiply(i2 Pol_Type) Pol_Type {
    i1.Check("*",i2)
    i3 := unGeneric(i2)
    a := i1.makeIntoInt()
    b := i3.makeIntoInt()
    return NewIntFromInt(a * b)
}

func (i1 Pol_Integer) Divide(i2 Pol_Type) Pol_Type{
    i1.Check("/",i2)
    i3 := unGeneric(i2)
    a := i1.makeIntoInt()
    b := i3.makeIntoInt()
    return NewIntFromInt(a/b)
}

func (i1 Pol_Integer) Exponent(i2 Pol_Type) Pol_Type{
    //printError("Operator not '!' is undefined for an integer")
    //return NewInt("1")
    i1.Check("^",i2)
    i3 := unGeneric(i2)
    var res Pol_Type = i1
    if i3.negative{
        printError("Unable to raise an integer to a negative power at ")
        return NewIntFromInt(0)
    }else if i3.value == 0 {
        return NewIntFromInt(1)
    }else{
        for i := 0; i < i3.makeIntoInt()-1; i++{
            res = i1.Multiply(res)
        }
        return i1
    }
}

func (i1 Pol_Integer) Mod(i2 Pol_Type) Pol_Type{
    i1.Check("/",i2)
    i3 := unGeneric(i2)
    a := i1.makeIntoInt()
    b := i3.makeIntoInt()
    return NewIntFromInt(a%b)
}

func (i1 Pol_Integer) Equals(i2 Pol_Type) Pol_Bool {
    i1.Check("==",i2)
    i3 := unGeneric(i2)
    return NewBool(i1.value == i3.value && i1.negative == i3.negative)
}

func (i1 Pol_Integer) NotEquals(i2 Pol_Type) Pol_Bool {
    return (i1.Equals(i2)).Not()
}

func (i1 Pol_Integer) Less(i2 Pol_Type) Pol_Bool {
    i1.Check("<",i2)
    i3 := unGeneric(i2)
    a := i1.makeIntoInt()
    b := i3.makeIntoInt()
    return NewBool(a < b)
}

func (i1 Pol_Integer) Greater(i2 Pol_Type) Pol_Bool {
    i1.Check(">",i2)
    i3 := unGeneric(i2)
    a := i1.makeIntoInt()
    b := i3.makeIntoInt()
    return NewBool(a > b)

}

func (i1 Pol_Integer) LessEquals(i2 Pol_Type) Pol_Bool {
    return (i1.Greater(i2)).Not()
}

func (i1 Pol_Integer) GreaterEquals(i2 Pol_Type) Pol_Bool {
    return (i1.Less(i2)).Not()
}

func (i1 Pol_Integer) Not() Pol_Bool {
    printError("Operator not '!' is undefined for an integer")
    return NewBool(false)
}

func (i1 Pol_Integer) And(i2 Pol_Type) Pol_Bool {
    printError("Operator not '&&' is undefined for an integer")
    return NewBool(false)
}

func (i1 Pol_Integer) Or(i2 Pol_Type) Pol_Bool {
    printError("Operator not '||' is undefined for an integer")
    return NewBool(false)
}
