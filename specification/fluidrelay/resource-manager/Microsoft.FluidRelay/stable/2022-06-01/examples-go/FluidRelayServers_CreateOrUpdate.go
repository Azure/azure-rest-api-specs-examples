package armfluidrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fluidrelay/armfluidrelay"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayServers_CreateOrUpdate.json
func ExampleServersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armfluidrelay.NewServersClient("xxxx-xxxx-xxxx-xxxx", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"myResourceGroup",
		"myFluidRelayServer",
		armfluidrelay.Server{
			Location: to.Ptr("west-us"),
			Tags: map[string]*string{
				"Category": to.Ptr("sales"),
			},
			Identity: &armfluidrelay.Identity{
				Type: to.Ptr(armfluidrelay.ResourceIdentityTypeSystemAssigned),
			},
			Properties: &armfluidrelay.ServerProperties{
				Storagesku: to.Ptr(armfluidrelay.StorageSKUBasic),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
