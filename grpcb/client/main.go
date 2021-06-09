package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/404th/grpcb/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:3537", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	cl := proto.NewContactServiceClient(conn)

	g := gin.Default()
	g.GET("/add/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "error not valid A"})
			return
		}
		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "error not valid B"})
			return
		}

		req := &proto.Req{A: int32(a), B: int32(b)}
		res, err := cl.Plus(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(res.Res),
			})
		}
	})

	if err := g.Run(); err != nil {
		panic(err)
	}

}
