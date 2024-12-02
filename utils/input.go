package utils

import (
    "bufio"
    "os"
)

func OpenFileScanner(path string) (*bufio.Scanner, func(), error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, nil, err
    }

    cleanup := func() {
        file.Close()
    }

    scanner := bufio.NewScanner(file)
    return scanner, cleanup, nil
}

