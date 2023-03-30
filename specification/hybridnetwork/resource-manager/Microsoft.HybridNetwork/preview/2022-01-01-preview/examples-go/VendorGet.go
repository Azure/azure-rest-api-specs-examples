package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/792db17291c758b2bfdbbc0d35d0e2f5b5a1bd05/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/VendorGet.json
func ExampleVendorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVendorsClient().Get(ctx, "TestVendor", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Vendor = armhybridnetwork.Vendor{
	// 	Name: to.Ptr("TestVendor"),
	// 	Type: to.Ptr("Microsoft.HybridNetwork/vendors"),
	// 	ID: to.Ptr("/subscriptions/subid/providers/Microsoft.HybridNetwork/vendors/TestVendor"),
	// 	Properties: &armhybridnetwork.VendorPropertiesFormat{
	// 		ProvisioningState: to.Ptr(armhybridnetwork.ProvisioningStateSucceeded),
	// 		SKUs: []*armhybridnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/providers/Microsoft.HybridNetwork/vendors/TestVendor/skus/TestVendorSku"),
	// 		}},
	// 	},
	// }
}
