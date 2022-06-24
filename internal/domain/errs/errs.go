package errs

import (
	"github.com/rendau/dop/dopErrs"
)

const (
	PhonesRequired  = dopErrs.Err("phones_required")
	MessageRequired = dopErrs.Err("message_required")
)
