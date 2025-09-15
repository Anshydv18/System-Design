package subjects

import "observer/observer"

type Subjects interface {
	Register(o observer.Observer)
	UnRegister(o observer.Observer)
	NotifyObserver()
}
