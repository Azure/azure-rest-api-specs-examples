package armfluidrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fluidrelay/armfluidrelay"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayContainers_ListByFluidRelayServer.json
func ExampleContainersClient_NewListByFluidRelayServersPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfluidrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContainersClient().NewListByFluidRelayServersPager("myResourceGroup", "myFluidRelayServer", nil)
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
		// page.ContainerList = armfluidrelay.ContainerList{
		// 	Value: []*armfluidrelay.Container{
		// 		{
		// 			Name: to.Ptr("myFluidRelayContainer"),
		// 			Type: to.Ptr("Microsoft.FluidRelay/fluidRelayServers/fluidRelayContainers"),
		// 			ID: to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/resourceGroups/myResourceGroup/Microsoft.FluidRelay/fluidRelayServers/myFluidRelayServer/fluidRelayContainers/myFluidRelayContainer"),
		// 			Properties: &armfluidrelay.ContainerProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-21T02:17:01.164Z"); return t}()),
		// 				FrsContainerID: to.Ptr("xxxx-yyyy-xxxxx-yyyy"),
		// 				FrsTenantID: to.Ptr("yyyy-yyyy-yyyyy-yyyy"),
		// 				LastAccessTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-21T02:17:01.164Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armfluidrelay.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
