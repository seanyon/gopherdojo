// 奇数偶数判定関数を作成してください
package main

func main() {
	for i := 1; i <= 100; i++ {
		print(i)
		judge(i)
	}
}

func judge(x int) {
	if x%2 == 0 {
		println("-偶数")
	} else {
		println("-奇数")
	}
}
