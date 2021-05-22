package main

import "math"

type Vehicle struct {
	id               int
	maxSpeed         int
	maxLoad          int
	nextDeliveryTime float32
}

func NewVehicle(id, maxSpeed, maxLoad int) *Vehicle {
	return &Vehicle{
		id:       id,
		maxSpeed: maxSpeed,
		maxLoad:  maxLoad,
	}
}

func (me *Vehicle) SetNextDeliveryTime(totalDeliveryTime float32) {
	me.nextDeliveryTime = roundoff(float32(2) * totalDeliveryTime)
}

type Vehicles []*Vehicle

// Implementing sort.Sort interface
func (me Vehicles) Len() int           { return len(me) }
func (me Vehicles) Swap(i, j int)      { me[i], me[j] = me[j], me[i] }
func (me Vehicles) Less(i, j int) bool { return me[i].nextDeliveryTime < me[j].nextDeliveryTime }

func roundoff(f float32) float32 {
	return float32(math.Floor(float64(f)*float64(100)) / 100)
}

func NewVehicles(vehicleCount, vehicleMaxSpeed, vehicleMaxLoad int) Vehicles {
	vehicles := Vehicles{}
	for i := 0; i < vehicleCount; i++ {
		vehicles = append(vehicles, NewVehicle(i, vehicleMaxSpeed, vehicleMaxLoad))
	}
	return vehicles
}
