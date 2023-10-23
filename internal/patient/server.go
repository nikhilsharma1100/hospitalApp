package patient

// Server struct for account server. Boot handler creates instance of server.
type Server struct {
	Core ICore
}

// NewServer Boot handler calls this method to create a new account server
func NewServer() *Server {
	return &Server{
		Core: NewCore(),
	}
}
