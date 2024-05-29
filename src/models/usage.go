package models

type ServerVcpuUtil struct {
  ID  uint `json:"id"`
  ServerID string `json:"server_id"`
  VcpuUtilSum string `json:"vcpu_util_sum"`
}
