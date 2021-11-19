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

func ToDoListToResponse(toDoList model.ToDoList) response.ToDoListResponse {
	return response.ToDoListResponse{
		ID:     toDoList.ID,
		ToDo:   ToDoToResponse(toDoList.ToDo),
		Task:   toDoList.Task,
		Status: toDoList.Status.Name,
	}
}

func ToDoListsToResponses(toDoLists []model.ToDoList) (responses []response.ToDoListResponse) {
	for _, toDoList := range toDoLists {
		responses = append(responses, ToDoListToResponse(toDoList))
	}
	return responses
}
