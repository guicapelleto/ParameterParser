package main

import (
	"os"
	"paramparser/Helper"
	"paramparser/custom"
	"paramparser/parser"
)

var (
	Verbose   = false
	Salvar    = false
	Recursivo = false
)

func funcoes() {
	Helper.ModoUso = "*Modo de uso*\nparamparser [URL] --opcoes\n\nEx: paramparser https://meualvo.com -r -s"
	Helper.CadastrarHelper("\nOPÇÃO", "   ALIAS", " ", "   DESCRIÇÃO")
	Helper.CadastrarHelper("--recursivo", "-r", " ", "Faz a verificação em todas as páginas recursivamente")
	Helper.CadastrarHelper("--verbose", "-v", " ", "Mostra todos os resultados na tela")
	Helper.CadastrarHelper("--salvar", "-s", " ", "Gera um relatório com os resultados obtidos")
	Helper.CadastrarHelper("--help", "-h", " ", "Mostra o menu de ajuda")
}

func executarOpcao(opcao string) {
	switch opcao {
	case "1":
		parser.ShowSubdomains()
	case "2":
		parser.ShowUrls()
	case "3":
		parser.ShowBadUrls()
	case "4":
		parser.ShowAllJsFiles()
	case "5":
		parser.ShowGetParametersByUrl()
	case "6":
		parser.ShowGetParameters()
	case "7":
		parser.ShowForms()
	case "8":
		parser.ShowComments()
	case "9":
		parser.ShowAll()
	case "s":
		parser.GerarRelatorio()
	case "q":
		os.Exit(0)
	}
}

func menuResultado() {
	custom.MiniBanner("Listagem de Opçoes:")
	opcoes := `
1) Mostrar Subdomínios
2) Mostrar URLs relacionadas
3) Mostrar outras URLs
4) Mostrar arquivos JavaScript
5) Mostrar parâmetros GET por URL
6) Mostrar todos parâmetros GET obtidos
7) Mostrar formulários encontrados
8) Mostrar comentários obtidos por página
9) Mostrar todos os resultados
s) Salvar resultados
q) Sair da aplicação
`
	for {
		custom.Print(opcoes, custom.RandomColorPicker())
		opt := custom.GetInput("\nOpção: ")
		executarOpcao(opt)
	}
}

func argParser() {
	funcoes()
	if len(os.Args) == 1 {
		Helper.MostrarAjuda()
	}
	for _, arg := range os.Args[1:] {
		switch arg {
		case "-h", "--help":
			Helper.MostrarAjuda()
		case "-v", "--verbose":
			Verbose = true
		case "-s", "--salvar":
			Salvar = true
		case "-r", "--recursivo":
			Recursivo = true
		}
	}
}

func main() {
	custom.Showbanner()
	argParser()
	parser.MainUrl = os.Args[1]
	maindomain, err := custom.StractDomain(parser.MainUrl)
	custom.ExitOnError("Um erro ocorreu ao processar a URL: ", err, true)
	parser.MainDomain = maindomain
	custom.Println("Domínio atual: "+parser.MainDomain, "green")
	parser.AttributesParser(Recursivo)
	parser.ShowResults()
	if Verbose {
		parser.ShowAll()
	} else {
		menuResultado()
	}
	if Salvar {
		parser.GerarRelatorio()
	}
}
