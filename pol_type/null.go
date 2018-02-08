package pol_type

func (n Pol_Null) GetType() Type_identifier {
    return NULL
}

func (n Pol_Null) String() string{
    return "NULL"
}

type Pol_Null struct {}

func NewNull() *Pol_Null {
    return &Pol_Null{}
}
