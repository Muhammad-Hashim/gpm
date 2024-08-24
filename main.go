// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"strings"
// 	"golang.org/x/net/html"
// )

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Usage: gpm <package-name>")
// 		os.Exit(1)
// 	}

// 	packageName := os.Args[1]
// 	fmt.Printf("Package Name: %s\n", packageName)

// 	searchURL := "https://www.gpm.com/search?q=" + packageName
// 	fmt.Printf("Searching at URL: %s\n", searchURL)

// 	// Step 1: Search the website for the package
// 	resp, err := http.Get(searchURL)
// 	if err != nil {
// 		fmt.Printf("Error searching for package: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer resp.Body.Close()
//     func() {

//    }()
// 	fmt.Printf("Received HTTP response, status code: %d\n", resp.StatusCode)

// 	// Step 2: Parse the HTML to find the download link
// 	doc := html.NewTokenizer(resp.Body)
// 	var downloadURL string

// 	for {
// 		tokenType := doc.Next()
// 		switch tokenType {
// 		case html.ErrorToken:
// 			if downloadURL == "" {
// 				fmt.Println("Download link not found")
// 				os.Exit(1)
// 			}
// 			goto Download
// 		case html.StartTagToken, html.SelfClosingTagToken:
// 			token := doc.Token()
// 			if token.Data == "a" {
// 				for _, attr := range token.Attr {
// 					if attr.Key == "href" && strings.Contains(attr.Val, packageName) {
// 						downloadURL = attr.Val
// 						fmt.Printf("Found download URL: %s\n", downloadURL)
// 						break
// 					}
// 				}
// 			}
// 		}
// 	}

// Download:
// 	if downloadURL == "" {
// 		fmt.Println("No download URL found")
// 		os.Exit(1)
// 	}

// 	// Step 3: Download the package using the retrieved URL
// 	fmt.Printf("Downloading from %s...\n", downloadURL)

// 	resp, err = http.Get(downloadURL)
// 	if err != nil {
// 		fmt.Printf("Error downloading package: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer resp.Body.Close()

// 	out, err := os.Create(packageName + ".tar.gz")
// 	if err != nil {
// 		fmt.Printf("Error creating file: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer out.Close()

// 	_, err = io.Copy(out, resp.Body)
// 	if err != nil {
// 		fmt.Printf("Error saving file: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println("Package downloaded successfully!")
// }

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var packageMap = map[string]string{
	// Web Frameworks
	"gin":   "github.com/gin-gonic/gin",
	"echo":  "github.com/labstack/echo",
	"fiber": "github.com/gofiber/fiber",
	"chi":   "github.com/go-chi/chi",
	"mux":   "github.com/gorilla/mux",

	// Database
	"gorm": "gorm.io/gorm",
	"sqlx": "github.com/jmoiron/sqlx",
	"ent":  "github.com/ent/ent",
	"pgx":  "github.com/jackc/pgx/v4",

	// Authentication
	"jwt-go": "github.com/dgrijalva/jwt-go",
	"oauth2": "golang.org/x/oauth2",
	"casbin": "github.com/casbin/casbin",

	// HTTP Clients
	"axios": "github.com/imroc/req",
	"resty": "github.com/go-resty/resty/v2",

	// Utilities
	"viper":  "github.com/spf13/viper",
	"zap":    "go.uber.org/zap",
	"logrus": "github.com/sirupsen/logrus",
	"cobra":  "github.com/spf13/cobra",
	"uuid":   "github.com/google/uuid",

	// Testing
	"testify": "github.com/stretchr/testify",
	"gock":    "github.com/h2non/gock",
	"mockery": "github.com/vektra/mockery/v2",
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gpm <package-name>")
		os.Exit(1)
	}

	packageName := os.Args[1]
	fmt.Printf("Package Name: %s\n", packageName)

	// Step 1: Lookup the URL for the package
	urlPath, found := packageMap[packageName]
	if !found {
		fmt.Printf("Package %s not found in mapping\n", packageName)
		os.Exit(1)
	}

	// Channel to signal when the go get command is done
	done := make(chan bool)

	// Goroutine for showing the loading animation
	go func() {
		for {
			select {
			case <-done:
				fmt.Printf("\r%s\r", "                ") // Clear the line after loading
				return
			default:
				for _, r := range `-\|/` {
					fmt.Printf("\r%c Fetching package...", r)
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	}()

	// Run `go get` command
	goGetCmd := exec.Command("go", "get", urlPath)
	var out bytes.Buffer
	var stderr bytes.Buffer
	goGetCmd.Stdout = &out
	goGetCmd.Stderr = &stderr
	err := goGetCmd.Run()

	// Signal the loading animation to stop
	done <- true

	if err != nil {
		fmt.Printf("\nError running go get command: %v\n", err)
		fmt.Printf("Stdout: %s\n", out.String())
		fmt.Printf("Stderr: %s\n", stderr.String())
		os.Exit(1)
	}

	fmt.Println("Package successfully fetched using go get.")
	fmt.Println("You can now import the package in your Go project using the package name.")
}
