package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v12"
)

// Generated from example definition: 2026-03-01/ExpressRouteLagAuthorizationList.json
func ExampleExpressRouteLagAuthorizationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExpressRouteLagAuthorizationsClient().NewListPager("rg1", "expressRouteLagName", nil)
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
		// page = armnetwork.ExpressRouteLagAuthorizationsClientListResponse{
		// 	ExpressRouteLagAuthorizationListResult: armnetwork.ExpressRouteLagAuthorizationListResult{
		// 		Value: []*armnetwork.ExpressRouteLagAuthorization{
		// 			{
		// 				Name: to.Ptr("authorizationName"),
		// 				Type: to.Ptr("Microsoft.Network/expressRouteLags/authorizations"),
		// 				Etag: to.Ptr("W/\"e33c875f-48df-4a91-b7d3-eb95b5ddbb89\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/expressRouteLagName/authorizations/authorizationName"),
		// 				Properties: &armnetwork.ExpressRouteLagAuthorizationPropertiesFormat{
		// 					AuthorizationUseStatus: to.Ptr(armnetwork.ExpressRouteLagAuthorizationUseStatusAvailable),
		// 					CircuitResourceURI: to.Ptr(""),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
