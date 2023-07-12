package main

import "fmt"

func arrayList() [5]string {

	// array com tamanho pré determinado
	var array [5]string
	array[0] = "item 1"
	array[1] = "item 2"
	array[2] = "item 3"
	array[3] = "item 4"
	array[4] = "item 5"
	return array
}

func sliceList() []string {

	// array sem tamanho definido
	itens := []string{"item 10", "item 20", "item 30"}

	// adicionando um item no array
	fmt.Println(append(itens, "item 40"))

	// verificando a quantidade de itens dentro do array
	fmt.Println(len(itens))

	// verificando a capacidade do array. PS: Quando o array "estoura", é dobrado a quantidade de espaços internamente
	fmt.Println(cap(itens))
	return itens
}

func forArray() {

	itens := []string{"item 10", "item 20", "item 30"}

	// for tradicional. PS: Não utiliza parênteses
	for i := 0; i < len(itens); i++ {
		fmt.Println(itens[i])
	}

	// for range. semelhante ao forEach
	for i, item := range itens {
		fmt.Println("meu index é:", i, "e meu item é:", item)
	}
}
