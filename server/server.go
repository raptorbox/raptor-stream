package server

import (
	"github.com/gin-gonic/gin"
	"github.com/raptorbox/raptor-sdk-go/models"
	"github.com/raptorbox/raptor-stream/influx"
)

//Start a server
func Start(port string) error {

	r := gin.Default()

	r.PUT("/:objectId/:streamId", func(c *gin.Context) {

		dev := models.NewDevice()
		dev.ID = c.Param("objectId")
		stream := models.NewStream(dev)
		stream.Name = c.Param("stremId")

		record := stream.CreateRecord()
		r := make([]*models.Record, 0)

		err := c.BindJSON(record)
		if err != nil {
			// try with an array
			err = c.BindJSON(r)
			if err != nil {
				c.JSON(400, gin.H{
					"message": err.Error(),
					"code":    400,
				})
				return
			}
		} else {
			r = append(r, record)
		}

		err = influx.Write(r)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
				"code":    400,
			})
			return
		}

		c.Status(202)
	})

	return r.Run(port)
}
