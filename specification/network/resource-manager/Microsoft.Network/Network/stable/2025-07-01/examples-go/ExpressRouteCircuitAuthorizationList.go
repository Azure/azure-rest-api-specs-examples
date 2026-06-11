package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ExpressRouteCircuitAuthorizationList.json
func ExampleExpressRouteCircuitAuthorizationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExpressRouteCircuitAuthorizationsClient().NewListPager("rg1", "circuitName", nil)
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
		// page = armnetwork.ExpressRouteCircuitAuthorizationsClientListResponse{
		// 	AuthorizationListResult: armnetwork.AuthorizationListResult{
		// 		Value: []*armnetwork.ExpressRouteCircuitAuthorization{
		// 			{
		// 				Name: to.Ptr("MyAuthorization1"),
		// 				Etag: to.Ptr("W/\"e33c875f-48df-4a91-b7d3-eb95b5ddbb89\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuitName/authorizations/MyAuthorization1"),
		// 				Properties: &armnetwork.AuthorizationPropertiesFormat{
		// 					AuthorizationKey: to.Ptr("authKey"),
		// 					AuthorizationUseStatus: to.Ptr(armnetwork.AuthorizationUseStatusAvailable),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
