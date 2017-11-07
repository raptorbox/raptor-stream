package api

import (
	"fmt"

	"github.com/raptorbox/raptor-sdk-go/models"
	"github.com/raptorbox/raptor-stream/errors"
	"github.com/raptorbox/raptor-stream/influx"
)

//Write store records to db
func Write(r []*models.Record) *errors.HTTPError {
	err := influx.Write(r)
	if err != nil {
		return errors.InternalServerError(err)
	}
	return nil
}

//Delete records
func Delete(deviceID string, streamID string) error {
	panic(fmt.Errorf("Not implemented"))
}

//Search records
func Search(q models.DataQuery) error {
	panic(fmt.Errorf("Not implemented"))
}

//List records
func List(from int, to int) error {
	panic(fmt.Errorf("Not implemented"))
}
