package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v2"
)

// Generated from example definition: 2025-03-01/Operations_List_MaximumSet_Gen.json
func ExampleOperationsClient_NewListPager_listTheOperationsForTheProviderGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armcontainerservicefleet.OperationsClientListResponse{
		// 	OperationListResult: armcontainerservicefleet.OperationListResult{
		// 		Value: []*armcontainerservicefleet.Operation{
		// 			{
		// 				Display: &armcontainerservicefleet.OperationDisplay{
		// 					Operation: to.Ptr("Get Operation"),
		// 					Provider: to.Ptr("Microsoft Container Service"),
		// 					Resource: to.Ptr("Operation"),
		// 					Description: to.Ptr("yvruoknlkuvuqxsodwkgqznxaig"),
		// 				},
		// 				Name: to.Ptr("Microsoft.ContainerService/locations/operations/read"),
		// 				Origin: to.Ptr(armcontainerservicefleet.OriginUserSystem),
		// 				IsDataAction: to.Ptr(true),
		// 				ActionType: to.Ptr(armcontainerservicefleet.ActionTypeInternal),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("http://nextlink.contoso.com"),
		// 	},
		// }
	}
}
