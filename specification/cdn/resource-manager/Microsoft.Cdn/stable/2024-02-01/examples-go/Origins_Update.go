package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Origins_Update.json
func ExampleOriginsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOriginsClient().BeginUpdate(ctx, "RG", "profile1", "endpoint1", "www-someDomain-net", armcdn.OriginUpdateParameters{
		Properties: &armcdn.OriginUpdatePropertiesParameters{
			Enabled:          to.Ptr(true),
			HTTPPort:         to.Ptr[int32](42),
			HTTPSPort:        to.Ptr[int32](43),
			OriginHostHeader: to.Ptr("www.someDomain2.net"),
			Priority:         to.Ptr[int32](1),
			PrivateLinkAlias: to.Ptr("APPSERVER.d84e61f0-0870-4d24-9746-7438fa0019d1.westus2.azure.privatelinkservice"),
			Weight:           to.Ptr[int32](50),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Origin = armcdn.Origin{
	// 	Name: to.Ptr("www-someDomain-net"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles/endpoints/origins"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/www-someDomain-net"),
	// 	Properties: &armcdn.OriginProperties{
	// 		Enabled: to.Ptr(true),
	// 		HostName: to.Ptr("www.someDomain.net"),
	// 		HTTPPort: to.Ptr[int32](42),
	// 		HTTPSPort: to.Ptr[int32](43),
	// 		OriginHostHeader: to.Ptr("www.someDomain2.net"),
	// 		Priority: to.Ptr[int32](1),
	// 		PrivateLinkAlias: to.Ptr("APPSERVER.d84e61f0-0870-4d24-9746-7438fa0019d1.westus2.azure.privatelinkservice"),
	// 		PrivateLinkApprovalMessage: to.Ptr("Please approve the connection request for this Private Link"),
	// 		Weight: to.Ptr[int32](50),
	// 		PrivateEndpointStatus: to.Ptr(armcdn.PrivateEndpointStatusPending),
	// 		ProvisioningState: to.Ptr(armcdn.OriginProvisioningStateSucceeded),
	// 		ResourceState: to.Ptr(armcdn.OriginResourceStateActive),
	// 	},
	// }
}
