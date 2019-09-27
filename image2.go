// https://gist.github.com/sergiotapia/7882944
package main

import (
    "fmt"
    "image"
    "os"
    _ "image/jpeg"
    _ "image/png"
)

func main() {
    width, height := getImageDimension("rainy.jpg")
    fmt.Println("Width:", width, "Height:", height)
}

func getImageDimension(imagePath string) (int, int) {
    file, err := os.Open(imagePath)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
    }

    image, _, err := image.DecodeConfig(file)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
    }
    return image.Width, image.Height
}
