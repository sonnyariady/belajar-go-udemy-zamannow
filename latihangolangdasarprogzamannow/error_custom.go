package main

import "fmt"

type errorValidasi struct {
	Pesan string
}

func (v *errorValidasi) Error() string {
	return v.Pesan
}

type errorNotFound struct {
	Pesan string
}

func (notfound *errorNotFound) Error() string {
	return notfound.Pesan
}

func SaveData(id string, data any) error {
	if id == "" {
		return &errorValidasi{Pesan: "Error validasi"}
	}
	if id == "sonny" {
		return &errorNotFound{Pesan: "data tidak ketemu"}
	}
	return nil
}

func main() {
	err := SaveData("sonny", nil)

	if err != nil {
		/*
				if validationErr, ok := err.(*errorValidasi); ok {
					fmt.Println("validation error:", validationErr.Error())
				} else if notFoundErr, ok := err.(*errorNotFound); ok {
					fmt.Println("not found error:", notFoundErr.Error())
				} else {
					fmt.Println("unknown error:", err.Error())
				}
			} else {
				fmt.Println("Simpan sukses")
			}
		*/
		switch finalError := err.(type) {
		case *errorValidasi:
			fmt.Println("validation error:", finalError.Error())
		case *errorNotFound:
			fmt.Println("not found error:", finalError.Error())
		default:
			fmt.Println("unknown error:", finalError.Error())
		}

	} else {
		// sukse
		fmt.Println("Sukses menyimpan")
	}
}
