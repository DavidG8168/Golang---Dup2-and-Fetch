// EX 1.9 -- START
// Fetch prints the content found at a URL.
package main
import (
"fmt"
// Use this package instead.
"io"
"net/http"
"os"
)
func main() {
    // Keep this part the same.
    // Iterate command line argument urls.
    for _, url := range os.Args[1:] {
        resp, err := http.Get(url)
        // Print errors if encountered.
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
	// Add printing the HTTP status code here.
	fmt.Printf("The HTTP status is: %s\n", resp.Status)
        // Modify function here.
        // Copy response body to stdout.
        // io.Copy(dst, src)
        _, err = io.Copy(os.Stdout, resp.Body)
        resp.Body.Close()
        if err != nil {
            // Modify to print only error.
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", err)
            os.Exit(1)
        }
    }
}
// EX 1.9 -- END