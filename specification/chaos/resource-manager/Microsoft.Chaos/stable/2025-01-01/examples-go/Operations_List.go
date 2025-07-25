package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2025-01-01/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armchaos.OperationsClientListResponse{
		// 	OperationListResult: armchaos.OperationListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/providers/Microsoft.Chaos/operations?continuationToken=myContinuationToken&api-version=2024-11-01-preview"),
		// 		Value: []*armchaos.Operation{
		// 			{
		// 				Name: to.Ptr("Microsoft.Chaos/experiments/read"),
		// 				Display: &armchaos.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft Chaos"),
		// 					Resource: to.Ptr("Chaos Experiment"),
		// 					Operation: to.Ptr("Gets all Chaos Experiments."),
		// 					Description: to.Ptr("Gets all Chaos Experiments in a resource group."),
		// 				},
		// 				IsDataAction: to.Ptr(false),
		// 				Origin: to.Ptr(armchaos.OriginUserSystem),
		// 			},
		// 		},
		// 	},
		// }
	}
}
