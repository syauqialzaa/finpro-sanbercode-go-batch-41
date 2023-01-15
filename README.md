# API Endpoints Quiz Week-3

### Bangun Datar
```go
router.GET("/bangun-datar/luas-segitiga", controllers.LuasSegitigaSamaSisi)
router.GET("/bangun-datar/keliling-segitiga", controllers.KelilingSegitigaSamaSisi)
router.GET("/bangun-datar/luas-persegi", controllers.LuasPersegi)
router.GET("/bangun-datar/keliling-persegi", controllers.KelilingPersegi)
router.GET("/bangun-datar/luas-persegi-panjang", controllers.LuasPersegiPanjang)
router.GET("/bangun-datar/keliling-persegi-panjang", controllers.KelilingPersegiPanjang)
router.GET("/bangun-datar/luas-lingkaran", controllers.LuasLingkaran)
router.GET("/bangun-datar/keliling-lingkaran", controllers.KelilingLingkaran)
```

### Book Category
```go
router.GET("/categories", controllers.GetAllCategory)
router.POST("/categories", middleware.BasicAuth, controllers.AddCategory)
router.PUT("/categories/:id", middleware.BasicAuth, controllers.UpdateCategory)
router.DELETE("/categories/:id", middleware.BasicAuth, controllers.DeleteCategory)
router.GET("/categories/:id/books", controllers.GetBooksByCategoryID)
```

### Book
```go
router.GET("/books", controllers.GetAllBook)
router.POST("/books", middleware.BasicAuth, controllers.AddBook)
router.PUT("books/:id", middleware.BasicAuth, controllers.UpdateBook)
router.DELETE("books/:id", middleware.BasicAuth, controllers.DeleteBook)
```

**Note:** Endpoints yang diawali dengan `/categories` dan `/books` dengan method `POST`, `PUT`, `DELETE` memerlukan autentikasi dengan `BasicAuth` untuk dapat digunakan. (check: `/middleware/basic_auth.go`)

## Cara Penggunaan
Untuk menggunakan endpoint, dapat mengirimkan request HTTP dengan method yang sesuai ke endpoint yang ditentukan. Untuk query parameters, dapat diisi sesuai dengan kebutuhan dan validasi yang ditentukan.

Contoh penggunaan endpoint `/bangun-datar/luas-segitiga`:
```
GET http://localhost:8080/bangun-datar/luas-segitiga?hitung=luas&alas=7&tinggi=10
```

**Note:** Perhatikan request parameter yang dibutuhkan endpoint.

Untuk melihat format json secara lebih jelas (dalam menginputkan data dalam bentuk json di postman) check file structs: `/structs/`

## Validasi
Untuk endpoint `/books`, memiliki beberapa validasi ketika melakukan method `POST` dan `PUT` seperti:
- `image_url` harus berupa url `http://` atau `https://`
- `release_year` antara tahun 1980 sampai 2021

## Deployment
Quiz Week-3 (quiz-3 branch) sudah di deploy di railway\
Check `quiz-3` branch untuk melihat project Quiz Week-3\
Deployment => ()