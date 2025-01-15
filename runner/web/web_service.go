package web

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

type Router interface {
	chi.Router
}

type WebModuleRegistry interface {
	RegisterRoutesTo(router chi.Router) error
}

type WebModuleRegistryFunc func(router chi.Router) error

func (f WebModuleRegistryFunc) RegisterRoutesTo(router chi.Router) error {
	return f(router)
}

type BeforeStartHook interface {
	BeforeStart(ctx context.Context, ws *WebService) error
}

type BeforeStartHookFunc func(ctx context.Context, ws *WebService) error

func (f BeforeStartHookFunc) BeforeStart(ctx context.Context, ws *WebService) error {
	return f(ctx, ws)
}

type WebServiceHooks struct {
	BeforeStart BeforeStartHook
}

type HTTPServer interface {
	SetAddr(addr string)
	SetHandler(handler http.Handler)
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type HTTPServerGo struct {
	*http.Server
}

func NewHTTPServerGo(server *http.Server) *HTTPServerGo {
	return &HTTPServerGo{
		Server: server,
	}
}

func (s *HTTPServerGo) ListenAndServe() error {
	return s.Server.ListenAndServe()
}

func (s *HTTPServerGo) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}

func (s *HTTPServerGo) SetAddr(addr string) {
	s.Server.Addr = addr
}

func (s *HTTPServerGo) SetHandler(handler http.Handler) {
	s.Server.Handler = handler
}

type WebServiceOptions struct {
	router     chi.Router
	address    string
	httpServer HTTPServer
	hooks      WebServiceHooks
}

type WebServiceOptionSetter func(opts *WebServiceOptions)

func WithRouter(router chi.Router) WebServiceOptionSetter {
	return func(opts *WebServiceOptions) {
		opts.router = router
	}
}

func WithAddress(address string) WebServiceOptionSetter {
	return func(opts *WebServiceOptions) {
		opts.address = address
	}
}

func WithHTTPServer(httpServer HTTPServer) WebServiceOptionSetter {
	return func(opts *WebServiceOptions) {
		opts.httpServer = httpServer
	}
}

func OnBeforeStart(s BeforeStartHook) WebServiceOptionSetter {
	return func(opts *WebServiceOptions) {
		opts.hooks.BeforeStart = s
	}
}

func defaultSettings() WebServiceOptions {
	return WebServiceOptions{
		router:     chi.NewRouter(),
		address:    ":8080",
		httpServer: NewHTTPServerGo(&http.Server{}),
	}
}

type WebService struct {
	opts             WebServiceOptions
	moduleRegistries []WebModuleRegistry
}

func NewWebService(optSetters ...WebServiceOptionSetter) *WebService {
	opt := defaultSettings()
	for _, setter := range optSetters {
		setter(&opt)
	}

	return &WebService{
		opts:             opt,
		moduleRegistries: []WebModuleRegistry{},
	}
}

func (ws *WebService) Run(ctx context.Context) error {

	if err := ws.runHooks(ctx); err != nil {
		return err
	}

	if err := ws.applyModuleRegistries(); err != nil {
		return err
	}

	httpServer := ws.opts.httpServer

	httpServer.SetAddr(ws.opts.address)
	httpServer.SetHandler(ws.opts.router)

	go func() {
		<-ctx.Done()
		logrus.Debug("Shutting down web service")
		httpServer.Shutdown(context.Background())
	}()

	logrus.Info(map[string]interface{}{
		"address": ws.opts.address,
		"event":   "web_service_started"})

	return httpServer.ListenAndServe()
}

func (ws *WebService) Router() chi.Router {
	return ws.opts.router
}

func (ws *WebService) RegisterModuleRegistry(registry ...WebModuleRegistry) {
	ws.moduleRegistries = append(ws.moduleRegistries, registry...)
}

func (ws *WebService) runHooks(ctx context.Context) error {

	logrus.Debug("Running before start hooks")

	if ws.opts.hooks.BeforeStart != nil {
		logrus.Debug("Running before start hooks")
		return ws.opts.hooks.BeforeStart.BeforeStart(ctx, ws)
	}

	logrus.Debug("All hooks run")
	return nil
}

func (ws *WebService) applyModuleRegistries() error {

	logrus.Debug(logrus.Fields{"registriesLen": len(ws.moduleRegistries), "event": "applying_module_registries"})

	for _, registry := range ws.moduleRegistries {
		if err := registry.RegisterRoutesTo(ws.opts.router); err != nil {
			logrus.Error(logrus.Fields{"event": "failed_to_register_routes", "error": err.Error()})
			return err
		}
	}

	logrus.Debug("All module registries applied")
	return nil
}
