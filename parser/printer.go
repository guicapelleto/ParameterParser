package parser

import (
	"os"
	"paramparser/custom"
	"strconv"
	"strings"
)

var (
	Relatorio []string
)

func ShowResults() {
	qtd_subdomain := len(Subdomains)
	qtd_url := len(Urls)
	qtd_url_outras := len(BadUrls)
	qtd_js := len(JavaScripts)
	qtd_get := len(Gets)
	qtd_form := len(Forms)
	qtd_comment := len(Comentarios)
	custom.MiniBanner("Resultados")
	custom.Aviso("Encontrado " + strconv.Itoa(qtd_subdomain) + " subdomínios")
	custom.Aviso("Encontrado " + strconv.Itoa(qtd_url) + " URLs correlacionadas")
	custom.Aviso("Encontrado " + strconv.Itoa(qtd_url_outras) + " outras URLs")
	custom.Aviso("Encontrado " + strconv.Itoa(qtd_js) + " arquivos JavaScript")
	custom.Aviso("Encontrado " + strconv.Itoa(qtd_get) + " parâmetros GET em URL")
	custom.Aviso("Encontrado " + strconv.Itoa(qtd_form) + " formulário em páginas")
	custom.Aviso("Encontrado " + strconv.Itoa(qtd_comment) + " comentários")
}

func ShowAll() {
	ShowSubdomains()
	ShowUrls()
	ShowBadUrls()
	ShowAllJsFiles()
	ShowGetParametersByUrl()
	ShowGetParameters()
	ShowForms()
	ShowComments()
}

func ShowGetParametersByUrl() {
	custom.MiniBanner("GET Parameters por URL:")
	for _, getvar := range Gets {
		custom.Println("\n[*]"+getvar.Url, "yellow")
		//		custom.Println("Parâmetros GET:", "magenta")
		for _, param := range getvar.Param {
			custom.Println(param, "green")
		}
	}
}

func ShowGetParameters() {
	var parameters []string
	custom.MiniBanner("Todos os GET Parameters:")
	for _, getvar := range Gets {
		//		custom.Println("\n[*]"+getvar.Url, "yellow")
		//		custom.Println("Parâmetros GET:", "magenta")
		for _, param := range getvar.Param {
			if !custom.SliceStrContains(parameters, param) {
				custom.Println(param, "green")
				parameters = append(parameters, param)
			}
		}
	}
}

func ShowAllJsFiles() {
	custom.MiniBanner("Arquivos JavaScripts:")
	for _, js := range JavaScripts {
		custom.Println(js, "green")
	}
}

func ShowUrls() {
	custom.MiniBanner("Urls Relacionadas:")
	for _, url := range Urls {
		custom.Println(url, "green")
	}
}

func ShowBadUrls() {
	custom.MiniBanner("Outras urls encontradas:")
	for _, url := range BadUrls {
		custom.Println(url, "green")
	}
}

func ShowSubdomains() {
	custom.MiniBanner("Subdomains encontrados:")
	for _, url := range Subdomains {
		custom.Println(url, "green")
	}
}

func ShowComments() {
	custom.MiniBanner("Comments encontrados:")
	for _, comentario := range Comentarios {
		custom.Println(comentario.Url, "green")
		for _, comment := range comentario.Comentarios {
			custom.Println("<!-- "+comment+" -->", "magenta")
		}
	}
}

func ShowForms() {
	custom.MiniBanner("Forms encontrados:")
	for _, getvar := range Forms {
		custom.Println("\n[*]"+getvar.Url, "yellow")
		custom.Println("Nome:"+getvar.Name, "magenta")
		custom.Println("Method:"+getvar.Method, "magenta")
		custom.Println("Action:"+getvar.Action, "magenta")
		custom.Println("Parâmetros de input:", "magenta")
		for _, param := range getvar.Inputs {
			custom.Println(param, "green")
		}
	}
}

func addrelatorio(texto string) {
	Relatorio = append(Relatorio, texto)
}

func GerarRelatorio() {
	var parameters []string
	nome_arquivo := "Report-" + MainDomain + ".txt"
	addrelatorio("\n\nSubdomains encontrados:")
	for _, url := range Subdomains {
		addrelatorio(url)
	}
	addrelatorio("\n\nUrls Relacionadas:")
	for _, url := range Urls {
		addrelatorio(url)
	}
	addrelatorio("\n\nOutras urls encontradas:")
	for _, url := range BadUrls {
		addrelatorio(url)
	}
	addrelatorio("\n\nArquivos JavaScripts:")
	for _, js := range JavaScripts {
		addrelatorio(js)
	}
	addrelatorio("\n\nTodos os GET Parameters:")
	for _, getvar := range Gets {
		for _, param := range getvar.Param {
			if !custom.SliceStrContains(parameters, param) {
				addrelatorio(param)
				parameters = append(parameters, param)
			}
		}
	}
	addrelatorio("\n\nGET Parameters por URL:")
	for _, getvar := range Gets {
		addrelatorio("\n[*]" + getvar.Url)
		for _, param := range getvar.Param {
			addrelatorio(param)
		}
	}
	addrelatorio("\n\nComments encontrados:")
	for _, comentario := range Comentarios {
		addrelatorio(comentario.Url)
		for _, comment := range comentario.Comentarios {
			addrelatorio("<!-- " + comment + " -->")
		}
	}
	relfinal := strings.Join(Relatorio, "\n")
	os.Create(nome_arquivo)
	arquivo, err := os.OpenFile(nome_arquivo, os.O_WRONLY, 660)
	custom.ExitOnError("Erro ao gerar relatório.", err, true)
	defer arquivo.Close()
	arquivo.WriteString(relfinal)
	arquivo.Close()
}
