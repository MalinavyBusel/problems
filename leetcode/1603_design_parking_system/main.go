package _1603_design_parking_system

type ParkingSystem struct {
	slots map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{slots: map[int]int{1: big, 2: medium, 3: small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.slots[carType] > 0 {
		this.slots[carType] -= 1
		return true
	}
	return false
}
