package tag_api

import (
	"github.com/alexedwards/scs"
	"github.com/julienschmidt/httprouter"
)

func NewData(host, port string) (data *ApiData) {

	data = &ApiData{apiUrl: host + ":" + port}
	data.InitSessions()

	d = data
	return
}

// Local data - most functions are methods of this
type ApiData struct {
	apiUrl         string
	router         *httprouter.Router
	sessionManager *scs.Manager
}

// For session access -- TODO: create session service, add to handler wrappers
var d *ApiData

type UserMap map[int64]User

type GroupMap map[int64]Group

type ImageMap map[int64]Image

type ImagesGroupsMap map[int64]bool

type JwtPayload struct {
	UserId int64  `json:"id"`
	Guid   string `json:"guid"`
}

type ResponseMessage struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type QueueMessage struct {
	Command string `json:"command"`
}
