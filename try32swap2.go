/*
値を入れ替えるswap2関数を実装してください
次のコードが正しく動作するように実装してください
*/
package main

func main() {
	n, m := 10, 20
	swap2(&n, &m)
	println(n, m)
}

func swap2(np, mp *int) (int, int) {
	*np, *mp = *mp, *np
	return *np, *mp
}

/*下の関数だとうまく行かないのは、*np, *mpに対し破壊的操作が加わっていないから
func swap2 (np, mp *int)(int, int){
	return  *mp, *np
}
*/
