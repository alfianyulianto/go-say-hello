package go_say_hello

//func SayHello() string {
//	return "Hello World"
//}

// semua yang memanggil function SayHello akan rusak karena yang seblumnya tidak ada param menjadi ada param
// inikah yang disebut major changes
func SayHello(name string) string {
	return "Hello " + name
}
