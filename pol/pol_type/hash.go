package pol_type

import "fmt"

const LOAD_FACTOR = 0.8

type Pol_Hash struct {
    id Type_identifier
    size int
    count int
    collisions int
    buckets []*bucket
}

type bucket struct {
    key Pol_Type
    value Pol_Type
    next *bucket
}

/*
func (h *Pol_Hash) Set(key Pol_Type, value Pol_Type){
    index := int(key.Hash()) % h.size
    b := h.buckets[index]
    b.value = value
    for ; b.next != nil; b = b.next {
        if b.key == key {
            b.value = value
            return
        }
    }

}

func (h *Pol_Hash) Get(key Pol_Type) Pol_Type {
    index := int(key.Hash()) % h.size
    return h.buckets[index].value
}
*/


func (h *Pol_Hash) Set(key Pol_Type, value Pol_Type){
    //fmt.Println("SETTING Item: ", key.String())
    index := int(key.Hash()) % h.size
    //fmt.Println("Hashed Value: ", (int(key.Hash())), " ", h.size)
    b := &h.buckets[index]

    for ; *b != nil; {
    //    fmt.Println("Search for item ", key, " ", (*b).key)
        if (*b).key == key {
    //        fmt.Println("Found item")
            (*b).value = value
            return
        }
        b = &((*b).next);
    }
    //fmt.Println("Did not find item")
    *b = &bucket{key,value,nil}
    //fmt.Println("Set bucket to: ")
    //fmt.Println(h.buckets[index])
    h.count++
    if float64(h.count) / float64(h.size) > LOAD_FACTOR {
        h.Resize()
    }
}

func (h *Pol_Hash) Get(key Pol_Type) Pol_Type{
    //fmt.Println("Getting Item: ", key.String())
    index := int(key.Hash()) % h.size
    //fmt.Println("Hashed Value: ", (int(key.Hash())), " ", h.size)
    b := h.buckets[index]
    //fmt.Println("Bucket value: ", b.value)

    for ; b != nil; b = b.next {
        //fmt.Println("Search for item ", key, " ", b.key)
        if  b.key == key {
            return b.value
        }
    }
    return NewNothing()
}


func (s Pol_Hash) Convert(c Pol_Type) Pol_Type{
    return NewNothing()
}

func (h Pol_Hash) Resize() {
    printError("Congrats you reached the maxium size of a hash table. Currently growing hashes is not implemented. Your program will now stop")
}

func (h Pol_Hash) Size() int {
    return h.size
}

func (h Pol_Hash) GetType() Type_identifier {
    return h.id
}

func (a Pol_Hash) Check(op string, b Pol_Hash) {
    if(a.GetType() != b.GetType()){
        printError(fmt.Sprintf("left and right hand side of operator '%s' must be of the same type given: %s and %s",op,a.GetType(),b.GetType()))
    }
}

func UnGenericHash(pt Pol_Type) Pol_Hash {
    retVal, isHash:= pt.(Pol_Hash)
    if !isHash{
        printError(fmt.Sprintf("%s is not an Hash",pt.String()))
    }
    return retVal
}

func (h Pol_Hash) String() string{
    return "{ HASH }"
}

func NewHash() Pol_Hash {
    hash := make([]*bucket,100)
    /*for i := 0; i < 100; i++ {
        hash[i] = &bucket{}
    }*/
    return Pol_Hash{HASH,100,0,0, hash}
}

func (h Pol_Hash) Hash() uint {
    return 1237
}

func (b1 Pol_Hash) Add(b2 Pol_Type) Pol_Type {
    printError("Operator binary add '+' is undefined for type hash")
    return NewNothing()
}

func (b1 Pol_Hash) Subtract(b2 Pol_Type) Pol_Type {
    printError("Operator binary subtract '-' is undefined for type hash")
    return NewNothing()
}

func (b1 Pol_Hash) Multiply(b2 Pol_Type) Pol_Type {
    printError("Operator binary multiply '*' is undefined for type hash")
    return NewNothing()
}

func (b1 Pol_Hash) Divide(b2 Pol_Type) Pol_Type{
    printError("Operator binary divide '/' is undefined for type hash")
    return NewNothing()
}

func (b1 Pol_Hash) Exponent(b2 Pol_Type) Pol_Type{
    printError("Operator binary exponent '^' is undefined for type hash")
    return NewNothing()
}

func (b1 Pol_Hash) Mod(b2 Pol_Type) Pol_Type {
    printError("Operator binary mod '%' is undefined for type hash")
    return NewNothing()
}

func (b1 Pol_Hash) Equals(b2 Pol_Type) Pol_Bool {
    printError("Operator binary equals '==' is undefined for type hash")
    return NewBool(false)
}

func (b1 Pol_Hash) NotEquals(b2 Pol_Type) Pol_Bool {
    return NewBool(false)
}

func (b1 Pol_Hash) Less(b2 Pol_Type) Pol_Bool {
    printError("Operator binary less '<' is undefined for type hash")
    return NewBool(false)
}

func (b1 Pol_Hash) Greater(b2 Pol_Type) Pol_Bool {
    printError("Operator binary greater '>' is undefined for type hash")
    return NewBool(false)
}

func (b1 Pol_Hash) LessEquals(b2 Pol_Type) Pol_Bool {
    printError("Operator binary lessEquals '<=' is undefined for type hash")
    return NewBool(false)
}

func (b1 Pol_Hash) GreaterEquals(b2 Pol_Type) Pol_Bool {
    printError("Operator binary greaterEquals '>=' is undefined for type hash")
    return NewBool(false)
}

func (b1 Pol_Hash) Not() Pol_Bool {
    printError("Operator binary greaterEquals '!' is undefined for type hash")
    return NewBool(false)
}

func (b1 Pol_Hash) And(b2 Pol_Type) Pol_Bool {
    printError("Operator binary greaterEquals '&&' is undefined for type hash")
    return NewBool(false)
}

func (b1 Pol_Hash) Or(b2 Pol_Type) Pol_Bool {
    printError("Operator binary or '||' is undefined for type hash")
    return NewBool(false)
}
