package main

func ChooseBestSum(t, k int, ls []int) int {
	if k > len(ls) {
		return -1
	}

	combinations := [][]int{}

	combination := []int{};
	for i:= 0; i < k; i++{
		combination = append(combination, i)
	}
	combinations = append(combinations, combination)

	j:= k - 1
	for j > 0{
		
	}


	}

	// your code
}
