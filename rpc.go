package summerrpc

import (
	"encoding/json"

	"github.com/Meduzz/rpc"
	"github.com/Meduzz/summer/api"
	"github.com/Meduzz/summer/errors"
	"github.com/Meduzz/summer/framework"
)

func RpcProxy(conn *rpc.RPC, topic string, timeout int) api.Handler {
	if timeout == 0 {
		timeout = 15
	}

	return func(r *api.Request) *api.Response {
		var res json.RawMessage
		err := conn.Request(topic, r.Params, &res, timeout)

		if err != nil {
			return framework.ErrorResponse(r.ID, errors.InternalError(err))
		}

		return framework.ResultResponse(r.ID, res)
	}
}
