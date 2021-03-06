//example from this websites
//https://www.bogotobogo.com/GoLang/GoLang_Web_Application_1.php
//https://www.bogotobogo.com/GoLang/GoLang_Web_Application_2.php

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Page struct {
	Title string
	Body  []byte
}

func (page *Page) writeOutPage() error {
	filename := page.Title + ".txt"
	return ioutil.WriteFile(filename, page.Body, 0600)
}
func readInPage(title string) (*Page, error) {
	filename := title + ".txt"

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, err
	}
	fmt.Println(dir)
	path := filepath.Join(dir, filename)

	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err

	}
	return &Page{Title: title, Body: body}, nil
}

// test case for readInPage() and writeOutPage()
func testFunctions() {
	page1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	page1.writeOutPage()
	page2, _ := readInPage("TestPage")
	fmt.Println(string(page2.Body))
}

// web server functions:
func handler(response http.ResponseWriter, request *http.Request) {
	//read in page
	testpage, err := readInPage("TestPage")
	if err != nil {
		//show err in console
		fmt.Println("\n", "readError: ", err)
		//print err to browser
		fmt.Fprintf(response, "<html>sh-app<h1>readError: %s</h1></html>", err)
	} else {
		//show page in console
		fmt.Println("\n", string(testpage.Title), string(testpage.Body))
		//print page to browser
		fmt.Fprintf(response, "<html>sh-app<h1>%s</h1><div>%s</div></html>", testpage.Title, testpage.Body)
	}

}
func whenLoadBrowserURL() {
	//when browser URL is loaded it starts handler function
	http.HandleFunc("/", handler)
}
func startWebServer() {
	log.Fatal(http.ListenAndServe(":8080", nil)) //listening to http://127.0.0.1:8080 bzw. http://localhost:8080
}
func shutdownWebServer() {
	//stop debugger or call
	os.Exit(0)
}

func main() {
	whenLoadBrowserURL()
	startWebServer()
}
