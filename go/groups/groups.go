package groups

import (
	"fmt"
	"strconv"
	"strings"
)

func fillGroups(passengersGroups string) []int {
	passengersStr := strings.Split(passengersGroups, ",")
	var passengersInt []int
	for i := 0; i < len(passengersStr); i++ {
		p, _ := strconv.Atoi(passengersStr[i])
		passengersInt = append(passengersInt, p)
	}
	return passengersInt
}

func DeterminePossibleSizes(passengersGroups string) string {
	passengersInt := fillGroups(passengersGroups)
	totalPassengers := 0
	for i := 0; i < len(passengersInt); i++ {
		totalPassengers += passengersInt[i]
	}
	var answer []int
	for i := 0; i < totalPassengers; i++ {
		canCarryEveryone := true
		busCapacity := i + 1
		for j := 0; j < len(passengersInt); j++ {
			pasajerosActuales := passengersInt[j]
			if busCapacity < pasajerosActuales {
				canCarryEveryone = false
				break
			}
			if !checkGroups(fillGroups(passengersGroups), busCapacity) || totalPassengers%busCapacity != 0 {
				canCarryEveryone = false
				break
			}
		}
		if canCarryEveryone {
			answer = append(answer, busCapacity)
		}
	}
	result := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(answer)), ","), "[]")
	return result
}

func checkGroups(groups []int, busCapacity int) bool {
	for i := 0; i < len(groups); i++ {
		grupoActual := groups[i]
		if grupoActual > busCapacity {
			return false
		}
		if grupoActual == busCapacity {
			groups = append(groups[:i], groups[i+1:]...)
			i--
		} else if grupoActual < busCapacity {
			passengers := grupoActual
			groups = append(groups[:i], groups[i+1:]...)
			i--
			for passengers < busCapacity {
				if i+1 < len(groups) {
					passengers += groups[i+1]
					groups = append(groups[:i+1], groups[i+2:]...)
				} else {
					break
				}
			}
			if passengers > busCapacity {
				return false
			}
		}
	}
	return len(groups) == 0
}