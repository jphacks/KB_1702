package resource

import (
	. "app/design/constant"
	"app/design/media"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("rooms", func() {
	DefaultMedia(media.RoomType)
	BasePath("/api/rooms")
	Action("create", func() {
		Description("ルームを作成する")
		Routing(
			POST(""),
		)
		Response(SeeOther)
		UseTrait(GeneralUserTrait)
	})
	Action("show", func() {
		Description("ルームを作成する")
		Routing(
			POST("/:id"),
		)
		Response(OK, Any)
		UseTrait(GeneralUserTrait)
	})
})
