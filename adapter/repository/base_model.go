package repository

import "time"

type BaseModel interface {
	SetLastUpdated(time.Time)
	SetCreatedTime(time.Time)
}

type BaseModelImpl struct {
	LastUpdated time.Time `gorethink:"lastUpdated" json:"lastUpdated"`
	CreatedTime time.Time `gorethink:"createdTime" json:"createdTime"`
}

func (b *BaseModelImpl) SetLastUpdated(t time.Time) {
	b.LastUpdated = t
}

func (b *BaseModelImpl) SetCreatedTime(t time.Time) {
	b.CreatedTime = t
}
