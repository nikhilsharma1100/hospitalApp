package doctor

// Server struct for account server. Boot handler creates instance of server.
type Server struct {
	core ICore
}

// NewServer Boot handler calls this method to create a new account server
func NewServer(core ICore) *Server {
	return &Server{
		core: core,
	}
}
