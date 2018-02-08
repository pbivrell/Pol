package common

var stdoutBuffer = make([]string,0)
var flushed = false

func WriteSTDOut(output string){
    stdoutBuffer = append(stdoutBuffer, output)
}

func Flush(){
    flushed = true
}

func IsFlushed() bool{
    return flushed
}

func GetSTDOut() string{
    retVal := ""
    for _, out := range stdoutBuffer{
        retVal += out + "\n"
    }
    return retVal
}
