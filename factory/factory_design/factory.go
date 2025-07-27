package factorydesign


func GetGun(gunType string) IGun {
	switch gunType {
	case "AK47":
		return NewAK47()
	case "SimpleGun":
		return NewSimpleGun()
	default:
		return nil
	}
}