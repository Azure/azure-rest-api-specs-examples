package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/PrivateLinkScopePrivateLinkResourceGet.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().Get(ctx, "MyResourceGroup", "MyPrivateLinkScope", "azuremonitor", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResource = armmonitor.PrivateLinkResource{
	// 	Name: to.Ptr("azuremonitor"),
	// 	Type: to.Ptr("Microsoft.Insights/privateLinkScopes/privateLinkResources"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/privateLinkScopes/MyPrivateLinkScope/privateLinkResources/azuremonitor"),
	// 	Properties: &armmonitor.PrivateLinkResourceProperties{
	// 		GroupID: to.Ptr("azuremonitor"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("api"),
	// 			to.Ptr("global.in.ai"),
	// 			to.Ptr("profiler"),
	// 			to.Ptr("live"),
	// 			to.Ptr("snapshot"),
	// 			to.Ptr("agentsolutionpackstore"),
	// 			to.Ptr("dce-global")},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.monitor.azure.com"),
	// 				to.Ptr("privatelink.oms.opinsights.azure.com"),
	// 				to.Ptr("privatelink.ods.opinsights.azure.com"),
	// 				to.Ptr("privatelink.agentsvc.azure-automation.net"),
	// 				to.Ptr("privatelink.blob.core.windows.net")},
	// 			},
	// 		}
}
