package _1603_design_parking_system

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_isValid(t *testing.T) {
	t.Run("test_adding_cars_to_parking_spaces", func(t *testing.T) {
		parking := Constructor(1, 1, 0)
		want := []bool{true, true, false, false}
		result := []bool{
			parking.AddCar(1),
			parking.AddCar(2),
			parking.AddCar(3),
			parking.AddCar(1),
		}
		require.Equal(t, want, result, "you cannot add more or less then 2 cars, because there is only 2 spaces")
	})
}
