package main

import (
	"fmt"
	"log"
	"net/http"
)

type Kitap struct {
	ID     int
	Baslik string
	Yazar  string
	Yil    int
}

var kitaplar = []Kitap{
	{ID: 1, Baslik: "Sefiller", Yazar: "Victor Hugo", Yil: 1862},
	{ID: 2, Baslik: "Nutuk", Yazar: "Mustafa Kemal Atatürk", Yil: 1927},
	{ID: 3, Baslik: "1984", Yazar: "George Orwell", Yil: 1949},
}

func anasayfa(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Kutuphane Sistemi</h1><p><a href='/kitaplar'>Kitaplari Listele</a></p>")
}

func tumKitaplariGetir(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	html := `
	<html>
	<head>
		<style>
			table { width: 50%; border-collapse: collapse; font-family: Arial; }
			th, td { border: 1px solid #ddd; padding: 8px; text-align: left; }
			th { background-color: #4CAF50; color: white; }
			tr:nth-child(even) { background-color: #f2f2f2; }
			h1 { font-family: Arial; color: #333; }
		</style>
	</head>
	<body>
		<h1>📚 Kitap Listesi</h1>
		<table>
			<tr>
				<th>ID</th>
				<th>Kitap Adi</th>
				<th>Yazar</th>
				<th>Yil</th>
			</tr>`

	for _, kitap := range kitaplar {
		html += fmt.Sprintf("<tr><td>%d</td><td>%s</td><td>%s</td><td>%d</td></tr>", kitap.ID, kitap.Baslik, kitap.Yazar, kitap.Yil)
	}

	html += `</table></body></html>`

	fmt.Fprintf(w, html)
}

func main() {
	http.HandleFunc("/", anasayfa)
	http.HandleFunc("/kitaplar", tumKitaplariGetir)

	fmt.Println("Sunucu http://localhost:8080 adresinde calisiyor...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}