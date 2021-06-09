package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/404th/grpcc/protos"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:3433", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error while dialing")
	}

	client := protos.NewContactServiceClient(conn)

	g := gin.Default()
	g.POST("/create/:name/:username", func(ctx *gin.Context) {

		var (
			n, u string
		)
		nq, kel := ctx.Params.Get("name")
		if kel {
			n = nq
		}
		uq, kel := ctx.Params.Get("username")
		if kel {
			u = uq
		}
		req := protos.Contact{
			Name:     n,
			Username: u,
		}

		res, err := client.Create(ctx, &req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"It's name": fmt.Sprint(res.GetName()),
			})
		}
	})

	g.GET("/contactget/:user", func(ctx *gin.Context) {
		req := protos.WithName{
			Name: ctx.Param("user"),
		}

		res, err := client.Get(ctx, &req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"It's name": fmt.Sprint(res.GetName()),
			})
		}
	})

	g.DELETE("/delete/:name", func(ctx *gin.Context) {
		req := protos.WithName{
			Name: ctx.Param("name"),
		}

		res, err := client.Delete(ctx, &req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"It's name": fmt.Sprint(res),
			})
		}
	})

	if err := g.Run(); err != nil {
		panic(err)
	}

}
