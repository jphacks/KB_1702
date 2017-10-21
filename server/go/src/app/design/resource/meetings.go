package resource

import (
	. "app/design/constant"
	"app/design/media"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("meetings", func() {
	DefaultMedia(media.MeetingType)
	BasePath("/meetings")
	Action("meetings", func() {
		Description("http://kokkai.ndl.go.jp/api.html")
		Routing(
			GET(""),
		)
		Response(OK, CollectionOf(media.MeetingType))
		UseTrait(GeneralUserTrait)
	})
})
