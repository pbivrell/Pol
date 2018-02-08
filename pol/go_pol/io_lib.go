package go_pol

import "../pol_type"
import "fmt"

type Go_Pol struct{}

const MAXBUFFERSIZE = 400

var stdoutbuffer = make([]string,0)

func WriteBuffer(){
    for _, s := range stdoutbuffer {
        fmt.Print(s)
    }
    stdoutbuffer = stdoutbuffer[0:0]
}

func (gp *Go_Pol) Flush(args []pol_type.Pol_Type){
    WriteBuffer()
}

func (gp *Go_Pol) Println(args []pol_type.Pol_Type){
    gp.Print(args)
    stdoutbuffer = append(stdoutbuffer,"\n")
    if len(stdoutbuffer) > MAXBUFFERSIZE {
        WriteBuffer()
    }
}

func (gp *Go_Pol) Print(args []pol_type.Pol_Type){
    val := ""
    for _, arg := range args {
        val += arg.String()
    }
    stdoutbuffer = append(stdoutbuffer,val)

    if len(stdoutbuffer) > MAXBUFFERSIZE {
        WriteBuffer()
    }
}
