package main

import (
	"fmt"
	"time"

	"github.com/blaskovicz/go-nest"
)

// list cameras and their states
// Set NEST_ACCESS_TOKEN env var before running this
// See https://developers.nest.com/documentation/cloud/how-to-auth
func main() {
	c := nest.New()
	cams, err := c.ListCameras()
	if err != nil {
		panic(err)
	}

	for _, cam := range cams {
		if cam.IsStreaming {
			continue
		}
		fmt.Printf("[%s] camera=%s name=%s online=%t streaming=%t\n", time.Now(), cam.ID, cam.Name, cam.IsOnline, cam.IsStreaming)
		fmt.Printf("Turning on camera.")
		cam, err = c.UpdateCameraIsStreaming(cam.ID, true)
		if err != nil {
			panic(err)
		}
		fmt.Printf("...%t\n", cam.IsStreaming)
	}
}
