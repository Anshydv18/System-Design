package adaptor

import "fmt"

type WindowsLightningAdapter struct {
	Windows *Windows
}

func (wa *WindowsLightningAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning connector to USB for Windows machine.")
	wa.Windows.InsertIntoUSBPort()
}
