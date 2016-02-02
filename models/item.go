package models
import "fmt"

type Item struct {
	Parent *Category
	Name string
	UOM *Unit
}

type Unit struct {
	Name string
	Ratio float32
}

func (i *Item) Move(newCat *Category) {
	i.Parent = newCat
}

func (i *Item) Stock(s []*Stock) []*Stock {
	r := make([]*Stock,0)
	for _, v := range s {
		if v.item == i {
			r = append(r, v)
		}
	}
//	fmt.Printf("r=", r)
	return r
}

func (i *Item) Calc(st []*Stock, tx []*Trans)  {
	for _, t := range tx {
		if t.item == i {
			for _, s := range st {
				if s.item == i && s.loc == t.locIn {
					s.bal += t.qty
				}
				if s.item == i && s.loc == t.locOut {
					s.bal -= t.qty
				}
				fmt.Printf("Item %v Loc: %v qty: %v bal: %v\n", s.item.Name, s.loc.Code, t.qty, s.bal)
			}
		}
	}
}