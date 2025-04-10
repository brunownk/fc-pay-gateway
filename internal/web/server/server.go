package server

import (
	"net/http"

	"github.com/brunownk/fc-pay-gateway/internal/service"
	"github.com/brunownk/fc-pay-gateway/internal/web/handlers"
	"github.com/brunownk/fc-pay-gateway/internal/web/middleware"
	"github.com/go-chi/chi/v5"
)

// Server gerencia as configurações e rotas da API
type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	invoiceService *service.InvoiceService
	port           string
}

// NewServer inicializa um novo servidor com suas dependências
func NewServer(accountService *service.AccountService, invoiceService *service.InvoiceService, port string) *Server {
	s := &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		invoiceService: invoiceService,
		port:           port,
	}
	s.ConfigureRoutes()
	return s
}

// ConfigureRoutes define todas as rotas da API e seus middlewares
func (s *Server) ConfigureRoutes() {
	accountHandler := handlers.NewAccountHandler(s.accountService)
	invoiceHandler := handlers.NewInvoiceHandler(s.invoiceService)
	authMiddleware := middleware.NewAuthMiddleware(s.accountService)

	// Rotas públicas - criação e consulta de contas
	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)

	// Grupo de rotas protegidas por autenticação via API Key
	s.router.Group(func(r chi.Router) {
		r.Use(authMiddleware.Authenticate)
		// Rotas de faturas - criação, consulta e listagem
		r.Post("/invoice", invoiceHandler.Create)
		r.Get("/invoice/{id}", invoiceHandler.GetByID)
		r.Get("/invoice", invoiceHandler.ListByAccount)
	})
}

// Start inicia o servidor HTTP na porta configurada
func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}
	return s.server.ListenAndServe()
}
