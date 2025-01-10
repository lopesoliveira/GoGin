There are several frameworks in Go that make it easier to build web APIs, similar to how Micronaut in Java or ASP.NET in C# work. Some of the most popular frameworks for building web APIs in Go include:

Gin: A high-performance HTTP web framework.
Echo: A fast and minimalist web framework.
Fiber: An Express-inspired web framework.
Revel: A high-productivity web framework.

mkdir GoGin
cd GoGin

Install Gin
$ go get -u github.com/gin-gonic/gin
This command will update your go.mod and go.sum files with the Gin package and its dependencies.

Create a file named main.go in your project directory and add the following code:
```
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    r.GET("/hello", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, World!",
        })
    })

    r.Run(":8080") // Run on port 8080
}
```

Run the application
go run main.go

The application will start, and you can access the API at http://localhost:8080/hello.




