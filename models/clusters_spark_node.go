/*
 * Databricks
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package models

type ClustersSparkNode struct {
	PrivateIp string `json:"private_ip,omitempty"`

	PublicDns string `json:"public_dns,omitempty"`

	NodeId string `json:"node_id,omitempty"`

	InstanceId string `json:"instance_id,omitempty"`

	StartTimestamp int64 `json:"start_timestamp,omitempty"`

	HostPrivateIp string `json:"host_private_ip,omitempty"`
}
