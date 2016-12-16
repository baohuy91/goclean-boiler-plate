package rethinkdbHelper

import (
	"fmt"
	rdb "github.com/dancannon/gorethink"
	"github.com/stretchr/testify/assert"
	"goclean/interfaceadapter/repository"
	"testing"
	"time"
)

const (
	t_DB_NAME = "unit_test"
	t_TB_NAME = "t_common_test"
)

func init() {
	initDB()
}

func initDB() {
	// TODO: pass these in env
	// Create connection
	session, err := rdb.Connect(rdb.ConnectOpts{
		Address: "localhost:28015",
	})
	if err != nil {
		fmt.Printf("Couldn't connect to rethinkdb, please start rethinkdb first. Error: %s", err.Error())
		return
	}
	defer session.Close()

	// Create db
	respR, err := rdb.DBList().Run(session)
	defer respR.Close()
	dbNames := []string{}
	_ = respR.All(&dbNames)

	dbExist := false
	for _, dbName := range dbNames {
		if dbName == t_DB_NAME {
			dbExist = true
		}
	}
	if dbExist {
		_, err = rdb.DBDrop(t_DB_NAME).RunWrite(session)
		if err != nil {
			fmt.Print(err)
		}
	}
	resp, err := rdb.DBCreate(t_DB_NAME).RunWrite(session)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%d DB created\n", resp.DBsCreated)

	// Create table & index
	_, err = rdb.DB(t_DB_NAME).TableCreate(t_TB_NAME).RunWrite(session)
	if err != nil {
		fmt.Print(err)
	}
	_, err = rdb.DB(t_DB_NAME).Table(t_TB_NAME).IndexCreate("createdTime").RunWrite(session)
	_, err = rdb.DB(t_DB_NAME).Table(t_TB_NAME).IndexWait().Run(session)
}

func t_connect() *rdb.Session {
	sess, _ := rdb.Connect(rdb.ConnectOpts{
		Address:  "localhost:28015",
		Database: t_DB_NAME,
	})
	return sess
}

type DataStruct struct {
	Data string `gorethink:"data"`
	repository.BaseModelImpl
}

func TestRdbHandler_Create(t *testing.T) {
	session := t_connect()
	defer session.Close()
	dbHandler := rdbHandler{
		session:   session,
		TableName: t_TB_NAME,
	}

	data := DataStruct{
		Data: "data",
	}
	now := time.Now()
	data.CreatedTime = now

	id, err := dbHandler.Create(&data)
	assert.NoError(t, err)
	assert.NotEmpty(t, id)

	resp, err := rdb.Table(t_TB_NAME).Get(id).Run(dbHandler.session)
	assert.NoError(t, err)
	defer resp.Close()

	dataInDB := DataStruct{}
	err = resp.One(&dataInDB)
	assert.NoError(t, err)
	assert.Equal(t, "data", dataInDB.Data)
	assert.Equal(t, now.Unix(), dataInDB.CreatedTime.Unix())
}

func TestRdbHandler_Get(t *testing.T) {
	session := t_connect()
	defer session.Close()
	dbHandler := rdbHandler{
		session:   session,
		TableName: t_TB_NAME,
	}

	data := DataStruct{
		Data: "data",
	}
	now := time.Now()
	data.CreatedTime = now

	resp, err := rdb.Table(t_TB_NAME).Insert(data).RunWrite(dbHandler.session)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(resp.GeneratedKeys))

	dataInDB := DataStruct{}
	err = dbHandler.Get(&dataInDB, resp.GeneratedKeys[0])

	assert.NoError(t, err)
	assert.NotEmpty(t, dataInDB)
	assert.Equal(t, "data", dataInDB.Data)
	assert.Equal(t, now.Unix(), dataInDB.CreatedTime.Unix())

	// Error
	err = dbHandler.Get(nil, resp.GeneratedKeys[0])
	assert.Error(t, err)

	// Session close
	session.Close()
	err = dbHandler.Get(&dataInDB, resp.GeneratedKeys[0])
	assert.Error(t, err)
}

func TestRdbHandler_GetList(t *testing.T) {
	session := t_connect()
	dbHandler := rdbHandler{
		session:   session,
		TableName: t_TB_NAME,
	}

	rdb.Table(t_TB_NAME).Delete().RunWrite(session)

	data := DataStruct{
		Data: "data",
	}
	now := time.Now()
	data.CreatedTime = now

	_, err := rdb.Table(t_TB_NAME).Insert(data).RunWrite(dbHandler.session)
	_, err = rdb.Table(t_TB_NAME).Insert(data).RunWrite(dbHandler.session)
	_, err = rdb.Table(t_TB_NAME).Insert(data).RunWrite(dbHandler.session)
	_, err = rdb.Table(t_TB_NAME).Insert(data).RunWrite(dbHandler.session)
	assert.NoError(t, err)

	dataInDB := []*DataStruct{}
	err = dbHandler.GetList(&dataInDB, "createdTime", now)
	assert.NoError(t, err)
	assert.Equal(t, 4, len(dataInDB))

	dataInDB = []*DataStruct{}
	err = dbHandler.GetList(&dataInDB, "createdTime", now.AddDate(0, 0, 1))
	assert.NoError(t, err)
	assert.Equal(t, 0, len(dataInDB))

	dataInDB = []*DataStruct{}
	err = dbHandler.GetList(&dataInDB, "createdTime", "abc")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(dataInDB))

	dataInDB = []*DataStruct{}
	err = dbHandler.GetList(&dataInDB, "not_a_index", now)
	assert.Error(t, err)

	// Error
	panicFunc := func() {
		dbHandler.GetList(nil, "createdTime", now)
	}
	assert.Panics(t, panicFunc)

	// Session close
	session.Close()
	err = dbHandler.GetList(&dataInDB, "createdTime", now)
	assert.Error(t, err)
}

func TestRdbHandler_GetPartOfTable(t *testing.T) {
	session := t_connect()
	defer session.Close()
	dbHandler := rdbHandler{
		session:   session,
		TableName: t_TB_NAME,
	}

	rdb.Table(t_TB_NAME).Delete().RunWrite(session)

	data := DataStruct{
		Data: "data",
	}
	now := time.Date(2016, 7, 6, 10, 15, 0, 0, time.UTC)
	data.CreatedTime = now.Add(1 * time.Minute)
	_, err := rdb.Table(t_TB_NAME).Insert(data).RunWrite(dbHandler.session)
	data.CreatedTime = now.Add(2 * time.Minute)
	_, err = rdb.Table(t_TB_NAME).Insert(data).RunWrite(dbHandler.session)
	data.CreatedTime = now.Add(3 * time.Minute)
	_, err = rdb.Table(t_TB_NAME).Insert(data).RunWrite(dbHandler.session)
	data.CreatedTime = now.Add(4 * time.Minute)
	_, err = rdb.Table(t_TB_NAME).Insert(data).RunWrite(dbHandler.session)
	assert.NoError(t, err)

	dataInDB := []*DataStruct{}
	err = dbHandler.GetPartOfTable(&dataInDB, now.Add(3*time.Minute), 2, map[string][]string{})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(dataInDB))

	dataInDB = []*DataStruct{}
	err = dbHandler.GetPartOfTable(&dataInDB, now.Add(4*time.Minute), 4, map[string][]string{})
	assert.NoError(t, err)
	assert.Equal(t, 3, len(dataInDB))

	dataInDB = []*DataStruct{}
	err = dbHandler.GetPartOfTable(&dataInDB, now.Add(5*time.Minute), 10, map[string][]string{})
	assert.NoError(t, err)
	assert.Equal(t, 4, len(dataInDB))

	dataInDB = []*DataStruct{}
	err = dbHandler.GetPartOfTable(&dataInDB, now.Add(1*time.Minute), 10, map[string][]string{})
	assert.Equal(t, 0, len(dataInDB))

	// Error
	panicFunc := func() {
		dbHandler.GetPartOfTable(nil, now.Add(5*time.Minute), 1, map[string][]string{})
	}
	assert.Panics(t, panicFunc)

	// Session close
	session.Close()
	err = dbHandler.GetPartOfTable(nil, now.Add(5*time.Minute), 1, map[string][]string{})
	assert.Error(t, err)
}

func TestRdbHandler_Update(t *testing.T) {
	session := t_connect()
	defer session.Close()
	dbHandler := rdbHandler{
		session:   session,
		TableName: t_TB_NAME,
	}

	data := DataStruct{
		Data: "data",
	}
	now := time.Now()
	data.CreatedTime = now

	respW, _ := rdb.Table(t_TB_NAME).Insert(data).RunWrite(dbHandler.session)
	dataId1 := respW.GeneratedKeys[0]

	// Normal case
	data1 := &DataStruct{
		Data: "data1",
	}
	err := dbHandler.Update(data1, dataId1)
	assert.NoError(t, err)
	resp, err := rdb.Table(t_TB_NAME).Get(dataId1).Run(dbHandler.session)
	assert.NoError(t, err)
	defer resp.Close()
	dataInDB := DataStruct{}
	err = resp.One(&dataInDB)
	assert.NoError(t, err)
	assert.Equal(t, "data1", dataInDB.Data)
	assert.NotEqual(t, now.Unix(), dataInDB.CreatedTime.Unix())

	// Session close
	session.Close()
	err = dbHandler.Update(data1, "123")
	assert.Error(t, err)
}

func TestRdbHandler_Delete(t *testing.T) {
	session := t_connect()
	defer session.Close()
	dbHandler := rdbHandler{
		session:   session,
		TableName: t_TB_NAME,
	}

	data := DataStruct{
		Data: "data",
	}
	now := time.Now()
	data.CreatedTime = now

	respW, _ := rdb.Table(t_TB_NAME).Insert(data).RunWrite(dbHandler.session)
	dataId1 := respW.GeneratedKeys[0]

	// Normal case
	err := dbHandler.Delete(dataId1)
	assert.NoError(t, err)
	resp, err := rdb.Table(t_TB_NAME).Get(dataId1).Run(dbHandler.session)
	assert.NoError(t, err)
	defer resp.Close()
	assert.True(t, resp.IsNil())

	// Session close
	respW, _ = rdb.Table(t_TB_NAME).Insert(data).RunWrite(dbHandler.session)
	dataId1 = respW.GeneratedKeys[0]

	session.Close()
	err = dbHandler.Delete(dataId1)
	assert.Error(t, err)
}
