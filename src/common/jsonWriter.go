package common

import "encoding/json"
import "os"

//TODO Make this generic so it can turn any slice of structs into JSON
func WriteJSON(toks []Token, file string)  {
	//Convert Exported struct Token to Json
	jsonData, err := json.Marshal(toks)
	if err != nil {
		panic(err)
	}

	fh, errw := os.Create(file)
    //Defer = do this code when function returns
    defer fh.Close()
	if err != nil {
		panic(errw)
	}
	_, write_err := fh.Write(jsonData)
	if write_err != nil {
		panic(write_err)
	}

}
