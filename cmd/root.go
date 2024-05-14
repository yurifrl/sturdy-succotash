package cmd

import (
	"context"
	"net/http"
	"os"

	"github.com/yurifrl/sturdy-succotash/pkg/api"
	"github.com/yurifrl/sturdy-succotash/pkg/config"
	"github.com/yurifrl/sturdy-succotash/pkg/worker"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var log = logrus.New()

var rootCmd = &cobra.Command{
	Use:   "lb-logs",
	Short: "Load Balancer Logs Processor",
	Long:  `The Load Balancer Logs Processor is a tool for processing and analyzing logs from load balancers such as Cloudflare.`,
}

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "TODO",
	Long:  `TODO.`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfg config.Config = &config.EnvConfig{}  // Create an instance of EnvConfig
		workerInstance, err := worker.NewWorker(cfg) // Create a new worker instance
		if err != nil {
			log.Fatalf("Failed to create worker: %v", err)
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go workerInstance.Run(ctx)

		apiInstance, err := api.NewApi(cfg) // Create a new API instance
		if err != nil {
			log.Fatalf("Failed to create API: %v", err)
		}
		apiInstance.RegisterEndpoints() // Register API endpoints

		http.Handle("/metrics", promhttp.Handler())
		http.HandleFunc("/liveness", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		http.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) {
			if err := workerInstance.CheckReadiness(ctx); err != nil {
				log.Printf("Readiness check failed: %v", err) // Changed from log.Errorf to log.Printf for standard library log package
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
		})

		log.Printf("Starting server on :8080") // Changed from log.Info to log.Printf for standard library log package
		if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(appCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
