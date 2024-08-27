package main

type Goods interface {
	GetPrice() float64
	GetNum() int
}

type Book struct {
	Name  string
	Price float64
	Num   int
}

func (b *Book) GetPrice() float64 {//最终实现接口的数据类型是结构体指针
	return b.Price
}

func (b *Book) GetNum() int {
	return b.Num
}

type Phone struct {
	Brand    string
	Discount float64
	Price    float64
	Num      int
}

func (p *Phone) GetPrice() float64 {
	return p.Price * p.Discount
}

func (p *Phone) GetNum() int {
	return p.Num
}

type Cart struct {
	Goods []Goods
}

func (c *Cart) Add(g Goods) {// 只有实现了Goods接口的商品才允许被添加到购物车中
	c.Goods = append(c.Goods, g)
}

func (c *Cart) TotalPrice() float64 {
	//计算购物车物品价格...
	var totalPrice float64
	for _, g := range c.Goods {
		totalPrice += float64(g.GetNum()) * g.GetPrice()
	}
	return totalPrice
}

func main() {
	//图书类商品
	b1 := Book{Name: "Go从入门到精通", Price: 50, Num: 2}
	b2 := Book{Name: "Go Action", Price: 60, Num: 1}
	//手机
	p1 := Phone{Brand: "华为", Price: 6000, Num: 1, Discount: 0.8}

	c := &Cart{}
	c.Add(&b1) //添加到购物车
	c.Add(&b2)
	c.Add(&p1)
	//fmt.Println(c.TotalPrice())//计算价格
}