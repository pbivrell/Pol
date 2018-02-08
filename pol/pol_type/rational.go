package pol_type

import "strconv"
import "fmt"

func (r Pol_Rational) GetType() Type_identifier {
    return r.id
}

func (r Pol_Rational) String() string {
    return fmt.Sprintf("%f",r.num)
}

func (r Pol_Rational) Convert(c Pol_Type) Pol_Type{
    if c.GetType() == STRING {
        return NewString(r.String())
    }

    return NewNothing()
}

type Pol_Rational struct {
    id Type_identifier
    num float64
}

func NewRational(constant string) Pol_Rational {
    i, _ := strconv.ParseFloat(constant,64)
    return Pol_Rational{RATIONAL,i}
}

func (r Pol_Rational) Hash() uint{
    return 37
}

func UnGenericRational(pt Pol_Type) Pol_Rational {
    retVal, isRat:= pt.(Pol_Rational)
    if !isRat{
        printError(fmt.Sprintf("%s is not an Rational",pt.String()))
    }
    return retVal
}

func (r1 Pol_Rational) Add(r2 Pol_Type) Pol_Type {
    r3 := UnGenericRational(r2)
    return Pol_Rational{RATIONAL,(r1.num + r3.num)}
}

func (r1 Pol_Rational) Subtract(r2 Pol_Type) Pol_Type {
    r3 := UnGenericRational(r2)
    return Pol_Rational{RATIONAL,(r1.num - r3.num)}
}

func (r1 Pol_Rational) Multiply(r2 Pol_Type) Pol_Type {
    r3 := UnGenericRational(r2)
    return Pol_Rational{RATIONAL,(r1.num * r3.num)}
}

func (r1 Pol_Rational) Divide(r2 Pol_Type) Pol_Type{
    r3 := UnGenericRational(r2)
    return Pol_Rational{RATIONAL,(r1.num / r3.num)}
}

func (r1 Pol_Rational) Exponent(r2 Pol_Type) Pol_Type{
    fmt.Println("Operator binary exponent '^' is undefined for type string")
    return NewRational("1.0")
}

func (r1 Pol_Rational) Mod (r2 Pol_Type) Pol_Type {
    fmt.Println("Operator binary mod '%' is undefined for type string")
    return NewRational("1.0")
}

func (r1 Pol_Rational) Equals(r2 Pol_Type) Pol_Bool {
    return NewBool(r1 == r2)
}

func (r1 Pol_Rational) NotEquals(r2 Pol_Type) Pol_Bool {
    return (r1.Equals(r2)).Not()
}

func (r1 Pol_Rational) Less(r2 Pol_Type) Pol_Bool {
    fmt.Println("Operator binary less '<' is undefined for type string")
    return NewBool(true)
}

func (r1 Pol_Rational) Greater(r2 Pol_Type) Pol_Bool {
    fmt.Println("Operator binary greater '>' is undefined for type string")
    return NewBool(true)
}

func (r1 Pol_Rational) LessEquals(r2 Pol_Type) Pol_Bool {
    fmt.Println("Operator binary lessEquals '<=' is undefined for type string")
    return NewBool(true)
}

func (r1 Pol_Rational) GreaterEquals(r2 Pol_Type) Pol_Bool {
    fmt.Println("Operator binary greaterEquals '>=' is undefined for type string")
    return NewBool(true)
}

func (r1 Pol_Rational) Not() Pol_Bool {
    fmt.Println("Operator unary not '!' is undefined for type string")
    return NewBool(true)

}

func (r1 Pol_Rational) And(Pol_Type) Pol_Bool{
    return NewBool(true)
}

func (r1 Pol_Rational) Or(Pol_Type) Pol_Bool{
    return NewBool(true)
}
