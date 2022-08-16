package core

import (
	"net/url"

	"github.com/rendau/dop/adapters/client/httpc"
	"github.com/rendau/dop/dopErrs"
	"github.com/rendau/sms_smsc/internal/domain/errs"
	"github.com/rendau/sms_smsc/internal/domain/types"
)

func (c *St) Send(pars *types.SendReqSt) error {
	var err error

	err = c.validateValues(pars)
	if err != nil {
		return err
	}

	repObj := &types.SendRepSt{}

	repData, err := c.httpc.SendRecvJson(nil, repObj, httpc.OptionsSt{
		Method: "GET",
		Path:   "send.php",
		Params: url.Values{
			"phones":  {pars.To},
			"mes":     {pars.Message},
			"charset": {"utf-8"},
			"fmt":     {"3"},
		},
	})
	if err != nil {
		return err
	}

	if (repObj.ErrorCode != 0) || (repObj.Error != "") {
		c.lg.Errorw(
			"Bad response smsc.kz", nil,
			"error_code", repObj.ErrorCode,
			"error", repObj.Error,
			"phone", pars.To,
			"rep_body", string(repData),
		)
		return dopErrs.ServiceNA
	}

	return nil
}

func (c *St) validateValues(pars *types.SendReqSt) error {
	if len(pars.To) == 0 {
		c.lg.Warnw("To is empty", errs.PhonesRequired)
		return errs.PhonesRequired
	}

	if len(pars.Message) == 0 {
		c.lg.Warnw("Message is empty", errs.MessageRequired)
		return errs.MessageRequired
	}

	return nil
}
