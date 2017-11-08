package api

import (
	"fmt"

	"github.com/raptorbox/raptor-sdk-go/models"
	"github.com/raptorbox/raptor-stream/errors"
)

//Sort informations
type Sort struct {
	Field     string `json:"field"`
	Direction string `json:"direction"`
}

//Pager paged result
type Pager struct {
	Content []models.Record `json:"content"`
	Total   int             `json:"total"`
	Size    int             `json:"size"`
	Sort    Sort            `json:"sort"`
}

//Save store records to db
func Save(r []*models.Record) *errors.HTTPError {
	err := WriteRecord(r)
	if err != nil {
		return errors.InternalServerError(err)
	}
	return nil
}

//Delete records
func Delete(f RecordQuery) *errors.HTTPError {
	err := DeleteRecord(f)
	if err != nil {
		return errors.InternalServerError(err)
	}
	return nil
}

//Search records
func Search(f RecordQuery, q *models.DataQuery) (*Pager, *errors.HTTPError) {
	pager := Pager{}

	res, err := Read(f)
	if err != nil {
		return nil, errors.InternalServerError(err)
	}
	// TODO
	return &pager, nil
}

//List records
func List(from int, to int) (*Pager, *errors.HTTPError) {
	panic(fmt.Errorf("Not implemented"))
}
