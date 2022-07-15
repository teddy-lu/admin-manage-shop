package address

import "admin-manage-shop/pkg/database"

func (model *Address) Create() {
	database.DB.Create(&model)
}
