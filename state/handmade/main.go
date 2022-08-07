package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	OffHook State = iota
	Connecting
	Connected
	OnHold
	OnHook
)

const (
	CallDialed Trigger = iota
	HungUp
	CallConnected
	PlacedOnHold
	TakenOffHold
	LeftMessage
)

var rules = map[State][]TriggerResult{
	OffHook: {
		{
			Trigger: CallDialed,
			State:   Connecting,
		},
	},
	Connecting: {
		{
			Trigger: HungUp,
			State:   OnHook,
		},
		{
			Trigger: CallConnected,
			State:   Connected,
		},
	},
	Connected: {
		{
			Trigger: LeftMessage,
			State:   OnHook,
		},
		{
			Trigger: HungUp,
			State:   OnHook,
		},
		{
			Trigger: PlacedOnHold,
			State:   OnHold,
		},
	},
	OnHold: {
		{
			Trigger: TakenOffHold,
			State:   Connected,
		},
		{
			Trigger: HungUp,
			State:   OnHook,
		},
	},
}

type State int

func (s State) String() string {
	switch s {
	case OffHook:
		return "OffHook"
	case Connecting:
		return "Connecting"
	case Connected:
		return "Connected"
	case OnHold:
		return "OnHold"
	case OnHook:
		return "OnHook"
	default:
		return "Unknown"
	}
}

type Trigger int

func (t Trigger) String() string {
	switch t {
	case CallDialed:
		return "CallDialed"
	case HungUp:
		return "HungUp"
	case CallConnected:
		return "CallConnected"
	case PlacedOnHold:
		return "PlacedOnHold"
	case TakenOffHold:
		return "TakenOffHold"
	case LeftMessage:
		return "LeftMessage"
	default:
		return "Unknown"
	}
}

type TriggerResult struct {
	Trigger Trigger
	State   State
}

func main() {
	state, exitState := OffHook, OnHook
	for ok := true; ok; ok = state != exitState {
		fmt.Println("The phone is currently", state)
		fmt.Println("Select a trigger:")

		for i := 0; i < len(rules[state]); i++ {
			tr := rules[state][i]
			fmt.Println(strconv.Itoa(i), "-", tr.Trigger)
		}

		input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		i, _ := strconv.Atoi(string(input))

		state = rules[state][i].State
	}
	fmt.Println("We are done using the phone")
}
