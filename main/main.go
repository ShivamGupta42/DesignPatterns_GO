package main

import "fmt"

func main() {
	apple := Product{"apple", 100, green}
	watch := Product{"watch", 2100, red}
	laptop := Product{"laptop", 45000, blue}

	n := NamePredicate{"apple"}
	products := []*Product{&apple, &watch, &laptop}
	predicates := []ProductPredicate{&n}

	res := FilterProducts(products, predicates)
	fmt.Println(*res[0])
}
