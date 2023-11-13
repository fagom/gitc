package internal

import (
	"fmt"
	"net/url"
	"os"
	"path"
)

func ParseUrl(rootPath string, endpoints ...string) string {
	rootAddr, err := url.Parse(rootPath)
	if err != nil {
		fmt.Println("Unable to parse URL. Exiting...")
		os.Exit(1)
	}

	for _, val := range endpoints {
		rootAddr.Path = path.Join(rootAddr.Path, val)
	}
	return rootAddr.String()
}
