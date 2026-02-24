package main

import (
	"errors"
	"fmt"
)

var (
	ErrUnsupported = errors.New("обновление недоступно")
)

type Device interface {
	UpdateOS(version string) error
	GetInfo() string
}

type Smartphone struct {
	OSVersion string
	Model     string
}

type Laptop struct {
	OSVersion string
	Model     string
}

type Smartwatch struct {
	OSVersion string
	Model     string
}

func (s *Smartphone) UpdateOS(version string) error {
	if s.OSVersion >= "12.0" {
		return ErrUnsupported
	}
	s.OSVersion = version
	return nil
}

func (s *Smartphone) GetInfo() string {
	return fmt.Sprintf("Модель: %s, OC: %s", s.Model, s.OSVersion)
}

func (l *Laptop) UpdateOS(version string) error {
	if len(version) < len("Windows") || version[:len("Windows")] != "Windows" {
		return ErrUnsupported
	}
	l.OSVersion = version
	return nil

}

func (l *Laptop) GetInfo() string {
	return fmt.Sprintf("Модель: %s, OC: %s", l.Model, l.OSVersion)
}

func (sw *Smartwatch) UpdateOS(version string) error {
	if len(version) < 5 {
		return ErrUnsupported
	}
	sw.OSVersion = version
	return nil
}

func (sw *Smartwatch) GetInfo() string {
	return fmt.Sprintf("Модель: %s, OC: %s", sw.Model, sw.OSVersion)
}

// Пример использования опционален
/*
func main() {
	device := []Device{
		&Smartphone{Model: "iPhone 13", OSVersion: "11.0"},
		&Laptop{Model: "Notebook E2", OSVersion: "Windows 10 Pro"},
		&Smartwatch{Model: "Watch 3", OSVersion: "2.0"},
}

for _, device := range devices {
	fmt.Println(device.GetInfo())
	err := device.UpdatedOS("new_version")
	if err != nil {
		fmt.Printf("Ошибка обновления: %v\n", err)
	} else {
		fmt.Println("Обновление успешно!")
}
		fmt.Println(device.GetInfo())
		fmt.Println("---")
	}
}
*/
