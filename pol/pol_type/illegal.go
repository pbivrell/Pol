package pol_type

func (i POL_Illegal) GetType() Type_identifier {
    return ILLEGAL
}

func (i POL_Illegal) String() string {
    return "ILLEGAL"
}

type POL_Illegal struct{}

func NewIllegal() POL_Illegal{
    return POL_Illegal{}
}
