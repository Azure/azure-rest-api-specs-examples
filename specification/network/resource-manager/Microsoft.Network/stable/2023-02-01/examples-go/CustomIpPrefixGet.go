package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9975d3476c05bcc6bd9535ad3dfb564e6a168fa5/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/CustomIpPrefixGet.json
func ExampleCustomIPPrefixesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomIPPrefixesClient().Get(ctx, "rg1", "test-customipprefix", &armnetwork.CustomIPPrefixesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomIPPrefix = armnetwork.CustomIPPrefix{
	// 	Name: to.Ptr("test-customipprefix"),
	// 	Type: to.Ptr("Microsoft.Network/customIpPrefixes"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/customIpPrefixes/test-customipprefix"),
	// 	Location: to.Ptr("westus"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 	Properties: &armnetwork.CustomIPPrefixPropertiesFormat{
	// 		AuthorizationMessage: to.Ptr("authorizationMessage"),
	// 		ChildCustomIPPrefixes: []*armnetwork.SubResource{
	// 		},
	// 		Cidr: to.Ptr("0.0.0.0/24"),
	// 		CommissionedState: to.Ptr(armnetwork.CommissionedStateCommissioned),
	// 		ExpressRouteAdvertise: to.Ptr(false),
	// 		FailedReason: to.Ptr(""),
	// 		NoInternetAdvertise: to.Ptr(false),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		PublicIPPrefixes: []*armnetwork.SubResource{
	// 		},
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
	// 		SignedMessage: to.Ptr("signedMessage"),
	// 	},
	// }
}
