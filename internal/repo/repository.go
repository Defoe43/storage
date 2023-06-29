package repo

type Repository interface {
	PutFile(filename string, data []byte)
	GetFile(filename string)
	DeleteFile(filename string)
}
