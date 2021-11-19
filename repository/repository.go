package repository

import (
	"fmt"
	"log"
	"todo-go-rest/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepository     UserRepository
	RoleRepository     RoleRepository
	StatusRepository   StatusRepository
	ToDoRepository     ToDoRepository
	ToDoListRepository ToDoListRepository
}

func CreateRepository(useSeed bool) *Repository {
	dsn := "root:@tcp(127.0.0.1:3306)/todo-go-rest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("COnnected to db")
	if err != nil {
		log.Fatal("DB connection failed")
	}

	db.AutoMigrate(&model.Role{},
		&model.User{},
		&model.Status{},
		&model.ToDo{},
		&model.ToDoList{})

	if useSeed {
		seeding(db)
	}

	return &Repository{
		UserRepository:     NewUserRepository(db),
		RoleRepository:     NewRoleRepository(db),
		StatusRepository:   NewStatusRepository(db),
		ToDoRepository:     NewToDoRepository(db),
		ToDoListRepository: NewToDoListRepository(db),
	}
}

func seeding(db *gorm.DB) {
}
