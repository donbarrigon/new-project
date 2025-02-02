package user

import "github.com/donbarrigon/new-project/internal/orm"

type User struct {
	orm.Model
}

func NewModel() *User {
	model := &User{}
	model.Table("user")
	model.Fillable("name", "email", "phone")
	model.Guarded("password")
	return model
}
