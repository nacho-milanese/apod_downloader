package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func main() {
	destPath := "image_download"
	var apod_url string = "https://apod.nasa.gov/"
	// Gets the html from the site, to be parsed
	wholeHtml := apod_fetch_html(apod_url)
	// regex out the image url + the prefix apod site
	var re = regexp.MustCompile(`IMG SRC="((\\")|[^"(\\")])+"`)
	matches := re.FindStringSubmatch(string(wholeHtml))
	// gets the image url
	imageUrl := apod_url + matches[0][9:len(matches[0])-1]

	apod_fetch_html(apod_url)
	url2file(imageUrl, destPath)
}

func apod_fetch_html(imageUrl string) []byte {
	// Gets a url and returns html in byte code
	resp, err := http.Get(imageUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return html
}

func url2file(apod_img string, destPath string) string {
	// gets a url and a subdir name in local path, downloads it's content to a file and returns the filename

	// Build fileName from fullPath
	fileURL, err := url.Parse(apod_img)
	if err != nil {
		panic(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]
	fmt.Printf("%s created \n", fileName)

	// Create blank file
	file, err := os.Create(destPath + "/" + fileName)
	if err != nil {
		panic(err)
	}

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(apod_img)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.Copy(file, resp.Body)
	defer file.Close()

	return fileName
}
