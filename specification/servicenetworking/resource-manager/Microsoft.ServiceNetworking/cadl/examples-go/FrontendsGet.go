package armservicenetworking_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/71121282e39bccae590462648e77bca283df6d2b/specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/FrontendsGet.json
func ExampleFrontendsInterfaceClient_NewListByTrafficControllerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicenetworking.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFrontendsInterfaceClient().NewListByTrafficControllerPager("rg1", "TC1", nil)
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
		// page.FrontendListResult = armservicenetworking.FrontendListResult{
		// 	Value: []*armservicenetworking.Frontend{
		// 		{
		// 			Name: to.Ptr("publicIp1"),
		// 			Type: to.Ptr("Microsoft.ServiceNetworking/trafficControllers/frontends"),
		// 			ID: to.Ptr("resourceUriAsString"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armservicenetworking.FrontendProperties{
		// 				IPAddressVersion: to.Ptr(armservicenetworking.FrontendIPAddressVersionIPv4),
		// 				Mode: to.Ptr("public"),
		// 				ProvisioningState: to.Ptr(armservicenetworking.ProvisioningStateSucceeded),
		// 				PublicIPAddress: &armservicenetworking.FrontendPropertiesIPAddress{
		// 					ID: to.Ptr("resourceUriAsString"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
