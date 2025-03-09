package tasks

import (
	"fmt"
	"time"
)

//list of random funtions to simulate complex tasks

func Download() {
	fmt.Println("downloading files")
	time.Sleep(3 * (time.Second))
	fmt.Println("files downloaded")
}

func ProcessImage() {
	fmt.Println("Processing image...")
	time.Sleep(3 * time.Second)
	fmt.Println("Image processed.")
}

func SendEmail() {
	fmt.Println("Sending email...")
	time.Sleep(1 * time.Second)
	fmt.Println("Email sent.")
}

func CleanTempFiles() {
	fmt.Println("Cleaning temporary files...")
	time.Sleep(2 * time.Second)
	fmt.Println("Cleanup completed.")
}

func ConvertFile() {
	fmt.Println("Converting file format...")
	time.Sleep(3 * time.Second)
	fmt.Println("Conversion completed.")
}
