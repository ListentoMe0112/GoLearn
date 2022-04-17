package dimensionArray

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func Test01() {
	var scores [3][5]float64
	var scores_sum [3]float64
	var scores_avg [3]float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scanln(&scores[i][j])
			scores_sum[i] += scores[i][j]
		}
		scores_avg[i] = scores_sum[i] / 5.0
	}
	var total float64
	var average float64
	for i := 0; i < 3; i++ {
		total += scores_sum[i]
	}
	average = total / 15
	fmt.Println("total average ", average)

	for i := 0; i < 3; i++ {
		fmt.Printf("class %d: %f\n", i, scores_avg[i])
	}
}

func Test02() {
	var nums []int
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Intn(100)+1)
	}

	sort.Ints(nums)

	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Println(nums[i])
	}

	fmt.Println(sort.SearchInts(nums, 55))
}
