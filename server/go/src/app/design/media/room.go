package media

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var Agenda = Type("agenda", func() {
	Attribute("id", String, "議題番号", func() {
		Example("dwadjjdwa")
	})
	Attribute("title", String, "議題", func() {
		Example("うぇいサウンド")
	})
	Attribute("goal", String, "議題のゴール", func() {
		Example("こんなことしたいな")
	})
	Attribute("decision", String, "結論", func() {
		Example("結論こんなん")
	})
	Attribute("time", Integer, "時間制限", func() {
		Example(10)
	})
	Attribute("start_at", DateTime, "開始時間")
	Attribute("end_at", DateTime, "終了時間")
	Attribute("child", ArrayOf(Child))
	Required(
		"id",
		"title",
		"goal",
		"decision",
		"time",
		"time",
		"child",
	)
})

var Child = Type("child", func() {
	Attribute("id", String, "議題番号", func() {
		Example("dwadjjdwa")
	})
	Attribute("title", String, "議題", func() {
		Example("うぇいサウンド")
	})
	Attribute("goal", String, "議題のゴール", func() {
		Example("こんなことしたいな")
	})
	Attribute("decision", String, "結論", func() {
		Example("結論こんなん")
	})
	Attribute("time", Integer, "時間制限", func() {
		Example(10)
	})
	Attribute("start_at", DateTime, "開始時間")
	Attribute("end_at", DateTime, "終了時間")
	Required(
		"id",
		"title",
		"goal",
		"decision",
		"time",
		"time",
	)
})

var Room = Type("room", func() {
	Attribute("name", String, "部屋名", func() {
		Example("うぇいサウンド部屋")
	})
	Attribute("agenda", ArrayOf(Agenda))
	Required(
		"agenda",
		"name",
	)
})

var RoomMongoType = MediaType("room_mongo", func() {
	Description("")
	Attributes(func() {
		Attribute("room", Room)
	})
	Required(
		"room",
	)
	View("default", func() {
		Attribute("room")
		Required(
			"room",
		)
	})
})
