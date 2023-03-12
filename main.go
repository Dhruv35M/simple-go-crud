package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/saurabh419/simple-go-crud/db"
)

func init() {
	godotenv.Load()
}

func main() {
	fmt.Println("I love golang!")
	authorName := os.Getenv("AUTHOR_NAME")
	fmt.Println("author name: ", authorName)

	r := gin.Default()
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "this works",
			"author":  authorName,
		})
	})

	books := r.Group("books")
	books.GET("/all", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    db.BookList,
		})
	})

	books.POST("new", func(ctx *gin.Context) {
		var newBook db.Book
		err := ctx.Bind(&newBook)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "invalid details",
			})
			return
		}

		db.Insert(newBook)
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "book stored successfully",
			"data":    newBook,
		})
	})

	books.DELETE(":id", func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "invalid book id",
			})
			return
		}
		err = db.Delete(uint(id))
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "failed to delete",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "book deleted successfully",
		})

	})

	log.Fatalln(r.Run(":4000"))
}
