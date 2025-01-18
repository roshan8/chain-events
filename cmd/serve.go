/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/roshan8/chain-events/api"
	"github.com/roshan8/chain-events/config"
	db "github.com/roshan8/chain-events/internal/database"
	"github.com/roshan8/chain-events/internal/handler"
	"github.com/roshan8/chain-events/internal/repository"
	"github.com/roshan8/chain-events/internal/service"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "API Server to collect and store events",
	Long: `Provides a set of APIs to collect events from different sources 
and store them in a database. It supports multiple event sources and ensures efficient storage and retrieval..`,
	Run: func(cmd *cobra.Command, args []string) {

		// initialize config
		config.Init()

		// initialize db
		db := db.InitializeDB()

		fmt.Println("Starting api server...")
		r := chi.NewRouter()

		// Middleware
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)

		kubernetesService := service.NewKubernetesService(
			repository.NewKubernetesRepository(db),
		)
		server := handler.NewServer(kubernetesService)

		// Register handlers
		handler := api.HandlerFromMux(server, r)

		log.Printf("Starting server on :8080")
		if err := http.ListenAndServe(":8080", handler); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
