package slashred

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

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
