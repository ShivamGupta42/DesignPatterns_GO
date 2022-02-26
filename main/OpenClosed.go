package main

type Color int

const (
	red Color = iota
	green
	blue
)

type Product struct {
	name  string
	price int
	color Color
}

type ProductPredicate interface {
	test(p *Product) bool
}

type NamePredicate struct {
	name string
}

type PricePredicate struct {
	price int
}

type ColorPredicate struct {
	color Color
}

func (n *NamePredicate) test(p *Product) bool {
	return n.name == p.name
}

func (pp *PricePredicate) test(p *Product) bool {
	return pp.price == p.price
}

func (c *ColorPredicate) test(p *Product) bool {
	return c.color == p.color
}

func FilterProducts(p []*Product, pr []ProductPredicate) []*Product {

	res := make([]*Product, 0)

	for _, v := range p {
		for _, t := range pr {
			if t.test(v) {
				res = append(res, v)
			}
		}
	}
	return res
}
