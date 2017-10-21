package controller

import (
	"app/app"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/goadesign/goa"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

type Room struct {
	RoomID   string    `json:"room_id"`
	Name     string    `json:"name"`
	Progress int       `json:"progress"`
	StartAt  time.Time `json:"start_at"`
	EndAt    time.Time `json:"end_at"`
	Agenda   []Agenda  `json:"agenda"`
}

type Agenda struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Goal    string    `json:"goal"`
	Time    int       `json:"time"`
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
	Child   []Agenda  `json:"child"`
}

type InsertAgenda struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	Room   Room          `json:"room"`
	Agenda []struct {
		ID      int       `json:"id"`
		Title   string    `json:"title"`
		Goal    string    `json:"goal"`
		Time    int       `json:"time"`
		StartAt time.Time `json:"start_at"`
		EndAt   time.Time `json:"end_at"`
		Child   []Agenda  `json:"child"`
	} `json:"agenda"`
}

// Create runs the create action.
func (c *RoomsController) Create(ctx *app.CreateRoomsContext) error {
	// RoomsController_Create: start_implement

	// Put your logic here
	jsonBytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return goa.ErrBadRequest(err)
	}
	data := new(InsertAgenda)
	err = json.Unmarshal(jsonBytes, data)
	if err != nil {
		return goa.ErrBadRequest(err)
	}
	data.ID = bson.NewObjectId()
	mongo := c.Mgo.DB("test").C("jphacks")
	err = mongo.Insert(&data)
	if err != nil {
		return goa.ErrBadRequest(err)
	}
	var agenda InsertAgenda
	if err := mongo.FindId(data.ID).One(&agenda); err != nil {
		return goa.ErrBadRequest(err)
	}

	// RoomsController_Create: end_implement
	ctx.ResponseData.Header().Set("Location", fmt.Sprintf("/rooms/%s", data.ID.Hex()))
	return ctx.SeeOther()
}

// Show runs the show action.
func (c *RoomsController) Show(ctx *app.ShowRoomsContext) error {
	// RoomsController_Show: start_implement

	// Put your logic here
	if !bson.IsObjectIdHex(ctx.ID) {
		return goa.ErrBadRequest(fmt.Errorf("idの形式がmongoに合っていません"))
	}
	id := bson.ObjectIdHex(ctx.ID)

	mongo := c.Mgo.DB("test").C("jphacks")
	var agenda InsertAgenda
	if err := mongo.FindId(id).One(&agenda); err != nil {
		return goa.ErrBadRequest(err)
	}

	// RoomsController_Show: end_implement
	return ctx.OK(agenda)
}
