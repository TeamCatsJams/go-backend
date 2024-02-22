package db

import "telmed_backend/pkgs/user"

var (
	UserSvc user.Service = nil
	// services
)

func InitServices() {
	db := GetDB()
	usersRepo := user.NewPostgresRepo(db)
	UserSvc = user.NewService(usersRepo)
}
