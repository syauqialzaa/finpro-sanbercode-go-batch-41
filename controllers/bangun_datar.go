package controllers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LuasSegitigaSamaSisi(ctx *gin.Context) {
	hitung := ctx.Query("hitung")
	alas, _ := strconv.Atoi(ctx.Query("alas"))
	tinggi, _ := strconv.Atoi(ctx.Query("tinggi"))

	if hitung == "luas" {
		luas := (alas * tinggi) / 2
		ctx.JSON(http.StatusOK, gin.H {
			"luas_segitiga_sama_sisi": luas,
		})
	}
}

func KelilingSegitigaSamaSisi(ctx *gin.Context) {
	hitung := ctx.Query("hitung")
	alas, _ := strconv.Atoi(ctx.Query("alas"))

	if hitung == "keliling" {
		keliling := alas * 3 
		ctx.JSON(http.StatusOK, gin.H {
			"keliling_segitiga_sama_sisi": keliling,
		})
	}
}

func LuasPersegi(ctx *gin.Context) {
	hitung := ctx.Query("hitung")
	sisi, _ := strconv.Atoi(ctx.Query("sisi"))

	if hitung == "luas" {
		luas := sisi * sisi
		ctx.JSON(http.StatusOK, gin.H {
			"luas_persegi": luas,
		})
	}
}

func KelilingPersegi(ctx *gin.Context) {
	hitung := ctx.Query("hitung")
	sisi, _ := strconv.Atoi(ctx.Query("sisi"))

	if hitung == "luas" {
		keliling := sisi + sisi + sisi + sisi
		ctx.JSON(http.StatusOK, gin.H {
			"keliling_persegi": keliling,
		})
	}
}

func LuasPersegiPanjang(ctx *gin.Context) {
	hitung := ctx.Query("hitung")
	panjang, _ := strconv.Atoi(ctx.Query("panjang"))
	lebar, _ := strconv.Atoi(ctx.Query("lebar"))

	if hitung == "luas" {
		luas := panjang * lebar
		ctx.JSON(http.StatusOK, gin.H {
			"luas_persegi_panjang": luas,
		})
	}
}

func KelilingPersegiPanjang(ctx *gin.Context) {
	hitung := ctx.Query("hitung")
	panjang, _ := strconv.Atoi(ctx.Query("panjang"))
	lebar, _ := strconv.Atoi(ctx.Query("lebar"))

	if hitung == "luas" {
		keliling := 2 * (panjang + lebar)
		ctx.JSON(http.StatusOK, gin.H {
			"keliling_persegi_panjang": keliling,
		})
	}
}

func LuasLingkaran(ctx *gin.Context) {
	hitung := ctx.Query("hitung")
	pi := math.Pi
	r, _ := strconv.ParseFloat(ctx.Query("tinggi"), 64)

	if hitung == "luas" {
		luas := pi * math.Pow(r, 2)
		ctx.JSON(http.StatusOK, gin.H {
			"luas_lingkaran": luas,
		})
	}
}

func KelilingLingkaran(ctx *gin.Context) {
	hitung := ctx.Query("hitung")
	pi := math.Pi
	r, _ := strconv.ParseFloat(ctx.Query("tinggi"), 64)

	if hitung == "luas" {
		keliling := 2 * pi * r
		ctx.JSON(http.StatusOK, gin.H {
			"keliling_lingkaran": keliling,
		})
	}
}