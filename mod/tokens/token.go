package tokens

type Token struct {
	value	string
}

func (t Token) Repr() string {
	return "TAG:" + t.value
}