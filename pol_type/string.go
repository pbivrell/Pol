package pol_type

import "fmt"

type Pol_String struct {
    id Type_identifier
    str string
}

func (s Pol_String) GetType() Type_identifier {
    return s.id
}

func (s Pol_String) String() string {
    return s.str
}

func (s Pol_String) Convert(c Pol_Type) Pol_Type{
    return NewNothing()
}


func UnGenericString(pt Pol_Type) Pol_String {
    retVal, isString:= pt.(Pol_String)
    if !isString{
        printError(fmt.Sprintf("%s is not an String",pt.String()))
    }
    return retVal
}


func NewString(constant string) Pol_String {
    retVal := Pol_String{STRING,constant}
    return retVal
}

func (s Pol_String) Hash() uint {
    var retVal uint = 0
    for _, next := range s.str {
        retVal += uint(next) * 31
    }
    return retVal
}

func (s1 Pol_String) Add(s2 Pol_Type) Pol_Type {
    s3 := UnGenericString(s2)
    return NewString(s1.str + s3.str)
}

func (s1 Pol_String) Subtract(s2 Pol_Type) Pol_Type {
    printError("Operator binary subtract '-' is undefined for type string")
    return NewString("")
}

func (s1 Pol_String) Multiply(s2 Pol_Type) Pol_Type {
    printError("Operator binary multiply '*' is undefined for type string")
    return NewString("")
}

func (s1 Pol_String) Divide(s2 Pol_Type) Pol_Type{
    printError("Operator binary divide '/' is undefined for type string")
    return NewString("")
}

func (s1 Pol_String) Exponent(s2 Pol_Type) Pol_Type{
    printError("Operator binary exponent '^' is undefined for type string")
    return NewString("")
}

func (s1 Pol_String) Mod(s2 Pol_Type) Pol_Type {
    printError("Operator binary mod '%' is undefined for type string")
    return NewString("")
}

func (s1 Pol_String) Equals(s2 Pol_Type) Pol_Bool {
    s3 := UnGenericString(s2)
    return NewBool(s1.str == s3.str)
}

func (s1 Pol_String) NotEquals(s2 Pol_Type) Pol_Bool {
    return (s1.Equals(s2)).Not()
}

func (s1 Pol_String) Less(s2 Pol_Type) Pol_Bool {
    printError("Operator binary less '<' is undefined for type string")
    return NewBool(false)
}

func (s1 Pol_String) Greater(s2 Pol_Type) Pol_Bool {
    printError("Operator binary greater '>' is undefined for type string")
    return NewBool(false)
}

func (s1 Pol_String) LessEquals(s2 Pol_Type) Pol_Bool {
    printError("Operator binary lessEquals '<=' is undefined for type string")
    return NewBool(false)
}

func (s1 Pol_String) GreaterEquals(s2 Pol_Type) Pol_Bool {
    printError("Operator binary greaterEquals '>=' is undefined for type string")
    return NewBool(false)
}

func (s1 Pol_String) Not() Pol_Bool {
    printError("Operator unary not '!' is undefined for type string")
    return NewBool(false)
}

func (s1 Pol_String) And(s2 Pol_Type) Pol_Bool {
    printError("Operator And '&&' is undefined for type string")
    return NewBool(false)
}

func (s1 Pol_String) Or(s2 Pol_Type) Pol_Bool {
    printError("Operator Or '||' is undefined for type string")
    return NewBool(false)
}
