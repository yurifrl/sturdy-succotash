package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/yurifrl/sturdy-succotash/pkg/config"
)

type Worker struct {
	logger      *log.Logger
	workerCount int
}

func NewWorker(cfg config.Config) (*Worker, error) {
	log.SetOutput(os.Stdout)

	return &Worker{
		logger: log.New(),
	}, nil
}

func (w *Worker) Run(ctx context.Context) {
	w.logger.Info("Starting worker...")

	// Create a ticker to simulate periodic work
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			w.logger.Info("Shutting down worker...")
			return
		case <-ticker.C:
			// Simulate pulling data from a fake endpoint
			err := w.pullDataFromFakeEndpoint()
			if err != nil {
				w.logger.Error("Error pulling data: ", err)
			}
		}
	}
}

func (w *Worker) pullDataFromFakeEndpoint() error {
	url := "https://jsonplaceholder.typicode.com/todos/1" // Example fake endpoint

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	// Check for a specific value in the response (e.g., "completed": true)
	if completed, ok := result["completed"].(bool); ok && completed {
		w.logger.Info("Task is completed")
	} else {
		w.logger.Info("Task is not completed yet")
	}

	return nil
}

func (w *Worker) CheckReadiness(ctx context.Context) error {

	// TODO

	return nil
}
