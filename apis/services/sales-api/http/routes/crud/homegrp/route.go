package homegrp

import (
	"net/http"

	"github.com/ardanlabs/service/app/core/crud/homeapp"
	"github.com/ardanlabs/service/business/api/auth"
	midhttp "github.com/ardanlabs/service/business/api/mid/http"
	"github.com/ardanlabs/service/business/core/crud/home"
	"github.com/ardanlabs/service/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	HomeCore *home.Core
	Auth     *auth.Auth
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := midhttp.Authenticate(cfg.Auth)
	ruleAny := midhttp.Authorize(cfg.Auth, auth.RuleAny)
	ruleUserOnly := midhttp.Authorize(cfg.Auth, auth.RuleUserOnly)
	ruleAuthorizeHome := midhttp.AuthorizeHome(cfg.Auth, cfg.HomeCore)

	hdl := new(homeapp.New(cfg.HomeCore))
	app.Handle(http.MethodGet, version, "/homes", hdl.query, authen, ruleAny)
	app.Handle(http.MethodGet, version, "/homes/{home_id}", hdl.queryByID, authen, ruleAuthorizeHome)
	app.Handle(http.MethodPost, version, "/homes", hdl.create, authen, ruleUserOnly)
	app.Handle(http.MethodPut, version, "/homes/{home_id}", hdl.update, authen, ruleAuthorizeHome)
	app.Handle(http.MethodDelete, version, "/homes/{home_id}", hdl.delete, authen, ruleAuthorizeHome)
}
