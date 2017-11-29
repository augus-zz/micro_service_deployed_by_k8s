package model

import (
	"fmt"
	db "micro_service_deployed_by_k8s/shared/db"
)

func ClearDB() {
	models := []interface{}{User{}, Customer{}, Store{}}
	for _, m := range models {
		db.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %v;", db.DB.NewScope(m).TableName()))
	}
}
