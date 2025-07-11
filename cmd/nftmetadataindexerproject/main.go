// cmd/nftmetadataindexerproject/main.go
package main

import (
"flag"
"log"
"os"

"nftmetadataindexerproject/internal/nftmetadataindexerproject"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nftmetadataindexerproject.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
