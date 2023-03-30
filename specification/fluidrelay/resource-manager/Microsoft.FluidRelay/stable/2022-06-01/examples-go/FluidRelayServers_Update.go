package armfluidrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fluidrelay/armfluidrelay"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayServers_Update.json
func ExampleServersClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfluidrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServersClient().Update(ctx, "myResourceGroup", "myFluidRelayServer", armfluidrelay.ServerUpdate{
		Tags: map[string]*string{
			"Category": to.Ptr("sales"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Server = armfluidrelay.Server{
	// 	Name: to.Ptr("myFluidRelayServer"),
	// 	Type: to.Ptr("Microsoft.FluidRelay/fluidRelayServers"),
	// 	ID: to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/resourceGroups/myResourceGroup/Microsoft.FluidRelay/fluidRelayServers/myFluidRelayServer"),
	// 	Location: to.Ptr("west-us"),
	// 	Tags: map[string]*string{
	// 		"Category": to.Ptr("sales"),
	// 	},
	// 	Properties: &armfluidrelay.ServerProperties{
	// 		FluidRelayEndpoints: &armfluidrelay.Endpoints{
	// 			OrdererEndpoints: []*string{
	// 				to.Ptr("https://www.contoso.org/orderer")},
	// 				ServiceEndpoints: []*string{
	// 					to.Ptr("https://www.afd.contoso.org")},
	// 					StorageEndpoints: []*string{
	// 						to.Ptr("https://www.contoso.org/storage")},
	// 					},
	// 					FrsTenantID: to.Ptr("yyyy-yyyy-yyyyy-yyyy"),
	// 					ProvisioningState: to.Ptr(armfluidrelay.ProvisioningStateSucceeded),
	// 				},
	// 			}
}
