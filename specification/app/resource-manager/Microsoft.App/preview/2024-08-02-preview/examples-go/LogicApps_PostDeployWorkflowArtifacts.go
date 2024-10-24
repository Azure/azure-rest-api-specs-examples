package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/LogicApps_PostDeployWorkflowArtifacts.json
func ExampleLogicAppsClient_DeployWorkflowArtifacts_deploysWorkflowArtifacts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewLogicAppsClient().DeployWorkflowArtifacts(ctx, "testrg123", "testapp2", "testapp2", &armappcontainers.LogicAppsClientDeployWorkflowArtifactsOptions{WorkflowArtifacts: &armappcontainers.WorkflowArtifacts{
		AppSettings: map[string]any{
			"eventHub_connectionString": "Endpoint=sb://example.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=EXAMPLE1a2b3c4d5e6fEXAMPLE=",
		},
		Files: map[string]any{
			"connections.json": map[string]any{
				"managedApiConnections": map[string]any{},
				"serviceProviderConnections": map[string]any{
					"eventHub": map[string]any{
						"displayName": "example1",
						"parameterValues": map[string]any{
							"connectionString": "@appsetting('eventHub_connectionString')",
						},
						"serviceProvider": map[string]any{
							"id": "/serviceProviders/eventHub",
						},
					},
				},
			},
			"test1/workflow.json": map[string]any{
				"definition": map[string]any{
					"$schema":        "https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#",
					"actions":        map[string]any{},
					"contentVersion": "1.0.0.0",
					"outputs":        map[string]any{},
					"triggers": map[string]any{
						"When_events_are_available_in_Event_hub": map[string]any{
							"type": "ServiceProvider",
							"inputs": map[string]any{
								"parameters": map[string]any{
									"eventHubName": "test123",
								},
								"serviceProviderConfiguration": map[string]any{
									"operationId":       "receiveEvents",
									"connectionName":    "eventHub",
									"serviceProviderId": "/serviceProviders/eventHub",
								},
							},
							"splitOn": "@triggerOutputs()?['body']",
						},
					},
				},
				"kind": "Stateful",
			},
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
