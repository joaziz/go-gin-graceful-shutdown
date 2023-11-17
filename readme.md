# Gin Graceful Shutdown Package
This package provides a simple and convenient way to implement graceful shutdown for Gin HTTP servers in Golang. It ensures that ongoing requests are completed before the server is gracefully terminated in response to OS signals.

## Installation
To use this package, you can simply run:

```bash
go get -u github.com/joaziz/go-gin-graceful-shutdown
```



```go

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joaziz/go-gin-graceful-shutdown"
	"time"
)

func main() {
	
	server := gin.Default()
	
	
	server.GET("/", "....")

	go_gin_graceful_shutdown.Serve(&go_gin_graceful_shutdown.Options{
		Port:        8080,
		Engin:       server,
		WaitTimeout: 20 * time.Second,
	})

}

```

## Contributing
We welcome contributions! Feel free to open issues, submit pull requests, or provide feedback.

## License
This project is licensed under the MIT License - see the LICENSE file for details.