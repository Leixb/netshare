package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/akamensky/argparse"
	"github.com/gin-gonic/gin"
)

func main() {

	parser := argparse.NewParser("", "Share folder on local network")

	allowUpload := parser.Flag("u", "upload", &argparse.Options{
		Help: "Allow upload to folder",
	})
	folder := parser.String("f", "folder", &argparse.Options{
		Default: ".",
		Help:    "Folder to share",
	})

	port := parser.Int("p", "port", &argparse.Options{
		Default: 8080,
		Help:    "Port to use",
	})

	debug := parser.Flag("d", "debug", &argparse.Options{
		Help: "Gin in debug mode",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	if !*debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	if *allowUpload {

		r.StaticFS("/browse", http.Dir(*folder))

		r.GET("/", func(c *gin.Context) {
			c.Data(http.StatusOK, "text/html", []byte(frame))
		})

		r.GET("/upload", func(c *gin.Context) {
			c.Data(http.StatusOK, "text/html", []byte(upload))
		})
		r.POST("/upload", func(c *gin.Context) {

			file, err := c.FormFile("file")
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}

			filename := filepath.Join(*folder, filepath.Base(file.Filename))

			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}

			fmt.Printf("uploaded file: %s\n", filename)
			c.String(http.StatusOK, fmt.Sprintf("Uploaded %s", filename))

		})
	} else {
		r.StaticFS("/", http.Dir(*folder))
	}

	fmt.Printf("Serving on: 127.0.0.1:%d\n", *port)

	err = r.Run(fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
