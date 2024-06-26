package armnotificationhubs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/85cfba195a19120f309bd292c4261aa53a586adb/specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/NHOperationsList.json
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
		// 				Description: to.Ptr("Registers the subscription for the NotifciationHubs resource provider and enables the creation of Namespaces and NotificationHubs"),
		// 				Operation: to.Ptr("Registers the NotificationHubs Provider"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Microsoft Azure Notification Hub"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/unregister/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Unregisters the subscription for the NotifciationHubs resource provider and enables the creation of Namespaces and NotificationHubs"),
		// 				Operation: to.Ptr("Unregisters the NotificationHubs Provider"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Microsoft Azure Notification Hub"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/operations/read"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Returns a list of supported operations for Notification Hubs provider"),
		// 				Operation: to.Ptr("Notification Hubs provider operations"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Microsoft Azure Notification Hub"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/operationResults/read"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Returns operation results for Notification Hubs provider"),
		// 				Operation: to.Ptr("Operation results for Notification Hubs provider"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Microsoft Azure Notification Hub"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Get Namespace diagnostic settings"),
		// 				Operation: to.Ptr("Read diagnostics setting"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Namespace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Create or Update Namespace diagnostic settings"),
		// 				Operation: to.Ptr("Write diagnostic settings"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Namespace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/providers/Microsoft.Insights/logDefinitions/read"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Gets the available logs for Namespace"),
		// 				Operation: to.Ptr("Read Namespace log definitions"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("The log definition of Namespace"),
		// 			},
		// 			Properties: &armnotificationhubs.OperationProperties{
		// 				ServiceSpecification: &armnotificationhubs.ServiceSpecification{
		// 					LogSpecifications: []*armnotificationhubs.LogSpecification{
		// 						{
		// 							Name: to.Ptr("OperationalLogs"),
		// 							BlobDuration: to.Ptr("PT1H"),
		// 							DisplayName: to.Ptr("Operational Logs"),
		// 					}},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/CheckNamespaceAvailability/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Checks whether or not a given Namespace resource name is available within the NotificationHub service."),
		// 				Operation: to.Ptr("Check Namespace name availability."),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Namespace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/write"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Create a Namespace Resource and Update its properties. Tags and Capacity of the Namespace are the properties which can be updated."),
		// 				Operation: to.Ptr("Create Or Update Namespace"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Namespace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/read"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Get the list of Namespace Resource Description"),
		// 				Operation: to.Ptr("Get Namespace Resource"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Namespace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/delete"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Delete Namespace Resource"),
		// 				Operation: to.Ptr("Delete Namespace"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Namespace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/authorizationRules/write"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Create a Namespace level Authorization Rules and update its properties. The Authorization Rules Access Rights, the Primary and Secondary Keys can be updated."),
		// 				Operation: to.Ptr("Create or Update Namespace Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Namespace Authorization Rule"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/authorizationRules/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Get the list of Namespaces Authorization Rules description."),
		// 				Operation: to.Ptr("Get Namespace Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Namespace Authorization Rule"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/authorizationRules/read"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Get the list of Namespaces Authorization Rules description."),
		// 				Operation: to.Ptr("Get Namespace Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Namespace Authorization Rule"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/authorizationRules/delete"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Delete Namespace Authorization Rule. The Default Namespace Authorization Rule cannot be deleted."),
		// 				Operation: to.Ptr("Delete Namespace Authorization Rule"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Namespace Authorization Rule"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/authorizationRules/listkeys/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Get the Connection String to the Namespace"),
		// 				Operation: to.Ptr("Get Namespace Listkeys"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Namespace Authorization Rule"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/authorizationRules/regenerateKeys/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Namespace Authorization Rule Regenerate Primary/SecondaryKey, Specify the Key that needs to be regenerated"),
		// 				Operation: to.Ptr("Namespace Authorization Rule Regenerate Keys"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Namespace Authorization Rule"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/CheckNotificationHubAvailability/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Checks whether or not a given NotificationHub name is available inside a Namespace."),
		// 				Operation: to.Ptr("Check NotificationHub name availability."),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Namespace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/NotificationHubs/write"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Create a Notification Hub and Update its properties. Its properties mainly include PNS Credentials. Authorization Rules and TTL"),
		// 				Operation: to.Ptr("Create or Update Notification Hub"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("NotificationHub"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/NotificationHubs/vapidkeys/read"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Get new pair of VAPID keys for a Notification Hub"),
		// 				Operation: to.Ptr("Get VAPID keys for a Notification Hub"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("NotificationHub"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/NotificationHubs/read"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Get list of Notification Hub Resource Descriptions"),
		// 				Operation: to.Ptr("Get Notification Hub"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("NotificationHub"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/NotificationHubs/delete"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Delete Notification Hub Resource"),
		// 				Operation: to.Ptr("Delete Notification Hub"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("NotificationHub"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/NotificationHubs/authorizationRules/write"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Create Notification Hub Authorization Rules and Update its properties. The Authorization Rules Access Rights, the Primary and Secondary Keys can be updated."),
		// 				Operation: to.Ptr("Create or Update Notification hub Authorization Rule"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("NotificationHub Authorization Rule"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/NotificationHubs/authorizationRules/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Get the list of Notification Hub Authorization Rules"),
		// 				Operation: to.Ptr("Get Notification Hub Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("NotificationHub Authorization Rule"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/NotificationHubs/authorizationRules/read"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Get the list of Notification Hub Authorization Rules"),
		// 				Operation: to.Ptr("Get Notification Hub Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("NotificationHub Authorization Rule"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/NotificationHubs/authorizationRules/delete"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Delete Notification Hub Authorization Rules"),
		// 				Operation: to.Ptr("Delete Notification Hub Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("NotificationHub Authorization Rule"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/NotificationHubs/authorizationRules/listkeys/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Get the Connection String to the Notification Hub"),
		// 				Operation: to.Ptr("Get Notification Hub Listkeys"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("NotificationHub Authorization Rule"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/NotificationHubs/authorizationRules/regenerateKeys/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Notification Hub Authorization Rule Regenerate Primary/SecondaryKey, Specify the Key that needs to be regenerated"),
		// 				Operation: to.Ptr("Notification Hub Authorization Rule Regenerate Keys"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("NotificationHub Authorization Rule"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/NotificationHubs/pnsCredentials/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Get All Notification Hub PNS Credentials. This includes, WNS, MPNS, APNS, GCM, Baidu and FcmV1 credentials"),
		// 				Operation: to.Ptr("Get Notification Hub PNS Credentials"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("NotificationHub PnsCredential"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/NotificationHubs/debugSend/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Send a test push notification to 10 matched devices."),
		// 				Operation: to.Ptr("Send a test push notification."),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("NotificationHub"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/Namespaces/NotificationHubs/metricDefinitions/read"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Get list of Namespace metrics Resource Descriptions"),
		// 				Operation: to.Ptr("Get Namespace metrics"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Namespace metrics"),
		// 			},
		// 			Properties: &armnotificationhubs.OperationProperties{
		// 				ServiceSpecification: &armnotificationhubs.ServiceSpecification{
		// 					MetricSpecifications: []*armnotificationhubs.MetricSpecification{
		// 						{
		// 							Name: to.Ptr("registration.all"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of all successful registration operations (creations updates queries and deletions). "),
		// 							DisplayName: to.Ptr("Registration Operations"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("registration.create"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of all successful registration creations."),
		// 							DisplayName: to.Ptr("Registration Create Operations"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("registration.update"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of all successful registration updates."),
		// 							DisplayName: to.Ptr("Registration Update Operations"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("registration.get"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of all successful registration queries."),
		// 							DisplayName: to.Ptr("Registration Read Operations"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("registration.delete"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of all successful registration deletions."),
		// 							DisplayName: to.Ptr("Registration Delete Operations"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("incoming"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of all successful send API calls. "),
		// 							DisplayName: to.Ptr("Incoming Messages"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("incoming.scheduled"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Scheduled Push Notifications Sent"),
		// 							DisplayName: to.Ptr("Scheduled Push Notifications Sent"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("incoming.scheduled.cancel"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Scheduled Push Notifications Cancelled"),
		// 							DisplayName: to.Ptr("Scheduled Push Notifications Cancelled"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("scheduled.pending"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Pending Scheduled Notifications"),
		// 							DisplayName: to.Ptr("Pending Scheduled Notifications"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("installation.all"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Installation Management Operations"),
		// 							DisplayName: to.Ptr("Installation Management Operations"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("installation.get"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Get Installation Operations"),
		// 							DisplayName: to.Ptr("Get Installation Operations"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("installation.upsert"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Create or Update Installation Operations"),
		// 							DisplayName: to.Ptr("Create or Update Installation Operations"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("installation.patch"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Patch Installation Operations"),
		// 							DisplayName: to.Ptr("Patch Installation Operations"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("installation.delete"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Delete Installation Operations"),
		// 							DisplayName: to.Ptr("Delete Installation Operations"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.allpns.success"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of all successful notifications."),
		// 							DisplayName: to.Ptr("Successful notifications"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.allpns.invalidpayload"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the PNS returned a bad payload error."),
		// 							DisplayName: to.Ptr("Payload Errors"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.allpns.pnserror"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because there was a problem communicating with the PNS (excludes authentication problems)."),
		// 							DisplayName: to.Ptr("External Notification System Errors"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.allpns.channelerror"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the channel was invalid not associated with the correct app throttled or expired."),
		// 							DisplayName: to.Ptr("Channel Errors"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.allpns.badorexpiredchannel"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the channel/token/registrationId in the registration was expired or invalid."),
		// 							DisplayName: to.Ptr("Bad or Expired Channel Errors"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.wns.success"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of all successful notifications."),
		// 							DisplayName: to.Ptr("WNS Successful Notifications"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.wns.invalidcredentials"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the PNS did not accept the provided credentials or the credentials are blocked. (Windows Live does not recognize the credentials)."),
		// 							DisplayName: to.Ptr("WNS Authorization Errors (Invalid Credentials)"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.wns.badchannel"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the ChannelURI in the registration was not recognized (WNS status: 404 not found)."),
		// 							DisplayName: to.Ptr("WNS Bad Channel Error"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.wns.expiredchannel"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the ChannelURI is expired (WNS status: 410 Gone)."),
		// 							DisplayName: to.Ptr("WNS Expired Channel Error"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.wns.throttled"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because WNS is throttling this app (WNS status: 406 Not Acceptable)."),
		// 							DisplayName: to.Ptr("WNS Throttled Notifications"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.wns.tokenproviderunreachable"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Windows Live is not reachable."),
		// 							DisplayName: to.Ptr("WNS Authorization Errors (Unreachable)"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.wns.invalidtoken"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The token provided to WNS is not valid (WNS status: 401 Unauthorized)."),
		// 							DisplayName: to.Ptr("WNS Authorization Errors (Invalid Token)"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.wns.wrongtoken"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The token provided to WNS is valid but for another application (WNS status: 403 Forbidden). This can happen if the ChannelURI in the registration is associated with another app. Check that the client app is associated with the same app whose credentials are in the notification hub."),
		// 							DisplayName: to.Ptr("WNS Authorization Errors (Wrong Token)"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.wns.invalidnotificationformat"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The format of the notification is invalid (WNS status: 400). Note that WNS does not reject all invalid payloads."),
		// 							DisplayName: to.Ptr("WNS Invalid Notification Format"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.wns.invalidnotificationsize"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The notification payload is too large (WNS status: 413)."),
		// 							DisplayName: to.Ptr("WNS Invalid Notification Size Error"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.wns.channelthrottled"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The notification was dropped because the ChannelURI in the registration is throttled (WNS response header: X-WNS-NotificationStatus:channelThrottled)."),
		// 							DisplayName: to.Ptr("WNS Channel Throttled"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.wns.channeldisconnected"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The notification was dropped because the ChannelURI in the registration is throttled (WNS response header: X-WNS-DeviceConnectionStatus: disconnected)."),
		// 							DisplayName: to.Ptr("WNS Channel Disconnected"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.wns.dropped"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The notification was dropped because the ChannelURI in the registration is throttled (X-WNS-NotificationStatus: dropped but not X-WNS-DeviceConnectionStatus: disconnected)."),
		// 							DisplayName: to.Ptr("WNS Dropped Notifications"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.wns.pnserror"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Notification not delivered because of errors communicating with WNS."),
		// 							DisplayName: to.Ptr("WNS Errors"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.wns.authenticationerror"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Notification not delivered because of errors communicating with Windows Live invalid credentials or wrong token."),
		// 							DisplayName: to.Ptr("WNS Authentication Errors"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.apns.success"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of all successful notifications."),
		// 							DisplayName: to.Ptr("APNS Successful Notifications"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.apns.invalidcredentials"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the PNS did not accept the provided credentials or the credentials are blocked."),
		// 							DisplayName: to.Ptr("APNS Authorization Errors"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.apns.badchannel"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the token is invalid (APNS status code: 8)."),
		// 							DisplayName: to.Ptr("APNS Bad Channel Error"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.apns.expiredchannel"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of token that were invalidated by the APNS feedback channel."),
		// 							DisplayName: to.Ptr("APNS Expired Channel Error"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.apns.invalidnotificationsize"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the payload was too large (APNS status code: 7)."),
		// 							DisplayName: to.Ptr("APNS Invalid Notification Size Error"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.apns.pnserror"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because of errors communicating with APNS."),
		// 							DisplayName: to.Ptr("APNS Errors"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.gcm.success"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of all successful notifications."),
		// 							DisplayName: to.Ptr("GCM Successful Notifications"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.gcm.invalidcredentials"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the PNS did not accept the provided credentials or the credentials are blocked."),
		// 							DisplayName: to.Ptr("GCM Authorization Errors (Invalid Credentials)"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.gcm.badchannel"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the registrationId in the registration was not recognized (GCM result: Invalid Registration)."),
		// 							DisplayName: to.Ptr("GCM Bad Channel Error"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.gcm.expiredchannel"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the registrationId in the registration was expired (GCM result: NotRegistered)."),
		// 							DisplayName: to.Ptr("GCM Expired Channel Error"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.gcm.throttled"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because GCM throttled this app (GCM status code: 501-599 or result:Unavailable)."),
		// 							DisplayName: to.Ptr("GCM Throttled Notifications"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.gcm.invalidnotificationformat"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the payload was not formatted correctly (GCM result: InvalidDataKey or InvalidTtl)."),
		// 							DisplayName: to.Ptr("GCM Invalid Notification Format"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.gcm.invalidnotificationsize"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the payload was too large (GCM result: MessageTooBig)."),
		// 							DisplayName: to.Ptr("GCM Invalid Notification Size Error"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.gcm.wrongchannel"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the registrationId in the registration is not associated to the current app (GCM result: InvalidPackageName)."),
		// 							DisplayName: to.Ptr("GCM Wrong Channel Error"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.gcm.pnserror"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because of errors communicating with GCM."),
		// 							DisplayName: to.Ptr("GCM Errors"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.gcm.authenticationerror"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the PNS did not accept the provided credentials the credentials are blocked or the SenderId is not correctly configured in the app (GCM result: MismatchedSenderId)."),
		// 							DisplayName: to.Ptr("GCM Authentication Errors"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.mpns.success"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of all successful notifications."),
		// 							DisplayName: to.Ptr("MPNS Successful Notifications"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.mpns.invalidcredentials"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the PNS did not accept the provided credentials or the credentials are blocked."),
		// 							DisplayName: to.Ptr("MPNS Invalid Credentials"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.mpns.badchannel"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the ChannelURI in the registration was not recognized (MPNS status: 404 not found)."),
		// 							DisplayName: to.Ptr("MPNS Bad Channel Error"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.mpns.throttled"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because MPNS is throttling this app (WNS MPNS: 406 Not Acceptable)."),
		// 							DisplayName: to.Ptr("MPNS Throttled Notifications"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.mpns.invalidnotificationformat"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the payload of the notification was too large."),
		// 							DisplayName: to.Ptr("MPNS Invalid Notification Format"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.mpns.channeldisconnected"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the ChannelURI in the registration was disconnected (MPNS status: 412 not found)."),
		// 							DisplayName: to.Ptr("MPNS Channel Disconnected"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.mpns.dropped"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that were dropped by MPNS (MPNS response header: X-NotificationStatus: QueueFull or Suppressed)."),
		// 							DisplayName: to.Ptr("MPNS Dropped Notifications"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.mpns.pnserror"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because of errors communicating with MPNS."),
		// 							DisplayName: to.Ptr("MPNS Errors"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("outgoing.mpns.authenticationerror"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("The count of pushes that failed because the PNS did not accept the provided credentials or the credentials are blocked."),
		// 							DisplayName: to.Ptr("MPNS Authentication Errors"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("notificationhub.pushes"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("All outgoing notifications of the notification hub"),
		// 							DisplayName: to.Ptr("All Outgoing Notifications"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("incoming.all.requests"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Total incoming requests for a notification hub"),
		// 							DisplayName: to.Ptr("All Incoming Requests"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("incoming.all.failedrequests"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Availabilities: []*armnotificationhubs.Availability{
		// 								{
		// 									BlobDuration: to.Ptr("P30D"),
		// 									TimeGrain: to.Ptr("PT1M"),
		// 							}},
		// 							DisplayDescription: to.Ptr("Total incoming failed requests for a notification hub"),
		// 							DisplayName: to.Ptr("All Incoming Failed Requests"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 					}},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/privateEndpointConnectionsApproval/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Approve Private Endpoint Connection"),
		// 				Operation: to.Ptr("Approve Private Endpoint Connection"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Namespace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/privateEndpointConnectionProxies/validate/action"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Validate Private Endpoint Connection Proxy"),
		// 				Operation: to.Ptr("Validate Private Endpoint Connection Proxy"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/privateEndpointConnectionProxies/read"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Get Private Endpoint Connection Proxy"),
		// 				Operation: to.Ptr("Get Private Endpoint Connection Proxy"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/privateEndpointConnectionProxies/write"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Create Private Endpoint Connection Proxy"),
		// 				Operation: to.Ptr("Create Private Endpoint Connection Proxy"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/privateEndpointConnectionProxies/delete"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Delete Private Endpoint Connection Proxy"),
		// 				Operation: to.Ptr("Delete Private Endpoint Connection Proxy"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/privateEndpointConnectionProxies/operationstatus/read"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Get the status of an asynchronous private endpoint operation"),
		// 				Operation: to.Ptr("Private endpoint operation status"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/privateEndpointConnections/read"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Get Private Endpoint Connection"),
		// 				Operation: to.Ptr("Get Private Endpoint Connection"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Private Endpoint Connection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/privateEndpointConnections/write"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Create or Update Private Endpoint Connection"),
		// 				Operation: to.Ptr("Create or Update Private Endpoint Connection"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Private Endpoint Connection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/privateEndpointConnections/delete"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Removes Private Endpoint Connection"),
		// 				Operation: to.Ptr("Removes Private Endpoint Connection"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Private Endpoint Connection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NotificationHubs/namespaces/privateEndpointConnections/operationstatus/read"),
		// 			Display: &armnotificationhubs.OperationDisplay{
		// 				Description: to.Ptr("Removes Private Endpoint Connection"),
		// 				Operation: to.Ptr("Removes Private Endpoint Connection"),
		// 				Provider: to.Ptr("Microsoft Azure Notification Hub"),
		// 				Resource: to.Ptr("Private Endpoint Connection"),
		// 			},
		// 	}},
		// }
	}
}
