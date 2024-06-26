package armnotificationhubs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NHOperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnotificationhubs.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armnotificationhubs.OperationListResult{
		// 	Value: []*armnotificationhubs.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/register/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Registers the NotificationHubs Provider"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Microsoft Azure Notification Hub"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/unregister/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Unregisters the NotificationHubs Provider"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Microsoft Azure Notification Hub"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/operationResults/read"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Operation results for Notification Hubs provider"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Microsoft Azure Notification Hub"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/CheckNamespaceAvailability/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Get namespace availability."),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Microsoft Azure Notification Hub"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/write"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Create Or Update Namespace "),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Namespace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/read"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Get Namespace Resource"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("Namespace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/Delete"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Delete Namespace"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("Namespace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/authorizationRules/write"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Namespace Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/authorizationRules/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Get Namespace Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/authorizationRules/read"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Get Namespace Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/authorizationRules/delete"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Delete Namespace Authorization Rule"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/authorizationRules/listkeys/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Get Namespace Listkeys"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/authorizationRules/regenerateKeys/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Resource Regeneratekeys"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/CheckNotificationHubAvailability/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("CheckNotificationHubAvailability"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/notificationHubs/write"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update notification hub"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("NotificationHub"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/notificationHubs/read"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Get notification hub"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("NotificationHub"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/notificationHubs/Delete"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Delete notification hub"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("NotificationHub"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/notificationHubs/authorizationRules/write"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update NotificationHub Authorization Rule"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("NotificationHub Authorization Rule"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/notificationHubs/authorizationRules/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr(" Get NotificationHub Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("NotificationHub AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/notificationHubs/authorizationRules/read"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr(" Get NotificationHub Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("NotificationHub AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/notificationHubs/authorizationRules/delete"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Delete NotificationHub Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("NotificationHub AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/notificationHubs/authorizationRules/listkeys/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("List NotificationHub keys"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("NotificationHub AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/notificationHubs/authorizationRules/regenerateKeys/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Resource Regeneratekeys"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("NotificationHub AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/NotificationHubs/pnsCredentials/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Resource Get Notification Hub PNS Credentials"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("NotificationHub PnsCredential"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/NotificationHubs/debugSend/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Send a test push notification"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("NotificationHub resource"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/NotificationHubs/metricDefinitions/read"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Operation: to.Ptr("Get NotificationHub metrics"),
		// 				Provider: to.Ptr("Microsoft Azure NotificationHubs"),
		// 				Resource: to.Ptr("NotificationHub metrics"),
		// 			},
		// 	}},
		// }
	}
}
