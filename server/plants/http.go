package plants

import (
	"encoding/json"
	"lab3/server/tools"
	"log"
	"net/http"
)

// Plants HTTP handler.
type HttpHandlerFunc http.HandlerFunc

// HttpHandler creates a new instance of channels HTTP handler.
func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListPlants(store, rw)
		} else if r.Method == "POST" {
			handleUpdatePlant(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleUpdatePlant(r *http.Request, rw http.ResponseWriter, store *Store) {
	var p Plant
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Printf("Error decoding plant input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	if p.SoilMoistureLevel > 1 || p.SoilMoistureLevel < 0 {
		log.Printf("Wrong soil moisture level. Must be in interval (0; 1)")
		tools.WriteJsonBadRequest(rw, "Wrong soil moisture level. Must be in interval (0; 1)")
		return
	}
	err := store.UpdatePlant(p.Id, p.SoilMoistureLevel)
	if err == nil {
		tools.WriteJsonOk(rw, "Successfully updated")
	} else if err.Error() == "Record with given id does not exist" {
		tools.WriteJsonBadRequest(rw, err.Error())
	} else {
		log.Printf("Error updating record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

func handleListPlants(store *Store, rw http.ResponseWriter) {
	res, err := store.ListPlants()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}
