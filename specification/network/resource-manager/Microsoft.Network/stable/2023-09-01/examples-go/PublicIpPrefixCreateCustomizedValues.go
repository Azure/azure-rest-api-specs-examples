package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/network/resource-manager/Microsoft.Network/stable/2023-09-01/examples/PublicIpPrefixCreateCustomizedValues.json
func ExamplePublicIPPrefixesClient_BeginCreateOrUpdate_createPublicIpPrefixAllocationMethod() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPublicIPPrefixesClient().BeginCreateOrUpdate(ctx, "rg1", "test-ipprefix", armnetwork.PublicIPPrefix{
		Location: to.Ptr("westus"),
		Properties: &armnetwork.PublicIPPrefixPropertiesFormat{
			PrefixLength:           to.Ptr[int32](30),
			PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		},
		SKU: &armnetwork.PublicIPPrefixSKU{
			Name: to.Ptr(armnetwork.PublicIPPrefixSKUNameStandard),
			Tier: to.Ptr(armnetwork.PublicIPPrefixSKUTierRegional),
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
	// res.PublicIPPrefix = armnetwork.PublicIPPrefix{
	// 	Name: to.Ptr("test-ipprefix"),
	// 	Type: to.Ptr("Microsoft.Network/publicIPPrefixes"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/test-ipprefix"),
	// 	Location: to.Ptr("westus"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 	Properties: &armnetwork.PublicIPPrefixPropertiesFormat{
	// 		IPPrefix: to.Ptr("192.168.254.2/30"),
	// 		IPTags: []*armnetwork.IPTag{
	// 		},
	// 		PrefixLength: to.Ptr[int32](30),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
	// 	},
	// 	SKU: &armnetwork.PublicIPPrefixSKU{
	// 		Name: to.Ptr(armnetwork.PublicIPPrefixSKUNameStandard),
	// 		Tier: to.Ptr(armnetwork.PublicIPPrefixSKUTierRegional),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1")},
	// 	}
}
