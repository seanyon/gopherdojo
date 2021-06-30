/*
3つの変数しか使わないように修正してください
プログラムの動作はそのままにすること

package main
func main() {
	n1 := 19
	n2 := 86
	n3 := 1
	n4 := 12

	sum := n1 + n2 + n3 + n4
	println(sum)
}
*/

package main

func main() {
	ns := [...]int{19, 86, 1, 12}
	sum := ns[0] + ns[1] + ns[2] + ns[3]
	println(sum)
}
