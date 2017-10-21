package controller

import (
	"app/app"
	"github.com/goadesign/goa"
)

// MeetingsController implements the meetings resource.
type MeetingsController struct {
	*goa.Controller
}

// NewMeetingsController creates a meetings controller.
func NewMeetingsController(service *goa.Service) *MeetingsController {
	return &MeetingsController{Controller: service.NewController("MeetingsController")}
}

// Meetings runs the meetings action.
func (c *MeetingsController) Meetings(ctx *app.MeetingsMeetingsContext) error {
	// MeetingsController_Meetings: start_implement

	// Put your logic here

	// MeetingsController_Meetings: end_implement
	res := app.MeetingtypeCollection{}
	return ctx.OK(res)
}
