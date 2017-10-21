package media

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var SpeechRecord = Type("speech-record", func() {
	Attribute("speech-order", Integer, "発言順(speechOrder)", func() {
		Default(0)
		Example(3)
	})
	Attribute("speaker", String, "発言者名(speaker)", func() {
		Default("")
		Example("安倍晋三")
	})
	Attribute("speech", String, "発言(speech)", func() {
		Default("")
		Example("○内閣総理大臣（安倍晋三君）　政権交代後、")
	})
	Required(
		"speech-order",
		"speaker",
		"speech",
	)
})

var SpeechType = MediaType("application/vnd.speechType+json", func() {
	Description("http://kokkai.ndl.go.jp/api.html")
	Attributes(func() {
		Attribute("id", Integer, "id", func() {
			Default(0)
			Example(1)
		})
		Attribute("session", Integer, "回次(session)", func() {
			Default(0)
			Example(193)
		})
		Attribute("name-of-house", String, "院名(nameOfHouse)", func() {
			Default("")
			Example("参議院")
		})
		Attribute("name-of-meeting", String, "会議名(nameOfMeeting)", func() {
			Default("")
			Example("決算委員会")
		})
		Attribute("issue", String, "号数(issue)", func() {
			Default("")
			Example("10号")
		})
		Attribute("date", DateTime, "開会日付(date)")
		Attribute("speech-record", ArrayOf(SpeechRecord))
		Attribute("meeting-url", String, "選択閲覧ページURL(meetingURL)", func() {
			Default("")
			Example("http://kokkai.ndl.go.jp/SENTAKU/sangiin/193/0015/19306050015010a.html<")
		})
		Attribute("pdf-url", String, "PDF画像URL(pdfURL)", func() {
			Default("")
			Example("http://kokkai.ndl.go.jp/SENTAKU/sangiin/193/0015/19306050015010.pdf")
		})
	})
	Required(
		"id",
		"session",
		"name-of-house",
		"name-of-meeting",
		"issue",
		"date",
		"speech-record",
		"meeting-url",
		"pdf-url",
	)
	View("default", func() {
		Attribute("id")
		Attribute("session")
		Attribute("name-of-house")
		Attribute("name-of-meeting")
		Attribute("issue")
		Attribute("date")
		Attribute("speech-record")
		Attribute("meeting-url")
		Attribute("pdf-url")
		Required(
			"id",
			"session",
			"name-of-house",
			"name-of-meeting",
			"issue",
			"date",
			"speech-record",
			"meeting-url",
			"pdf-url",
		)
	})
})
