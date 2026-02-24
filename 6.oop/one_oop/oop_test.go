package main

import (
	"strings"
	"testing"
)

//Тесты для Car

func TestCar_StartEngine(t *testing.T) {
	car := NewCar("Toyota")

	// Первый запуск должен быть успешным
	if err := car.StartEngine(); err != nil {
		t.Errorf("StartEngine() returned error: %v", err)
	}

	// Проверка статуса
	if !car.GetEngineStatus() {
		t.Error("GetEngineStatus() should return true after StartEngine()")
	}

	// Повторный запуск должен вернуть ошибку
	if err := car.StartEngine(); err != ErrEngineAlreadyRunning {
		t.Errorf("StartEngine() should return ErrEngineAlreadyRunning, got: %v", err)
	}
}

func TestCar_StopEngine(t *testing.T) {
	car := NewCar("Honda")

	// Остановка заглушенного двигателя должна вернуть ошибку
	if err := car.StopEngine(); err != ErrEngineOff {
		t.Errorf("StopEngine() should return ErrEngineOff, got: %v", err)
	}

	// Запускаем и останавливаем
	car.StartEngine()
	if err := car.StopEngine(); err != nil {
		t.Errorf("StopEngine() returned error: %v", err)
	}

	// Проверка статуса
	if car.GetEngineStatus() {
		t.Error("GetEngineStatus() should return false after StopEngine()")
	}
}

func TestCar_Honk(t *testing.T) {
	car := NewCar("BMW")
	expected := "Beep beep!"
	got := car.Honk()

	if got != expected {
		t.Errorf("Honk() = %q; want %q", got, expected)
	}
}

func TestCar_GetInfo(t *testing.T) {
	car := NewCar("Mercedes")
	info := car.GetInfo()

	if !strings.Contains(info, "Mercedes") {
		t.Errorf("GetInfo() should contain brand name, got: %q", info)
	}
	if !strings.Contains(info, "Автомобиль") {
		t.Errorf("GetInfo() should contain type, got: %q", info)
	}
}

//Тесты для Truck

func TestTruck_StartEngine(t *testing.T) {
	truck := NewTruck("Volvo", 25.0)

	if err := truck.StartEngine(); err != nil {
		t.Errorf("StartEngine() returned error: %v", err)
	}

	if !truck.GetEngineStatus() {
		t.Error("GetEngineStatus() should return true")
	}
}

func TestTruck_Honk(t *testing.T) {
	truck := NewTruck("Scania", 30.0)
	expected := "Honk Honk!"
	got := truck.Honk()

	if got != expected {
		t.Errorf("Honk() = %q; want %q", got, expected)
	}
}

func TestTruck_GetCargoCapacity(t *testing.T) {
	truck := NewTruck("MAN", 18.5)
	expected := 18.5
	got := truck.GetCargoCapacity()

	if got != expected {
		t.Errorf("GetCargoCapacity() = %.2f; want %.2f", got, expected)
	}
}

func TestTruck_GetInfo(t *testing.T) {
	truck := NewTruck("Volvo", 20.0)
	info := truck.GetInfo()

	if !strings.Contains(info, "Грузовик") {
		t.Errorf("GetInfo() should contain type, got: %q", info)
	}
	if !strings.Contains(info, "20.0") {
		t.Errorf("GetInfo() should contain cargo capacity, got: %q", info)
	}
}

// Тесты для ElectricCar

func TestElectricCar_StartEngine_LowBattery(t *testing.T) {
	eCar := NewElectricCar("Tesla", 5)

	err := eCar.StartEngine()
	if err != ErrLowBattery {
		t.Errorf("StartEngine() should return ErrLowBattery, got: %v", err)
	}

	if eCar.GetEngineStatus() {
		t.Error("Engine should not start with low battery")
	}
}

func TestElectricCar_StartEngine_SufficientBattery(t *testing.T) {
	eCar := NewElectricCar("Tesla", 10)

	err := eCar.StartEngine()
	if err != nil {
		t.Errorf("StartEngine() returned error: %v", err)
	}

	if !eCar.GetEngineStatus() {
		t.Error("Engine should start with sufficient battery")
	}
}

func TestElectricCar_GetBatteryLevel(t *testing.T) {
	eCar := NewElectricCar("Nissan", 75)
	expected := 75
	got := eCar.GetBatteryLevel()

	if got != expected {
		t.Errorf("GetBatteryLevel() = %d; want %d", got, expected)
	}
}

func TestElectricCar_Charge(t *testing.T) {
	eCar := NewElectricCar("Tesla", 50)

	eCar.Charge(30)
	if eCar.GetBatteryLevel() != 80 {
		t.Errorf("Charge(30) should result in 80%%, got %d%%", eCar.GetBatteryLevel())
	}

	eCar.Charge(50) // Должно ограничиться 100%
	if eCar.GetBatteryLevel() != 100 {
		t.Errorf("Charge should cap at 100%%, got %d%%", eCar.GetBatteryLevel())
	}
}

func TestElectricCar_GetInfo(t *testing.T) {
	eCar := NewElectricCar("Tesla", 90)
	info := eCar.GetInfo()

	if !strings.Contains(info, "Электрокар") {
		t.Errorf("GetInfo() should contain type, got: %q", info)
	}
	if !strings.Contains(info, "90%") {
		t.Errorf("GetInfo() should contain battery level, got: %q", info)
	}
}

// Тесты на полиморфизм

func TestPolymorphism_VehicleInterface(t *testing.T) {
	vehicles := []Vehicle{
		NewCar("Toyota"),
		NewTruck("Volvo", 20.0),
		NewElectricCar("Tesla", 80),
	}

	for i, v := range vehicles {
		// Все должны реализовывать интерфейс Vehicle
		if v == nil {
			t.Errorf("Vehicle %d is nil", i)
			continue
		}

		// Тест StartEngine
		if err := v.StartEngine(); err != nil {
			t.Errorf("Vehicle %d StartEngine() error: %v", i, err)
		}

		// Тест GetInfo
		info := v.GetInfo()
		if info == "" {
			t.Errorf("Vehicle %d GetInfo() returned empty string", i)
		}

		// Тест StopEngine
		if err := v.StopEngine(); err != nil {
			t.Errorf("Vehicle %d StopEngine() error: %v", i, err)
		}
	}
}

func TestPolymorphism_HonkMethods(t *testing.T) {
	car := NewCar("BMW")
	truck := NewTruck("Scania", 25.0)
	eCar := NewElectricCar("Tesla", 90)

	// Проверка уникальности сигналов
	carHonk := car.Honk()
	truckHonk := truck.Honk()
	eCarHonk := eCar.Honk()

	if carHonk != "Beep beep!" {
		t.Errorf("Car Honk() = %q; want %q", carHonk, "Beep beep!")
	}
	if truckHonk != "Honk Honk!" {
		t.Errorf("Truck Honk() = %q; want %q", truckHonk, "Honk Honk!")
	}
	if eCarHonk != "Beep beep!" {
		t.Errorf("ElectricCar Honk() = %q; want %q", eCarHonk, "Beep beep!")
	}

	// Грузовик должен иметь уникальный сигнал
	if carHonk == truckHonk {
		t.Error("Car and Truck should have different honk sounds")
	}
}

// Тесты на инкапсуляцию

func TestEncapsulation_UnexportedFields(t *testing.T) {
	car := NewCar("Toyota")
	truck := NewTruck("Volvo", 20.0)
	eCar := NewElectricCar("Tesla", 75)

	// Проверяем, что поля доступны только через методы
	if car.GetEngineStatus() != false {
		t.Error("Initial engine status should be false")
	}
	if truck.GetCargoCapacity() != 20.0 {
		t.Error("Cargo capacity should be accessible via getter")
	}
	if eCar.GetBatteryLevel() != 75 {
		t.Error("Battery level should be accessible via getter")
	}
}

// Бенчмарки

func BenchmarkCar_StartEngine(b *testing.B) {
	car := NewCar("Toyota")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		car.StartEngine()
		car.StopEngine()
	}
}

func BenchmarkTruck_StartEngine(b *testing.B) {
	truck := NewTruck("Volvo", 20.0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		truck.StartEngine()
		truck.StopEngine()
	}
}

func BenchmarkElectricCar_StartEngine(b *testing.B) {
	eCar := NewElectricCar("Tesla", 80)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eCar.StartEngine()
		eCar.StopEngine()
	}
}
