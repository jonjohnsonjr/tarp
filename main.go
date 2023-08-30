package main

import (
	"archive/tar"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
)

func main() {
	if err := mainE(); err != nil {
		log.Fatal(err)
	}
}

func mainE() error {
	tr := tar.NewReader(os.Stdin)

	enc := json.NewEncoder(os.Stdout)
	for {
		hdr, err := tr.Next()
		if errors.Is(err, io.EOF) {
			return nil
		} else if err != nil {
			return err
		}

		if err := enc.Encode(hdr); err != nil {
			return err
		}
	}
}
