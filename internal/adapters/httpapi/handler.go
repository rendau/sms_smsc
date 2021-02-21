package httpapi

import (
	"net/http"

	"github.com/rendau/sms/internal/domain/entities"
)

func (a *St) hSend(w http.ResponseWriter, r *http.Request) {
	reqObj := &entities.SendReqSt{}

	if !a.uParseRequestJSON(w, r, reqObj) {
		return
	}

	if reqObj.Sync {
		sendErr := a.cr.Send(reqObj)
		if sendErr != nil {
			a.uHandleError(sendErr, w)
			return
		}
	} else {
		go func() { _ = a.cr.Send(reqObj) }()
	}

	w.WriteHeader(200)
}

func (a *St) hBcast(w http.ResponseWriter, r *http.Request) {
	reqObj := &entities.SendReqSt{}

	if !a.uParseRequestJSON(w, r, reqObj) {
		return
	}

	if reqObj.Sync {
		sendErr := a.cr.Bcast(reqObj)
		if sendErr != nil {
			a.uHandleError(sendErr, w)
			return
		}
	} else {
		go func() { _ = a.cr.Bcast(reqObj) }()
	}

	w.WriteHeader(200)
}

func (a *St) hGetBalance(w http.ResponseWriter, r *http.Request) {
	balance := a.cr.GetBalance()

	a.uRespondJSON(w, map[string]float64{
		"balance": balance,
	})
}

func (a *St) hCronCheckBalance(w http.ResponseWriter, r *http.Request) {
	a.cr.CheckBalance()

	w.WriteHeader(200)
}
