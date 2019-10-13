package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/akamensky/argparse"
	"github.com/gin-contrib/static"
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

	r.Use(static.Serve("/", static.LocalFile(*folder, true)))

	if *allowUpload {

		r.SetHTMLTemplate(
			template.Must(template.New("upload").Parse(`<html><head><title>Upload File</title></head>
<body><form enctype="multipart/form-data" action="#" method="POST">
  	<input type="file" name="file"/>
	<input type="submit" value="upload"/>
</form></body></html> `)))

		r.GET("/upload", func(c *gin.Context) {
			c.HTML(http.StatusOK, "upload", gin.H{})
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
	}

	fmt.Printf("Serving on: 127.0.0.1:%d\n", *port)

	err = r.Run(fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
