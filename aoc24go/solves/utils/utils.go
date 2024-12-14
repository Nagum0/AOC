package utils

import (
	"fmt"
	"os"
)

func ReadFileToString(path string) string {
    data, err := os.ReadFile(path)
    
    if err != nil {
        fmt.Printf("Error while reading file contents: %v\n", err.Error())
        os.Exit(1)
    }

    return string(data)
}
