package mongo

import "net/http"

type Handler struct {
}

func (h *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	uploadFile(w, r)
}

func (h *Handler) DownloadFile(w http.ResponseWriter, r *http.Request) {
	downloadFile(w, r)
}

func (h *Handler) DeleteFile(w http.ResponseWriter, r *http.Request) {
	deleteFile(w, r)
}
