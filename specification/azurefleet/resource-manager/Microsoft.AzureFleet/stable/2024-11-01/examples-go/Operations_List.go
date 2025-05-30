package armcomputefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computefleet/armcomputefleet"
)

// Generated from example definition: D:/w/Azure/azure-sdk-for-go/sdk/resourcemanager/computefleet/armcomputefleet/TempTypeSpecFiles/AzureFleet.Management/examples/2024-11-01/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputefleet.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armcomputefleet.OperationsClientListResponse{
		// 	OperationListResult: armcomputefleet.OperationListResult{
		// 		Value: []*armcomputefleet.Operation{
		// 			{
		// 				Origin: to.Ptr(armcomputefleet.OriginUserSystem),
		// 				Name: to.Ptr("Microsoft.AzureFleet/fleets/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armcomputefleet.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft Azure Fleet"),
		// 					Resource: to.Ptr("Fleets"),
		// 					Operation: to.Ptr("Get Azure Fleet"),
		// 					Description: to.Ptr("Get properties of Azure Fleet resource"),
		// 				},
		// 			},
		// 			{
		// 				Origin: to.Ptr(armcomputefleet.OriginUserSystem),
		// 				Name: to.Ptr("Microsoft.AzureFleet/fleets/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armcomputefleet.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft Azure Fleet"),
		// 					Resource: to.Ptr("Fleets"),
		// 					Operation: to.Ptr("Create or Update Azure Fleet"),
		// 					Description: to.Ptr("Creates a new Azure Fleet resource or updates an existing one"),
		// 				},
		// 			},
		// 			{
		// 				Origin: to.Ptr(armcomputefleet.OriginUserSystem),
		// 				Name: to.Ptr("Microsoft.AzureFleet/fleets/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armcomputefleet.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft Azure Fleet"),
		// 					Resource: to.Ptr("Fleets"),
		// 					Operation: to.Ptr("Delete Virtual Machine and Virtual Machine scale sets in a Azure Fleet resource"),
		// 					Description: to.Ptr("Deletes all compute resources of Azure Fleet resource"),
		// 				},
		// 			},
		// 			{
		// 				Origin: to.Ptr(armcomputefleet.OriginUserSystem),
		// 				Name: to.Ptr("Microsoft.AzureFleet/register/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armcomputefleet.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft Azure Fleet"),
		// 					Resource: to.Ptr("Subscription"),
		// 					Operation: to.Ptr("Register subscription for Microsoft.AzureFleet"),
		// 					Description: to.Ptr("Registers Subscription with Microsoft.AzureFleet resource provider"),
		// 				},
		// 			},
		// 			{
		// 				Origin: to.Ptr(armcomputefleet.OriginUserSystem),
		// 				Name: to.Ptr("Microsoft.AzureFleet/unregister/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armcomputefleet.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft Azure Fleet"),
		// 					Resource: to.Ptr("Subscription"),
		// 					Operation: to.Ptr("Unregister Subscription for Microsoft.AzureFleet"),
		// 					Description: to.Ptr("Unregisters Subscription with Microsoft.AzureFleet resource provider"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
