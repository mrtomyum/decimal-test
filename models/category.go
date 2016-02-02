package models
import "strings"

type Category struct {
	Parent *Category
	Name string
}

func (c *Category) New() {
	c.Name = "New Cat"
}

func FindByName(cats []*Category, n string) (*Category){
	for i, c  := range cats {
		if strings.Contains(c.Name, n) {
			return cats[i]
		}
	}
	return nil
}

func (c *Category) Move(newCat *Category) {
	c.Parent = newCat
}