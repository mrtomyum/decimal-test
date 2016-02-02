package models
import (
	"testing"
	"fmt"
)

func setup() ([]*Category, []*Item, []*Location, []*Stock, []*Trans){
	c := make([]*Category, 5)
	c[0] = &Category{Name: "root"}
	c[1] = &Category{Parent: c[0], Name: "Beverage"}
	c[2] = &Category{Parent: c[0], Name: "Snack"}
	c[3] = &Category{Parent: c[0], Name: "Fruit"}
	c[4] = &Category{Parent: c[1], Name: "Soda"}

	i := make([]*Item, 10)
	i[0] = &Item{Parent: c[0], Name: "Dummy Item"}
	i[1] = &Item{Parent: c[4], Name: "Coke Can 325ml"}
	i[2] = &Item{Parent: c[4], Name: "Pepsi Can 325ml"}
	i[3] = &Item{Parent: c[1], Name: "Ichitan Can 325ml"}
	i[4] = &Item{Parent: c[1], Name: "Coke Can 325ml"}
	i[5] = &Item{Parent: c[1], Name: "Coke Can 325ml"}
	i[6] = &Item{Parent: c[1], Name: "Coke Can 325ml"}
	i[7] = &Item{Parent: c[1], Name: "Coke Can 325ml"}
	i[8] = &Item{Parent: c[1], Name: "Coke Can 325ml"}
	i[9] = &Item{Parent: c[1], Name: "Coke Can 325ml"}

	l := make([]*Location, 15)
	l[0] = &Location{LocType: ROOT, Code: "0"}
	l[1] = &Location{Parent: l[0], LocType: STORE, Code:"S1"}
	l[2] = &Location{Parent: l[1], LocType: BUY, Code:"B1"}
	l[3] = &Location{Parent: l[1], LocType: BUY, Code:"V2"}
	l[4] = &Location{Parent: l[1], LocType: SALE, Code:"SA1"}
	l[5] = &Location{Parent: l[1], LocType: SALE, Code:"SA2"}
	l[6] = &Location{Parent: l[1], LocType: VEHICLE, Code:"V1"}
	l[7] = &Location{Parent: l[1], LocType: VEHICLE, Code:"V2"}
	l[8] = &Location{Parent: l[1], LocType: MACHINE, Code:"C1"}
	l[9] = &Location{Parent: l[1], LocType: MACHINE, Code:"C2"}
	l[10] = &Location{Parent: l[8], LocType: COLUMN, Code:"01"} // Column in Machine
	l[11] = &Location{Parent: l[8], LocType: COLUMN, Code:"02"}
	l[12] = &Location{Parent: l[8], LocType: COLUMN, Code:"03"}
	l[13] = &Location{Parent: l[8], LocType: COLUMN, Code:"04"}
	l[14] = &Location{Parent: l[8], LocType: COLUMN, Code:"05"}

	s := make([]*Stock,5)
	s[0] = &Stock{item: i[1], loc: l[1], bal: 0}
	s[1] = &Stock{item: i[1], loc: l[2], bal: 0}
	s[2] = &Stock{item: i[1], loc: l[3], bal: 0}
	s[3] = &Stock{item: i[2], loc: l[1], bal: 0}
	s[4] = &Stock{item: i[2], loc: l[2], bal: 0}

	t := make([]*Trans, 10)
	t[0] = &Trans{ ID: 1, item: i[1], locOut: l[0], locIn: l[1], qty: 100}
	t[1] = &Trans{ ID: 2, item: i[1], locOut: l[0], locIn: l[1], qty: 10}
	t[2] = &Trans{ ID: 3, item: i[1], locOut: l[0], locIn: l[1], qty: 20}
	t[3] = &Trans{ ID: 4, item: i[1], locOut: l[1], locIn: l[2], qty: 50}
	t[4] = &Trans{ ID: 5, item: i[1], locOut: l[2], locIn: l[3], qty: 10}
	t[5] = &Trans{ ID: 6, item: i[1], locOut: l[1], locIn: l[3], qty: 10}
	t[6] = &Trans{ ID: 7, item: i[1], locOut: l[1], locIn: l[4], qty: 10}
	t[7] = &Trans{ ID: 8, item: i[1], locOut: l[3], locIn: l[5], qty: 1}
	t[8] = &Trans{ ID: 9, item: i[1], locOut: l[4], locIn: l[1], qty: 1}
	t[9] = &Trans{ ID: 10, item: i[1], locOut: l[3], locIn: l[1], qty: 1}
	return c, i, l, s, t
}

func TestPrintCat(t *testing.T) {
//	c := Category{Name: "น้ำอัดลม" }
	cats, _, _, _, _ := setup()
	for i, v := range cats {
		fmt.Printf("\nItem: %v %v ", i, v)
	}
}

// SearchCat where name match c.Name must return array index of [1]
func TestCat_FindByName(t *testing.T) {
	cats, _, _, _, _ := setup()
	c := FindByName(cats, "Bev")
	if c != cats[1] {
		t.Error("Expected return index 1 but =", c)
	}

	x := FindByName(cats, "Alcohol")
	if x != nil {
		t.Error("Expected no matched category but =", x)
	}
	fmt.Println("Pass! TestCat_FindByName return index =", c)
}

func TestCat_MoveNode(t *testing.T) {
	cats, _, _, _, _ := setup()
	c := cats[4]
	newCat := cats[2]
	c.Move(newCat)
	if c.Parent != newCat {
		t.Errorf("Expected c.Parent: = %v but =", newCat, c.Parent)
	}
}

func TestItem_MoveCat(t *testing.T) {
	cats, items, _, _, _ := setup()
	i := items[1]
	c := cats[4]
	i.Move(c)
	if i.Parent != c {
		t.Errorf("Expected % = %", i.Parent, c)
	}
	fmt.Printf("Pass! MoveCat [ %v ]with=> %v\n", i.Name, i.Parent.Name)
//	fmt.Println(i.Name)
}

func TestStock_MoveLoc(t *testing.T) {
	_, _, locs, stocks, _ := setup()
	newLoc := locs[3]
	s := stocks[0]
	fmt.Println("s.loc = ", s.loc.Code)
	s.Move(newLoc)
	if s.loc != locs[3] {
		t.Error("Expected s.loc = %v but = %v",locs[3], s.loc)
	}
	fmt.Println("s.loc = ", s.loc.Code)
}

func TestStock_CalcStockBalanceFromTrans(t *testing.T) {
	_, items , _, stocks, trans := setup()

	i := items[1]
	fmt.Println("A Stock:", stocks)
	a := i.Stock(stocks)
	fmt.Println("B Stock:", i.Stock(a))

	i.Calc(stocks, trans)
	for _, s := range a {
		fmt.Println(">>", s.item.Name, s.loc.Code, "Balance=", s.bal)
	}
}