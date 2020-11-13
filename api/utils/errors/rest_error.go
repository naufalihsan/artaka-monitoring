package errors

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// func (gh getHandler) Handle(w http.ResponseWriter, r *http.Request) {
// 	id := gh.rr.GetRouteParam(r, "id")
// 	err, resp := gh.uc.Get(id)
// 	if err != nil {
// 		httplib.ResponseException(w, err, 422)
// 		return
// 	}
// 	httplib.ResponseData(w, resp)
// }

// func (gh getHandler) Handle(w http.ResponseWriter, r *http.Request) {
// 	id := gh.rr.GetRouteParam(r, "id")
// 	err, resp := gh.uc.Get(id)
// 	if err != nil {
// 	httplib.ResponseException(w, err, 422)
// 	return
// 	}
// 	httplib.ResponseData(w, resp)
// 	}
