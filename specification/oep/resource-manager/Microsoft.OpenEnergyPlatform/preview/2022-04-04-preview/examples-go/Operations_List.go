package armoep_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oep/armoep"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/oep/resource-manager/Microsoft.OpenEnergyPlatform/preview/2022-04-04-preview/examples/Operations_List.json
func ExampleOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoep.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsClient().List(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationListResult = armoep.OperationListResult{
	// 	Value: []*armoep.Operation{
	// 		{
	// 			Name: to.Ptr("Microsoft.OpenEnergyPlatform/register/action"),
	// 			Display: &armoep.OperationDisplay{
	// 				Description: to.Ptr("Register the subscription for Microsoft.OpenEnergyPlatform"),
	// 				Operation: to.Ptr("Register the Microsoft.OpenEnergyPlatform"),
	// 				Provider: to.Ptr("Microsoft.OpenEnergyPlatform"),
	// 				Resource: to.Ptr("Microsoft.OpenEnergyPlatform"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.OpenEnergyPlatform/unregister/action"),
	// 			Display: &armoep.OperationDisplay{
	// 				Description: to.Ptr("Unregister the subscription for Microsoft.OpenEnergyPlatform"),
	// 				Operation: to.Ptr("Unregister the Microsoft.OpenEnergyPlatform"),
	// 				Provider: to.Ptr("Microsoft.OpenEnergyPlatform"),
	// 				Resource: to.Ptr("Microsoft.OpenEnergyPlatform"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.OpenEnergyPlatform/Locations/OperationStatuses/read"),
	// 			Display: &armoep.OperationDisplay{
	// 				Description: to.Ptr("read OperationStatuses"),
	// 				Operation: to.Ptr("read_OperationStatuses"),
	// 				Provider: to.Ptr("Microsoft.OpenEnergyPlatform"),
	// 				Resource: to.Ptr("Locations/OperationStatuses"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.OpenEnergyPlatform/Locations/OperationStatuses/write"),
	// 			Display: &armoep.OperationDisplay{
	// 				Description: to.Ptr("write OperationStatuses"),
	// 				Operation: to.Ptr("write_OperationStatuses"),
	// 				Provider: to.Ptr("Microsoft.OpenEnergyPlatform"),
	// 				Resource: to.Ptr("Locations/OperationStatuses"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.OpenEnergyPlatform/energyServices/read"),
	// 			Display: &armoep.OperationDisplay{
	// 				Description: to.Ptr("Returns list of oep resources.."),
	// 				Operation: to.Ptr("OepResources_ListByResourceGroup"),
	// 				Provider: to.Ptr("Microsoft.OpenEnergyPlatform"),
	// 				Resource: to.Ptr("energyServices"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.OpenEnergyPlatform/energyServices/read"),
	// 			Display: &armoep.OperationDisplay{
	// 				Description: to.Ptr("Lists a collection of oep resources under the given Azure Subscription ID."),
	// 				Operation: to.Ptr("OepResources_ListBySubscriptionId"),
	// 				Provider: to.Ptr("Microsoft.OpenEnergyPlatform"),
	// 				Resource: to.Ptr("energyServices"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.OpenEnergyPlatform/energyServices/read"),
	// 			Display: &armoep.OperationDisplay{
	// 				Description: to.Ptr("Returns oep resource for a given name."),
	// 				Operation: to.Ptr("OepResources_Get"),
	// 				Provider: to.Ptr("Microsoft.OpenEnergyPlatform"),
	// 				Resource: to.Ptr("energyServices"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.OpenEnergyPlatform/energyServices/write"),
	// 			Display: &armoep.OperationDisplay{
	// 				Description: to.Ptr("Method that gets called if subscribed for ResourceCreationBegin trigger."),
	// 				Operation: to.Ptr("OepResources_Create"),
	// 				Provider: to.Ptr("Microsoft.OpenEnergyPlatform"),
	// 				Resource: to.Ptr("energyServices"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.OpenEnergyPlatform/energyServices/delete"),
	// 			Display: &armoep.OperationDisplay{
	// 				Description: to.Ptr("Deletes oep resource"),
	// 				Operation: to.Ptr("OepResources_Delete"),
	// 				Provider: to.Ptr("Microsoft.OpenEnergyPlatform"),
	// 				Resource: to.Ptr("energyServices"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.OpenEnergyPlatform/energyServices/write"),
	// 			Display: &armoep.OperationDisplay{
	// 				Description: to.Ptr("write energyServices"),
	// 				Operation: to.Ptr("OepResources_Update"),
	// 				Provider: to.Ptr("Microsoft.OpenEnergyPlatform"),
	// 				Resource: to.Ptr("energyServices"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.OpenEnergyPlatform/Operations/read"),
	// 			Display: &armoep.OperationDisplay{
	// 				Description: to.Ptr("read Operations"),
	// 				Operation: to.Ptr("read_Operations"),
	// 				Provider: to.Ptr("Microsoft.OpenEnergyPlatform"),
	// 				Resource: to.Ptr("Operations"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.OpenEnergyPlatform/checkNameAvailability/action"),
	// 			Display: &armoep.OperationDisplay{
	// 				Description: to.Ptr("action checkNameAvailability"),
	// 				Operation: to.Ptr("action_checkNameAvailability"),
	// 				Provider: to.Ptr("Microsoft.OpenEnergyPlatform"),
	// 				Resource: to.Ptr("checkNameAvailability"),
	// 			},
	// 			IsDataAction: to.Ptr(false),
	// 	}},
	// }
}
