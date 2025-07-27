package main

import adaptor "adapter/adapter-classes"

func main() {

	client := &adaptor.Client{}
	mac := &adaptor.Mac{}

	client.InsertLighteningPortIntoComputer(mac)

	windowsMachine := &adaptor.Windows{}
	windowsMachineAdapter := &adaptor.WindowsLightningAdapter{
		Windows: windowsMachine,
	}

	client.InsertLighteningPortIntoComputer(windowsMachineAdapter)

}
