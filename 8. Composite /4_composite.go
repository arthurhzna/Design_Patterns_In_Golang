package main

import "fmt"

// Component
type Menu interface {
	Show(level int)
}

// Leaf
type MenuItem struct {
	Name string
}

func (m MenuItem) Show(level int) {
	fmt.Printf("%s├── %s\n", spaces(level), m.Name)
}

// Composite
type MenuGroup struct {
	Name     string
	Children []Menu
}

func (g MenuGroup) Show(level int) {
	if level == 0 {
		fmt.Println(g.Name)
	} else {
		fmt.Printf("%s└── %s\n", spaces(level), g.Name)
	}

	for _, child := range g.Children {
		child.Show(level + 1)
	}
}

// helper
func spaces(level int) string {
	result := ""
	for i := 0; i < level; i++ {
		result += "    "
	}
	return result
}

func main() {
	settings := MenuGroup{
		Name: "Settings",
		Children: []Menu{
			MenuItem{Name: "Profile"},
			MenuItem{Name: "Security"},
			MenuItem{Name: "Notifications"},
		},
	}

	settings.Show(0)
}
