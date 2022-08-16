package types

type SendReqSt struct {
	To      string `schema:"to"`
	Message string `schema:"text"`
	Sync    bool   `schema:"sync"`
}

type SendRepSt struct {
	ID  uint64 `json:"id"`
	CNT int    `json:"cnt"`
	SmscErrSt
}

type SmscErrSt struct {
	ErrorCode int    `json:"error_code"`
	Error     string `json:"error"`
}
