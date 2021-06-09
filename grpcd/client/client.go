package main

import (
	"fmt"
	"net/http"

	"github.com/404th/grpcd/protos"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:3435", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := protos.NewServiceNameClient(conn)

	g := gin.Default()

	g.POST("/task/:title/:body", func(ctx *gin.Context) {

		res, err := client.Create(ctx, &protos.Task{
			Id:    "",
			Title: ctx.Param("title"),
			Body:  ctx.Param("body"),
		})
		if err != nil {
			fmt.Printf("error while creating: %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"id": res.GetId(),
			})
		}
	})

	g.GET("/task/:id", func(ctx *gin.Context) {

		res, err := client.Get(ctx, &protos.WithID{
			Id: ctx.Param("id"),
		})

		if err != nil {
			fmt.Printf("error while getting: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"id":    res.GetId(),
				"title": res.GetTitle(),
				"body":  res.GetBody(),
			})
		}
	})

	g.GET("/task", func(ctx *gin.Context) {

		res, err := client.GetAll(ctx, &protos.Empty{})

		if err != nil {
			fmt.Printf("error while getting all: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"tasks": res.GetTasks(),
			})
		}
	})

	g.PUT("/task/:id/:title/:body", func(ctx *gin.Context) {

		res, err := client.Update(ctx, &protos.Task{
			Id:    ctx.Param("id"),
			Title: ctx.Param("title"),
			Body:  ctx.Param("body"),
		})

		if err != nil {
			fmt.Printf("error while updating: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"id": res.GetId(),
			})
		}
	})

	g.DELETE("/task/:id", func(ctx *gin.Context) {
		res, err := client.Delete(ctx, &protos.WithID{
			Id: ctx.Param("id"),
		})
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()},
			)
		} else {
			ctx.JSON(
				http.StatusOK,
				gin.H{"id": res.GetId()},
			)
		}
	})

	if err := g.Run(); err != nil {
		panic(err)
	}
}
