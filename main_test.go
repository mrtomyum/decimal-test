package main
import (
	"testing"
	"fmt"
)

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

//เทสการบวกเลข
func TestAdd(t *testing.T) {
	a := Add(1000, 250)
	fmt.Println(a)
}

//เทสบวกเลขมากเกินจน Overflow


//เทสการหารเลขมีเศษทศนิยมไม่รู้จบ จะต้องปัดเศษให้กองสุดท้าย
a
        ggbaaaaaaaaaazaaaaaaaa
//type money struct {aa
//}
// ทดสอบการหาร Decimal
//func TestDecimal(t *testing.T) {
//	a := money{}0, -2)
//	c.amount = a.amount + b.amount
//	fmt.Printf("a = %v b = %v
//	b := money{a
//	c := money{}
//	a.amount = decimal.New(10000, -2)
//	b.amount = decimal.New(20", a.amount , b.amount)
//	fmt.Printf("a + b = %v", a.amount + b.amount)
//	fmt.Printf("c = ", c.amount)
//}
// ลองโหลดข้อมูลสินค้า และ stockcard
