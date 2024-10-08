package armfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fabric/armfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e79d9ef3e065f2dcb6bd1db51e29c62a99dff5cb/specification/fabric/resource-manager/Microsoft.Fabric/stable/2023-11-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfabric.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armfabric.OperationListResult{
		// 	Value: []*armfabric.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Fabric/capacities/read"),
		// 			Display: &armfabric.OperationDisplay{
		// 				Description: to.Ptr("Retrieves the information of the specified Fabric Capacity."),
		// 				Operation: to.Ptr("Read Fabric Capacity"),
		// 				Provider: to.Ptr("Microsoft Fabric"),
		// 				Resource: to.Ptr("capacities"),
		// 			},
		// 			Origin: to.Ptr(armfabric.OriginUserSystem),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Fabric/capacities/write"),
		// 			Display: &armfabric.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the specified Fabric Capacity."),
		// 				Operation: to.Ptr("Create/Update Fabric Capacity"),
		// 				Provider: to.Ptr("Microsoft Fabric"),
		// 				Resource: to.Ptr("capacities"),
		// 			},
		// 			Origin: to.Ptr(armfabric.OriginUserSystem),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Fabric/capacities/delete"),
		// 			Display: &armfabric.OperationDisplay{
		// 				Description: to.Ptr("Deletes the Fabric Capacity."),
		// 				Operation: to.Ptr("Delete the Fabric Capacity"),
		// 				Provider: to.Ptr("Microsoft Fabric"),
		// 				Resource: to.Ptr("capacities"),
		// 			},
		// 			Origin: to.Ptr(armfabric.OriginUserSystem),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Fabric/capacities/suspend/action"),
		// 			Display: &armfabric.OperationDisplay{
		// 				Description: to.Ptr("Suspend the specified Fabric capacity"),
		// 				Operation: to.Ptr("Suspend the specified Fabric capacity"),
		// 				Provider: to.Ptr("Microsoft Fabric"),
		// 				Resource: to.Ptr("capacities"),
		// 			},
		// 			Origin: to.Ptr(armfabric.OriginUserSystem),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Fabric/capacities/resume/action"),
		// 			Display: &armfabric.OperationDisplay{
		// 				Description: to.Ptr("Resume the specified Fabric capacity"),
		// 				Operation: to.Ptr("Resume the specified Fabric capacity"),
		// 				Provider: to.Ptr("Microsoft Fabric"),
		// 				Resource: to.Ptr("capacities"),
		// 			},
		// 			Origin: to.Ptr(armfabric.OriginUserSystem),
		// 	}},
		// }
	}
}
