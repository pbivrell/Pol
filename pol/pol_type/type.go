package pol_type

import "os"
import "fmt"

type Type_identifier string

const (
    INT Type_identifier = "INT"
    STRING = "STRING"
    RATIONAL = "RATIONAL"
    BOOL = "BOOL"
    HASH = "HASH"
    ARRAY = "ARRAY"
    ILLEGAL = "ILLEGAL"
    NULL = "NULL"
    NOTHING = "NOTHING"
)

func printError(msg string){
    fmt.Printf("[Runtime Error] %s\n",msg)
    os.Exit(8)
}

/*func (a Pol_Type) Check(op string, b Pol_Type) {
    if(a.GetType() != b.GetType()){
        fmt.Printf("[Runtime Error] left and right hand side of operator '%s' must be of the same type given: %s and %s",op,a.GetType(),b.GetType())
        os.Exit(-1)
    }
}*/

type Pol_Type interface{
    Convert(Pol_Type) Pol_Type
    GetType() Type_identifier
    String() string
    Add(Pol_Type) Pol_Type
    Subtract(Pol_Type) Pol_Type
    Multiply(Pol_Type) Pol_Type
    Divide(Pol_Type) Pol_Type
    Exponent(Pol_Type) Pol_Type
    Mod(Pol_Type) Pol_Type
    Equals(Pol_Type) Pol_Bool
    NotEquals(Pol_Type) Pol_Bool
    Not() Pol_Bool
    Greater(Pol_Type) Pol_Bool
    Less(Pol_Type) Pol_Bool
    GreaterEquals(Pol_Type) Pol_Bool
    LessEquals(Pol_Type) Pol_Bool
    Or(Pol_Type) Pol_Bool
    And(Pol_Type) Pol_Bool
    Hash() uint
}
