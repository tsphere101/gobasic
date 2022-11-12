package main

type Bill struct {
	name  string
	items []string
	tip   float64
}

// Make new Bill
func NewBill(name string) Bill {
	var b Bill = Bill{
		name:  name,
		items: []string{},
		tip:   0,
	}
	return b
}

// Reciever Function
func (b *Bill) Format() string {
	return b.name
}

func (b *Bill) UpdateTip(value float64) {
	(*b).tip = value
}

// For struct pointer is automatically derefferencing.
func (b *Bill) AddItem(item string) {
	b.items = append(b.items, item)
}

func (b *Bill) UpdateName(name string) {
	b.name = name
}
