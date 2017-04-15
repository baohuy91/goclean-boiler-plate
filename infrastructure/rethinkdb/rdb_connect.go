package rethinkdbHelper

import (
	rdb "github.com/dancannon/gorethink"
)

func RethinkdbConnect(address string, db string, authKey string) (*rdb.Session, error) {
	session, err := rdb.Connect(rdb.ConnectOpts{
		Address:  address,
		Database: db,
		AuthKey:  authKey,
	})
	if err != nil {
		return nil, err
	}

	return session, nil
}
