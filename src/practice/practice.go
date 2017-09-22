package main

import "fmt"

func isP(i int) bool{
    return i > 0
}

func main(){
    i := 8
    for char := isP(i); char; char = isP(i) {
        fmt.Println("TEST")
        i = i -1
    }
}
