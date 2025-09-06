package http

// const (
// 	serviceName = "softswitch-backend"
// 	version     = "beta"
// )

// type Server struct {
// 	Engine *gin.Engine
// }

// func NewServer() *Server {
// 	engine := gin.New()
// 	engine.Use(gin.Recovery())
// 	engine.Use(CORSMiddleware())
// 	engine.Use(allowOptionsMethod())
// 	// engine.Use(responseTimeMdw.Handler)
// 	engine.GET("/", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, map[string]interface{}{
// 			"service": serviceName,
// 			"version": version,
// 			"time":    time.Now().Unix(),
// 		})
// 	})
// 	server := &Server{Engine: engine}
// 	return server
// }

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
// 		c.Next()
// 	}
// }

// func allowOptionsMethod() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(http.StatusOK)
// 			return
// 		}
// 		c.Next()
// 	}
// }

// func (server *Server) Start(ctx context.Context, host string, port int) {
// 	go func() {
// 		if err := server.Engine.Run(fmt.Sprintf("%s:%d", host, port)); err != nil {
// 			panic(err)
// 		}
// 	}()
// 	logger.Infof("service %v listening on port %v", serviceName, port)
// 	<-ctx.Done()
// 	logger.Info("Gin stop")
// }
