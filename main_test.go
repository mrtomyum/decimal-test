package main
import (
	"testing"
	"github.com/shopspring/decimal"
	"fmt"
)

type money struct {
	amount decimal.Decimal
}
// ทดสอบการหาร Decimal
func TestDecimal(t *testing.T) {
	a := money{}
	b := money{}
	c := money{}
	a.amount = decimal.New(10000, -2)
	b.amount = decimal.New(200, -2)
	c.amount = a.amount + b.amount
	fmt.Printf("a = %v b = %v", a.amount , b.amount)
//	fmt.Printf("a + b = %v", a.amount + b.amount)
	fmt.Printf("c = ", c.amount)
}
// ลองโหลดข้อมูลสินค้า และ stockcard
