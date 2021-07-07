/*
値を入れ替えるswap関数を実装してください
次のコードが正しく動作するように実装してください
*/
package main

func main() {
	n, m := swap(10, 20)
	println(n, m)
}

func swap(a, b int) (int, int) {
	return b, a
}

/*名前付き戻り値
func swap(a, b int) (a1, b1 int) {
	a1, b1 = b, a
	return
}
*/
