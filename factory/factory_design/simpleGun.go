package factorydesign

type SimpleGun struct {
	Gun
}

func NewSimpleGun() *SimpleGun {
	simpleGun := &SimpleGun{
		Gun: Gun{
			name:  "Simple Gun",
			power: 50,
		},
	}

	return simpleGun
}

func (s *SimpleGun) SetRace(race string) {}

func (s *SimpleGun) GetRace() string {
	return ""
}