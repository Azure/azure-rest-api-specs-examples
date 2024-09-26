package armfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fabric/armfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e79d9ef3e065f2dcb6bd1db51e29c62a99dff5cb/specification/fabric/resource-manager/Microsoft.Fabric/stable/2023-11-01/examples/FabricCapacities_ListSkus.json
func ExampleCapacitiesClient_NewListSKUsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCapacitiesClient().NewListSKUsPager(nil)
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
		// page.RpSKUEnumerationForNewResourceResult = armfabric.RpSKUEnumerationForNewResourceResult{
		// 	Value: []*armfabric.RpSKUDetailsForNewResource{
		// 		{
		// 			Name: to.Ptr("F8"),
		// 			Locations: []*string{
		// 				to.Ptr("West Europe")},
		// 				ResourceType: to.Ptr("Capacities"),
		// 			},
		// 			{
		// 				Name: to.Ptr("F64"),
		// 				Locations: []*string{
		// 					to.Ptr("West Europe")},
		// 					ResourceType: to.Ptr("Capacities"),
		// 				},
		// 				{
		// 					Name: to.Ptr("F128"),
		// 					Locations: []*string{
		// 						to.Ptr("West Europe")},
		// 						ResourceType: to.Ptr("Capacities"),
		// 					},
		// 					{
		// 						Name: to.Ptr("F512"),
		// 						Locations: []*string{
		// 							to.Ptr("West Europe")},
		// 							ResourceType: to.Ptr("Capacities"),
		// 					}},
		// 				}
	}
}
