package adaptor

import "fmt"

type Client struct {
}

func (c *Client) InsertLighteningPortIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
