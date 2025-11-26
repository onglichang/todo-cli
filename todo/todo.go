package todo

import "time"

type Todo struct {
	ID          int
	Title       string
	Done        bool
	DateCreated time.Time
}

type List struct {
	items []Todo
}

func NewList() *List {
	return &List{items: []Todo{}}
}

func (l *List) Add(title string) {
	t := Todo{
		ID:          len(l.items) + 1,
		Title:       title,
		Done:        false,
		DateCreated: time.Now(),
	}
	l.items = append(l.items, t)
}

func (l *List) All() []Todo {
	return l.items
}

func (l *List) Complete(id int) bool {
	for i := range l.items {
		if l.items[i].ID == id {
			l.items[i].Done = true
			return true
		}
	}
	return false
}
