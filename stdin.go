package main

import (
    // "bufio"
    "fmt"
    // "io"
    // "log"
    "os"
)

func YesNO() {
    userFile := "/dev/tty"
here:
    fmt.Printf("Continue (y/n):")
    fin, err := os.Open(userFile)
    defer fin.Close()
    if err != nil {
        fmt.Println(userFile, err)
        return
    }
    buf := make([]byte, 1024)

    for {
        n, _ := fin.Read(buf)
        if 0 == n || n > 1 {
            break
        }
    }
    switch string(buf[:2]) {
    case "y\n":
        fmt.Println("go on ...")
    case "n\n":
        fmt.Println("exit ...")
        os.Exit(1)
    default:
        goto here
    }
}