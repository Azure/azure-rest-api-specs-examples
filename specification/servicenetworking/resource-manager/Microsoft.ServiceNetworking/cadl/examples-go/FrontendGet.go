package armservicenetworking_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a7af0df86022e5e6cc6e8f40ca1981c4557a4bc/specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/FrontendGet.json
func ExampleFrontendsInterfaceClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicenetworking.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFrontendsInterfaceClient().Get(ctx, "rg1", "tc1", "fe1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Frontend = armservicenetworking.Frontend{
	// 	Name: to.Ptr("fe1"),
	// 	Type: to.Ptr("Microsoft.ServiceNetworking/trafficControllers/frontends"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ServiceNetworking/trafficControllers/tc1/frontends/fe1"),
	// 	Location: to.Ptr("NorthCentralUS"),
	// 	Properties: &armservicenetworking.FrontendProperties{
	// 		Fqdn: to.Ptr("test.net"),
	// 		ProvisioningState: to.Ptr(armservicenetworking.ProvisioningStateSucceeded),
	// 	},
	// }
}
