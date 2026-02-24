package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var (
	ErrInvalidAmount       = errors.New("некорректная сумма платежа")
	ErrProviderUnavailable = errors.New("провайдер недоступен")
)

type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

type Sberbank struct {
	APIKey string
}

type Tbank struct {
	APIKey string
}

type Alfabank struct {
	APIKey string
}

func isProviderAvailable() bool {
	return rand.Intn(100) > 10
}

func (s Sberbank) ProcessPayment(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if !isProviderAvailable() {
		return ErrProviderUnavailable
	}
	return nil

}

func (t Tbank) ProcessPayment(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if !isProviderAvailable() {
		return ErrProviderUnavailable
	}
	return nil

}

func (a Alfabank) ProcessPayment(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if !isProviderAvailable() {
		return ErrProviderUnavailable
	}
	return nil
}

func main() {

	sber := Sberbank{APIKey: "sber-api-key-1"}
	tink := Tbank{APIKey: "tink-api-key-2"}
	alfa := Alfabank{APIKey: "alfa-api-key-3"}

	err1 := sber.ProcessPayment(100)
	if err1 != nil {
		fmt.Println("Ошибка:", err1)
	} else {
		fmt.Println("Успешно!")
	}

	err2 := tink.ProcessPayment(-100)
	if err2 != nil {
		fmt.Println("Ошибка:", err2)
	} else {
		fmt.Println("Успешно!")
	}

	err3 := alfa.ProcessPayment(50)
	if err3 != nil {
		fmt.Println("Ошибка:", err3)
	} else {
		fmt.Println("Успешно!")
	}
}
