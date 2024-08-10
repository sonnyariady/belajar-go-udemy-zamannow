package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Karyawan struct {
	Nama       string `required:"true" max:"10"`
	Jabatan    string `required:"true" max:"30"`
	Umur       int    `required:"true"`
	Tanggungan string `max:"30"`
	KodeCabang string `max:"4"`
}

type BukuAlamat struct {
	Nama   string `required:"true" max:"20"`
	Alamat string `required:"true" max:"50"`
	NoTelp string `required:"true" max:"12"`
}

func bacaField(value any) {
	var nilaitipe = reflect.TypeOf(value)
	fmt.Println("Nilai tipe: ", nilaitipe)
	fmt.Println("Nilai tipe name: ", nilaitipe.Name())
	for i := 0; i < nilaitipe.NumField(); i++ {
		var structfield reflect.StructField = nilaitipe.Field(i)
		fmt.Println(structfield.Name, "dengan tipe", structfield.Type)
		fmt.Println("Tag Required: ", structfield.Tag.Get("required"))
		fmt.Println("Tag Max: ", structfield.Tag.Get("max"))
		if structfield.Tag.Get("max") != "" {
			databaca := reflect.ValueOf(value).Field(i).Interface()
			dataStr := fmt.Sprintf("%s", databaca)
			fmt.Printf("Panjangnya %s %d : ", dataStr, len(dataStr))
		}

	}
}

func IsValid(value any) bool {
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			return data != ""
		}

		if f.Tag.Get("max") != "" {
			databaca := reflect.ValueOf(value).Field(i).Interface()
			dataStr := fmt.Sprintf("%s", databaca)
			pjg, _ := strconv.Atoi(f.Tag.Get("max"))
			return len(dataStr) <= pjg
		}
	}
	return true
}

func validateMaxLength(s interface{}, maxLength int) {
	v := reflect.ValueOf(s)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		// Hanya memproses field dengan tipe string
		if field.Kind() == reflect.String {
			fieldName := t.Field(i).Name
			fieldValue := field.String()

			// Periksa panjang string
			if len(fieldValue) > maxLength {
				fmt.Printf("Field '%s' memiliki panjang %d, melebihi maksimum %d\n", fieldName, len(fieldValue), maxLength)
			} else {
				fmt.Printf("Field '%s' memiliki panjang %d, sesuai dengan maksimum %d\n", fieldName, len(fieldValue), maxLength)
			}
		}
	}
}

func main() {
	karyawan1 := Karyawan{
		Nama:       "Sonny",
		Jabatan:    "Backend",
		Umur:       41,
		KodeCabang: "2001",
	}

	alamat1 := BukuAlamat{"Budi", "Kebon Jeruk", "0817123"}
	fmt.Println("Cetak karyawan:", karyawan1)
	bacaField(karyawan1)
	fmt.Println("Cetak bukualamat:", alamat1)
	bacaField(alamat1)

	karyawan2 := Karyawan{
		Jabatan: "Frontend",
	}

	karyawan3 := Karyawan{
		Nama:       "Abdul",
		Jabatan:    "QA",
		Umur:       43,
		KodeCabang: "2012312345",
	}

	fmt.Println("IsValid Karyawan1: ", IsValid(karyawan1))
	fmt.Println("IsValid Karyawan2: ", IsValid(karyawan2))
	fmt.Println("IsValid Karyawan3: ", IsValid(karyawan3))
}
