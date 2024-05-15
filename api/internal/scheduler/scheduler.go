package scheduler

import (
	"log"
	"newsletter/internal/handlers"
	"time"
)

func StartScheduler() {
	for {
		config, err := handlers.ReadConfig()

		if err != nil {
			log.Printf("Erro ao ler a configuração: %v", err)
			time.Sleep(1 * time.Hour)
			continue
		}

		frequency, err := time.ParseDuration(config.Frequency)

		if err != nil {
			log.Printf("Formato de frequência inválido: %v", err)
			time.Sleep(1 * time.Hour)
			continue
		}

		tricker := time.NewTicker(frequency)
		defer tricker.Stop()

		for {
			select {
			case <-tricker.C:
				log.Printf("Enviando newsletter")
				handlers.SendNewsletter(nil, nil)
			}
		}
	}
}
