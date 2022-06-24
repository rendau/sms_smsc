package types

type SendReqSt struct {
	Phone   string `schema:"phone"`
	Message string `schema:"message"`
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
