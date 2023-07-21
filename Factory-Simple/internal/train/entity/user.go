package train

import (
	"factorySimple/internal/core"
)

// 使用者
type User struct {
	Id   core.UUID // 物件追蹤ID
	Name string
}

func CreateUser(name string) User {
	return User{
		Id:   core.NewUUID(),
		Name: name,
	}
}
