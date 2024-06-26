package armfluidrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fluidrelay/armfluidrelay"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayServerOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfluidrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationListResult = armfluidrelay.OperationListResult{
		// 	Value: []*armfluidrelay.OperationResult{
		// 		{
		// 			Name: to.Ptr("Microsoft.FluidRelay/fluidRelayServers/Read"),
		// 			Display: &armfluidrelay.OperationDisplay{
		// 				Description: to.Ptr("Read fluid relay server"),
		// 				Operation: to.Ptr("Get/List fluid relay server resources"),
		// 				Provider: to.Ptr("Microsoft.FluidRelay"),
		// 				Resource: to.Ptr("fluidRelayServers"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.FluidRelay/fluidRelayServers/Write"),
		// 			Display: &armfluidrelay.OperationDisplay{
		// 				Description: to.Ptr("Write fluid relay server"),
		// 				Operation: to.Ptr("Create/Update fluid relay server resources"),
		// 				Provider: to.Ptr("Microsoft.FluidRelay"),
		// 				Resource: to.Ptr("fluidRelayServers"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.FluidRelay/fluidRelayServers/Delete"),
		// 			Display: &armfluidrelay.OperationDisplay{
		// 				Description: to.Ptr("Delete fluid relay server"),
		// 				Operation: to.Ptr("Delete fluid relay server resources"),
		// 				Provider: to.Ptr("Microsoft.FluidRelay"),
		// 				Resource: to.Ptr("fluidRelayServers"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
