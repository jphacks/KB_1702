package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("swaggerui", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swaggerui/*filepath", "swaggerui/")
})

var _ = Resource("swagger", func() {
	Files("/swagger.json", "swagger/swagger.json")
})

var _ = Resource("front", func() {
	Files("/", "public/html/index.html")
	Files("/rooms/*", "public/html/room.html")
	Files("/rooms", "public/html/create.html")
	Files("/dist/*filepath", "public/dist/")
})
