package ToDoList

type ToDoList struct {
	items []string
}

func (tdl *ToDoList) AddItem(newItem string) {
	tdl.items = append(tdl.items, newItem)
}

func (tdl *ToDoList) DeleteItem(i int) {
	tdl.items = append(tdl.items[:i], tdl.items[i+1:]...)
}

func (tdl *ToDoList) GetAllItems() []string {
	return tdl.items
}

func (tdl *ToDoList) GetItem(i int) string {
	return tdl.items[i]
}
