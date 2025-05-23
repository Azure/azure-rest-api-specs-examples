package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5759c77eee2d57bdb9e47aa1805d0ffb61704f2d/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2024-05-01-preview/examples/EHOperations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armeventhub.OperationListResult{
		// 	Value: []*armeventhub.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/checkNameAvailability/action"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Get namespace availability."),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("Non Resource Operation"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/register/action"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Registers the EventHub Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("EventHub Resource Provider"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/write"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Create Or Update Namespace "),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("Namespace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/read"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Get Namespace Resource"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("Namespace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/Delete"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Delete Namespace"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("Namespace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/authorizationRules/read"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Get Namespace Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/authorizationRules/write"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Namespace Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/authorizationRules/delete"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Delete Namespace Authorization Rule"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/authorizationRules/listkeys/action"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Get Namespace Listkeys"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/authorizationRules/regenerateKeys/action"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Resource Regeneratekeys"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/eventhubs/write"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update EventHub"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("EventHub"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/eventhubs/read"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Get EventHub"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("EventHub"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/eventhubs/Delete"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Delete EventHub"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("EventHub"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/eventhubs/authorizationRules/read"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr(" Get EventHub Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("EventHub AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/eventhubs/authorizationRules/write"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update EventHub Authorization Rule"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("EventHub AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/eventhubs/authorizationRules/delete"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Delete EventHub Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("EventHub AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/eventhubs/authorizationRules/listkeys/action"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("List EventHub keys"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("EventHub AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/eventhubs/authorizationRules/regenerateKeys/action"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Resource Regeneratekeys"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("EventHub AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/eventHubs/consumergroups/write"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update ConsumerGroup"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("ConsumerGroup"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/eventHubs/consumergroups/read"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Get ConsumerGroup"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("ConsumerGroup"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/eventHubs/consumergroups/Delete"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Delete ConsumerGroup"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("ConsumerGroup"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/metricDefinitions/read"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Get Namespace metrics"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("Namespace metrics"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/diagnosticSettings/read"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Get Namespace diagnostic settings"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("Namespace diagnostic settings"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/diagnosticSettings/write"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Namespace diagnostic settings"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("Namespace diagnostic settings"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EventHub/namespaces/logDefinitions/read"),
		// 			Display: &armeventhub.OperationDisplay{
		// 				Operation: to.Ptr("Get Namespace logs"),
		// 				Provider: to.Ptr("Microsoft Azure EventHub"),
		// 				Resource: to.Ptr("Namespace logs"),
		// 			},
		// 	}},
		// }
	}
}
