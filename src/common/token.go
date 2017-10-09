package common

/* token is a struct containing token information
   - typ is the token type
   - value is the token value
*/

type Token struct {
	Type   string `json:"type"`
	Value  string `json:"value"`
	Lineno int
}

func (t Token) String() string {
	return t.Value
}
