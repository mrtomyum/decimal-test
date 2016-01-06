package main
import (
	"testing"
	"fmt"
)

// เทสการบวกเลข
func TestAdd(t *testing.T) {
	a := Add(1000, 250)
	if a != 1250 {
		t.Error("1000 + 250 Expected = 1250 but =", a)
	}
	fmt.Printf("2. a= %d Bath = %s", a, a.String() )
}

// เทสการหาร
func TestDivide(t *testing.T) {
	a := Divide(1000, 3)
	if a != 333 {
		t.Error("1000 / 3 Expected = 333 but =", a)
	}
	fmt.Printf("2. a= %d Bath = %s", a, a.String() )
}

// เทสการแบ่งรายได้ 999 จ่าย 3 คน A=33%, B=33%, C=34%
// ยอดเงินจ่ายบวกกลับต้องได้เท่ากับ 999
func TestShareProfit(t *testing.T) {
	a := ShareHolder{}
	b := ShareHolder{}
	c := ShareHolder{}
	a.stock = 33
	b.stock = 33
	c.stock = 34
	profit := Money(999)
	a.Share(profit)
	b.Share(profit)
	c.Share(profit)
	fmt.Printf("a.bath = %s\nb.bath = %s\nc.bath = %s a + b + c = %d", a.bath, b.bath, c.bath, profit)
	if profit != a.bath + b.bath + c.bath {
		t.Errorf("Bath a + b + c (%d + %d + %d) <> profit (%d)", a.bath, b.bath, c.bath, profit)
	}
}

// เทสบวกเลขมากเกินจน Overflow

//เทสการหารเลขมีเศษทศนิยมไม่รู้จบ จะต้องปัดเศษให้กองสุดท้าย

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

// กรณี นำส่งเงิน แล้วลืม/ไม่ได้ยิงเคาท์เตอร์ตู้ใดมาในวันนั้น ให้ถือว่าบันทึกรับเงินสด
// เข้าบัญชีตั้งพัก รอกระทบยอด แยกแต่ละตู้ไว้ได้ จนกว่าจะมียอดเคาท์เตอร์ในวันถ้ดไป