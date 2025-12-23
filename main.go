package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Kitap struct {
	ID     int    `json:"id"`
	Baslik string `json:"baslik"`
	Yazar  string `json:"yazar"`
	Yil    int    `json:"yil"`
}

var kitaplar = []Kitap{
	{ID: 1, Baslik: "Sefiller", Yazar: "Victor Hugo", Yil: 1862},
	{ID: 2, Baslik: "Nutuk", Yazar: "Mustafa Kemal Atatürk", Yil: 1927},
	{ID: 3, Baslik: "1984", Yazar: "George Orwell", Yil: 1949},
}

func anasayfa(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Merhaba! Kitap listesi icin '/kitaplar' adresine gidin.")
}

func tumKitaplariGetir(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kitaplar)
}

func main() {
	http.HandleFunc("/", anasayfa)
	http.HandleFunc("/kitaplar", tumKitaplariGetir)

	fmt.Println("Sunucu 8080 portunda calisiyor...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}