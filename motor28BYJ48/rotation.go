package motor28BYJ48

import (
	"machine"
	"time"
)

const (
	D1  = 5
	D2  = 4
	D3  = 0
	D4  = 2
	IN1 = machine.Pin(D1)
	IN2 = machine.Pin(D2)
	IN3 = machine.Pin(D3)
	IN4 = machine.Pin(D4)
)

func AntiClockWise() {
	IN1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	IN2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	IN3.Configure(machine.PinConfig{Mode: machine.PinOutput})
	IN4.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		Step0()
		time.Sleep(time.Millisecond * 5)
		Step1()
		time.Sleep(time.Millisecond * 5)
		Step2()
		time.Sleep(time.Millisecond * 5)
		Step3()
		time.Sleep(time.Millisecond * 5)
	}
}

func ClockWise() {
	IN1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	IN2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	IN3.Configure(machine.PinConfig{Mode: machine.PinOutput})
	IN4.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		Step3()
		time.Sleep(time.Millisecond * 5)
		Step2()
		time.Sleep(time.Millisecond * 5)
		Step1()
		time.Sleep(time.Millisecond * 5)
		Step0()
		time.Sleep(time.Millisecond * 5)
	}
}

func Step0() {
	IN1.High()
	IN2.High()
	IN3.Low()
	IN4.Low()
}

func Step1() {
	IN1.Low()
	IN2.High()
	IN3.High()
	IN4.Low()
}

func Step2() {
	IN1.Low()
	IN2.Low()
	IN3.High()
	IN4.High()
}

func Step3() {
	IN1.High()
	IN2.Low()
	IN3.Low()
	IN4.High()
}
