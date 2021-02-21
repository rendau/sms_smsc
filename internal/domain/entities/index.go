package entities

type SendReqSt struct {
	Phones  string `schema:"phones"`
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

type GetBalanceRepSt struct {
	Balance string `json:"balance"`
	SmscErrSt
}

type BalanceCacheSt struct {
	Balance float64 `json:"balance"`
}
