package main

func sum(a, b int) int {
	return a + b
}

func sumAsPointers(a, b *int) int {
	return *a + *b
}



// func main() {
// 	minesVar1 := 10
// 	minesVar2 := 10
// 	println(sumAsPointers(&minesVar1, &minesVar2))
// 	println(minesVar1)
// }
