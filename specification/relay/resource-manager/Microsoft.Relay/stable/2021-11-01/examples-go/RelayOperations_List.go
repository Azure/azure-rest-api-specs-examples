package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/RelayOperations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armrelay.OperationListResult{
		// 	Value: []*armrelay.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Relay/checkNameAvailability/action"),
		// 			Display: &armrelay.OperationDisplay{
		// 				Operation: to.Ptr("Get namespace availability."),
		// 				Provider: to.Ptr("Microsoft Azure Relay"),
		// 				Resource: to.Ptr("Non Resource Operation"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Relay/register/action"),
		// 			Display: &armrelay.OperationDisplay{
		// 				Operation: to.Ptr("Registers the Relay Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Azure Relay"),
		// 				Resource: to.Ptr("Relay Resource Provider"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Relay/namespaces/write"),
		// 			Display: &armrelay.OperationDisplay{
		// 				Operation: to.Ptr("Create Or Update Namespace "),
		// 				Provider: to.Ptr("Microsoft Azure Relay"),
		// 				Resource: to.Ptr("Namespace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Relay/namespaces/read"),
		// 			Display: &armrelay.OperationDisplay{
		// 				Operation: to.Ptr("Get Namespace Resource"),
		// 				Provider: to.Ptr("Microsoft Azure Relay"),
		// 				Resource: to.Ptr("Namespace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Relay/namespaces/Delete"),
		// 			Display: &armrelay.OperationDisplay{
		// 				Operation: to.Ptr("Delete Namespace"),
		// 				Provider: to.Ptr("Microsoft Azure Relay"),
		// 				Resource: to.Ptr("Namespace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Relay/namespaces/Relays/write"),
		// 			Display: &armrelay.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Relay"),
		// 				Provider: to.Ptr("Microsoft Azure Relay"),
		// 				Resource: to.Ptr("Relay"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Relay/namespaces/Relays/read"),
		// 			Display: &armrelay.OperationDisplay{
		// 				Operation: to.Ptr("Get Relay"),
		// 				Provider: to.Ptr("Microsoft Azure Relay"),
		// 				Resource: to.Ptr("Relay"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Relay/namespaces/Relays/Delete"),
		// 			Display: &armrelay.OperationDisplay{
		// 				Operation: to.Ptr("Delete Relay"),
		// 				Provider: to.Ptr("Microsoft Azure Relay"),
		// 				Resource: to.Ptr("Relay"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Relay/namespaces/metricDefinitions/read"),
		// 			Display: &armrelay.OperationDisplay{
		// 				Operation: to.Ptr("Get Namespace metrics"),
		// 				Provider: to.Ptr("Microsoft Azure Relay"),
		// 				Resource: to.Ptr("Namespace metrics"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Relay/namespaces/diagnosticSettings/read"),
		// 			Display: &armrelay.OperationDisplay{
		// 				Operation: to.Ptr("Get Namespace diagnostic settings"),
		// 				Provider: to.Ptr("Microsoft Azure Relay"),
		// 				Resource: to.Ptr("Namespace diagnostic settings"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Relay/namespaces/diagnosticSettings/write"),
		// 			Display: &armrelay.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Namespace diagnostic settings"),
		// 				Provider: to.Ptr("Microsoft Azure Relay"),
		// 				Resource: to.Ptr("Namespace diagnostic settings"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Relay/namespaces/logDefinitions/read"),
		// 			Display: &armrelay.OperationDisplay{
		// 				Operation: to.Ptr("Get Namespace logs"),
		// 				Provider: to.Ptr("Microsoft Azure Relay"),
		// 				Resource: to.Ptr("Namespace logs"),
		// 			},
		// 	}},
		// }
	}
}
