package controller

import (
	"app/app"

	"github.com/goadesign/goa"
	"gopkg.in/mgo.v2"
)

// RoomsController implements the rooms resource.
type RoomsController struct {
	*goa.Controller
	Mgo *mgo.Session
}

// NewRoomsController creates a rooms controller.
func NewRoomsController(service *goa.Service, mgo *mgo.Session) *RoomsController {
	return &RoomsController{
		Controller: service.NewController("RoomsController"),
		Mgo:        mgo,
	}
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
	return ctx.OK(``)
}
