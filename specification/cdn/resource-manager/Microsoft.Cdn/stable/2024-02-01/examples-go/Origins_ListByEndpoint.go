package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Origins_ListByEndpoint.json
func ExampleOriginsClient_NewListByEndpointPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOriginsClient().NewListByEndpointPager("RG", "profile1", "endpoint1", nil)
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
		// page.OriginListResult = armcdn.OriginListResult{
		// 	Value: []*armcdn.Origin{
		// 		{
		// 			Name: to.Ptr("www-someDomain-net"),
		// 			Type: to.Ptr("Microsoft.Cdn/profiles/endpoints/origins"),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/www-someDomain-net"),
		// 			Properties: &armcdn.OriginProperties{
		// 				Enabled: to.Ptr(true),
		// 				HostName: to.Ptr("www.someDomain.net"),
		// 				OriginHostHeader: to.Ptr("www.someDomain.net"),
		// 				Priority: to.Ptr[int32](1),
		// 				PrivateLinkAlias: to.Ptr("APPSERVER.d84e61f0-0870-4d24-9746-7438fa0019d1.westus2.azure.privatelinkservice"),
		// 				PrivateLinkApprovalMessage: to.Ptr("Please approve the connection request for this Private Link"),
		// 				Weight: to.Ptr[int32](50),
		// 				PrivateEndpointStatus: to.Ptr(armcdn.PrivateEndpointStatusPending),
		// 				ProvisioningState: to.Ptr(armcdn.OriginProvisioningStateSucceeded),
		// 				ResourceState: to.Ptr(armcdn.OriginResourceStateActive),
		// 			},
		// 	}},
		// }
	}
}
