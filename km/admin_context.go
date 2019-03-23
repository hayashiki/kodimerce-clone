package km

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hayashiki/kodimerce-clone/entities"
)

type AdminContext struct {
	*ServerContext
	User *entities.User
}

type ServerContext struct {
	Context context.Context
}

func (c *AdminContext) CreateCategory(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		c.ServeJson(w, http.StatusBadRequest, "Name cannot be empty")
		return
	}
	category, err := entities.CreateCategory(c.Context, name)
	if err != nil {
		c.ServeJson(w, http.StatusInternalServerError, "Unexpected error creating category.")
		return
	}
	c.ServeJson(w, http.StatusOK, category)
	return
}

func (c *AdminContext) ServeJson(w http.ResponseWriter, status int, value interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	bts, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s", bts)
}
