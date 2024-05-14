package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/yurifrl/sturdy-succotash/pkg/config"
)

type Api struct {
	logger *log.Logger
}

func NewApi(cfg config.Config) (*Api, error) {
	log.SetOutput(os.Stdout)

	return &Api{
		logger: log.New(),
	}, nil
}
func (a *Api) RegisterEndpoints() {
	http.HandleFunc("/hello", a.HelloHandler)
	http.HandleFunc("/update", a.updateHandler)
}

func (a *Api) HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, world!"}
	json.NewEncoder(w).Encode(response)
}

func (a *Api) updateHandler(w http.ResponseWriter, r *http.Request) {

	if err := updateStatus(""); err != nil {
		a.logger.Error(err)
	}
	response := map[string]string{"message": "Hello, world!"}
	json.NewEncoder(w).Encode(response)
}

func updateStatus(token string) (err error) {

	fmt.Println("Status updated successfully!")
}
