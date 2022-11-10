package app

import (
	"log"

	"github.com/ssibrahimbas/azure-k8s-demo.counter/src/config"
	"github.com/ssibrahimbas/azure-k8s-demo.counter/src/internal"
	cnf "github.com/ssibrahimbas/ssi-core/pkg/config"
	"github.com/ssibrahimbas/ssi-core/pkg/http"
	"github.com/ssibrahimbas/ssi-core/pkg/i18n"
	"github.com/ssibrahimbas/ssi-core/pkg/validator"
)

type App struct {
	Cnf *config.AppConfig
	I18n *i18n.I18n
	Http *http.Client
	Val *validator.Validator
	Repo *internal.Repo
	Srv *internal.Srv
	Hnd *internal.Handler
}

func New() *App {
	return &App{}
}

func (a *App) Init() *App {
	a.loadConfig()
	a.loadI18n()
	a.loadHttp()
	a.loadValidator()
	a.loadInternal()
	return a
}

func (a *App) Serve() {
	log.Fatal(a.Http.Listen(a.Cnf.Port))
}

func (a *App) loadConfig() {
	cnfg := config.AppConfig{}
	cnf.LoadConfig(&cnfg)
	a.Cnf = &cnfg
}

func (a *App) loadI18n() {
	i := i18n.New(a.Cnf.I18n.Fallback)
	i.LoadLanguages(a.Cnf.I18n.LocaleDir, a.Cnf.I18n.Locales...)
	a.I18n = i
}

func (a *App) loadHttp() {
	a.Http = http.New(a.I18n)
}

func (a *App) loadValidator() {
	val := validator.New(a.I18n)
	val.ConnectCustom()
	val.RegisterTagName()
	a.Val = val
}

func (a *App) loadInternal() {
	a.Repo = internal.NewRepo(&internal.RepoParams{})
	a.Srv = internal.NewSrv(&internal.SrvParams{
		Repo: a.Repo,
	})
	a.Hnd = internal.NewHandler(&internal.HandlerParams{
		Srv: a.Srv,
		Http: a.Http,
		I18n: a.I18n,
		Valid: a.Val,
	})
	a.Hnd.InitAllVersions()
}