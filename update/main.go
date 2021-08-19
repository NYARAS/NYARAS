package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/mmcdole/gofeed"
)

func makeReadme(filename string) error {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://calvineotieno010.medium.com/")
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
	}
	// Get the freshest item
	blogItem := feed.Items[0]

	wc, err := fp.ParseURL("https://calvineotieno010.medium.com/")
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
	}
	// Add this much magic
	wcItem := wc.Items[0]

	date := time.Now().Format("2 Jan 2006")

	// Whisk together static and dynamic content until stiff peaks form
	hello := "### Hello! I’m Calvine Otieno.\n\nI,m DevOps Engineer. I am involved in designing and implementing automation processes, CI/CD, Infrastructure as Code (IaC), monitoring, and logging. Am also interested in designing systems to take advantage of the automation and scalability of cloud-native, micro service-oriented, and containerized environments using technology such as Docker and Kubernetes " + wcItem.Description + " words I’ve written on [medium.com](https://https://calvineotieno010.medium.com/). I hope to empower people to learn openly and fearlessly through knowledge sharing and technology leadership."
	blog := "You might like my latest blog post: **[" + blogItem.Title + "](" + blogItem.Link + ")**. You can subscribe to my [**blog RSS**](https://calvineotieno010.medium.com/) or by email at [**medium.com**](https://https://calvineotieno010.medium.com/)."
	updated := "<sub>Last updated by magic on " + date + ".</sub>"
	data := fmt.Sprintf("%s\n\n%s\n\n%s\n", hello, blog, updated)

	// Prepare file with a light coating of os
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Bake at n bytes per second until golden brown
	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func main() {

	makeReadme("../README.md")

}
