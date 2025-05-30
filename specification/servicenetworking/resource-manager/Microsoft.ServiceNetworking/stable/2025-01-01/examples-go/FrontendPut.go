package armservicenetworking_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking"
)

// Generated from example definition: 2025-01-01/FrontendPut.json
func ExampleFrontendsInterfaceClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicenetworking.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFrontendsInterfaceClient().BeginCreateOrUpdate(ctx, "rg1", "tc1", "fe1", armservicenetworking.Frontend{
		Location:   to.Ptr("NorthCentralUS"),
		Properties: &armservicenetworking.FrontendProperties{},
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
	// res = armservicenetworking.FrontendsInterfaceClientCreateOrUpdateResponse{
	// 	Frontend: &armservicenetworking.Frontend{
	// 		Name: to.Ptr("fe1"),
	// 		Location: to.Ptr("NorthCentralUS"),
	// 		Type: to.Ptr("Microsoft.ServiceNetworking/trafficControllers/frontends"),
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ServiceNetworking/trafficControllers/tc1/frontends/fe1"),
	// 		Properties: &armservicenetworking.FrontendProperties{
	// 			Fqdn: to.Ptr("test.net"),
	// 			ProvisioningState: to.Ptr(armservicenetworking.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
