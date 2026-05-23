package armpowerbidedicated_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated/v2"
)

// Generated from example definition: 2021-01-01/listSKUsForNew.json
func ExampleCapacitiesClient_ListSKUs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpowerbidedicated.NewClientFactory("613192d7-503f-477a-9cfe-4efc3ee2bd60", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCapacitiesClient().ListSKUs(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpowerbidedicated.CapacitiesClientListSKUsResponse{
	// 	SKUEnumerationForNewResourceResult: &armpowerbidedicated.SKUEnumerationForNewResourceResult{
	// 		Value: []*armpowerbidedicated.CapacitySKU{
	// 			{
	// 				Name: to.Ptr("A1"),
	// 			},
	// 			{
	// 				Name: to.Ptr("A2"),
	// 			},
	// 			{
	// 				Name: to.Ptr("A3"),
	// 			},
	// 			{
	// 				Name: to.Ptr("A4"),
	// 			},
	// 			{
	// 				Name: to.Ptr("A5"),
	// 			},
	// 			{
	// 				Name: to.Ptr("A6"),
	// 			},
	// 		},
	// 	},
	// }
}
