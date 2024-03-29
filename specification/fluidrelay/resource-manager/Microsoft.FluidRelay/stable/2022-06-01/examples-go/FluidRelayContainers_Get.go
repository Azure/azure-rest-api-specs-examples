package armfluidrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fluidrelay/armfluidrelay"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayContainers_Get.json
func ExampleContainersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfluidrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainersClient().Get(ctx, "myResourceGroup", "myFluidRelayServer", "myFluidRelayContainer", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Container = armfluidrelay.Container{
	// 	Name: to.Ptr("myFluidRelayContainer"),
	// 	Type: to.Ptr("Microsoft.FluidRelay/fluidRelayServers/fluidRelayContainers"),
	// 	ID: to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/resourceGroups/myResourceGroup/Microsoft.FluidRelay/fluidRelayServers/myFluidRelayServer/fluidRelayContainers/myFluidRelayContainer"),
	// 	Properties: &armfluidrelay.ContainerProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-21T02:17:01.164Z"); return t}()),
	// 		FrsContainerID: to.Ptr("xxxx-yyyy-xxxxx-yyyy"),
	// 		FrsTenantID: to.Ptr("yyyy-yyyy-yyyyy-yyyy"),
	// 		LastAccessTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-21T02:17:01.164Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armfluidrelay.ProvisioningStateSucceeded),
	// 	},
	// }
}
