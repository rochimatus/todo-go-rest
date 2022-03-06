package service

import (
	"todo-go-rest/model"
	"todo-go-rest/model/request"
	"todo-go-rest/repository"
)

type AttachmentService interface {
	Create(attachmentReq request.AttachmentRequest, todolistId int, fileName string) (model.Attachment, error)
}

type attachmentService struct {
	repository repository.AttachmentRepository
}

func NewAttachmentService(repository repository.AttachmentRepository) *attachmentService {
	return &attachmentService{
		repository: repository,
	}
}

func (service *attachmentService) Create(req request.AttachmentRequest, todolistId int, fileName string) (model.Attachment, error) {
	uri := "http://localhost:8080/image/" + fileName
	attachment := model.Attachment{
		ToDoListId: todolistId,
		Caption:    req.Caption,
		Url:        uri,
	}
	attachment, err := service.repository.Create(attachment)
	if err != nil {
		return model.Attachment{}, err
	}
	return attachment, err
}
