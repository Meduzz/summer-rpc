package main

import (
	"github.com/Meduzz/disconnected"
	"github.com/Meduzz/disconnected/pkg/web"
	"github.com/Meduzz/helper/nuts"
	"github.com/Meduzz/rpc"
	"github.com/Meduzz/rpc/encoding"
	"github.com/Meduzz/summer"
	summerrpc "github.com/Meduzz/summer-rpc"
	"github.com/gin-gonic/gin"
)

func main() {
	conn, err := nuts.Connect()

	if err != nil {
		panic(err)
	}

	rpcConn := rpc.NewRpc(conn, encoding.Json())
	summer.Register("greet", summerrpc.RpcProxy(rpcConn, "test.greeting", 5))

	disconnected.HttpServer("/", func(s *web.Server) {
		s.WithRouter(func(e *gin.Engine) {
			e.POST("/api/rpc", summer.HTTP())
		})
	})
}
