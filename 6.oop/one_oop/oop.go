package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrEngineAlreadyRunning = errors.New("двигатель уже работает")
	ErrEngineOff            = errors.New("двигатель не запущен")
	ErrLowBattery           = errors.New("низкий заряд батареи")
)

type Vehicle interface {
	StartEngine() error
	StopEngine() error
	GetInfo() string
}

type BaseVehicle struct {
	brand    string
	engineOn bool
}

type Car struct {
	BaseVehicle
}

type Truck struct {
	Car
	CargoCapacity float64
}

type ElectricCar struct {
	Car
	batteryLevel int
}

func (b *BaseVehicle) StartEngine() error {
	if b.engineOn {
		return ErrEngineAlreadyRunning
	}
	b.engineOn = true
	return nil
}

func (b *BaseVehicle) StopEngine() error {
	if !b.engineOn {
		return ErrEngineOff
	}
	b.engineOn = false
	return nil
}

func (b *BaseVehicle) GetEngineStatus() bool {
	return b.engineOn
}

func (b *BaseVehicle) GetBrand() string {
	return b.brand
}

func NewCar(brand string) *Car {
	return &Car{
		BaseVehicle: BaseVehicle{
			brand:    brand,
			engineOn: false,
		},
	}
}

func (c *Car) Honk() string {
	return "Beep beep!"
}

func (c *Car) GetInfo() string {
	status := "Заглушен"
	if c.engineOn {
		status = "Заведен"
	}
	return fmt.Sprintf("Автомобиль %s: двигатель %s", c.brand, status)
}

func NewTruck(brand string, cargoCapacity float64) *Truck {
	return &Truck{
		Car: Car{
			BaseVehicle: BaseVehicle{
				brand:    brand,
				engineOn: false,
			},
		},
		CargoCapacity: cargoCapacity,
	}
}

func (t *Truck) Honk() string {
	return "Honk Honk!"
}

func (t *Truck) GetCargoCapacity() float64 {
	return t.CargoCapacity
}

func (t *Truck) GetInfo() string {
	status := "Заглушен"
	if t.engineOn {
		status = "Заведен"
	}
	return fmt.Sprintf("Грузовик %s: двигатель %s, грузоподъемность %.1f т", t.brand, status, t.CargoCapacity)

}

func NewElectricCar(brand string, batteryLevel int) *ElectricCar {
	return &ElectricCar{
		Car: Car{
			BaseVehicle: BaseVehicle{
				brand:    brand,
				engineOn: false,
			},
		},
		batteryLevel: batteryLevel,
	}
}

func (e *ElectricCar) StartEngine() error {
	if e.batteryLevel <= 5 {
		return ErrLowBattery
	}
	return e.BaseVehicle.StartEngine()
}

func (e *ElectricCar) GetBatteryLevel() int {
	return e.batteryLevel
}

func (e *ElectricCar) Charge(amount int) {
	e.batteryLevel += amount
	if e.batteryLevel > 100 {
		e.batteryLevel = 100
	}
}

func (e *ElectricCar) Honk() string {
	return "Beep beep!"
}

func (e *ElectricCar) GetInfo() string {
	status := "Заглушен"
	if e.engineOn {
		status = "Заведен"
	}
	return fmt.Sprintf("Электрокар %s: двигатель %s, заряд %d%%", e.brand, status, e.batteryLevel)
}

func StartAllVehicles(vehicles []Vehicle) {
	fmt.Println("Запуск всех транспортных средств:")
	for i, v := range vehicles {
		err := v.StartEngine()
		if err != nil {
			fmt.Printf("[%d] %s - ошибка: %v\n", i+1, v.GetInfo(), err)
		} else {
			fmt.Printf("[%d] %s - успешно\n", i+1, v.GetInfo())
		}
	}
}

func MakeSomeNoise(vehicles []Vehicle) {
	fmt.Println("Сигналы транспорта:")
	for i, v := range vehicles {

		var honk string
		switch vehicle := v.(type) {
		case *Truck:
			honk = vehicle.Honk()
		case *ElectricCar:
			honk = vehicle.Honk()
		case *Car:

			honk = vehicle.Honk()
		default:
			honk = "..."
		}
		fmt.Printf("[%d] %s\n", i+1, honk)
	}
}

func main() {
	fmt.Println("Система управления транспортом")
	fmt.Println("=" + strings.Repeat("=", 50))

	car := NewCar("Toyota")
	truck := NewTruck("Volvo", 20.5)
	electricCar := NewElectricCar("Tesla", 75)
	lowBatteryCar := NewElectricCar("Nissan", 3)

	fmt.Println("Информация о транспорте:")
	fmt.Println(car.GetInfo())
	fmt.Println(truck.GetInfo())
	fmt.Println(electricCar.GetInfo())
	fmt.Println(lowBatteryCar.GetInfo())

	fmt.Println("Тестирование запуска двигателя:")

	fmt.Printf("\n1. %s\n", car.GetInfo())
	fmt.Printf("   Запуск: %v\n", car.StartEngine())
	fmt.Printf("   Статус: %v\n", car.GetEngineStatus())
	fmt.Printf("   Повторный запуск: %v\n", car.StartEngine())
	fmt.Printf("   Остановка: %v\n", car.StopEngine())

	fmt.Printf("\n2. %s\n", truck.GetInfo())
	fmt.Printf("   Запуск: %v\n", truck.StartEngine())
	fmt.Printf("   Грузоподъёмность: %.1f т\n", truck.GetCargoCapacity())
	fmt.Printf("   Сигнал: %s\n", truck.Honk())

	fmt.Printf("\n3. %s\n", electricCar.GetInfo())
	fmt.Printf("   Запуск: %v\n", electricCar.StartEngine())
	fmt.Printf("   Заряд: %d%%\n", electricCar.GetBatteryLevel())

	fmt.Printf("\n4. %s\n", lowBatteryCar.GetInfo())
	fmt.Printf("   Запуск (ожидается ошибка): %v\n", lowBatteryCar.StartEngine())
	fmt.Printf("   Зарядка до 50%%\n")
	lowBatteryCar.Charge(47)
	fmt.Printf("   Новый заряд: %d%%\n", lowBatteryCar.GetBatteryLevel())
	fmt.Printf("   Запуск после зарядки: %v\n", lowBatteryCar.StartEngine())

	fmt.Println("\n" + strings.Repeat("=", 50))
	vehicles := []Vehicle{car, truck, electricCar}
	StartAllVehicles(vehicles)
	MakeSomeNoise(vehicles)
}
