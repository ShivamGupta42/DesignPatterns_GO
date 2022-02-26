package main

import "testing"

func TestFilterProducts(t *testing.T) {

	t.Run("Test Single Filters", func(t *testing.T) {
		apple := Product{"apple", 100, green}
		watch := Product{"watch", 2100, red}
		laptop := Product{"laptop", 45000, blue}

		n := NamePredicate{"apple"}
		products := []*Product{&apple, &watch, &laptop}
		predicates := []ProductPredicate{&n}

		res := FilterProducts(products, predicates)

		if len(res) != 1 {
			t.Errorf("Length of res is not 1")
		}

		if *res[0] != apple {
			t.Errorf("Same object is not returned")
		}
	})
}
