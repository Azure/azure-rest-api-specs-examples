package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c58fa033619b12c7cfa8a0ec5a9bf03bb18869ab/specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/PublicIpPrefixUpdateTags.json
func ExamplePublicIPPrefixesClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPublicIPPrefixesClient().UpdateTags(ctx, "rg1", "test-ipprefix", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PublicIPPrefix = armnetwork.PublicIPPrefix{
	// 	Name: to.Ptr("test-ipprefix"),
	// 	Type: to.Ptr("Microsoft.Network/publicIPPrefixes"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/test-ipprefix"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 	Properties: &armnetwork.PublicIPPrefixPropertiesFormat{
	// 		IPPrefix: to.Ptr("40.85.154.247/30"),
	// 		IPTags: []*armnetwork.IPTag{
	// 		},
	// 		PrefixLength: to.Ptr[int32](30),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
	// 	},
	// 	SKU: &armnetwork.PublicIPPrefixSKU{
	// 		Name: to.Ptr(armnetwork.PublicIPPrefixSKUNameStandard),
	// 	},
	// }
}
