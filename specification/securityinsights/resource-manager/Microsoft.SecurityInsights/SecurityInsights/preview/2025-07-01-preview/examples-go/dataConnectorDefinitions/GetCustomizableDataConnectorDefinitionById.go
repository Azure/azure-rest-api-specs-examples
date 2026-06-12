package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/dataConnectorDefinitions/GetCustomizableDataConnectorDefinitionById.json
func ExampleDataConnectorDefinitionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorDefinitionsClient().Get(ctx, "myRg", "myWorkspace", "763f9fa1-c2d3-4fa2-93e9-bccd4899aa12", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorDefinitionsClientGetResponse{
	// 	DataConnectorDefinitionClassification: &armsecurityinsights.CustomizableConnectorDefinition{
	// 		Name: to.Ptr("763f9fa1-c2d3-4fa2-93e9-bccd4899aa12"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectorDefinitions"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectorDefinitions/763f9fa1-c2d3-4fa2-93e9-bccd4899aa12"),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorDefinitionKindCustomizable),
	// 		Properties: &armsecurityinsights.CustomizableConnectorDefinitionProperties{
	// 			ConnectionsConfig: &armsecurityinsights.CustomizableConnectionsConfig{
	// 				TemplateSpecName: to.Ptr("templateNameMock"),
	// 				TemplateSpecVersion: to.Ptr("1.0.0"),
	// 			},
	// 			ConnectorUIConfig: &armsecurityinsights.CustomizableConnectorUIConfig{
	// 				Availability: &armsecurityinsights.ConnectorDefinitionsAvailability{
	// 					IsPreview: to.Ptr(false),
	// 					Status: to.Ptr[int32](1),
	// 				},
	// 				ConnectivityCriteria: []*armsecurityinsights.ConnectivityCriterion{
	// 					{
	// 						Type: to.Ptr("IsConnectedQuery"),
	// 						Value: []*string{
	// 							to.Ptr("GitHubAuditLogPolling_CL \n | summarize LastLogReceived = max(TimeGenerated)\n | project IsConnected = LastLogReceived > ago(30d)"),
	// 						},
	// 					},
	// 				},
	// 				DataTypes: []*armsecurityinsights.ConnectorDataType{
	// 					{
	// 						Name: to.Ptr("GitHubAuditLogPolling_CL"),
	// 						LastDataReceivedQuery: to.Ptr("GitHubAuditLogPolling_CL \n            | summarize Time = max(TimeGenerated)\n            | where isnotempty(Time)"),
	// 					},
	// 				},
	// 				DescriptionMarkdown: to.Ptr("The GitHub audit log connector provides the capability to ingest GitHub logs into Azure Sentinel. By connecting GitHub audit logs into Azure Sentinel, you can view this data in workbooks, use it to create custom alerts, and improve your investigation process."),
	// 				GraphQueries: []*armsecurityinsights.GraphQuery{
	// 					{
	// 						BaseQuery: to.Ptr("GitHubAuditLogPolling_CL"),
	// 						Legend: to.Ptr("GitHub audit log events"),
	// 						MetricName: to.Ptr("Total events received"),
	// 					},
	// 				},
	// 				InstructionSteps: []*armsecurityinsights.InstructionStep{
	// 					{
	// 						Description: to.Ptr("Enable GitHub audit Logs. \n Follow [this](https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token) to create or find your personal key"),
	// 						Instructions: []*armsecurityinsights.InstructionStepDetails{
	// 							{
	// 								Type: to.Ptr("OAuthForm"),
	// 								Parameters: map[string]any{
	// 									"clientIdLabel": "Client ID",
	// 									"clientSecretLabel": "Client Secret",
	// 									"connectButtonLabel": "Connect",
	// 									"disconnectButtonLabel": "Disconnect",
	// 								},
	// 							},
	// 						},
	// 						Title: to.Ptr("Connect GitHub Enterprise Audit Log to Azure Sentinel"),
	// 					},
	// 				},
	// 				Permissions: &armsecurityinsights.ConnectorDefinitionsPermissions{
	// 					Customs: []*armsecurityinsights.CustomPermissionDetails{
	// 						{
	// 							Name: to.Ptr("GitHub API personal token Key"),
	// 							Description: to.Ptr("You need access to GitHub personal token, the key should have 'admin:org' scope"),
	// 						},
	// 					},
	// 					ResourceProvider: []*armsecurityinsights.ConnectorDefinitionsResourceProvider{
	// 						{
	// 							PermissionsDisplayText: to.Ptr("read and write permissions are required."),
	// 							Provider: to.Ptr("Microsoft.OperationalInsights/workspaces"),
	// 							ProviderDisplayName: to.Ptr("Workspace"),
	// 							RequiredPermissions: &armsecurityinsights.ResourceProviderRequiredPermissions{
	// 								Action: to.Ptr(false),
	// 								Delete: to.Ptr(false),
	// 								Read: to.Ptr(false),
	// 								Write: to.Ptr(true),
	// 							},
	// 							Scope: to.Ptr(armsecurityinsights.ProviderPermissionsScopeWorkspace),
	// 						},
	// 					},
	// 				},
	// 				Publisher: to.Ptr("GitHub"),
	// 				Title: to.Ptr("GitHub Enterprise Audit Log"),
	// 			},
	// 		},
	// 	},
	// }
}
