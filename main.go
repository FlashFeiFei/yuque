package main

import (
	"encoding/json"
	"fmt"
	"github.com/FlashFeiFei/yuque/response"
	"io/ioutil"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	content, _ := ioutil.ReadAll(req.Body)

	doc := response.ResponseDocDetailSerializer{}

	json.Unmarshal(content, &doc)
	log.Println(string(content))
	log.Println("===============================================")
	log.Println(doc)
	log.Println(fmt.Sprintf("v=%v, t=%T", doc, doc))
	log.Println(doc.Data.ID)
	log.Println(doc.Data.Body)
	log.Println(doc.Data.DeletedAt)
}

func UserInfo(w http.ResponseWriter, req *http.Request) {
	client := http.Client{}

	creq, _ := http.NewRequest(http.MethodGet, "https://www.yuque.com/api/v2/users/262184", nil)
	creq.Header.Add("X-Auth-Token", "oJz8ZC7jfPD95cDxeTIRyBei6SuuUUrvfcOoDfue")
	resp, _ := client.Do(creq)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}
func main() {

	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/user", UserInfo)

	http.ListenAndServe(":12345", nil)
}
