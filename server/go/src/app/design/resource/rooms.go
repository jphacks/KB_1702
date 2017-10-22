package resource

import (
	. "app/design/constant"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("rooms", func() {
	BasePath("/api/rooms")
	Action("create", func() {
		Description("ルームを作成する")
		Routing(
			POST(""),
		)
		Response(SeeOther, String)
		UseTrait(GeneralUserTrait)
	})
	Action("show", func() {
		Description("ルームを作成する")
		Routing(
			POST("/:id"),
		)
		Params(func() {
			Param("id", String, "id")
			Required("id")
		})
		Response(OK, Any)
		UseTrait(GeneralUserTrait)
	})
})
