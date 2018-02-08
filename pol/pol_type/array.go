package pol_type
/*

import "fmt"

func (a Pol_Array) GetType() type_identifier {
    return a.id
}

func (a Pol_Array) String() string {
    retVal := "["
    for i, next := range a.data{
        if i == 0 {
            retVal += next.String()
        }else{
            retVal += "," + retVal
        }
    }
    return retVal + "]"
}

type Pol_Array struct {
    id type_identifier
    size int
    data []Pol_Type
}

func NewArray() *Pol_Array{
    return &Pol_Array{ARRAY,0,make([]Pol_Type,0)}
}

func (a Pol_Array) Resize(size int) {
    newSlice := make([]Pol_Type,size)
    copy(newSlice,a.data)
    a.data = newSlice
    a.size = size
}

func (a Pol_Array) Set(index Pol_Type, value Pol_Type){
    if index.GetType() != INT {
        fmt.Println("Attempting to index array with non-integer value")
    }
    var i Pol_Integer = index.value
    if i > a.size {
        a.Resize(i)
    }
    a.data[i] = value
}

func (a Pol_Array) Get(index Pol_Type) Pol_Type{
    if index.GetType() != INT {
        fmt.Println("Attempting to index array with non-integer value")
        return MakeNull()
    }
    i := index.value
    if i > a.size {
        return MakeNull()
    }
    return a.data[i]
}*/
