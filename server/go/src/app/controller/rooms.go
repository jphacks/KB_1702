package controller

import (
	"app/app"
	"encoding/json"
	"log"
	"time"

	"fmt"

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
	ID       bson.ObjectId `bson:"_id"`
	Name     string        `json:"name"`
	Progress int           `json:"progress"`
	StartAt  time.Time     `json:"start_at"`
	EndAt    time.Time     `json:"end_at"`
	Agenda   []Agenda      `json:"agenda"`
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
	Room   Room `json:"room"`
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

var jsonStr = `{
    "agenda": [
        {
            "id": 1,
            "title": "アイデア出し",
            "goal": "アイデアを10個出す",
            "time": 10,
            "start_at": "1995-01-11T06:25:13+09:00",
            "end_at": "1995-01-11T06:25:13+09:00",
            "child": [
                {
                    "id": 2,
                    "title": "テーマを考える",
                    "goal": "テーマ案を5個出す",
                    "time": 5,
                    "start_at": "1995-01-11T06:25:13+09:00",
                    "end_at": "1995-01-11T06:25:13+09:00",
                    "child": [
                        {
                            "id": 3,
                            "title": "ブレスト",
                            "goal": "アイデアを10個出す",
                            "time": 5,
                            "start_at": "1995-01-11T06:25:13+09:00",
                            "end_at": "1995-01-11T06:25:13+09:00"
                        }
                    ]
                },
                {
                    "id": 3,
                    "title": "ブレスト",
                    "goal": "アイデアを10個出す",
                    "time": 5,
                    "start_at": "1995-01-11T06:25:13+09:00",
                    "end_at": "1995-01-11T06:25:13+09:00"
                }
            ]
        },
        {
            "id": 4,
            "title": "アイデアを絞る",
            "goal": "アイデアを1個に絞る",
            "time": 10,
            "start_at": "1995-01-11T06:25:13+09:00",
            "end_at": "1995-01-11T06:25:13+09:00"
        }
    ]
}`

// Create runs the create action.
func (c *RoomsController) Create(ctx *app.CreateRoomsContext) error {
	// RoomsController_Create: start_implement

	// Put your logic here
	jsonBytes := ([]byte)(jsonStr)
	data := new(InsertAgenda)

	err := json.Unmarshal(jsonBytes, data)
	if err != nil {
		return goa.ErrBadRequest(err)
	}
	data.Room.ID = bson.NewObjectId()
	mongo := c.Mgo.DB("test").C("jphacks")
	err = mongo.Insert(&data)
	if err != nil {
		return goa.ErrBadRequest(err)
	}

	//idStr := data.Room.ID.String()
	//if !bson.IsObjectIdHex(idStr) {
	//	return goa.ErrBadRequest(goa.ErrBadRequest(err))
	//}
	//id := bson.ObjectIdHex(idStr)

	var men []InsertAgenda
	if err := mongo.Find(nil).All(&men); err != nil {
		log.Fatal(err)
	}
	//spew.Dump(men)
	fmt.Println(data.Room.ID)
	fmt.Println(bson.ObjectIdHex("59eb24144e25ba1f6a3d063a"))

	var result InsertAgenda
	if err := mongo.FindId(bson.ObjectIdHex("59eb24144e25ba1f6a3d063a")).One(&result); err != nil {
		return goa.ErrInternal(err)
	}
	//spew.Dump(result)

	//idStr := "565edd868bc93d268a13bc02"
	//if !bson.IsObjectIdHex(idStr) {
	//	log.Fatal("not objectId")
	//}
	//id := bson.ObjectIdHex(idStr)
	//
	//// Collection People
	//// Query One
	//result := InsertAgenda{}
	//
	//err = mongo.FindId(id).One(&result)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Phone", result)

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
