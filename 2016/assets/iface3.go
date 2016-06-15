package main

import "strconv"
import "fmt"
import "reflect"

type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(uint64(i), 2)
}

type Stringer interface {
	String() string
}

func main() {
	iface := Stringer(Binary(200))
	println(iface)
	typ := reflect.TypeOf(iface)
	fmt.Println("Name:", typ.Name())
	fmt.Println("Packge Path:", typ.PkgPath())
	fmt.Println("Size:", typ.Size())
	fmt.Println("Bits:", typ.Bits())
	fmt.Println("Align:", typ.Align())
	fmt.Println("Comparable:", typ.Comparable())
	fmt.Println("Kind:", typ.Kind())
	fmt.Println("NumMethod:", typ.NumMethod())
	fmt.Println("Method[0]:", typ.Method(0))
}
