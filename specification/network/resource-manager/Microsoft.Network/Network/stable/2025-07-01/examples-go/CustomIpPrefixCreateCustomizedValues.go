package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/CustomIpPrefixCreateCustomizedValues.json
func ExampleCustomIPPrefixesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCustomIPPrefixesClient().BeginCreateOrUpdate(ctx, "rg1", "test-customipprefix", armnetwork.CustomIPPrefix{
		Location: to.Ptr("westus"),
		Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
			Cidr: to.Ptr("0.0.0.0/24"),
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
	// res = armnetwork.CustomIPPrefixesClientCreateOrUpdateResponse{
	// 	CustomIPPrefix: armnetwork.CustomIPPrefix{
	// 		Name: to.Ptr("test-customipprefix"),
	// 		Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
	// 		Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/customIpPrefixes/test-customipprefix"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
	// 			AuthorizationMessage: to.Ptr("authorizationMessage"),
	// 			ChildCustomIPPrefixes: []*armnetwork.SubResource{
	// 			},
	// 			Cidr: to.Ptr("192.168.254.2/24"),
	// 			CommissionedState: to.Ptr(armnetwork.CommissionedStateProvisioning),
	// 			FailedReason: to.Ptr(""),
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			PublicIPPrefixes: []*armnetwork.SubResource{
	// 			},
	// 			ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
	// 			SignedMessage: to.Ptr("signedMessage"),
	// 		},
	// 		Zones: []*string{
	// 			to.Ptr("1"),
	// 		},
	// 	},
	// }
}
