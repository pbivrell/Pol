package interpreter

import "fmt"
//import "../common"
import "../pol_type"


type Scope struct {
    vars map[string]pol_type.Pol_Type
    prev *Scope
}

func MakeScope() *Scope{
    return &Scope{make(map[string]pol_type.Pol_Type),nil}
}

func (s *Scope) Dump() []pol_type.Pol_Type {
    i := 'a'
    ret := make([]pol_type.Pol_Type,0)
    for _,_ = range scope.vars {
        if val, hasValue := s.vars[string(i) + "_pol_variable"]; hasValue{
            ret = append(ret,val)
            i++
        }
    }
    return ret
}

func (s *Scope) Enter() *Scope {
    newScope := MakeScope()
    for key, value := range s.vars{
        newScope.vars[key] = value
    }
    newScope.prev = s
    return newScope
}

func (s *Scope) Exit() *Scope {
    if s.prev == nil {
        printError("Popping Empty Stack!")
    }
    return s.prev
}

func (s *Scope) Call() (*Scope, *Scope) {
    oldScope := *s
    globalScope := s
    for globalScope.prev != nil {
        globalScope = globalScope.Exit()
    }
    return globalScope.Enter(), &oldScope
}

func (s *Scope) Assign(identifer string, value pol_type.Pol_Type) {
    s.vars[identifer] = value
}

func (s *Scope) ValueOf(where string) pol_type.Pol_Type{
    return s.vars[where]
}

func (s *Scope) LookUp(where string, key pol_type.Pol_Type) pol_type.Pol_Type{
    item, hasItem := s.vars[where]
    if !hasItem{
        fmt.Println("Item was not found")
        return pol_type.NewNothing()
    }
    hash := pol_type.UnGenericHash(item)
    return (&hash).Get(key)
}

func (s *Scope) AssignCollection(where string, key pol_type.Pol_Type, value pol_type.Pol_Type){
    _, hasItem := s.vars[where]

    if !hasItem{
        s.vars[where] = pol_type.NewHash()
    }

    hash := pol_type.UnGenericHash(s.vars[where])
    (&hash).Set(key,value)
}
