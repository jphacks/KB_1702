package controller

import (
	"app/app"

	"github.com/goadesign/goa"
	"encoding/json"
)

// RoomsController implements the rooms resource.
type RoomsController struct {
	*goa.Controller
}

// NewRoomsController creates a rooms controller.
func NewRoomsController(service *goa.Service) *RoomsController {
	return &RoomsController{Controller: service.NewController("RoomsController")}
}

// Create runs the create action.
func (c *RoomsController) Create(ctx *app.CreateRoomsContext) error {
	// RoomsController_Create: start_implement

	// Put your logic here

	// RoomsController_Create: end_implement
	ctx.ResponseData.Header().Set("Location", "/rooms/1")
	return ctx.SeeOther()
}

// Show runs the show action.
func (c *RoomsController) Show(ctx *app.ShowRoomsContext) error {
	// RoomsController_Show: start_implement

	// Put your logic here

	// RoomsController_Show: end_implement
	json.NewEncoder(w).Encode()
	json.NewEncoder([]byte(`{"room": {"id": "dwabdhjwabkjdbadkad","name": "うぇいサウンド","progress": 1,"start_at": "1995-01-11T06:25:13+09:00","end_at": "1995-01-11T06:25:13+09:00","agenda": [{"id": 1,"title": "アイデア出し","goal": "アイデアを10個出す","time": 10,"start_at": "1995-01-11T06:25:13+09:00","end_at": "1995-01-11T06:25:13+09:00","child": [{"id": 2,"title": "アイデア出し","goal": "アイデアを10個出す","time": 10,"start_at": "1995-01-11T06:25:13+09:00","end_at": "1995-01-11T06:25:13+09:00"}]},{"id": 3,"title": "アイデア出し","goal": "アイデアを10個出す","time": 10,"start_at": "1995-01-11T06:25:13+09:00","end_at": "1995-01-11T06:25:13+09:00"}]}}`))
	return ctx.OK(``)
}
