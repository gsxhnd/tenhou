package middleware

type Middlewarer interface {
	// RequestLog(ctx *fiber.Ctx) error
	// Websocket(ctxc *fiber.Ctx) error
}
type middleware struct {
	// logger utils.Logger
	// tracer utils.Tracer
}

func NewMiddleware() Middlewarer {
	return &middleware{
		// logger: l,
		// tracer: t,
	}
}
