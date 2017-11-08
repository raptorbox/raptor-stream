package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/raptorbox/raptor-sdk-go/models"
	"github.com/raptorbox/raptor-stream/api"
)

//Start a server
func Start(port string) error {

	r := gin.Default()

	r.PUT("/:objectId", func(c *gin.Context) {

		dev := models.NewDevice()
		dev.ID = c.Param("objectId")
		stream := models.NewStream(dev)
		stream.Name = c.Param("streamId")
		r := make([]*models.Record, 0)

		err := c.BindJSON(r)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
				"code":    400,
			})
			return
		}

		err = api.Write(r)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
				"code":    400,
			})
			return
		}

		c.Status(202)
	})
	r.PUT("/:objectId/:streamId", func(c *gin.Context) {

		dev := models.NewDevice()
		dev.ID = c.Param("objectId")
		stream := models.NewStream(dev)
		stream.Name = c.Param("stremId")
		record := stream.CreateRecord()
		r := make([]*models.Record, 0)

		err := c.BindJSON(record)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
				"code":    400,
			})
		} else {
			r = append(r, record)
		}

		err = api.Save(r)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
				"code":    400,
			})
			return
		}

		c.Status(202)
	})

	// drop data
	r.DELETE("/:objectId", func(c *gin.Context) {
		f := api.RecordQuery{
			DeviceID: c.Param("objectId"),
		}
		err := api.Delete(f)
		if err != nil {
			c.JSON(err.Code, gin.H{
				"message": err.Message,
				"code":    err.Code,
			})
			return
		}
	})
	r.DELETE("/:objectId/:streamId", func(c *gin.Context) {
		f := api.RecordQuery{
			DeviceID: c.Param("objectId"),
			StreamID: c.Param("streamId"),
		}
		err := api.Delete(f)
		if err != nil {
			c.JSON(err.Code, gin.H{
				"message": err.Message,
				"code":    err.Code,
			})
			return
		}
	})

	// list paged data
	r.GET("/:objectId/:streamId", func(c *gin.Context) {
		log.Fatalf("Not implemented GET /:objectId/:streamId")
	})
	// search data
	r.POST("/:objectId/:streamId", func(c *gin.Context) {
		log.Fatalf("Not implemented POST /:objectId/:streamId")
	})

	return r.Run(port)
}
