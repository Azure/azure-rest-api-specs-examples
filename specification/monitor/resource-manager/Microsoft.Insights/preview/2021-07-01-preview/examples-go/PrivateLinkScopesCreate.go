package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/PrivateLinkScopesCreate.json
func ExamplePrivateLinkScopesClient_CreateOrUpdate_privateLinkScopeCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewPrivateLinkScopesClient("86dc51d3-92ed-4d7e-947a-775ea79b4919", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "my-resource-group", "my-privatelinkscope", armmonitor.AzureMonitorPrivateLinkScope{
		Location: to.Ptr("Global"),
		Properties: &armmonitor.AzureMonitorPrivateLinkScopeProperties{
			AccessModeSettings: &armmonitor.AccessModeSettings{
				Exclusions:          []*armmonitor.AccessModeSettingsExclusion{},
				IngestionAccessMode: to.Ptr(armmonitor.AccessModeOpen),
				QueryAccessMode:     to.Ptr(armmonitor.AccessModeOpen),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
