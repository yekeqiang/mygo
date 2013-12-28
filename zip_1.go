package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
)

func ZipFile() {
	file, _ := os.Create("test.zip")
	defer file.close()

	writer := zip.NewWriter(file)
	defer writer.Close()

	datas := []struct{ name, body string }{
		{"a.txt", "Hello,World"},
		{"b.txt", "112344......"},
	}

	for _, data := range datas {
		w, _ := writer.Create(data.name)
		w.Write([]byte(data.body))
	}

}

func UnZipFile() {
	zr, _ := zip.OpenReader("test.zip")
	defer zr.Close()

	for _, f := range zr.File {
		br, _ := f.Open()
		bs, _ := ioutil.ReadAll(br)

		fmt.Printf("%s: %s\n", f.Name, string(bs))
	}
}

func main() {
	ZipFile()
	UnZipFile()
}
