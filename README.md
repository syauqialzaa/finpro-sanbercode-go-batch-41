<h1 align="center"><b>Basic Project Management System</b></h1>

Backend dari Sistem Manajemen Project. Project ini dibuat sebagai Final Project Sanbercode Golang Bootcamp Backend Development Batch 41.

## Daftar Isi
- [Daftar Isi](#daftar-isi)
- [Persyaratan Final Project](#persyaratan-final-project)
- [Fitur](#fitur)
- [Tech Stack](#tech-stack)
- [Libraries](#libraries)
- [Struktur Project](#struktur-project)
- [Endpoints](#endpoints)
- [Cara Penggunaan](#cara-penggunaan)
- [Deployment](#deployment)

## Persyaratan Final Project
- [x] Wajib deploy (tidak terdeploy maka nilai akan langsung **dikurangi setengahnya**)
- [x] Database yang digunakan harus berelasi, tidak boleh NoSQL
- [x] Wajib menggunakan relasi database (melampirkan skema database di PPT)
- [x] Terdapat register, login dan wajib ada Basic Auth (nilai plus jika menggunakan JWT) pada method yang dibutuhkan
- [x] Wajib ada interaksi user dengan tema terkait, misal pada review game user dapat memberikan rating dan menulis review
- [x] Harus ada keterkaitan antara tabel satu sama lain dan **minimal ada 4 tabel**
- [x] Untuk data master tanpa relasi wajib minimal ada CRUD
- [x] Mengganti file readme.md pada github dengan penjelasan API dan penggunaan aplikasinya
- [x] Wajib membuat example request dan response API di postman

## Fitur
Berikut merupakan fitur-fitur yang diimplementasikan dalam project ini.
- Basic JWT Authentication
- Create a Project
- Show all Project
- Modify a Project
- Destroy a Project
- Create a Task
- Show all Task
- Show Tasks by Project ID
- Modify a Task
- Destroy a Task
- Add a Member
- Show all Member
- Show Members by Team ID
- Modify a Member
- Destroy a Member
- Add a Team
- Show all Team
- Modify a Team
- Destroy a Team

## Tech Stack
Berikut merupakan teknologi-teknologi yang digunakan dalam project ini.
| Technolgy                       | Used                          |
| ------------------------------- | ----------------------------- |
| Programming Language            | Go v1.19                      |
| Framework                       | Gin v1.8.2                    |
| Database                        | PostgreSQL                    |
| Deployment                      | Railway                       |

## Libraries
Berikut merupakan libraries yang digunakan dalam project ini.
| Library Urls                    |
| ------------------------------- |
| github.com/gin-gonic/gin        |
| github.com/lib/pq               |
| github.com/rubenv/sql-migrate   |
| github.com/gobuffalo/packr/v2   |
| github.com/joho/godotenv        |
| github.com/dgrijalva/jwt-go     |

## Struktur Project
```
├── config              => package config, untuk konfigurasi app
│   └── database.go
├── controllers         => package controllers, sebagai handlers endpoint
│   ├── member.go
│   ├── project.go
│   ├── task.go
│   └── team.go
├── database            => package database, untuk set up database dan migrasinya
│   ├── database.go
│   └── sql
│       └── 1_initiate.sql
├── dummy               => dummy data untuk melakukan interaksi data json
│   ├── admin.json
│   ├── member.json
│   ├── project.json
│   ├── task.json
│   └── team.json
├── go.mod
├── go.sum
├── main.go
├── middleware          => package middleware, handlers untuk authentikasi
│   ├── auth_middleware.go
│   ├── login.go
│   └── register.go
├── models              => package models, berisi structs sebagai model dan format json
│   ├── admin.go
│   ├── member.go
│   ├── project.go
│   ├── task.go
│   └── team.go
├── README.md
├── routers
│   └── routers.go
└── services            => package services, untuk manipulasi data memggunakan db query
    ├── member.go
    ├── project.go
    ├── task.go
    └── team.go
```

## Endpoints
Berikut merupakan endpoints dalam project ini.
```go
// AUTHENTICATION
router.POST("/register", middleware.Register)
router.POST("/login", middleware.Login)

// CRUD PROJECTS
router.GET("/projects", controllers.GetAllProject)
router.POST("/projects", middleware.AuthMiddlware(), controllers.CreateProject)
router.PUT("/projects/:id", middleware.AuthMiddlware(), controllers.UpdateProject)
router.DELETE("/projects/:id", middleware.AuthMiddlware(), controllers.DeleteProject)
router.GET("/projects/:id/tasks", controllers.GetTaskByProjectID)    // get task by projectID

// CRUD TASKS
router.GET("/tasks", controllers.GetAllTask)
router.POST("/tasks", middleware.AuthMiddlware(), controllers.CreateTask)
router.PUT("/tasks/:id", middleware.AuthMiddlware(), controllers.UpdateTask)
router.DELETE("/tasks/:id", middleware.AuthMiddlware(), controllers.DeleteTask)

// CRUD TEAM
router.GET("/teams", controllers.GetAllTeam)
router.POST("/teams", middleware.AuthMiddlware(), controllers.CreateTeam)
router.PUT("/teams/:id", middleware.AuthMiddlware(), controllers.UpdateTeam)
router.DELETE("/teams/:id", middleware.AuthMiddlware(), controllers.DeleteTeam)
router.GET("/teams/:id/members", controllers.GetMemberByTeamID)    // get member by teamID

// CRUD MEMBER
router.GET("/members", controllers.GetAllMember)
router.POST("/members", middleware.AuthMiddlware(), controllers.AddMember)
router.PUT("/members/:id", middleware.AuthMiddlware(), controllers.UpdateMember)
router.DELETE("/members/:id", middleware.AuthMiddlware(), controllers.DeleteMember)
```

## Cara Penggunaan
Berikut contoh penggunaan endpoint dengan Postman.
```go
https://finpro-sanbercode-go-batch-41-production.up.railway.app/members/1
```
**Note:**
- Endpoint dengan handlers `middleware.AuthMiddleware()` membutuhkan authentication terlebih dahulu sebelum direquest urlnya
- Authentication dengan register dan login akan men-genereate token yang nantinya digunakan untuk akses ke endpoint yang memiliki requirement authentikasi
- Perhatikan juga untuk setiap endpoint yang membutuhkan parameter id seperti method `PUT`, `DELETE`

## Deployment
https://finpro-sanbercode-go-batch-41-production.up.railway.app/members \
https://finpro-sanbercode-go-batch-41-production.up.railway.app/projects \
https://finpro-sanbercode-go-batch-41-production.up.railway.app/tasks \
https://finpro-sanbercode-go-batch-41-production.up.railway.app/teams