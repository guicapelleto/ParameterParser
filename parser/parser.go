package parser

import (
	"paramparser/custom"
	"strconv"
	"strings"
)

type GetParameter struct {
	Url   string
	Param []string
}

type FormsParameter struct {
	Url    string
	Name   string
	Method string
	Action string
	Inputs []string
}

type Comentario struct {
	Url         string
	Comentarios []string
}

var (
	Gets        []GetParameter
	Forms       []FormsParameter
	Comentarios []Comentario
	Urls        []string
	JavaScripts []string
	MainUrl     string
	MainDomain  string
	Subdomains  []string
	BodyLogs    []string
	BadUrls     []string
)

func bodySplitter(form, filtro, filtro_final, campo_divisor string) (results []string) {
	split_data := strings.Split(form, filtro)
	for _, data := range split_data {
		contents := strings.Split(data, campo_divisor)
		if len(contents) == 1 {
			continue
		}
		content := strings.Split(contents[1], filtro_final)[0]
		if content == "" {
			continue
		}
		results = append(results, content)
	}
	return
}

func AttributesParser(recursive bool) {
	totalPage := 1
	custom.Aviso("Coletando informações...")
	HttpParser(MainUrl)
	if recursive {
		custom.Aviso("Modo recursivo ativo, pode levar mais tempo que o normal\n\n")
		for _, url := range Urls {
			if !custom.SliceStrContains(BodyLogs, url) {
				totalPage++
				txt := "Total de páginas verificadas: " + strconv.Itoa(totalPage)
				custom.PrintRecursive(txt, custom.RandomColorPicker())
				HttpParser(url)
			}
		}
		custom.PrintRecursive("", "")
	}
	JsParser()            //Adiciona JS
	UrlParametersParser() // Adiciona os parametros get

}

func HttpParser(url string) {
	mainpage, err := custom.HttpGet(url)
	if err != nil {
		return
	}
	BodyLogs = StringUniqAppend(BodyLogs, url)
	//	custom.ExitOnError("Um erro ocorreu ao processar a página: ", err, true)
	hrefs := bodySplitter(mainpage, "<", "\"", "href=\"")
	srcs := bodySplitter(mainpage, "<", "\"", "src=\"")
	comments := bodySplitter(mainpage, "<", "-->", "!--")
	if len(comments) != 0 {
		comment := Comentario{Url: url, Comentarios: comments}
		Comentarios = append(Comentarios, comment)
	}
	for _, href := range hrefs {
		if len(href) > 1 && href[:1] == "/" {
			href = MainUrl + href
		}
		if len(href) > 1 && href[:1] == "#" {
			href = MainUrl + "/" + href
		}
		if len(href) > 4 && href[:4] != "http" {
			href = MainUrl + "/" + href
		}
		if Blacklist(href) {
			if !strings.Contains(href, MainDomain) {
				BadUrls = StringUniqAppend(BadUrls, href)
			}
			continue
		}
		if strings.Contains(href, MainDomain) && href[:4] == "http" {
			if !custom.SliceStrContains(Urls, href) {
				subdomain := subdomainExtract(href)
				if subdomain != "" {
					Subdomains = StringUniqAppend(Subdomains, subdomain) // Adiciona subdomains
					Urls = StringUniqAppend(Urls, href)                  // Adiciona URLS

				}
			}
		}
	}
	for _, href := range srcs {
		if len(href) > 1 && href[:1] == "/" {
			href = MainUrl + href
		}
		if len(href) > 1 && href[:1] == "#" {
			href = MainUrl + "/" + href
		}
		if len(href) > 4 && href[:4] != "http" {
			href = MainUrl + "/" + href
		}
		if Blacklist(href) {
			if !strings.Contains(href, MainDomain) {
				BadUrls = StringUniqAppend(BadUrls, href)
			}
			continue
		}
		if strings.Contains(href, MainDomain) && href[:4] == "http" {
			if !custom.SliceStrContains(Urls, href) {
				subdomain := subdomainExtract(href)
				if subdomain != "" { //realiza acao somente se for do domínio
					Subdomains = StringUniqAppend(Subdomains, subdomain) // Adiciona subdomains
					Urls = StringUniqAppend(Urls, href)                  // Adiciona URLS
				}
			}
		}
	}
	FormParser(mainpage, url) //Adiciona Forms
}

func subdomainExtract(url string) string {
	if strings.Contains(url, "?") {
		url = strings.Split(url, "?")[0]
	}
	splited_url := strings.Split(url, "/")
	if len(splited_url) < 3 {
		return ""
	}
	if splited_url[1] != "" || splited_url[2] == "" {
		return ""
	}
	if strings.Contains(splited_url[2], MainDomain) {
		return splited_url[2]
	}
	return ""
}

func JsParser() {
	for _, url := range Urls {
		if strings.Contains(url, ".js") {
			jss := strings.Split(url, ".js")
			jsfile := ""
			if url[len(url)-3:] == ".js" {
				jsfile = url
			}
			if len(jss) > 1 && url[len(url)-3:] != ".js" {
				for _, badpart := range strings.Split(url, subdomainExtract(url))[1:] {
					if strings.Contains(badpart, ".js") {
						badpart = strings.Split(badpart, ".js")[0]
						jsfile = strings.Split(url, badpart)[0] + badpart + ".js"
					}
				}
			}
			if jsfile != "" {
				JavaScripts = StringUniqAppend(JavaScripts, jsfile)
			}
		}
	}
}

func FormParser(body, url string) {
	forms := strings.Split(body, "</form>")
	for _, form := range forms {
		markup := false
		formParameter := FormsParameter{
			Url: url,
		}
		formname := bodySplitter(form, "<form", "\"", "name=\"")
		formaction := bodySplitter(form, "<form", "\"", "action=\"")
		formmethod := bodySplitter(form, "<form", "\"", "method=\"")
		if len(formname) != 0 {
			formParameter.Name = formname[0]
			markup = true
		}
		if len(formaction) != 0 {
			formParameter.Action = formaction[0]
			markup = true
		}
		if len(formmethod) != 0 {
			formParameter.Method = formmethod[0]
			markup = true
		}
		inputs := bodySplitter(form, "<input", "\"", "id=\"")
		for _, input := range inputs {
			formParameter.Inputs = append(formParameter.Inputs, input)
			markup = true
		}
		if markup {
			Forms = append(Forms, formParameter)
		}
	}

}

func UrlParametersParser() {
	for _, url := range Urls {
		first_parameter := bodySplitter(url, "/", "=", "?")
		if len(first_parameter) == 0 {
			continue
		}
		first_parameter = append(first_parameter, bodySplitter(url, "?", "=", "&")...)
		index, confirm := findParameter(url)
		if confirm {
			for _, p := range first_parameter {
				Gets[index].Param = StringUniqAppend(Gets[index].Param, p)
			}
		} else {
			getvar := GetParameter{
				Url:   treatUrl(url),
				Param: first_parameter,
			}
			Gets = append(Gets, getvar)
		}
	}
}

func StringUniqAppend(parameters []string, param string) []string {
	if custom.SliceStrContains(parameters, param) {
		return parameters
	} else {
		return append(parameters, param)
	}
}

func treatUrl(url string) string {
	url = strings.Split(url, "?")[0]
	return url
}

func findParameter(fullUrl string) (int, bool) {
	for n, p := range Gets {
		if treatUrl(fullUrl) == p.Url {
			return n, true
		}
	}
	return 0, false
}
