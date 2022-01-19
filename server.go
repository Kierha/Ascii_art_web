package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type Data struct {
	Value string
}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/respons.html", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		Value := r.FormValue("ascii")
		namefile := r.FormValue("namefileed")
		Aprint := Data{Value: asciiprint(Value, namefile)}
		templates, _ := template.ParseFiles("./static/respons.html")
		templates.ExecuteTemplate(w, "result", Aprint)
	})
	http.HandleFunc("/generascii.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./static/generascii.html")
		tmpl.Execute(w, "")
	})

	fmt.Printf("Démarage du serveur go sur le port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

//func verifPath(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path != "/verif" {
//		http.Error(w, "404 page not found!!", http.StatusNotFound)
//		return
//	}
//
//	if r.Method != "GET" {
//		http.Error(w, "Method is not supported.", http.StatusNotFound)
//		return
//	}
//
//}

//Code intégré du Ascii-Art

var v string  // valeur du resultat concaténé des tableau
var i int     // index tableau
var j int     // index element
var d int = 0 // distance entre les lettre
//var r int = 0

func asciiprint(s, namefile string) string {

	arg := s
	FormBase := []rune(arg)
	index := 0
	limit := 9

	//FormBase = append(FormBase, arg)
	var result []string
	var ConvertToAscii []string
	var StringLocation rune

	file, err := os.Open(namefile)
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	// var StringLocation rune

	for range arg {
		if FormBase[index] == 32 {
			StringLocation = 1
		} else {
			StringLocation = rune(((FormBase[index] - 32) * 9))
		}

		for counter := 0; counter < 9; counter++ {

			if err != nil {
				log.Fatalf("failed opening file: %s", err)
			}
			for scan.Scan() {

				ConvertToAscii = append(ConvertToAscii, scan.Text())
			}

			result = append(result, ConvertToAscii[StringLocation])

			file.Close()
			StringLocation = StringLocation + 1
		}

		if StringLocation < StringLocation+9 {
			index = index + 1
		}
	}

	var tabResult [][]string
	for i := 0; i < len(result); i += limit {
		tabResult = append(tabResult, result[i:min(i+limit, len(result))])
	}

	tabFinal := make([]string, 9)

	manageTable(tabResult)

	for i = 0; i < len(tabResult); {
		for j <= len(tabResult[i])-1 {
			// ici je recupére la valeur de la case
			s := tabResult[i][j]
			//s[len(v):len(s)] voila un exemple => imagine que s = "bonjour" avec (s[3:len(s)]) en va obtenir => jour et je fait  (s[3:len(s)-1]) je vais obtenir jou
			//en fait cette expréssion nous permet de couper la chaine de caractaire apartir de l'index quand va lui donnée
			v += s[len(v):len(s)]
			if i == len(tabResult)-1 {
				tabFinal[j] = v
				i = 0
				j++
				v = ""
			} else {
				i++
			}
			break
		}
		if j == len(tabResult[i]) {
			break
		}

	}

	var ascii_str string

	//for _, v := range tabFinal {
	//	ascii_str += v
	//	ascii_str += "\n"
	//}

	ascii_str = strings.Join(tabFinal, "\n")
	v = ""
	i = 0
	j = 0
	d = 0
	return ascii_str

}

// renvoie la limite du tableau
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// renvoie la longeur de la plus grande item de tableau d'une lettre
func getMaxLenght(t []string) int {
	var s string = t[0]
	for i := 1; i < len(t); i++ {
		if len(s) < len(t[i]) {
			s = t[i]
		}
	}
	return len(s)
}

func manageTable(tabResult [][]string) {
	s := "" // variable space
	for i := 1; i < len(tabResult); {
		d = getMaxLenght(tabResult[i-1])
		if manageSpace(tabResult[i]) {
			s = " "
		}
		for j := 0; j <= len(tabResult[i])-1; j++ {
			tabResult[i][j] = strings.Repeat(" ", d) + s + tabResult[i][j]
			if j == len(tabResult[i])-1 {
				i++
				break
			}
		}
	}
}

func manageSpace(t []string) bool {
	b := true
	for i := 0; i < len(t)-1; i++ {
		if t[i] != "" {
			b = false
			break
		}
	}
	return b
}
