package resource

import (
	. "app/design/constant"
	"app/design/media"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("speeches", func() {
	DefaultMedia(media.SpeechType)
	BasePath("/speeches")
	Action("speeches", func() {
		Description("http://kokkai.ndl.go.jp/api.html")
		Routing(
			GET(""),
		)
		Response(OK, CollectionOf(media.SpeechType))
		UseTrait(GeneralUserTrait)
	})
})
