package Helper

import (
	"os"
	"paramparser/custom"
	"strings"
)

var (
	ModoUso string
	Helpers []HelperAgent
)

type HelperAgent struct {
	Argumento string
	Alias     string
	Exemplo   string
	Descricao string
}

func CadastrarHelper(argumento, alias, exemplo, funcao string) {
	helper := HelperAgent{
		Argumento: argumento,
		Alias:     alias,
		Exemplo:   exemplo,
		Descricao: funcao,
	}
	Helpers = append(Helpers, helper)
}

func MostrarAjuda() {
	d0 := 15
	d1 := 9
	d2 := 3
	custom.MiniBanner("*Menu de ajuda*")
	custom.Println("\n\n"+ModoUso, "yellow")
	if len(Helpers) > 0 {
		custom.Println("\nOpções:", "magenta")
		for _, helper := range Helpers {
			argumento := helper.Argumento
			alias := helper.Alias
			exemplo := helper.Exemplo
			descricao := helper.Descricao
			space0 := d0 - len(argumento)
			space1 := d1 - len(alias)
			space2 := d2 - len(exemplo)
			texto := argumento + strings.Repeat(" ", space0) + alias + strings.Repeat(" ", space1) + exemplo + strings.Repeat(" ", space2) + descricao
			custom.Println(texto, "yellow")
		}
	}
	os.Exit(0)
}
