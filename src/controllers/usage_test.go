package controllers_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/AlissonBarbosa/shylockgo-report/src/controllers"
	"github.com/AlissonBarbosa/shylockgo-report/src/models"
	md "github.com/AlissonBarbosa/shylockgo-scraper/src/models"
)

func TestCalculateVcpuUtil(t *testing.T) {
  serverUsages := []md.ServerUsage{
    {ID: 1, Timestamp: 1716914203, ServerID: "server1", VcpuUsage: "5.00", RamUsage: "20.00", Domain: "instance-10001", HostID: "host1"},
    {ID: 2, Timestamp: 1716914203, ServerID: "server2", VcpuUsage: "7.00", RamUsage: "30.00", Domain: "instance-10002", HostID: "host1"},
    {ID: 3, Timestamp: 1716914203, ServerID: "server3", VcpuUsage: "15.00", RamUsage: "25.00", Domain: "instance-10003", HostID: "host1"},
    {ID: 4, Timestamp: 1716914203, ServerID: "server4", VcpuUsage: "35.00", RamUsage: "22.00", Domain: "instance-10004", HostID: "host1"},
    {ID: 5, Timestamp: 1716914503, ServerID: "server1", VcpuUsage: "15.00", RamUsage: "25.00", Domain: "instance-10001", HostID: "host1"},
    {ID: 6, Timestamp: 1716914503, ServerID: "server2", VcpuUsage: "7.00", RamUsage: "30.00", Domain: "instance-10002", HostID: "host1"},
    {ID: 7, Timestamp: 1716914503, ServerID: "server3", VcpuUsage: "5.00", RamUsage: "20.00", Domain: "instance-10003", HostID: "host1"},
    {ID: 8, Timestamp: 1716914503, ServerID: "server4", VcpuUsage: "None", RamUsage: "20.00", Domain: "instance-10004", HostID: "host1"},
  }

  expected := []models.ServerVcpuUtil{
    {ServerID: "server1", VcpuUtilSum: "10.00"},
    {ServerID: "server2", VcpuUtilSum: "7.00"},
    {ServerID: "server3", VcpuUtilSum: "10.00"},
    {ServerID: "server4", VcpuUtilSum: "35.00"},
  }

  result := controllers.SumServerVcpuUsage(serverUsages)

  sort.Slice(expected, func(i, j int) bool{
    return expected[i].ServerID < expected[j].ServerID
  })

  sort.Slice(result, func(i, j int) bool {
    return result[i].ServerID < result[j].ServerID
  })

  if !reflect.DeepEqual(result, expected) {
    t.Errorf("Expected %v, but got %v", expected, result)
  }
}
