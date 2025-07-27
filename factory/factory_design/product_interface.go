package factorydesign

type IGun interface {
	SetName(name string)
	GetName() string
	SetPower(power int)
	GetPower() int
	SetRace(race string)
	GetRace() string
}

func SetRace(race string){
	
}
func GetRace() string {
	return ""
}