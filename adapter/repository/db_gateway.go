package repository

import "time"

type DbGateway interface {
	Get(receiverObjPtr CommonModel, id string) error

	Create(dataObjPtr CommonModel) (string, error)

	GetList(receiverObjs interface{}, index string, val interface{}) error

	GetPartOfTable(receiverObjs interface{}, timeIndex time.Time, size int, filterMap map[string][]string) error

	Update(receiverObjsPtr CommonModel, id string) error

	Delete(id string) error
}
