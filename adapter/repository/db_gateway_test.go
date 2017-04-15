package repository

import "time"

type MockDbGateway struct {
	ModifiedParam1 interface{}
	Result1        interface{}
	Result2        interface{}
}

func (m *MockDbGateway) Get(receiverObjPtr CommonModel, id string) error {
	if m.ModifiedParam1 != nil {
		*receiverObjPtr.(*UserModel) = *m.ModifiedParam1.(*UserModel)
	}
	if m.Result1 != nil {
		return m.Result1.(error)
	}
	return nil
}

func (m *MockDbGateway) Create(dataObjPtr CommonModel) (string, error) {
	return m.Result1.(string), m.Result2.(error)
}

func (m *MockDbGateway) GetList(receiverObjs interface{}, index string, val interface{}) error {
	return m.Result1.(error)
}

func (m *MockDbGateway) GetPartOfTable(receiverObjs interface{}, timeIndex time.Time, size int, filterMap map[string][]string) error {
	return m.Result1.(error)
}

func (m *MockDbGateway) Update(receiverObjsPtr CommonModel, id string) error {
	return m.Result1.(error)
}

func (m *MockDbGateway) Delete(id string) error {
	return m.Result1.(error)
}
