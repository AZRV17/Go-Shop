package app

import (
	"github.com/AZRV17/Go-Shop/internal/config"
	"github.com/AZRV17/Go-Shop/pkg/db/mongo"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func Run() {
	cfg, err := config.New("internal/config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	db, err := mongo.NewClient(cfg.Mongo.User, cfg.Mongo.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = mongo.CloseClient(db); err != nil {
			log.Fatal(err)
		}
	}()

	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, map[string]string{
			"message": "Mongo Connected!",
		})
	})

	log.Println("Listening...")
	err = http.ListenAndServe(cfg.HTTP.Host+":"+cfg.HTTP.Port, r)
	if err != nil {
		log.Fatal(err)
	}
}
