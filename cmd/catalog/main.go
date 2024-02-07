package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/vitor-porfirio/imersao17/goapi/internal/database"
	"github.com/vitor-porfirio/imersao17/goapi/internal/service"
	"github.com/vitor-porfirio/imersao17/goapi/internal/webserver"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:3711521@tcp(localhost:3306)/imersao17")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)

	WebCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	WebProductHandler := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.Post("/category", WebCategoryHandler.CreateCategory)
	c.Get("/category", WebCategoryHandler.GetCategories)
	c.Get("/category/{id}", WebCategoryHandler.GetCategory)

	c.Post("/product", WebProductHandler.CreateProduct)
	c.Get("/product", WebProductHandler.GetProducts)
	c.Get("/product/{id}", WebProductHandler.GetProduct)
	c.Get("/product/category/{categoryID}", WebProductHandler.GetProductByCategoryID)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)

}
