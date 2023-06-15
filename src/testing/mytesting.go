package mytesting

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(val float64) (float64, error) {
	if val < 0 {
		return 0.0, errors.New("negative number")
	}
	return math.Sqrt(val), nil
}

func main() {
	res, err := Sqrt(2.)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	
}