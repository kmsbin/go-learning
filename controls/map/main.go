package main

import "fmt"

func main() {
	// maps devem ser inicializados, se não ocorrerá um erro
	// e assim como o array ele tem tamanho fixo então a melho alternatica é usar o make()
	// var aprov map[int]string
	aprov := make(map[int]string)

	aprov[13047347905] = "kaulindo"
	aprov[12341346354] = "geraldin"
	aprov[12342322435] = "alaska"

	fmt.Println(aprov)

	for cpf, name := range aprov {
		fmt.Printf("%s CPF: %d\n", name, cpf)
	}

	fmt.Println(aprov[13047347905])
	delete(aprov, 12341346354)
	fmt.Println(len(aprov), aprov)
}
