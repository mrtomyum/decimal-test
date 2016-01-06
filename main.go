package main
import "fmt"

type Money int64

func Add(x, y Money) Money {
	r := x + y
	if y >= 0 {
		if r < x {
			panic("Money Add: overflow")
		}
	}
	return x + y
}

func Sub(x, y Money) Money {
	return x - y
}

func (m Money) String() string{
	return fmt.Sprintf("%d.%02d", int64(m)/100, int64(m)%100)
}

func Divide(x, y Money) Money{
	return x / y
}

type ShareHolder struct {
	stock int64   //เปอร์เซนต์แบ่งจ่าย
	bath  Money // ยอดเงินรับจริง
}

func (s *ShareHolder) Share(profit Money) {
	a := int64(profit) * s.stock /100
	s.bath = Money(a)
}
