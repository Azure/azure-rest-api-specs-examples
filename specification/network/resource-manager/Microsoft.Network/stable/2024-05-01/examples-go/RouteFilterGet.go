package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ab04533261eff228f28e08900445d0edef3eb70c/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/RouteFilterGet.json
func ExampleRouteFiltersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRouteFiltersClient().Get(ctx, "rg1", "filterName", &armnetwork.RouteFiltersClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RouteFilter = armnetwork.RouteFilter{
	// 	Name: to.Ptr("filterName"),
	// 	Type: to.Ptr("Microsoft.Network/routeFilters"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/routeFilters/filterName"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.RouteFilterPropertiesFormat{
	// 		Peerings: []*armnetwork.ExpressRouteCircuitPeering{
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		Rules: []*armnetwork.RouteFilterRule{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/routeFilters/filterName/routeFilterRules/ruleName"),
	// 				Name: to.Ptr("ruleName"),
	// 				Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 				Properties: &armnetwork.RouteFilterRulePropertiesFormat{
	// 					Access: to.Ptr(armnetwork.AccessAllow),
	// 					Communities: []*string{
	// 						to.Ptr("12076:5030"),
	// 						to.Ptr("12076:5040")},
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						RouteFilterRuleType: to.Ptr(armnetwork.RouteFilterRuleTypeCommunity),
	// 					},
	// 			}},
	// 		},
	// 	}
}
