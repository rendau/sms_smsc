package types

type SendReqSt struct {
	To      string `json:"to"`
	Message string `json:"text"`
	Sync    bool   `json:"sync"`
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
