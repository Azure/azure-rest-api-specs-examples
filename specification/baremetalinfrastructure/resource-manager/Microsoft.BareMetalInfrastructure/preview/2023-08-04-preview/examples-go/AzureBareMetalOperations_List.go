package armbaremetalinfrastructure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/baremetalinfrastructure/armbaremetalinfrastructure/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/90115af9fda46f323e5c42c274f2b376108d1d47/specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/preview/2023-08-04-preview/examples/AzureBareMetalOperations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbaremetalinfrastructure.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armbaremetalinfrastructure.OperationListResult{
		// 	Value: []*armbaremetalinfrastructure.Operation{
		// 		{
		// 			Name: to.Ptr("AzureBareMetalOp1"),
		// 			Display: &armbaremetalinfrastructure.OperationDisplay{
		// 				Description: to.Ptr("AzureBareMetalOp1Description"),
		// 				Operation: to.Ptr("AzureBareMetalOp1OperationName"),
		// 				Provider: to.Ptr("AzureBareMetalOp1ProviderName"),
		// 				Resource: to.Ptr("AzureBareMetalOp1ResourceName"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("AzureBareMetalOp2"),
		// 			Display: &armbaremetalinfrastructure.OperationDisplay{
		// 				Description: to.Ptr("AzureBareMetalOp2Description"),
		// 				Operation: to.Ptr("AzureBareMetalOp2OperationName"),
		// 				Provider: to.Ptr("AzureBareMetalOp2ProviderName"),
		// 				Resource: to.Ptr("AzureBareMetalOp2ResourceName"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 	}},
		// }
	}
}
