package armstorageactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storageactions/armstorageactions"
)

// Generated from example definition: 2023-01-01/misc/OperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorageactions.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armstorageactions.OperationsClientListResponse{
		// 	OperationListResult: armstorageactions.OperationListResult{
		// 		Value: []*armstorageactions.Operation{
		// 			{
		// 				Name: to.Ptr("Microsoft.StorageActions/storageTasks/read"),
		// 				Display: &armstorageactions.OperationDisplay{
		// 					Description: to.Ptr("Gets or Lists existing StorageTask resource(s)."),
		// 					Operation: to.Ptr("Get or List StorageTask resource(s)."),
		// 					Provider: to.Ptr("Microsoft StorageActions"),
		// 					Resource: to.Ptr("StorageTasks"),
		// 				},
		// 				IsDataAction: to.Ptr(false),
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.StorageActions/storageTasks/write"),
		// 				Display: &armstorageactions.OperationDisplay{
		// 					Description: to.Ptr("Creates or Updates StorageTask resource."),
		// 					Operation: to.Ptr("Create or Update StorageTask resource."),
		// 					Provider: to.Ptr("Microsoft StorageActions"),
		// 					Resource: to.Ptr("StorageTasks"),
		// 				},
		// 				IsDataAction: to.Ptr(false),
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.StorageActions/storageTasks/delete"),
		// 				Display: &armstorageactions.OperationDisplay{
		// 					Description: to.Ptr("Deletes StorageTask resource."),
		// 					Operation: to.Ptr("Delete StorageTask resource."),
		// 					Provider: to.Ptr("Microsoft StorageActions"),
		// 					Resource: to.Ptr("StorageTasks"),
		// 				},
		// 				IsDataAction: to.Ptr(false),
		// 			},
		// 		},
		// 	},
		// }
	}
}
