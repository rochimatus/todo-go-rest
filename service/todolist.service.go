package service

import (
	"errors"
	"strconv"
	"todo-go-rest/model"
	"todo-go-rest/model/request"
	"todo-go-rest/repository"
)

type ToDoListService interface {
	FindAll() ([]model.ToDoList, error)
	FindByID(ID int) (model.ToDoList, error)
	Create(req request.ToDoListRequest) (model.ToDoList, error)
	Delete(ID int) (model.ToDoList, error)
	Update(ID int, req request.ToDoListRequest) (model.ToDoList, error)
	AddFile(req request.AttachmentRequest, todolistId int, fileName string) (model.Attachment, error)
}

type toDoListService struct {
	repository           repository.ToDoListRepository
	attachmentRepository repository.AttachmentRepository
	toDoService          ToDoService
	statusService        StatusService
}

func NewToDoListService(repository repository.ToDoListRepository, attachmentRepository repository.AttachmentRepository, toDoService ToDoService, statusService StatusService) *toDoListService {
	return &toDoListService{
		repository:           repository,
		attachmentRepository: attachmentRepository,
		toDoService:          toDoService,
		statusService:        statusService,
	}
}

func (service *toDoListService) FindAll() ([]model.ToDoList, error) {
	return service.repository.FindAll()
}

func (service *toDoListService) FindByID(ID int) (model.ToDoList, error) {
	return service.repository.FindByID(ID)
}

func (service *toDoListService) Create(req request.ToDoListRequest) (model.ToDoList, error) {
	status, err := service.statusService.FindByID(req.StatusID)
	if err != nil {
		return model.ToDoList{}, err
	}

	toDo, err := service.toDoService.FindByID(req.ToDoID)
	if err != nil {
		return model.ToDoList{}, err
	}

	toDoList := model.ToDoList{
		ToDo:     toDo,
		ToDoID:   toDo.ID,
		Status:   status,
		StatusID: status.ID,
		Task:     req.Task,
	}

	return service.repository.Create(toDoList)
}

func (service *toDoListService) Delete(ID int) (model.ToDoList, error) {
	toDoList, err := service.repository.FindByID(ID)

	if err != nil {
		return toDoList, errors.New("ToDoList with ID " + strconv.Itoa(ID) + " is not found")
	}

	return service.repository.Delete(toDoList)
}

func (service *toDoListService) Update(ID int, req request.ToDoListRequest) (model.ToDoList, error) {
	toDoList, err := service.repository.FindByID(ID)

	if err != nil {
		return toDoList, errors.New("ToDoList with ID " + strconv.Itoa(ID) + " is not found")
	}

	if toDoList.StatusID != req.StatusID {
		status, err := service.statusService.FindByID(req.StatusID)
		if err != nil {
			return model.ToDoList{}, err
		}
		toDoList.Status = status
		toDoList.StatusID = status.ID
	}

	if toDoList.ToDoID != req.ToDoID {
		toDo, err := service.toDoService.FindByID(req.StatusID)
		if err != nil {
			return model.ToDoList{}, err
		}
		toDoList.ToDo = toDo
		toDoList.ToDoID = toDo.ID
	}

	toDoList.Task = req.Task

	return service.repository.Update(toDoList)
}

func (service *toDoListService) AddFile(req request.AttachmentRequest, todolistId int, fileName string) (model.Attachment, error) {
	uri := "http://localhost:8080/image/" + fileName
	attachment := model.Attachment{
		ToDoListId: todolistId,
		Caption:    req.Caption,
		Url:        uri,
	}
	attachment, err := service.attachmentRepository.Create(attachment)
	if err != nil {
		return model.Attachment{}, err
	}
	return attachment, err
}
