package armresourcemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/OperationsDiscovery_Get.json
func ExampleOperationsDiscoveryClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsDiscoveryClient().Get(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationsDiscoveryCollection = armresourcemover.OperationsDiscoveryCollection{
	// 	Value: []*armresourcemover.OperationsDiscovery{
	// 		{
	// 			Name: to.Ptr("Read"),
	// 			Display: &armresourcemover.Display{
	// 				Description: to.Ptr("Reads the move collection."),
	// 				Operation: to.Ptr("Read"),
	// 				Provider: to.Ptr("Microsoft.Migrate"),
	// 				Resource: to.Ptr("Microsoft.Migrate/moveCollections"),
	// 			},
	// 			Origin: to.Ptr("user"),
	// 		},
	// 		{
	// 			Name: to.Ptr("Write"),
	// 			Display: &armresourcemover.Display{
	// 				Description: to.Ptr("Creates or updates a move collection."),
	// 				Operation: to.Ptr("Write"),
	// 				Provider: to.Ptr("Microsoft.Migrate"),
	// 				Resource: to.Ptr("Microsoft.Migrate/moveCollections"),
	// 			},
	// 			Origin: to.Ptr("user"),
	// 	}},
	// }
}
