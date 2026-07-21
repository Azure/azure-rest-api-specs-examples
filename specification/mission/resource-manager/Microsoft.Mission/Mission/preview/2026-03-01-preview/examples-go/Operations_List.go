package armenclave_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/enclave/armenclave"
)

// Generated from example definition: 2026-03-01-preview/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armenclave.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armenclave.OperationsClientListResponse{
		// 	OperationListResult: armenclave.OperationListResult{
		// 		Value: []*armenclave.Operation{
		// 			{
		// 				Name: to.Ptr("kribdpoznbvrjci"),
		// 				IsDataAction: to.Ptr(true),
		// 				Display: &armenclave.OperationDisplay{
		// 					Provider: to.Ptr("usvnnruysydqdpwj"),
		// 					Resource: to.Ptr("lrdjzltcbeoljosqrw"),
		// 					Operation: to.Ptr("vdnitexr"),
		// 					Description: to.Ptr("valid description goes here"),
		// 				},
		// 				Origin: to.Ptr(armenclave.OriginUser),
		// 				ActionType: to.Ptr(armenclave.ActionTypeInternal),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
