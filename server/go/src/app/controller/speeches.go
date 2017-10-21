package controller

import (
	"app/app"
	"github.com/goadesign/goa"
)

// SpeechesController implements the speeches resource.
type SpeechesController struct {
	*goa.Controller
}

// NewSpeechesController creates a speeches controller.
func NewSpeechesController(service *goa.Service) *SpeechesController {
	return &SpeechesController{Controller: service.NewController("SpeechesController")}
}

// Speeches runs the speeches action.
func (c *SpeechesController) Speeches(ctx *app.SpeechesSpeechesContext) error {
	// SpeechesController_Speeches: start_implement

	// Put your logic here

	// SpeechesController_Speeches: end_implement
	res := app.SpeechtypeCollection{}
	return ctx.OK(res)
}
