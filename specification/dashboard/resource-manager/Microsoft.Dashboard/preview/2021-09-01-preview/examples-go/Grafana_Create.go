package armdashboard_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dashboard/resource-manager/Microsoft.Dashboard/preview/2021-09-01-preview/examples/Grafana_Create.json
func ExampleGrafanaClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdashboard.NewGrafanaClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"myResourceGroup",
		"myWorkspace",
		armdashboard.ManagedGrafana{
			Identity: &armdashboard.ManagedIdentity{
				Type: to.Ptr(armdashboard.IdentityTypeSystemAssigned),
			},
			Location: to.Ptr("West US"),
			Properties: &armdashboard.ManagedGrafanaProperties{
				ProvisioningState: to.Ptr(armdashboard.ProvisioningStateAccepted),
				ZoneRedundancy:    to.Ptr(armdashboard.ZoneRedundancyEnabled),
			},
			SKU: &armdashboard.ResourceSKU{
				Name: to.Ptr("Standard"),
			},
			Tags: map[string]*string{
				"Environment": to.Ptr("Dev"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
