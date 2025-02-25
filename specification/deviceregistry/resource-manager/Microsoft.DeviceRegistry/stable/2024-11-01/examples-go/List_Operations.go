package armdeviceregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry"
)

// Generated from example definition: 2024-11-01/List_Operations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armdeviceregistry.OperationsClientListResponse{
		// 	OperationListResult: armdeviceregistry.OperationListResult{
		// 		Value: []*armdeviceregistry.Operation{
		// 			{
		// 				Name: to.Ptr("Microsoft.DeviceRegistry/assets/write"),
		// 				Display: &armdeviceregistry.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft Azure Device Registry"),
		// 					Resource: to.Ptr("Asset"),
		// 					Operation: to.Ptr("write"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
