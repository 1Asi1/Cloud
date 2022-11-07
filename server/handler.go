package server

import (
	"encoding/json"
	"log"
	"net/http"
	"solution/model"
	"solution/service"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	service service.Service
}

func NewHandler(s service.Service) *handler {
	return &handler{service: s}
}

func (h *handler) InitHandlers(r *chi.Mux) {
	r.Post("/config", h.сreate)
	r.Get("/config", h.read)
	r.Put("/config", h.update)
	r.Delete("/config", h.delete)
	r.Get("/version", h.getServiceVersion)
}

func (h *handler) сreate(w http.ResponseWriter, r *http.Request) {
	conf := model.NewConfig()
	err := json.NewDecoder(r.Body).Decode(conf)
	if err != nil {
		log.Println("query deserialization error", err)
		return
	}

	if conf.Service == "" {
		log.Println("incorrect data")
		return
	}

	for _, data := range conf.Data {
		if len(data) == 0 {
			log.Println("incorrect data")
			return
		}

		for key := range data {
			if data[key] == "" {
				log.Println("incorrect data")
				return
			}
		}
	}

	h.service.CreateConfig(conf)
}

func (h *handler) read(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("service")
	if s == "" {
		log.Println("incorrect data")
		return
	}

	conf := model.NewConfig()
	h.service.ReadConfig(conf, s)

	data, err := json.Marshal(conf.Data[0])
	if err != nil {
		log.Println("serialization error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (h *handler) update(w http.ResponseWriter, r *http.Request) {
	conf := model.NewConfig()
	err := json.NewDecoder(r.Body).Decode(conf)
	if err != nil {
		log.Println("query deserialization error", err)
	}

	h.service.UpdateConfig(conf)
}

func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("service")
	if s == "" {
		log.Println("incorrect data")
		return
	}

	v := r.URL.Query().Get("version")
	if v == "" {
		log.Println("incorrect data")
		return
	}

	ver, err := strconv.ParseFloat(v, 64)
	if err != nil {
		log.Println("not parse", err)
	}

	h.service.DeleteConfig(s, ver)
}

func (h *handler) getServiceVersion(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("service")
	if s == "" {
		log.Println("incorrect data")
		return
	}

	json.NewEncoder(w).Encode(h.service.GetVersion())
}
