package pol_type

func (i Pol_Nothing) GetType() Type_identifier {
    return "NOTHING"
}

func (i Pol_Nothing) String() string {
    return ""
}

type Pol_Nothing struct {}

func NewNothing() Pol_Nothing {
    return Pol_Nothing{}
}

func (n Pol_Nothing) Convert(Pol_Type) Pol_Type {
    return NewNothing()
}

func (i Pol_Nothing) Hash() uint {
    return 0
}

func (i1 Pol_Nothing) Add(i2 Pol_Type) Pol_Type {
    printError("Operator not '+' is undefined for nothing")
    return NewNothing()

}

func (i1 Pol_Nothing) Subtract(i2 Pol_Type) Pol_Type {
    printError("Operator not '-' is undefined for nothing")
    return NewNothing()
}

func (i1 Pol_Nothing) Multiply(i2 Pol_Type) Pol_Type {
    printError("Operator not '*' is undefined for nothing")
    return NewNothing()

}

func (i1 Pol_Nothing) Divide(i2 Pol_Type) Pol_Type{
    printError("Operator not '/' is undefined for nothing")
    return NewNothing()
}

func (i1 Pol_Nothing) Exponent(i2 Pol_Type) Pol_Type{
    printError("Operator not '^' is undefined for nothing")
    return NewNothing()
}

func (i1 Pol_Nothing) Mod(i2 Pol_Type) Pol_Type{
    printError("Operator not '%' is undefined for nothing")
    return NewNothing()
}

func (i1 Pol_Nothing) Equals(i2 Pol_Type) Pol_Bool {
    printError("Operator not '==' is undefined for nothing")
    return NewBool(false)
}

func (i1 Pol_Nothing) NotEquals(i2 Pol_Type) Pol_Bool {
    printError("Operator not '!=' is undefined for nothing")
    return NewBool(false)
}

func (i1 Pol_Nothing) Less(i2 Pol_Type) Pol_Bool {
    printError("Operator not '<' is undefined for nothing")
    return NewBool(false)
}

func (i1 Pol_Nothing) Greater(i2 Pol_Type) Pol_Bool {
    printError("Operator not '>' is undefined for nothing")
    return NewBool(false)
}

func (i1 Pol_Nothing) LessEquals(i2 Pol_Type) Pol_Bool {
    printError("Operator not '<=' is undefined for nothing")
    return NewBool(false)
}

func (i1 Pol_Nothing) GreaterEquals(i2 Pol_Type) Pol_Bool {
    printError("Operator not '>=' is undefined for nothing")
    return NewBool(false)
}

func (i1 Pol_Nothing) Not() Pol_Bool {
    printError("Operator not '!' is undefined for nothing")
    return NewBool(false)
}

func (i1 Pol_Nothing) And(i2 Pol_Type) Pol_Bool {
    printError("Operator not '&&' is undefined for nothing")
    return NewBool(false)
}

func (i1 Pol_Nothing) Or(i2 Pol_Type) Pol_Bool {
    printError("Operator not '||' is undefined for an integer")
    return NewBool(false)
}
