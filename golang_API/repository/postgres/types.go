package postgresrepository

import "gorm.io/gorm"

type UserDB struct {
	repoPrefix string
	db         *gorm.DB
}
