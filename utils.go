package slashred

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func PrintBytes(body io.Reader) {
	printBytes(body)
}
func printBytes(body io.Reader) {

	bodyBytes, err := ioutil.ReadAll(body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bodyBytes))
}
func dumpHeader(req *http.Request) {
	requestDump, err := httputil.DumpRequest(req, true)
if err != nil {
  fmt.Println(err)
}
fmt.Println(string(requestDump))
}
func respError(path string) {
	log.Fatalf("Error in getting reponse from path : %s\n", path)
}
func PrintHeader(resp *http.Response) {
	fmt.Println(resp.Status)

	fmt.Println("--------REQUEST HEADER--------")
	fmt.Printf("Ratelimit Used : %s\n", resp.Header.Get("X-Ratelimit-Used"))
	fmt.Printf("Ratelimit Remaining : %s\n", resp.Header.Get("X-Ratelimit-Remaining"))
	fmt.Printf("Ratelimit Reset : %s\n", resp.Header.Get("X-Ratelimit-Reset"))
	fmt.Println("--------REQUEST HEADER--------")

}

func SaveResponse(r io.Reader, filepath string) {

	f, err := os.Create(filepath)

	if err != nil {
		log.Fatal("Error in creating file")
	}

	defer f.Close()

	written, copyErr := io.Copy(f, r)

	if copyErr != nil {
		log.Fatal("Error in writing bytes to file")
	}

	fmt.Printf("Wrote %d bytes at %s\n", written, filepath)
}
