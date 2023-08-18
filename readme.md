# Srv
A library for communicate with HTTP protocol.

## Install

## Getting Started 
1. Prepair a customizer for register a route
```
type ServerCustomizer struct {

}

func (c *ServerCustomizer) Register(e *gin.Engine) {
    e.GET("/ping", func(c *gin.Context) {
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
| SERVER_ADDR | Port to start HTTP server | :8080

