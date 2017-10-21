package design

import (
	"os"

	. "app/design/constant"
	_ "app/design/models"
	_ "app/design/resource"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

const REPO = "parliament"

var _ = API(REPO, func() {
	Title(REPO)
	Description(REPO)
	Host(func() string {
		switch os.Getenv("Op") {
		case "develop":
			return "localhost:8080"
		case "staging":
			return "staging.com"
		case "production":
			return "production.com"
		}
		return "localhost:8080"
	}())
	Scheme(func() string {
		switch os.Getenv("Op") {
		case "develop":
			return "http"
		case "staging":
			return "https"
		case "production":
			return "https"
		}
		return "http"
	}())
	BasePath("/")
	Trait(AdminUserTrait, func() {
		Security(AdminAuth)
		Response(Unauthorized, ErrorMedia)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
	Trait(GeneralUserTrait, func() {
		Security(GeneralAuth)
		Response(Unauthorized, ErrorMedia)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
})
