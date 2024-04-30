type Carro struct {
	Nome String
}

func (c Carro) andar() {
	fmt.Println(c.Nome, "andou")
}

func (c Carro) parar() {
	fmt.Println(c.Nome, "parou")
}

func main() {

	carro := Carro{
		Nome: "BMWx",
	}

	carro.andar()

	carro.parou()
}