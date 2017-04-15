package repository

import "time"

type CommonModel interface {
	SetLastUpdated(time.Time)
	SetCreatedTime(time.Time)
}

type CommonModelImpl struct {
	LastUpdated time.Time `gorethink:"lastUpdated" json:"lastUpdated"`
	CreatedTime time.Time `gorethink:"createdTime" json:"createdTime"`
}

func (b *CommonModelImpl) SetLastUpdated(t time.Time) {
	b.LastUpdated = t
}

func (b *CommonModelImpl) SetCreatedTime(t time.Time) {
	b.CreatedTime = t
}
