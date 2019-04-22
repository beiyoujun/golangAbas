package main

import "fmt"

const (
	PAY_WAIT    = 0 //待支付
	PAY_SUCCESS = 1 //支付成功
	PAY_STAY    = 2 //支付中
	PAY_FAIL    = 3 //支付失败
)

func main() {
	o := NewOrderInfo(0)
	o.Refund()
}

type OrderDo interface {
	Pay()
	Delete()
	Cancel()
	Refund()
}

type OrderInfo struct {
	orderDo    OrderDo
	orderState int
}

func NewOrderInfo(state int) *OrderInfo {
	o := &OrderInfo{
		orderState: state,
	}
	o.changeState()
	return o
}

func (o *OrderInfo) changeState() {
	switch o.orderState {
	case PAY_WAIT:
		o.orderDo = &PayWaitDo{}
		break
	case PAY_SUCCESS:
		o.orderDo = &PaySuccessDo{}
		break
	case PAY_STAY:
		o.orderDo = &PayStayDo{}
		break
	case PAY_FAIL:
		o.orderDo = &PayFailDo{}
	}
}

func (o *OrderInfo) Pay() {
	o.orderDo.Pay()
}
func (o *OrderInfo) Delete() {
	o.orderDo.Delete()
}
func (o *OrderInfo) Cancel() {
	o.orderDo.Cancel()
}
func (o *OrderInfo) Refund() {
	o.orderDo.Refund()
}

type PayWaitDo struct {
}

func (p *PayWaitDo) Pay() {
	fmt.Println("可以支付")
}

func (p *PayWaitDo) Delete() {
	fmt.Println("不可删除")
}

func (p *PayWaitDo) Refund() {
	fmt.Println("没有支付 不可以退款")
}

func (p *PayWaitDo) Cancel() {
	fmt.Println("可以取消")
}

type PaySuccessDo struct {
}

func (*PaySuccessDo) Pay() {
	fmt.Println("支付成功，不能再支付")
}

func (*PaySuccessDo) Delete() {
	fmt.Println("支付成功，不可删除")
}

func (*PaySuccessDo) Cancel() {
	fmt.Println("支付成功，不可取消")
}

func (*PaySuccessDo) Refund() {
	fmt.Println("可以退款")
}

type PayStayDo struct {
}

func (*PayStayDo) Pay() {
	fmt.Println("支付中，请等待")
}

func (*PayStayDo) Delete() {
	fmt.Println("支付中，不可删除")
}

func (*PayStayDo) Cancel() {
	fmt.Println("支付中，不可取消")
}

func (*PayStayDo) Refund() {
	fmt.Println("支付中，不可退款")
}

type PayFailDo struct {
}

func (*PayFailDo) Pay() {
	fmt.Println("可以支付")
}

func (*PayFailDo) Delete() {
	fmt.Println("不可删除")
}

func (*PayFailDo) Cancel() {
	fmt.Println("可以取消")
}

func (*PayFailDo) Refund() {
	fmt.Println("不可退款")
}
