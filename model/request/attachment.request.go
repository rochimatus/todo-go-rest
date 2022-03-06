package request

import "mime/multipart"

type AttachmentRequest struct {
	File    *multipart.FileHeader `json:"file" binding:"required"`
	Caption string                `json:"caption"`
}
