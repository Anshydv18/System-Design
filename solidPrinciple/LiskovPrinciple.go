// the superclass of the subclass should be easily replacable by its subclass without breaking the behaviour
// package solidprinciple_test

// type Bird interface {
// 	Fly() string
// }

// type Pigeon struct{}

// func (p *Pigeon) Fly() string {
// 	return "Pigeon is flying"
// }

// type Penguin struct{}

// func (p *Penguin) Fly() string {
// 	return "penguin is flying"
// }

//the penguin cannot fly so the, penguin cannot fly

package main

type Bird interface{
	MakeSound() string
}

type FlyingBird interface{
	Bird
	Fly() string
}

type Pigeon struct{}

func(p *Pigeon) MakeSound() string{
	return "coo"
}

func(p *Pigeon) Fly() string{
	return "flying in air"
}

type Penguin struct{}
func(pen *Penguin) MakeSound() string{
	return "sqwak"
}

//pigeon can be easily replace the penguin super class