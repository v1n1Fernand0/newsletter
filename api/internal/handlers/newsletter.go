package handlers

import (
	"net/http"
	"newsletter/internal/collector"
	"newsletter/internal/email"
	"newsletter/internal/formatter"
	"newsletter/internal/models"
	"os"

	"gopkg.in/yaml.v3"
)

// SendNewsletter godoc
// @Summary Send newsletter
// @Description Send the current newsletter to the configured email
// @Tags newsletter
// @Success 200
// @Router /api/newsletter [post]
func SendNewsletter(w http.ResponseWriter, r *http.Request) {
	config, err := ReadConfig()
	if err != nil {
		if w != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	content := collector.CollectContent(config.Sites)
	newsletter := formatter.FormatNewsletter(content)
	email.SendEmail(config.Email.SMTPServer, config.Email.SMTPPort, config.Email.Username, config.Email.Password, config.Email.From, config.Email.To, "Daily Newsletter", newsletter)

	if w != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Newsletter sent!"))
	}
}

func ReadConfig() (models.Config, error) {
	var config models.Config

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	return config, err
}
