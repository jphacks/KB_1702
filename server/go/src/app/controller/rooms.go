package controller

import (
	"app/app"
	"time"

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

type Room struct {
	Room struct {
		ID       string    `json:"id"`
		Name     string    `json:"name"`
		Progress int       `json:"progress"`
		StartAt  time.Time `json:"start_at"`
		EndAt    time.Time `json:"end_at"`
		Agenda   []Agenda  `json:"agenda"`
	} `json:"room"`
}

type Agenda struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Goal    string    `json:"goal"`
	Time    int       `json:"time"`
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
	Child   *[]Agenda `json:"child"`
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
                    "end_at": "1995-01-11T06:25:13+09:00"
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
	//con := c.Mgo.DB("test").C("people")
	//byteArray, _ := ioutil.ReadAll(ctx.Request.Body)
	//jsonBytes := ([]byte)(byteArray)
	//fmt.Println(jsonBytes)

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
