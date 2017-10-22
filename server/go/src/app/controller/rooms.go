package controller

import (
	"app/app"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"app/util"

	"app/design/constant"

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
	RoomID   string `json:"room_id" bson:"room_id"`
	Name     string `json:"name" bson:"name"`
	Progress int    `json:"progress" bson:"progress"`
	//StartAt  time.Time `json:"start_at" bson:"start_at"`
	//EndAt    time.Time `json:"end_at" bson:"end_at"`
	Agenda []Agenda `json:"agenda" bson:"agenda"`
}

type Agenda struct {
	ID       int    `json:"id" bson:"id"`
	Title    string `json:"title" bson:"title"`
	Goal     string `json:"goal" bson:"goal"`
	Decision string `json:"decision" bson:"decision"`
	Time     int    `json:"time" bson:"time"`
	//StartAt  time.Time `json:"start_at" bson:"start_at"`
	//EndAt    time.Time `json:"end_at" bson:"end_at"`
	Child []Agenda `json:"child" bson:"child"`
}

type InsertAgenda struct {
	ID   bson.ObjectId `json:"-" bson:"_id"`
	Room Room          `json:"room" bson:"room"`
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

	mongo := c.Mgo.DB(constant.DBName).C(constant.CollectionRoom)
	for {
		roomID := util.CreateTokenHash(data.Room.Name)
		fmt.Println(roomID)
		count, err := mongo.Find(bson.M{"room.room_id": roomID}).Count()
		if err != nil {
			return goa.ErrBadRequest(err)
		}
		fmt.Println(count)
		if count <= 0 {
			data.Room.RoomID = roomID
			break
		}
	}
	err = mongo.Insert(&data)
	if err != nil {
		return goa.ErrBadRequest(err)
	}
	// RoomsController_Create: end_implement
	ctx.ResponseData.Header().Set("Location", fmt.Sprintf("/rooms/%s", data.Room.RoomID))
	return ctx.SeeOther()
}

// Show runs the show action.
func (c *RoomsController) Show(ctx *app.ShowRoomsContext) error {
	// RoomsController_Show: start_implement

	// Put your logic here
	mongo := c.Mgo.DB(constant.DBName).C(constant.CollectionRoom)
	var agenda InsertAgenda
	err := mongo.Find(bson.M{"room.room_id": ctx.ID}).One(&agenda)
	if err != nil {
		return goa.ErrBadRequest(err)
	}

	// RoomsController_Show: end_implement
	return ctx.OK(agenda)
}
