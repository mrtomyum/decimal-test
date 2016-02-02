package models

type LocType int

const (
	ROOT LocType = iota
	BUY
	STORE
	VEHICLE
	MACHINE
	COLUMN
	SALE
)

type Location struct {
	Parent *Location
	LocType
	Code string
}

type Stock struct {
	item *Item
	loc  *Location
	bal  int64
}

func (s *Stock) Move(newLoc *Location) {
	s.loc = newLoc
}

type Trans struct {
	ID     uint64
	item   *Item
	locOut *Location
	locIn  *Location
	qty    int64
}
