package rethinkdbHelper

import (
	rdb "github.com/dancannon/gorethink"
)

type RethinkdbBuilder struct {
	Address  string
	Database string
	AuthKey  string
}

func (builder *RethinkdbBuilder) Create() (*rdb.Session, error) {
	session, err := rdb.Connect(rdb.ConnectOpts{
		Address:  builder.Address,
		Database: builder.Database,
		AuthKey:  builder.AuthKey,
	})
	if err != nil {
		return nil, err
	}

	return session, nil
}
