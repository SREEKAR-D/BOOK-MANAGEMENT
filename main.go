package main

import (
	"log"
	"net/http"

	"golang2/handler"
	auth "golang2/middleware"
	"golang2/model"
	"golang2/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	port := "8084"

	db := model.SetupDB()

	bookService := service.NewBookService(db)
	userService := service.NewUserService(db)

	bookHandler := handler.NewBookHandler(bookService)
	userHandler := handler.NewUserHandler(userService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/signup", userHandler.SignUp)
	r.Post("/login", userHandler.Login)

	r.Route("/protected", func(r chi.Router) {
		r.Use(auth.JWTAuth)

		r.Get("/books", bookHandler.GetAllBooks)
		r.Get("/books/{id}", bookHandler.GetBookByID)
		r.Post("/books", bookHandler.AddBook)
		r.Put("/books/{id}", bookHandler.UpdateBook)
		r.Delete("/books/{id}", bookHandler.DeleteBook)
	})

	log.Printf("Starting server at port:%v", port)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Welcome to my Website"))
		//log.Println(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":"+port, r))
}
