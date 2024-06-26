package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/sashabaranov/go-openai"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	c := openai.NewClient(os.Getenv("API_KEY"))

	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Promt >")
	if !s.Scan() {
		panic("Failed to get User Input")
	}

	req := openai.ImageRequest{
		Prompt:         s.Text(),
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
	}
	resp, err := c.CreateImage(context.Background(), req)
	if err != nil {
		panic(err)
	}
	b, err := base64.StdEncoding.DecodeString(resp.Data[0].B64JSON)
	if err != nil {
		panic(err)
	}

	d, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}
	defer d.Close()
	d.Write(b)
	fmt.Println(resp.Data)
}
