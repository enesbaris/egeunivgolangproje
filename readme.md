# Go (Golang) Programlama Dili:

![Go Logo](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png)

> **"Go, Google'ın yazılım mühendisliği problemlerine cevabıdır."** — Rob Pike

---

## 📑 İçindekiler

1. Giriş ve Tarihçe  
2. Go’nun Tasarım Felsefesi  
3. Neden Go? (Avantajları)  
4. Kurulum ve Çalışma Ortamı  
5. Temel Sözdizimi (Syntax)  
6. Kontrol Yapıları  
7. Veri Yapıları  
8. Fonksiyonlar ve Metotlar  
9. Pointerlar (İşaretçiler)  
10. Interface (Arayüz) Kavramı  
11. Eşzamanlılık (Concurrency)  
12. Hata Yönetimi (Error Handling)  
13. Paket Yönetimi (Go Modules)  
14. Sık Kullanılan Komutlar (Cheat Sheet)  
15. Go Proverbs (Atasözleri)  

---

## 1. Giriş ve Tarihçe

Go programlama dili, 2007 yılında Google’da **Robert Griesemer**, **Rob Pike** ve **Ken Thompson**
tarafından geliştirilmeye başlanmıştır.  
2009 yılında açık kaynak olarak yayınlanmış ve kısa sürede sistem programlama alanında
popüler hale gelmiştir.

Go’nun çıkış noktası şudur:

> Büyük ekiplerin, büyük sistemleri **basit**, **okunabilir** ve **bakımı kolay**
bir şekilde geliştirebilmesi.

Bugün Go; **Docker**, **Kubernetes**, **Terraform**, **Prometheus** gibi kritik altyapı
yazılımlarının temelini oluşturmaktadır.

---

## 2. Go’nun Tasarım Felsefesi

Go’nun temel yaklaşımı şunlara dayanır:

- Basit ve okunabilir sözdizimi
- Az ama güçlü anahtar kelimeler
- Açık ve tahmin edilebilir davranış
- Karmaşık dil özelliklerinden bilinçli olarak kaçınma

Go, geliştiriciyi “nasıl yazmalıyım?” sorusu yerine
“ne yazmalıyım?” sorusuna odaklamayı hedefler.

---

## 3. Neden Go? (Avantajları)

| Özellik | Açıklama |
|------|---------|
| Performans | Makine koduna derlenir, JVM veya interpreter kullanmaz |
| Hızlı Derleme | Büyük projelerde bile saniyeler içinde build |
| Concurrency | Goroutine ve Channel desteği dilin çekirdeğinde |
| Garbage Collector | Otomatik bellek yönetimi |
| Standart Kütüphane | Web, dosya, ağ işlemleri için yeterlidir |
| Tek Binary | Deploy süreci çok kolaydır |

---

## 4. Kurulum ve Çalışma Ortamı

1. https://go.dev/dl adresinden Go’yu indirin  
2. Kurulumu doğrulamak için:

```bash
go version
```

3. Editör olarak **VS Code** ve **Go eklentisi** önerilir

---

## 5. Temel Sözdizimi (Syntax)

### Değişkenler

```go
var isim string = "Ahmet"
yas := 25
aktif := true
```

Go, tip güvenli bir dildir ve kısa tanımlama (`:=`) yalnızca fonksiyon içinde kullanılabilir.

### Sabitler

```go
const Pi = 3.14
```

### Zero Values

Go’da değişkenler varsayılan değerle başlar:

- int → 0
- string → ""
- bool → false

---

## 6. Kontrol Yapıları

### For Döngüsü

Go’da **while** ve **do-while** yoktur.

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

```go
for sayi < 10 {
    sayi++
}
```

### If / Else

```go
if x := 10; x > 5 {
    fmt.Println("Büyük")
} else {
    fmt.Println("Küçük")
}
```

---

## 7. Veri Yapıları

### Array

```go
var notlar [3]int
notlar[0] = 100
```

### Slice

```go
meyveler := []string{"Elma", "Armut"}
meyveler = append(meyveler, "Muz")
```

### Map

```go
ogrenci := make(map[string]int)
ogrenci["Ali"] = 90
```

### Struct

```go
type Araba struct {
    Marka string
    Model int
}
```

---

## 8. Fonksiyonlar ve Metotlar

```go
func Hesapla(a int, b int) (int, string) {
    return a + b, "Başarılı"
}
```

Go fonksiyonları birden fazla değer döndürebilir.

---

## 9. Pointerlar (İşaretçiler)

```go
i := 10
p := &i
*p = 20
```

Pointer kullanımı performans ve bellek yönetimi açısından önemlidir.

---

## 10. Interface (Arayüz) Kavramı

```go
type Sekil interface {
    Alan() float64
}
```

Go’da interface’ler **örtük (implicit)** olarak uygulanır.

---

## 11. Eşzamanlılık (Concurrency)

### Goroutine

```go
go fonksiyon()
```

### Channel

```go
kanal := make(chan string)
kanal <- "Merhaba"
```

---

## 12. Hata Yönetimi (Error Handling)

```go
file, err := os.Open("dosya.txt")
if err != nil {
    log.Fatal(err)
}
```

Go’da hata bir değerdir, istisna değildir.

---

## 13. Paket Yönetimi (Go Modules)

```bash
go mod init github.com/kullanici/proje
go get github.com/gin-gonic/gin
go mod tidy
```

---

## 14. Sık Kullanılan Komutlar

| Komut | Açıklama |
|------|---------|
| go run | Kodu çalıştırır |
| go build | Binary üretir |
| go test | Testleri çalıştırır |
| go fmt | Kod formatlar |

---

## 15. Go Proverbs

> Don’t communicate by sharing memory; share memory by communicating.  
> Concurrency is not parallelism.  
> Make the zero value useful.

---

### Hazırlayan

Muhammed Enes Barış  
Roni Kılıç  