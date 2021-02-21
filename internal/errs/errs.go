package errs

type Err string

func (e Err) Error() string {
	return string(e)
}

const (
	PhonesRequired  = Err("phones_required")
	MessageRequired = Err("message_required")
	ServerNA        = Err("server_not_available")
)
