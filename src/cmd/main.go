package main

import (
	"flag"
	//"fmt"
	"log/slog"
	"os"

	"github.com/AlissonBarbosa/shylockgo-report/src/controllers"
	"github.com/AlissonBarbosa/shylockgo-scraper/src/models"
)

func main() {
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(l)

  models.ConnectDatabase()

  serverUsageFlag := flag.Bool("usage-servers", false, "Create server usage report")
  flag.Parse()

  if *serverUsageFlag {
    slog.Info("Creating server usage report")
    serverUsage, err := controllers.GetServersUsage()
    if err!= nil {
      slog.Error("Error getting server usage", err)
    }
    controllers.SumServerVcpuUsage(serverUsage)
  }
}
