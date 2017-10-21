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
changeAgenda
リクエストの中に、ルームのIDと次のアジェンダのID
プログレッシブをいまの議題のIDに書き換える

reaction
from to type
これをmongoにぶち込みたい

result
id, agenda_id, result
agenda_id.resultを書き換える
*/

const (
	// id
	wsNextAgenda = "nextAgenda"
	// from to type
	wsReaction = "reaction"
	// id, result
	wsResult = "result"
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

	// AgendaController_AddReaction: end_implement
	return nil
}

// AddResult runs the addResult action.
func (c *AgendaController) AddResult(ctx *app.AddResultAgendaContext) error {
	// AgendaController_AddResult: start_implement

	// Put your logic here

	// AgendaController_AddResult: end_implement
	return nil
}

// Next runs the next action.
func (c *AgendaController) Next(ctx *app.NextAgendaContext) error {
	// AgendaController_Next: start_implement

	// Put your logic here
	mongo := c.Mgo.DB(constant.DBName).C(constant.CollectionRoom)
	nextAgendaID := ctx.Payload.FinishAgendaID + 1
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
	c.ws.Send(ctx.ID, wsNextAgenda, n)

	// AgendaController_Next: end_implement
	return nil
}

// Websocket runs the websocket action.
func (c *AgendaController) Websocket(ctx *app.WebsocketAgendaContext) error {
	// AgendaController_Websocket: start_implement

	// Put your logic here
	c.ws.WebsocketServe(ctx.ID, ctx.ResponseWriter, ctx.Request)

	// AgendaController_Websocket: end_implement
	return nil
}
