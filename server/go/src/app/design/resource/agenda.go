package resource

import (
	. "app/design/constant"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("agenda", func() {
	BasePath("/api")
	Action("next", func() {
		Description("次のagendaIDを取得する（progressも書き換える）")
		Routing(
			POST("/agenda/:id/next"),
		)
		Params(func() {
			Param("id", Integer, "終了した議題ID", func() {
				Example(2)
			})
		})
		Payload(func() {
			Param("room_id", String, "ルームID", func() {
				Example("dwadadawdha")
			})
			Required("room_id")
		})
		Response(OK)
		UseTrait(GeneralUserTrait)
	})
	Action("addDecision", func() {
		Description("アジェンダの結論をかき出す")
		Routing(
			POST("/agenda/:id/decision"),
		)
		Params(func() {
			Param("id", Integer, "アジェンダID")
			Required("id")
		})
		Payload(func() {
			Param("room_id", String, "ルームID", func() {
				Example("dwadwadhlaw")
			})
			Param("decision", String, "結論", func() {
				Example("いい感じだった")
			})
			Required("room_id", "decision")
		})
		Response(OK)
		UseTrait(GeneralUserTrait)
	})
	Action("websocket", func() {
		Description("参加中インターンの新着メモをpush配信する。")
		Routing(GET("/rooms/ws"))
		Params(func() {
			Param("id", String, "id")
			Required("id")
		})
		UseTrait(WebsocketTrait)
		Response(SwitchingProtocols)
	})
	Action("addReaction", func() {
		Description("次のagendaIDを取得する（progressも書き換える）")
		Routing(
			POST("/users/:id/reaction"),
		)
		Params(func() {
			Param("id", String, "アジェンダID", func() {
				Example("dwuiadoawbdbawdba")
			})
			Required("id")
		})
		Payload(func() {
			Param("from", String, "リアクションを送った人", func() {
				Example("dwuiadoawbdbawdba")
			})
			Param("to", String, "リアクションを送られた人", func() {
				Example("dwuiadoawbdbawdba")
			})
			Required("from", "to")
		})
		Response(OK)
		UseTrait(GeneralUserTrait)
	})
})
