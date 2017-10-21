package media

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var RoomType = MediaType("room", func() {
	Description("")
	Attributes(func() {
		Attribute("id", Integer, "id", func() {
			Default(0)
			Example(1)
		})
	})
	Required(
		"id",
	)
	View("default", func() {
		Attribute("id")
		Required(
			"id",
		)
	})
})
