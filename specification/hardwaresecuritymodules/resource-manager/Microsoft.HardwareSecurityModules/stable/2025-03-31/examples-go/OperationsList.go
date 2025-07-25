package armhardwaresecuritymodules_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules/v2"
)

// Generated from example definition: 2025-03-31/OperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhardwaresecuritymodules.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armhardwaresecuritymodules.OperationsClientListResponse{
		// 	OperationListResult: armhardwaresecuritymodules.OperationListResult{
		// 		Value: []*armhardwaresecuritymodules.Operation{
		// 			{
		// 				Name: to.Ptr("Microsoft.HardwareSecurityModules/dedicatedhsms/read"),
		// 				Display: &armhardwaresecuritymodules.OperationDisplay{
		// 					Description: to.Ptr("View the properties of a DedicatedHsm"),
		// 					Operation: to.Ptr("View DedicatedHsm"),
		// 					Provider: to.Ptr("Microsoft Hardware Security Modules"),
		// 					Resource: to.Ptr("DedicatedHsm"),
		// 				},
		// 				IsDataAction: to.Ptr(false),
		// 				Origin: to.Ptr(armhardwaresecuritymodules.OriginSystem),
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.HardwareSecurityModules/cloudHsmClusters/read"),
		// 				Display: &armhardwaresecuritymodules.OperationDisplay{
		// 					Description: to.Ptr("View the properties of a CloudHsm"),
		// 					Operation: to.Ptr("View CloudHsm"),
		// 					Provider: to.Ptr("Microsoft Hardware Security Modules"),
		// 					Resource: to.Ptr("CloudHsm"),
		// 				},
		// 				IsDataAction: to.Ptr(false),
		// 				Origin: to.Ptr(armhardwaresecuritymodules.OriginSystem),
		// 			},
		// 		},
		// 	},
		// }
	}
}
