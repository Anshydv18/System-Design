package concretesubjects

import "observer/observer"

type Item struct {
	ObserverList []observer.Observer
	Name         string
	Instock      bool
}

func NewItem(name string) Item {
	return Item{
		Name: name,
	}
}

func (it *Item) Register(o observer.Observer) {
	it.ObserverList = append(it.ObserverList, o)
}

func (it *Item) UnRegister(o observer.Observer) {
	it.ObserverList = removeFromSlice(it.ObserverList, o)
}

func (i *Item) NotifyObserver() {
	for _, observer := range i.ObserverList {
		observer.Update(i.Name)
	}
}

func removeFromSlice(oblist []observer.Observer, o observer.Observer) []observer.Observer {
	len := len(oblist)
	for i, ob := range oblist {
		if ob.GetId() == o.GetId() {
			oblist[len-1], oblist[i] = oblist[i], oblist[len-1]
			return oblist[:len-1]
		}
	}

	return oblist
}
