package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/cassio/go-files/api/service"
)

func main() {

	service.StartServer()

}

func CriarAquivo() {

	file, err := os.Create("empty.txt")

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("file created")
}

func VerificarSeExiste() {
	_, err := os.Stat("words.txt")

	if errors.Is(err, os.ErrNotExist) {

		fmt.Println("file does not exist")
	} else {

		fmt.Println("file exists")
	}
}

func DeletarArquivo() {

	err := os.Remove("words.txt")

	if err != nil {

		log.Fatal(err)
	}

	fmt.Println("file deleted")

}

func lerUmAquivo() {

	content, err := ioutil.ReadFile("words.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))

}

func InverterBarra(palavra string) string {
	BarraInvertidaParaCaminhoWindows := strings.Replace(palavra, "/", "\\", 1)
	fmt.Print(BarraInvertidaParaCaminhoWindows)
	return BarraInvertidaParaCaminhoWindows
}
