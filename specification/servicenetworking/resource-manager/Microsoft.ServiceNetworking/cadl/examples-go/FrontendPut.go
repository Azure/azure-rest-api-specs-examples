package armservicenetworking_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/71121282e39bccae590462648e77bca283df6d2b/specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/FrontendPut.json
func ExampleFrontendsInterfaceClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicenetworking.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFrontendsInterfaceClient().BeginCreateOrUpdate(ctx, "rg1", "TC1", "publicIp1", armservicenetworking.Frontend{
		Location: to.Ptr("West US"),
		Properties: &armservicenetworking.FrontendProperties{
			IPAddressVersion: to.Ptr(armservicenetworking.FrontendIPAddressVersionIPv4),
			Mode:             to.Ptr("public"),
			PublicIPAddress: &armservicenetworking.FrontendPropertiesIPAddress{
				ID: to.Ptr("resourceUriAsString"),
			},
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
	// res.Frontend = armservicenetworking.Frontend{
	// 	Name: to.Ptr("publicIp1"),
	// 	Type: to.Ptr("Microsoft.ServiceNetworking/trafficControllers/frontends"),
	// 	ID: to.Ptr("resourceUriAsString"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armservicenetworking.FrontendProperties{
	// 		IPAddressVersion: to.Ptr(armservicenetworking.FrontendIPAddressVersionIPv4),
	// 		Mode: to.Ptr("public"),
	// 		ProvisioningState: to.Ptr(armservicenetworking.ProvisioningStateSucceeded),
	// 		PublicIPAddress: &armservicenetworking.FrontendPropertiesIPAddress{
	// 			ID: to.Ptr("resourceUriAsString"),
	// 		},
	// 	},
	// }
}
