package main

import (
	"fmt"
	"minitask/utils"
)

func main() {
	// Nomor 1
	var person = utils.BioData{
		Name:         "Aisha Nur Ramadhani",
		Age:          24,
		Phone:        "088899996666",
		Photo:        "https://media.hswstatic.com/eyJidWNrZXQiOiJjb250ZW50Lmhzd3N0YXRpYy5jb20iLCJrZXkiOiJnaWZcL3BsYXlcLzBiN2Y0ZTliLWY1OWMtNDAyNC05ZjA2LWIzZGMxMjg1MGFiNy0xOTIwLTEwODAuanBnIiwiZWRpdHMiOnsicmVzaXplIjp7IndpZHRoIjo4Mjh9fX0=",
		Email:        "admin@gmail.com",
		MarialStatus: false,
		School:       [2]utils.School{{Name: "SMAN 1 Kota Tangerang", Major: "Science"}, {Name: "Universitas Padjadjaran", Major: "Sarjana Terapan Kebidanan"}},
	}

	fmt.Println(`Biodata`)
	fmt.Printf("Name  : %s\n", person.Name)
	fmt.Printf("Age  : %d\n", person.Age)
	fmt.Printf("Phone : %s\n", person.Phone)
	fmt.Printf("Photo : %s\n", person.Photo)
	fmt.Printf("Email : %s\n", person.Email)
	fmt.Printf("Marial Status : %t\n", person.MarialStatus)
	fmt.Printf("School 1  : %s", person.School[1].Name)
	fmt.Printf("Major : %s\n", person.School[1].Major)
	fmt.Printf("School 2  : %s", person.School[0].Name)
	fmt.Printf("Major : %s\n", person.School[0].Major)

	// Nomor 2
	status, err := utils.Login("admin@gmail.com", "admin123")
	status1, err1 := utils.Login("admin@gmail.co11m", "admin0")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(status)
	}

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(status1)
	}

	// Nomor 3
	var numbers = [...]uint{2, 3, 2, 4, 3, 9, 10, 26, 11, 24, 67, 11, 11, 8}
	fmt.Println(utils.BigNumber(numbers[0:]))

	// Nomor 4
	// temperature := utils.Temperature(50, "C", "K")
	// fmt.Printf("Temperature = %.2f\n", temperature)

	fmt.Printf("Temperature = %.2f\n", utils.Temperature(50, "R", "C"))
	fmt.Printf("Temperature = %.2f\n", utils.Temperature(50, "R", "K"))
	fmt.Printf("Temperature = %.2f\n", utils.Temperature(50, "R", "F"))
	fmt.Printf("Temperature = %.2f\n", utils.Temperature(50, "R", "Re"))

	// Nomor 5
	utils.SliceInteger()
}
