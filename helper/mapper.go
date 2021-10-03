package helper

import (
	"todo-go-rest/model"
	"todo-go-rest/model/response"
)

func RoleToRoleResponse(role model.Role) response.RoleResponse {
	return response.RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
	}
}

func UserToCredentialResponse(user model.User) response.CredentialResponse {
	return response.CredentialResponse{
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role.ID,
	}
}

func ToDoToResponse(toDo model.ToDo) response.ToDoResponse {
	return response.ToDoResponse{
		ID:    toDo.ID,
		Title: toDo.Title,
		User:  UserToCredentialResponse(toDo.User),
	}
}

func ToDosToResponses(toDos []model.ToDo) (responses []response.ToDoResponse) {
	for _, toDo := range toDos {
		responses = append(responses, ToDoToResponse(toDo))
	}
	return responses
}
