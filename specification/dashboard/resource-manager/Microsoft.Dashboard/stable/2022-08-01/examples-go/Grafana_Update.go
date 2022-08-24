package armdashboard_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2022-08-01/examples/Grafana_Update.json
func ExampleGrafanaClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdashboard.NewGrafanaClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "myResourceGroup", "myWorkspace", armdashboard.ManagedGrafanaUpdateParameters{
		Properties: &armdashboard.ManagedGrafanaPropertiesUpdateParameters{
			APIKey:                  to.Ptr(armdashboard.APIKeyEnabled),
			DeterministicOutboundIP: to.Ptr(armdashboard.DeterministicOutboundIPEnabled),
			GrafanaIntegrations: &armdashboard.GrafanaIntegrations{
				AzureMonitorWorkspaceIntegrations: []*armdashboard.AzureMonitorWorkspaceIntegration{
					{
						AzureMonitorWorkspaceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.monitor/accounts/myAzureMonitorWorkspace"),
					}},
			},
		},
		Tags: map[string]*string{
			"Environment": to.Ptr("Dev 2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
