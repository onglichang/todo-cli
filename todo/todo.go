package todo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

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

// --- Load from text file ---
func (l *List) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		// If file doesn't exist, nothing to load — not an error
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) != 4 {
			continue
		}

		id, _ := strconv.Atoi(parts[0])
		title := parts[1]
		done, _ := strconv.ParseBool(parts[2])
		created, _ := time.Parse(time.RFC3339, parts[3])

		l.items = append(l.items, Todo{
			ID:          id,
			Title:       title,
			Done:        done,
			DateCreated: created,
		})
	}

	return scanner.Err()
}

// --- Save to text file ---
func (l *List) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, item := range l.items {
		line := fmt.Sprintf(
			"%d|%s|%v|%s\n",
			item.ID,
			item.Title,
			item.Done,
			item.DateCreated.Format(time.RFC3339),
		)
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
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
