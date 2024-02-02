# go-gin-graceful-shutdown
`go-gin-graceful-shutdown` is a lightweight, easy-to-use package that seamlessly integrates graceful shutdown capabilities into your Gin-based HTTP servers. Designed with simplicity and efficiency in mind, it ensures that your web applications can handle termination signals gracefully, allowing ongoing requests to complete before the server shuts down. This mechanism is crucial for maintaining high availability and reliability of services, especially in production environments.


## Features
- **Simple Integration**: Easily integrate with any Gin-based project with minimal setup.
- **Customizable Wait Timeout**: Configure how long the server should wait for active requests to complete before shutting down.
- **Automatic Signal Handling**: Automatically listens for termination signals (SIGINT/SIGTERM) to initiate the graceful shutdown process.
- **Zero Dependency on External Libraries**: Built to work with the standard Gin framework without needing extra dependencies.
- **Open-Source and Community-Driven**: Contributions, feedback, and improvements are welcome to make the package more robust and feature-rich.



## Installation
To integrate `go-gin-graceful-shutdown` into your project, execute the following command in your terminal:

```bash
go get -u github.com/joaziz/go-gin-graceful-shutdown
```
This command fetches the package and adds it to your project's dependencies.

## Quick Start
Here's a quick guide to get you started with integrating graceful shutdown capabilities into your Gin server:

```go

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joaziz/go-gin-graceful-shutdown"
	"time"
)

func main() {
	// Initialize your Gin router
	router := gin.Default()

	// Define your routes
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})

	// Configure and start the server with graceful shutdown
	go_gin_graceful_shutdown.Serve(&go_gin_graceful_shutdown.Options{
		Port:        8080,
		Engine:      router,
		WaitTimeout: 20 * time.Second,
	})
}


```

## Configuration Options
The Serve function accepts an Options struct where you can customize the behavior of your graceful shutdown:
- **Port**: The port number on which the server listens.
- **Engine**: Your Gin engine instance.
- **WaitTimeout**: Duration to wait for active requests to complete before shutting down.



## Contributing
We welcome contributions from the community! Whether it's submitting bugs, proposing new features, or improving documentation, every contribution is valuable. Here's how you can contribute:

- **Reporting Issues**: Use the GitHub Issues section to report bugs and feature requests.
- **Submitting Pull Requests**: Submit PRs for bug fixes or feature additions. Please ensure your code follows the existing style and has been tested.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments
- Special thanks to the Gin community for the fantastic web framework.
- Everyone who has contributed to making this package better!