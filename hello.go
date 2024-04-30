type Carro struct {
	Nome String
}

func (c Carro) andar() {
	fmt.Println(c.Nome, "andou")
}

func main() {

	carro := Carro{
		Nome: "BMW",
	}

	carro.andar()
}