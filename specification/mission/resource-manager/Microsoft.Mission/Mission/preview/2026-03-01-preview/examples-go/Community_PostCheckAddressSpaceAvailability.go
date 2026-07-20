package armenclave_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/enclave/armenclave"
)

// Generated from example definition: 2026-03-01-preview/Community_PostCheckAddressSpaceAvailability.json
func ExampleCommunityClient_CheckAddressSpaceAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armenclave.NewClientFactory("CA1CB369-DD26-4DB2-9D43-9AFEF0F22093", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCommunityClient().CheckAddressSpaceAvailability(ctx, "rgopenapi", "TestMyCommunity", armenclave.CheckAddressSpaceAvailabilityRequest{
		CommunityResourceID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestMyRg/providers/Microsoft.Mission/communities/TestMyCommunity"),
		EnclaveVirtualNetwork: &armenclave.VirtualNetworkModel{
			NetworkSize:     to.Ptr("small"),
			CustomCidrRange: to.Ptr("10.0.0.0/24"),
			SubnetConfigurations: []*armenclave.SubnetConfiguration{
				{
					SubnetName:        to.Ptr("test"),
					NetworkPrefixSize: to.Ptr[int32](26),
				},
			},
			AllowSubnetCommunication: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armenclave.CommunityClientCheckAddressSpaceAvailabilityResponse{
	// 	CheckAddressSpaceAvailabilityResponse: armenclave.CheckAddressSpaceAvailabilityResponse{
	// 		Value: to.Ptr(true),
	// 	},
	// }
}
