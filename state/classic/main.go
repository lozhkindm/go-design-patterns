package main

import "fmt"

type State interface {
	On(s *Switch)
	Off(s *Switch)
}

type Switch struct {
	State State
}

func (s *Switch) On() {
	s.State.On(s)
}

func (s *Switch) Off() {
	s.State.Off(s)
}

func NewSwitch() *Switch {
	return &Switch{
		State: NewOffState(),
	}
}

type BaseState struct{}

func (b *BaseState) On(s *Switch) {
	fmt.Println("Light is already on")
}

func (b *BaseState) Off(s *Switch) {
	fmt.Println("Light is already off")
}

type OnState struct {
	BaseState
}

func (o *OnState) Off(s *Switch) {
	fmt.Println("Turning the light off...")
	s.State = NewOffState()
}

func NewOnState() *OnState {
	fmt.Println("Light turned on")
	return &OnState{
		BaseState: BaseState{},
	}
}

type OffState struct {
	BaseState
}

func (o *OffState) On(s *Switch) {
	fmt.Println("Turning the light on...")
	s.State = NewOnState()
}

func NewOffState() *OffState {
	fmt.Println("Light turned off")
	return &OffState{
		BaseState: BaseState{},
	}
}

func main() {
	switcher := NewSwitch()
	switcher.On()
	switcher.Off()
	switcher.Off()
}
