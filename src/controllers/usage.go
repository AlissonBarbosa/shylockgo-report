package controllers

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/AlissonBarbosa/shylockgo-report/src/models"

	md "github.com/AlissonBarbosa/shylockgo-scraper/src/models"
)

func getLatestEpoch() (int64, error) {
  var maxEpoch int64
  if err := md.DB.Table("server_usages").Select("MAX(timestamp)").Row().Scan(&maxEpoch); err != nil {
    return 0, err
  }
  return maxEpoch, nil
}


func GetServersUsage() ([]md.ServerUsage, error){
  var serverUsage []md.ServerUsage
  query := md.DB

  //query = query.Where("timestamp >= ?", "1716489002")
  //query = query.Where("timestamp <= ?", "1716489596")


  if err := query.Find(&serverUsage).Error; err != nil {
    slog.Error("Error quering database")
    return nil, err
  }
  return serverUsage, nil
}

func SumServerVcpuUsage(serverUsage []md.ServerUsage) []models.ServerVcpuUtil{
  vcpuSum := make(map[string]float64)
  count := make(map[string]int)
  var serverVcpuUtil []models.ServerVcpuUtil

  for _, usage := range serverUsage{
    vcpu, err := strconv.ParseFloat(usage.VcpuUsage, 64)
    if err != nil {
      slog.Error("Error parsing VcpuUsage", err)
      continue
    }
    vcpuSum[usage.ServerID] += vcpu
    count[usage.ServerID]++
  }

  for serverID, sum := range vcpuSum {
    avg := sum / float64(count[serverID])
    serverVcpuUtil = append(serverVcpuUtil, models.ServerVcpuUtil{
      ServerID: serverID,
      VcpuUtilSum: fmt.Sprintf("%.2f", avg),
    })
  }

  return serverVcpuUtil
}
