package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/RouteFilterRuleListByRouteFilter.json
func ExampleRouteFilterRulesClient_NewListByRouteFilterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRouteFilterRulesClient().NewListByRouteFilterPager("rg1", "filterName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.RouteFilterRuleListResult = armnetwork.RouteFilterRuleListResult{
		// 	Value: []*armnetwork.RouteFilterRule{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/routeFilters/filterName/routeFilterRules/ruleName"),
		// 			Name: to.Ptr("ruleName"),
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.RouteFilterRulePropertiesFormat{
		// 				Access: to.Ptr(armnetwork.AccessAllow),
		// 				Communities: []*string{
		// 					to.Ptr("12076:5030"),
		// 					to.Ptr("12076:5040")},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					RouteFilterRuleType: to.Ptr(armnetwork.RouteFilterRuleTypeCommunity),
		// 				},
		// 		}},
		// 	}
	}
}
