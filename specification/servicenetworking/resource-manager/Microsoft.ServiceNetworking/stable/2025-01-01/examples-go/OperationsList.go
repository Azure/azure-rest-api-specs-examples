package armservicenetworking_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking"
)

// Generated from example definition: 2025-01-01/OperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicenetworking.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armservicenetworking.OperationsClientListResponse{
		// 	OperationListResult: armservicenetworking.OperationListResult{
		// 		Value: []*armservicenetworking.Operation{
		// 			{
		// 				Name: to.Ptr("Microsoft.ServiceNetworking/trafficControllers/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armservicenetworking.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft Service Networking"),
		// 					Resource: to.Ptr("Traffic Controller"),
		// 					Operation: to.Ptr("Get Traffic Controller configuration"),
		// 					Description: to.Ptr("Traffic Controller is a L7 Load Balancing solution for a Multi Cluster setup"),
		// 				},
		// 				Origin: to.Ptr(armservicenetworking.OriginUserSystem),
		// 			},
		// 		},
		// 	},
		// }
	}
}
