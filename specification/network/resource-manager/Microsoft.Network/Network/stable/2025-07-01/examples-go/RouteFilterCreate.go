package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/RouteFilterCreate.json
func ExampleRouteFiltersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRouteFiltersClient().BeginCreateOrUpdate(ctx, "rg1", "filterName", armnetwork.RouteFilter{
		Location: to.Ptr("West US"),
		Properties: &armnetwork.RouteFilterPropertiesFormat{
			Rules: []*armnetwork.RouteFilterRule{
				{
					Name: to.Ptr("ruleName"),
					Properties: &armnetwork.RouteFilterRulePropertiesFormat{
						Access: to.Ptr(armnetwork.AccessAllow),
						Communities: []*string{
							to.Ptr("12076:5030"),
							to.Ptr("12076:5040"),
						},
						RouteFilterRuleType: to.Ptr(armnetwork.RouteFilterRuleTypeCommunity),
					},
				},
			},
		},
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.RouteFiltersClientCreateOrUpdateResponse{
	// 	RouteFilter: armnetwork.RouteFilter{
	// 		Name: to.Ptr("filterName"),
	// 		Type: to.Ptr("Microsoft.Network/routeFilters"),
	// 		Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/routeFilters/filterName"),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armnetwork.RouteFilterPropertiesFormat{
	// 			Peerings: []*armnetwork.ExpressRouteCircuitPeering{
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			Rules: []*armnetwork.RouteFilterRule{
	// 				{
	// 					Name: to.Ptr("ruleName"),
	// 					Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/routeFilters/filterName/routeFilterRules/ruleName"),
	// 					Properties: &armnetwork.RouteFilterRulePropertiesFormat{
	// 						Access: to.Ptr(armnetwork.AccessAllow),
	// 						Communities: []*string{
	// 							to.Ptr("12076:5030"),
	// 							to.Ptr("12076:5040"),
	// 						},
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						RouteFilterRuleType: to.Ptr(armnetwork.RouteFilterRuleTypeCommunity),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 		},
	// 	},
	// }
}
