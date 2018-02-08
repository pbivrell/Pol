package common

//import "time"
import "fmt"
import "os"

var errorBuffer = make([]string,0)

func GetErrors() string{
    retVal := ""
    for _, error := range errorBuffer {
        retVal += error + "\n"
    }
    errorBuffer = errorBuffer[0:0]
    return retVal
}

func HasErrors() bool{
    return len(errorBuffer) != 0
}

func AppendError(error string){
    errorBuffer = append(errorBuffer,error)
}

func RuntimeError(error string){
    fmt.Println(error)
    os.Exit(8)
}
