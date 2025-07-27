package factorydesign

type AK47 struct {
	Gun
	Race string
}

func NewAK47() *AK47 {
	ak47 := &AK47{
		Gun: Gun{
			name:  "AK-47",
			power: 100,
		},
	}

	return ak47
}

func (a *AK47) SetRace(race string) {
	a.Race = race
}

func (a *AK47) GetRace() string {
	return a.Race
}
