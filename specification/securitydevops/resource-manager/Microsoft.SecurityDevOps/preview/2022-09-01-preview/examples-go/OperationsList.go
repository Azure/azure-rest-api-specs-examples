package armsecuritydevops_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c71a66dab813061f1d09982c2748a09317fe0860/specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/OperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecuritydevops.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armsecuritydevops.OperationListResult{
		// 	Value: []*armsecuritydevops.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityDevOps/gitHubConnectors/Read"),
		// 			Display: &armsecuritydevops.OperationDisplay{
		// 				Description: to.Ptr("Read gitHubConnector"),
		// 				Operation: to.Ptr("Get/List gitHubConnector resources"),
		// 				Provider: to.Ptr("Microsoft.SecurityDevOps"),
		// 				Resource: to.Ptr("gitHubConnectors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityDevOps/gitHubConnectors/Write"),
		// 			Display: &armsecuritydevops.OperationDisplay{
		// 				Description: to.Ptr("Write gitHubConnector"),
		// 				Operation: to.Ptr("Create/Update gitHubConnector resources"),
		// 				Provider: to.Ptr("Microsoft.SecurityDevOps"),
		// 				Resource: to.Ptr("gitHubConnectors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityDevOps/gitHubConnectors/Delete"),
		// 			Display: &armsecuritydevops.OperationDisplay{
		// 				Description: to.Ptr("Delete gitHubConnector"),
		// 				Operation: to.Ptr("Delete gitHubConnector resources"),
		// 				Provider: to.Ptr("Microsoft.SecurityDevOps"),
		// 				Resource: to.Ptr("gitHubConnectors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
