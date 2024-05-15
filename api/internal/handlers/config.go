package handlers

import (
	"encoding/json"
	"net/http"
	"newsletter/internal/models"
	"os"

	"gopkg.in/yaml.v3"
)

var configFilePath = "config/config.yaml"

// GetConfig godoc
// @Summary Get configuration
// @Description Get the current configuration for the newsletter
// @Tags config
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Config
// @Router /api/config [get]
func GetConfig(w http.ResponseWriter, r *http.Request) {

	data, err := os.ReadFile(configFilePath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var config models.Config
	err = yaml.Unmarshal(data, &config)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(config)
}

// UpdateConfig godoc
// @Summary Update configuration
// @Description Update the current configuration for the newsletter
// @Tags config
// @Accept  json
// @Produce  json
// @Param config body models.Config true "Configuration to update"
// @Success 200
// @Router /api/config [post]
func UpdateConfig(w http.ResponseWriter, r *http.Request) {
	var config models.Config
	err := json.NewDecoder(r.Body).Decode(&config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = os.WriteFile(configFilePath, data, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
