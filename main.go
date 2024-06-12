package main

import (
	"app/multi_cache"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Capacity of cache
const length = 3

func main() {
	// Importing the multi_cache package
	cache := multi_cache.NewMultiCache()

	//storing the values in the cache
	cache.Set("a", "1", length)
	cache.Set("b", "2", length)

	//intiliazes the restapi with gin
	r := gin.Default()

	// Endpoint to retrieve a value by key
	r.GET("/:key", func(ctx *gin.Context) {
		k := ctx.Param("key")
		ctx.JSON(http.StatusOK, cache.Get(k))
	})

	// Endpoint to delete a value by key
	r.DELETE("/:key", func(ctx *gin.Context) {
		//storing the key value
		k := ctx.Param("key")
		//calling the delete method
		cache.Del(k)
	})

	// Endpoint to set a key-value pair
	r.POST("/:key/:value", func(ctx *gin.Context) {
		//storing the key and value pair
		k := ctx.Param("key")
		v := ctx.Param("value")
		//calling the set methods and sending the key,value and length
		cache.Set(k, v, length)
	})

	// Endpoint to print the redis cache contents
	r.GET("/redis/print", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, cache.Print_redis())
	})

	// Endpoint to print the in-memory cache contents
	r.GET("/inmemory/print", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, cache.Print_in_mem())
	})

	// Endpoint to delete entire data
	r.DELETE("/all", func(ctx *gin.Context) {
		cache.Del_ALL()
	})

	// Start the Gin server
	r.Run()

}
