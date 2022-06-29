package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 5
	fmt.Printf("x: %v\n", x)
	var total float64
	for i := 0; i < len(x); i++ {
		total += float64(x[i])
	}
	fmt.Printf("total: %v\n", total)

	//slices

	y := make([]float64, 5)
	fmt.Printf("y: %v\n", y)

	z := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("z: %v\n", z)

	z_sliced := z[3:5]
	fmt.Printf("z_sliced: %v\n", z_sliced)
	z_sliced = append(z_sliced, 99, 100)
	fmt.Printf("z_sliced after append: %v\n", z_sliced)

	copy(z_sliced, z)
	fmt.Printf("z_sliced after copy: %v\n", z_sliced)

	//map
	//var x_map map[string]int  <---- error
	x_map := make(map[string]int)
	x_map["key"] = 12
	x_map["yek"] = 21
	delete(x_map, "key")
	fmt.Printf("x_map: %v\n", x_map)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])
	name, ok := elements["Un"] // <----- ok false here
	fmt.Println(name, ok)
}
