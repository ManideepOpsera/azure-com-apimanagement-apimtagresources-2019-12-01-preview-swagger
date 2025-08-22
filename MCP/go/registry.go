package main

import (
	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	tools_tagresource "github.com/apimanagementclient/mcp-server/tools/tagresource"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_tagresource.CreateTagresource_listbyserviceTool(cfg),
	}
}
