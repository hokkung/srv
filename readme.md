# Srv
A library for communicate with HTTP protocol.

## Installation
```
go get github.com/hokkung/srv
```

## Getting Started
1. Prepair a customizer to register a route
```
type ServerCustomizer struct {

}

func (c *ServerCustomizer) Register(s *server.Server) {
    s.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
```
2. Create a server and start
```
customizer := ServerCustomizer{}
server := server.NewServer(customizer)
s.Start()
```

## Environment Configuration ##
| Key | Description | Example | 
| --- | ----------- | ------- | 
| SERVER_ADDR | Port to start HTTP server | :8080 (default)
