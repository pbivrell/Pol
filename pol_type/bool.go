package pol_type

import "fmt"

func (b Pol_Bool) GetType() Type_identifier {
    return b.id
}

func (a Pol_Bool) Check(op string, b Pol_Type) {
    if(a.GetType() != b.GetType()){
        printError(fmt.Sprintf("left and right hand side of operator '%s' must be of the same type given: %s and %s",op,a.GetType(),b.GetType()))
    }
}

func UnGenericBool(pt Pol_Type) Pol_Bool {
    retVal, isInt:= pt.(Pol_Bool)
    if !isInt{
        printError(fmt.Sprintf("%s is not an Bool",pt.String()))
    }
    return retVal
}

func (b Pol_Bool) Convert(c Pol_Type) Pol_Type{
    if c.GetType() == STRING {
        return NewString(b.String())
    }

    return NewNothing()
}

func (b Pol_Bool) String() string{
    if b.value {
        return "true"
    }else{
        return "false"
    }
}

func (b Pol_Bool) GetValue() bool {
    return b.value
}

type Pol_Bool struct {
    id Type_identifier
    value bool
}

func NewBool(constant bool) Pol_Bool {
    return Pol_Bool{BOOL,constant}
}

func (b Pol_Bool) Hash() uint {
    if b.value{
        return 1231
    }else{
        return 1237
    }
}

func (b1 Pol_Bool) Add(b2 Pol_Type) Pol_Type {
    printError("Operator binary add '+' is undefined for type bool")
    return NewBool(false)
}

func (b1 Pol_Bool) Subtract(b2 Pol_Type) Pol_Type {
    printError("Operator binary subtract '-' is undefined for type bool")
    return NewBool(false)
}

func (b1 Pol_Bool) Multiply(b2 Pol_Type) Pol_Type {
    printError("Operator binary multiply '*' is undefined for type bool")
    return NewBool(false)
}

func (b1 Pol_Bool) Divide(b2 Pol_Type) Pol_Type{
    printError("Operator binary divide '/' is undefined for type bool")
    return NewBool(false)
}

func (b1 Pol_Bool) Exponent(b2 Pol_Type) Pol_Type{
    printError("Operator binary exponent '^' is undefined for type bool")
    return NewBool(false)
}

func (b1 Pol_Bool) Mod (b2 Pol_Type) Pol_Type {
    printError("Operator binary mod '%' is undefined for type bool")
    return NewBool(false)
}

func (b1 Pol_Bool) Equals(b2 Pol_Type) Pol_Bool {
    return NewBool(b1 == b2)
}

func (b1 Pol_Bool) NotEquals(b2 Pol_Type) Pol_Bool {
    return (b1.Equals(b2)).Not()
}

func (b1 Pol_Bool) Less(b2 Pol_Type) Pol_Bool {
    printError("Operator binary less '<' is undefined for type bool")
    return NewBool(false)
}

func (b1 Pol_Bool) Greater(b2 Pol_Type) Pol_Bool {
    printError("Operator binary greater '>' is undefined for type bool")
    return NewBool(false)
}

func (b1 Pol_Bool) LessEquals(b2 Pol_Type) Pol_Bool {
    printError("Operator binary lessEquals '<=' is undefined for type bool")
    return NewBool(false)
}

func (b1 Pol_Bool) GreaterEquals(b2 Pol_Type) Pol_Bool {
    printError("Operator binary greaterEquals '>=' is undefined for type bool")
    return NewBool(false)
}

func (b1 Pol_Bool) Not() Pol_Bool {
    return NewBool(!b1.value)
}

func (b1 Pol_Bool) And(b2 Pol_Type) Pol_Bool {
    b1.Check("||",b2)
    b3 := UnGenericBool(b2)
    return NewBool(b1.GetValue() && b3.GetValue())
}

func (b1 Pol_Bool) Or(b2 Pol_Type) Pol_Bool {
    b1.Check("||",b2)
    b3 := UnGenericBool(b2)
    return NewBool(b1.GetValue() || b3.GetValue())
}
