package controller

import (
	"app/app"
	"app/design/constant"
	"app/mywebsocket"

	"github.com/goadesign/goa"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
reaction
from to type
これをmongoにぶち込みたい
*/

const (
	wsNextAgenda = "nextAgenda"
	// from to type
	wsReaction       = "reaction"
	wsDecisionAgenda = "decisionAgenda"
)

// AgendaController implements the agenda resource.
type AgendaController struct {
	*goa.Controller
	Mgo *mgo.Session
	ws  *mywebsocket.Server
}

// NewAgendaController creates a agenda controller.
func NewAgendaController(service *goa.Service, mgo *mgo.Session) *AgendaController {
	ws := mywebsocket.NewServer()
	go ws.Start()
	return &AgendaController{
		Controller: service.NewController("AgendaController"),
		Mgo:        mgo,
		ws:         ws,
	}
}

// AddReaction runs the addReaction action.
func (c *AgendaController) AddReaction(ctx *app.AddReactionAgendaContext) error {
	// AgendaController_AddReaction: start_implement

	// Put your logic here

	// AgendaController_AddDecision: end_implement
	return nil
}

// AddDecision runs the addDecision action.
func (c *AgendaController) AddDecision(ctx *app.AddDecisionAgendaContext) error {
	// AgendaController_AddDecision: start_implement

	// Put your logic here
	query := bson.M{
		"room.room_id": ctx.Payload.RoomID,
		"room.agenda":  bson.M{"$elemMatch": bson.M{"id": ctx.ID}},
	}
	update := bson.M{
		"$set": bson.M{
			"room.agenda.$.decision": ctx.Payload.Decision,
		},
	}
	mongo := c.Mgo.DB(constant.DBName).C(constant.CollectionRoom)
	_, err := mongo.Upsert(query, update)
	if err != nil {
		return goa.ErrBadRequest(err)
	}
	type wsNextAgendaStruct struct {
		ID       int    `json:"id"`
		RoomID   string `json:"room_id"`
		Decision string `json:"decision"`
	}
	n := &wsNextAgendaStruct{
		ID:       ctx.ID,
		RoomID:   ctx.Payload.RoomID,
		Decision: ctx.Payload.Decision,
	}
	c.ws.Send(ctx.Payload.RoomID, wsDecisionAgenda, n)

	// AgendaController_AddResult: end_implement
	return nil
}

// Next runs the next action.
func (c *AgendaController) Next(ctx *app.NextAgendaContext) error {
	// AgendaController_Next: start_implement

	// Put your logic here
	mongo := c.Mgo.DB(constant.DBName).C(constant.CollectionRoom)
	//query := bson.M{
	//	"room.room_id": ctx.Payload.RoomID,
	//	"room.agenda":  bson.M{"$elemMatch": bson.M{"id": ctx.ID}},
	//}
	//update := bson.M{
	//	"$currentDate": bson.M{
	//		"lastModified":         true,
	//		"room.agenda.$.end_at": bson.M{"$type": "timestamp"},
	//	},
	//}
	//err := mongo.Update(query, update)
	//if err != nil {
	//	return goa.ErrBadRequest(err)
	//}

	nextAgendaID := ctx.ID + 1
	change := bson.M{"$set": bson.M{"room.progress": nextAgendaID}}
	find := bson.M{"room.room_id": ctx.ID}
	err := mongo.Update(find, change)
	if err != nil {
		return goa.ErrBadRequest(err)
	}

	type wsNextAgendaStruct struct {
		ID int `json:"id"`
	}
	n := &wsNextAgendaStruct{
		ID: nextAgendaID,
	}
	c.ws.Send(ctx.Payload.RoomID, wsNextAgenda, n)

	// AgendaController_Next: end_implement
	return nil
}

// Websocket runs the websocket action.
func (c *AgendaController) Websocket(ctx *app.WebsocketAgendaContext) error {
	// AgendaController_Websocket: start_implement

	// Put your logic here
	c.ws.WebsocketServe(ctx.ID, ctx.ResponseWriter, ctx.Request)
	c.ws.Send(ctx.ID, "hello", "world")

	// AgendaController_Websocket: end_implement
	return nil
}
