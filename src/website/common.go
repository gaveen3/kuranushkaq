package website

//Result *
type Result struct {
	Ok   bool        `json:"ok" xml:"ok"`
	Msg  string      `json:"msg" xml:"msg"`
	Data interface{} `json:"data" xml:"data"`
}

//User *
type User struct {
	User string `json:"user" xml:"user"`
	Pwd  string `json:"pwd" xml:"pwd"`
}

//WsParam *
type WsParam struct {
	Vars string `json:"vars" xml:"vars"`
	Cmd  string `json:"cmd" xml:"cmd"`
}
