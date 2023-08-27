# Tempo Media Web App - Golang 

tempo-news adalah aplikasi web sederhana yang menggunakan bahasa [Go](https://go.dev/),  dibuat dengan web framework [Gin](https://github.com/gin-gonic/gin).

**tempo-news menggunakan:**

- [Gin Web Framework](https://gin-gonic.com/).
- [GORM](https://gorm.io/)
- [Load .env](https://github.com/joho/godotenv)
- [Go Cryptography](https://pkg.go.dev/golang.org/x/crypto#section-readme)
- [JWT](https://github.com/golang-jwt/jwt)
- [Postgres Database](https://www.postgresql.org/)
- [Validator](https://pkg.go.dev/github.com/go-playground/validator/v10Z)

## Getting started

### Running Postgres

jalankan postgres dengan docker compose

```
docker-compose up
```

### Running tempo-news

jalankan aplikasi web server

```
go run main.go
```

### Middleware Auth
Middleware untuk otentikasi saya pasang di handler untuk setiap api yang memerlukan akses.
User yang ingin mengakses halaman-halaman yang memerlukan otentikasi, harus registrasi terlebih dahulu
```
curl --request POST \
  --url http://localhost:8000/api/v1/signup \
  --header 'Content-Type: application/json' \
  --data '{
	"email": "abdu@tempo.co.id",
	"password": "user_tempo"
}'
```
Setelah registrasi berhasil, user akan mendapatkan kunci dengan login
```
curl --request POST \
  --url http://localhost:8000/api/v1/login \
  --header 'Content-Type: application/json' \
  --data '{
	"email": "abdu@tempo.co.id",
	"password": "user_tempo"
}'
```
Di system ini, setelah berhasil login, system akan membuat cookies yang digunakan sebagai kunci yang akan expired dalam 24 jam, berikut adalah contoh cookies nya
```
cookie Authorization=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMyNDA2NTEsInN1YiI6ImFiZHVAdGVtcG8uY28uaWQifQ.UawBvevIDCOJA-EW6oClff4LMTEVo17X1PgW5q01dyw
```
Selama cookies tersebut masih ada, user dapat mengakses halaman berita dan detail berita
```
curl --request GET \
  --url http://localhost:8000/api/v1/articles \
  --cookie Authorization=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMyNDA2NTEsInN1YiI6ImFiZHVAdGVtcG8uY28uaWQifQ.UawBvevIDCOJA-EW6oClff4LMTEVo17X1PgW5q01dyw
```
```
curl --request GET \
  --url http://localhost:8000/api/v1/article/1 \
  --cookie Authorization=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMyNDA2NTEsInN1YiI6ImFiZHVAdGVtcG8uY28uaWQifQ.UawBvevIDCOJA-EW6oClff4LMTEVo17X1PgW5q01dyw
```

### Mengisi artikel Berita
Untuk mengisi artikel berita, user harus punya akses terlebih dahulu.  Setelah itu user dapat membuat artikel dengan API
```
curl --request POST \
  --url http://localhost:8000/api/v1/article \
  --header 'Content-Type: application/json' \
  --cookie Authorization=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMyNDA2NTEsInN1YiI6ImFiZHVAdGVtcG8uY28uaWQifQ.UawBvevIDCOJA-EW6oClff4LMTEVo17X1PgW5q01dyw \
  --data '{
	"title": "Sempat Disiapkan Tak Bermain sebagai Kiper di Final Piala AFF U-23 2023, Daffa Fasya: Deg-degan Parah",
	"head_line": "Penjaga gawang timnas U-23 Indonesia Daffa Fasya Sumawijaya sempat disiapkan pelatih Shin Tae-yong untuk bermain di posisi nonkiper di final Piala AFF U-23 2023. Ia terlihat duduk di bangku cadangan mengenakan seragam pemain yang berwarna merah.",
	"content": "TEMPO.CO, Jakarta - Penjaga gawang timnas U-23 Indonesia Daffa Fasya Sumawijaya sempat disiapkan pelatih Shin Tae-yong untuk bermain di posisi nonkiper di final Piala AFF U-23 2023. Ia terlihat duduk di bangku cadangan mengenakan seragam pemain yang berwarna merah. Hal itu dilakukan untuk mengantisipasi terbatasnya jumlah pemain Indonesia dalam pertandingan melawan Vietnam. Shin Tae-yong hanya bisa mendaftarkan 18 nama, 11 pemain utama, tujuh cadangan dan dua di antaranya kiper yang salah satunya adalah Daffa."
}'
```

### Sample Response List Berita
```
{
	"data": [
		{
			"id": 1,
			"title": "Honda Berikan Klarifikasi",
			"head_line": "Setelah puluhan rangka patah, akhirnya honda buka suara",
			"content": "Puluhan motor honda mengalami kerusakan yang sama, yaitu rangka patah. Akhirnya honda buka suara menjelaskan kuning dalam rangka. Ternyata klarifikasinya tersebut malah menjelaskan kelemahan honda itu sendiri.",
			"posted_at": "00:00:00",
			"created_at": "00:11:33.000902743",
			"updated_at": "00:11:33.000902743",
			"deleted_at": null
		},
		{
			"id": 2,
			"title": "Honda Melakukan Kebodohan",
			"head_line": "Mengumbar kurang baiknya dari sisi QA",
			"content": "Video klarifikasi menjadi senjata makan tuan",
			"posted_at": "00:00:00",
			"created_at": "00:11:33.000974692",
			"updated_at": "00:11:33.000974692",
			"deleted_at": null
		},
		{
			"id": 3,
			"title": "Tempo Media Terbaik 2023",
			"head_line": "Apakah anda tau siapa media terbaik 2023",
			"content": "Kini tempo semakin didepan",
			"posted_at": "23:52:30.000086242",
			"created_at": "23:52:30.000862457",
			"updated_at": "23:52:30.000862457",
			"deleted_at": null
		},
		{
			"id": 4,
			"title": "Tambahan Title ke 4",
			"head_line": "Headline ini tidak sepanjang content",
			"content": "Konten ya konten. Mau macam begini juga bisa dibilang content. Mau baik mau tidak, konten tetap konten.",
			"posted_at": "23:54:39.000509953",
			"created_at": "23:54:39.000510221",
			"updated_at": "23:54:39.000510221",
			"deleted_at": null
		},
		{
			"id": 5,
			"title": "Tambahan Title ke 5",
			"head_line": "Headline ini tidak sepanjang content",
			"content": "Konten ya konten. Mau macam begini juga bisa dibilang content. Mau baik mau tidak, konten tetap konten.",
			"posted_at": "23:55:24.000472816",
			"created_at": "23:55:24.000473106",
			"updated_at": "23:55:24.000473106",
			"deleted_at": null
		},
		{
			"id": 6,
			"title": "Tambahan Title ke 6",
			"head_line": "Headline ini tidak sepanjang content",
			"content": "Konten ya konten. Mau macam begini juga bisa dibilang content. Mau baik mau tidak, konten tetap konten.",
			"posted_at": "23:57:31.000263753",
			"created_at": "23:57:31.000264027",
			"updated_at": "23:57:31.000264027",
			"deleted_at": null
		},
		{
			"id": 7,
			"title": "Tambahan Title ke 7",
			"head_line": "Headline ini tidak sepanjang content",
			"content": "Konten ya konten. Mau macam begini juga bisa dibilang content. Mau baik mau tidak, konten tetap konten.",
			"posted_at": "00:01:10.000018626",
			"created_at": "00:01:10.000018626",
			"updated_at": "00:01:10.000018626",
			"deleted_at": null
		},
		{
			"id": 8,
			"title": "Tambahan Title ke 8",
			"head_line": "Headline ini tidak sepanjang content",
			"content": "Konten ya konten. Mau macam begini juga bisa dibilang content. Mau baik mau tidak, konten tetap konten.",
			"posted_at": "00:11:34.000156486",
			"created_at": "00:11:34.000156486",
			"updated_at": "00:11:34.000156486",
			"deleted_at": null
		},
		{
			"id": 9,
			"title": "Tambahan Title ke 9",
			"head_line": "Headline ini tidak sepanjang content",
			"content": "Konten ya konten. Mau macam begini juga bisa dibilang content. Mau baik mau tidak, konten tetap konten.",
			"posted_at": "00:11:34.000156539",
			"created_at": "00:11:34.000156539",
			"updated_at": "00:11:34.000156539",
			"deleted_at": null
		},
		{
			"id": 10,
			"title": "Sempat Disiapkan Tak Bermain sebagai Kiper di Final Piala AFF U-23 2023, Daffa Fasya: Deg-degan Parah",
			"head_line": "Penjaga gawang timnas U-23 Indonesia Daffa Fasya Sumawijaya sempat disiapkan pelatih Shin Tae-yong untuk bermain di posisi nonkiper di final Piala AFF U-23 2023. Ia terlihat duduk di bangku cadangan mengenakan seragam pemain yang berwarna merah.",
			"content": "TEMPO.CO, Jakarta - Penjaga gawang timnas U-23 Indonesia Daffa Fasya Sumawijaya sempat disiapkan pelatih Shin Tae-yong untuk bermain di posisi nonkiper di final Piala AFF U-23 2023. Ia terlihat duduk di bangku cadangan mengenakan seragam pemain yang berwarna merah. Hal itu dilakukan untuk mengantisipasi terbatasnya jumlah pemain Indonesia dalam pertandingan melawan Vietnam. Shin Tae-yong hanya bisa mendaftarkan 18 nama, 11 pemain utama, tujuh cadangan dan dua di antaranya kiper yang salah satunya adalah Daffa.",
			"posted_at": "00:11:34.000157768",
			"created_at": "00:11:34.000157768",
			"updated_at": "00:11:34.000157768",
			"deleted_at": null
		}
	],
	"message": "Success.",
	"success": true
}
```

### Sample response detail berita
```
{
	"data": {
		"id": 10,
		"title": "Sempat Disiapkan Tak Bermain sebagai Kiper di Final Piala AFF U-23 2023, Daffa Fasya: Deg-degan Parah",
		"head_line": "Penjaga gawang timnas U-23 Indonesia Daffa Fasya Sumawijaya sempat disiapkan pelatih Shin Tae-yong untuk bermain di posisi nonkiper di final Piala AFF U-23 2023. Ia terlihat duduk di bangku cadangan mengenakan seragam pemain yang berwarna merah.",
		"content": "TEMPO.CO, Jakarta - Penjaga gawang timnas U-23 Indonesia Daffa Fasya Sumawijaya sempat disiapkan pelatih Shin Tae-yong untuk bermain di posisi nonkiper di final Piala AFF U-23 2023. Ia terlihat duduk di bangku cadangan mengenakan seragam pemain yang berwarna merah. Hal itu dilakukan untuk mengantisipasi terbatasnya jumlah pemain Indonesia dalam pertandingan melawan Vietnam. Shin Tae-yong hanya bisa mendaftarkan 18 nama, 11 pemain utama, tujuh cadangan dan dua di antaranya kiper yang salah satunya adalah Daffa.",
		"posted_at": "00:11:34.000157768",
		"created_at": "00:11:34.000157768",
		"updated_at": "00:11:34.000157768",
		"deleted_at": null
	},
	"message": "Success.",
	"success": true
}
```
