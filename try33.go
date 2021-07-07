/*
次のプログラムを正しく動作するようにしてください
Incメソッドは自身を1ずつ加算する
今の実装だと正しく動かない
動かない理由を考え、意図通り動くように修正してください
*/
package main

type MyInt int

/*
func (n MyInt) Inc() { n++ }
func main() {
	var n MyInt
	println(n)
	//ここでコピーが発生する
	n.Inc()
	println(n)
}
*/

func (n *MyInt) Inc() { *n++ }
func main() {
	var n MyInt
	println(n)
	(&n).Inc()
	println(n)
}
