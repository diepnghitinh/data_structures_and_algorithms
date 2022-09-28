package main

import "fmt"

/*

Abstract Factory tạo ra các object mà không cần phải chỉ định rõ ràng Class

Output:
created OrderId from momo
created OrderId from zalopay

*/

type IPaymentFactory interface {
	createOrder()
	getOrderId(id string)
	getName() string
}

/* Momo Payment */
type PaymentMomo struct {
}

func (s *PaymentMomo) createOrder() {
	fmt.Println("created OrderId from momo")
}

func (s *PaymentMomo) getOrderId(id string) {
	fmt.Printf("get OrderId from momo %s\n", id)
}

func (s *PaymentMomo) getName() string {
	return "momo"
}

/* Momo Payment */
type PaymentZaloPay struct {
}

func (s *PaymentZaloPay) createOrder() {
	fmt.Println("created OrderId from zalopay")
}

func (s *PaymentZaloPay) getOrderId(id string) {
	fmt.Printf("get OrderId from zalopay %s\n", id)
}

func (s *PaymentZaloPay) getName() string {
	return "zalopay"
}

func getPaymentInstance(payment string) (IPaymentFactory, error) {
	switch payment {
	case "momo":
		return &PaymentMomo{}, nil
	case "zalopay":
		return &PaymentZaloPay{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}

func dp__creational__abstract_factory() {

	paymentFactory, _ := getPaymentInstance("momo")
	paymentFactory.createOrder()

	paymentFactory, _ = getPaymentInstance("zalopay")
	paymentFactory.createOrder()
}
