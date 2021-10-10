package handlers

import (
	"io/ioutil"
	"net/http"
	"os"
)

type ListHandler struct {}

func (l *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//switch r.FormValue("type") {
	//case "png":
		//name := r.FormValue("name")
		//fmt.Fprintf(w, "Parsed query-param with key \"type\": %s", name)
	//}

	//typeFile := r.FormValue("type")
	//fileLink := "http://localhost:8081/"
	//req, err := http.NewRequest(http.MethodHead, fileLink, nil)
	//if err != nil {
	//	log.Println(err)
	//	http.Error(w, "Unable to check file", http.StatusInternalServerError)
	//	return
	//}
	//cli := &http.Client{}
	//resp, err := cli.Do(req)
	//if err != nil {
	//	log.Println(err)
	//	http.Error(w, "Unable to check file", http.StatusInternalServerError)
	//	return
	//}
	//list, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Println(err)
	//	http.Error(w, "Unable to check file", http.StatusInternalServerError)
	//	return
	//}
	//
	//
	//fmt.Fprintln(w, string(list))

	//var url = r.URL.Path[1:]   // Упростим напсание кода
	////if url == "" || url == "/" { // Дабы разрешить "/" запросы
	//	url = "index.html"
	////}
	//
	//data := readFile(url) // Считываем файл
	//
	//// Отправляет файл в ответ на запрос
	//fmt.Fprintln(w, data)

	//req, err := http.NewRequest(http.MethodGet, "http://localhost:8081/", nil)
	//if err != nil {
	//	log.Println(err)
	//	http.Error(w, "Unable to check file", http.StatusInternalServerError)
	//	return
	//}
	//
	//cli := &http.Client{}
	//resp, err := cli.Do(req)
	//if err != nil {
	//	log.Println(err)
	//	http.Error(w, "Unable to check file", http.StatusInternalServerError)
	//	return
	//}
	//defer resp.Body.Close()
	//
	//body, _ := ioutil.ReadAll(resp.Body)
	//
	//
	//
	//fmt.Fprintln(w, string(body))

}

// Считыватель файлов. Принимает имя файла и выдаёт его содержимое
func readFile(iFileName string) string {
	// Считываем файл
	lData, err := ioutil.ReadFile(iFileName)
	var lOut string // Объявляем строчную переменную
	// Если файл существует - записываем его в переменную lOut
	if !os.IsNotExist(err) {
		lOut = string(lData)
	} else { // Иначе - отправляем стандартный текст
		lOut = "404"
	}
	return lOut // Возвращаем полученную информацию
}
