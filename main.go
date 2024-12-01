package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

const (
	defaultTemplate = `<!DOCTYPE html>
  <html>
	<head>
	  <meta http-equiv="content-type" content="text/html; charset=utf-8">
	  <title>Preview</title>
	</head>
	<body>
	  {{ . }}
	</body>
  </html>
  `
	fName = "preview_assets.html"
)

var ext = []string{".png", ".jpg", ".jpeg", ".gif", ".svg"}

func main() {

	path := flag.String("path", "", "Directory Path")
	flag.Parse()

	if *path == "" {
		curDir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		*path = curDir
	}

	if err := run(*path); err != nil {
		fmt.Println(err)
	}

}

func run(path string) error {

	/// 1. Read the directory path
	/// 2. Find all the assets directory
	/// 3. Add in template
	/// 4. Find all the .png, .jpg, .jpeg, .gif, .svg files in the assets directory
	/// 5. Add in template
	/// 6. Create a new html file
	/// 7. Open the html file in browser
	/// 8. Done

	htmlContent, err := parseContent(path)
	if err != nil {
		return err
	}

	err = saveHTML(fName, htmlContent)
	if err != nil {
		return err
	}

	return preview(fName)
}

func parseContent(path string) ([]byte, error) {

	// Parse the contents of the defaultTemplate const into a new Template
	t, err := template.New("preview").Parse(defaultTemplate)
	if err != nil {
		return nil, err
	}

	var body template.HTML

	err = filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		/// Check if the path is a directory with name assets
		if info.IsDir() && info.Name() == "assets" {
			// fmt.Println("Found assets directory", path)
			body += template.HTML(fmt.Sprintf("<h1>%s</h1>", path))

			// Read the directory
			err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}

				if !d.IsDir() {
					// Check if the file is an image
					for _, e := range ext {
						if filepath.Ext(path) == e {
							body += template.HTML(fmt.Sprintf("<img src=\"%s\" width=\"200\" height=\"200\" alt=\"%s\" />  %s  </br> ", path, path, d.Name()))
							break
						}
					}
				}

				return nil
			})

			if err != nil {
				return err
			}

		}

		return nil

	})

	if err != nil {
		return nil, err
	}

	// Create a buffer of bytes to write to file
	var buffer bytes.Buffer

	// Execute the template with the content type
	if err := t.Execute(&buffer, body); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func saveHTML(outFname string, data []byte) error {
	// Write the bytes to the file
	return os.WriteFile(outFname, data, 0644)
}

func preview(fname string) error {
	cName := ""
	cParams := []string{}

	// Define executable based on OS
	switch runtime.GOOS {
	case "linux":
		cName = "xdg-open"
	case "windows":
		cName = "cmd.exe"
		cParams = []string{"/C", "start"}
	case "darwin":
		cName = "open"
	default:
		return fmt.Errorf("OS not supported")
	}

	// Append filename to parameters slice
	cParams = append(cParams, fname)

	// Locate executable in PATH
	cPath, err := exec.LookPath(cName)
	if err != nil {
		return err
	}

	// Open the file using default program
	err = exec.Command(cPath, cParams...).Run()

	// Give the browser some time to open the file before deleting it
	time.Sleep(2 * time.Second)
	return err
}
