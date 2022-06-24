package rest

import (
	"github.com/gin-gonic/gin"
	dopHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/sms_smsc/internal/domain/types"
)

// @Router   /send [post]
// @Tags     general
// @Param    body  body  types.SendReqSt  false  "body"
// @Produce  json
// @Success  200
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hSend(c *gin.Context) {
	reqObj := &types.SendReqSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	if reqObj.Sync {
		dopHttps.Error(c, o.core.Send(reqObj))
	} else {
		go func() { _ = o.core.Send(reqObj) }()
	}
}
