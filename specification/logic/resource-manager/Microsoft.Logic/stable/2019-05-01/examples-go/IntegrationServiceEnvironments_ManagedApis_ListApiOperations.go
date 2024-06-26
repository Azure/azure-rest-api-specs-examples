package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_ManagedApis_ListApiOperations.json
func ExampleIntegrationServiceEnvironmentManagedAPIOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIntegrationServiceEnvironmentManagedAPIOperationsClient().NewListPager("testResourceGroup", "testIntegrationServiceEnvironment", "servicebus", nil)
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
		// page.APIOperationListResult = armlogic.APIOperationListResult{
		// 	Value: []*armlogic.APIOperation{
		// 		{
		// 			Name: to.Ptr("SendMessage"),
		// 			ID: to.Ptr("/subscriptions/80d4fe69-c95b-4dd2-a938-9250f1c8ab03/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment/managedApis/servicebus/apiOperations/SendMessage"),
		// 			Location: to.Ptr("brazilsouth"),
		// 			Properties: &armlogic.APIOperationPropertiesDefinition{
		// 				Description: to.Ptr("This operation sends a message to a queue or topic."),
		// 				Annotation: &armlogic.APIOperationAnnotation{
		// 					Family: to.Ptr("SendMessage"),
		// 					Revision: to.Ptr[int32](1),
		// 					Status: to.Ptr(armlogic.StatusAnnotationProduction),
		// 				},
		// 				API: &armlogic.APIReference{
		// 					Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments/managedApis"),
		// 					ID: to.Ptr("/subscriptions/80d4fe69-c95b-4dd2-a938-9250f1c8ab03/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment/managedApis/servicebus"),
		// 					Description: to.Ptr("Connect to Azure Service Bus to send and receive messages. You can perform actions such as send to queue, send to topic, receive from queue, receive from subscription, etc."),
		// 					BrandColor: to.Ptr("#c4d5ff"),
		// 					Category: to.Ptr(armlogic.APITierStandard),
		// 					DisplayName: to.Ptr("Service Bus"),
		// 					IconURI: to.Ptr("https://powerappsconnectorsdf.blob.core.windows.net/officialicons/servicebus/icon_1.0.1216.1605.png"),
		// 					IntegrationServiceEnvironment: &armlogic.ResourceReference{
		// 						Name: to.Ptr("testIntegrationServiceEnvironment"),
		// 						Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments"),
		// 						ID: to.Ptr("/subscriptions/80d4fe69-c95b-4dd2-a938-9250f1c8ab03/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment"),
		// 					},
		// 				},
		// 				IsNotification: to.Ptr(false),
		// 				IsWebhook: to.Ptr(false),
		// 				Pageable: to.Ptr(false),
		// 				Summary: to.Ptr("Send message"),
		// 				Visibility: to.Ptr("important"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SendMessages"),
		// 			ID: to.Ptr("/subscriptions/80d4fe69-c95b-4dd2-a938-9250f1c8ab03/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment/managedApis/servicebus/apiOperations/SendMessages"),
		// 			Location: to.Ptr("brazilsouth"),
		// 			Properties: &armlogic.APIOperationPropertiesDefinition{
		// 				Description: to.Ptr("This operation sends one or more messages to a queue or topic."),
		// 				Annotation: &armlogic.APIOperationAnnotation{
		// 					Family: to.Ptr("SendMessages"),
		// 					Revision: to.Ptr[int32](1),
		// 					Status: to.Ptr(armlogic.StatusAnnotationProduction),
		// 				},
		// 				API: &armlogic.APIReference{
		// 					Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments/managedApis"),
		// 					ID: to.Ptr("/subscriptions/80d4fe69-c95b-4dd2-a938-9250f1c8ab03/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment/managedApis/servicebus"),
		// 					Description: to.Ptr("Connect to Azure Service Bus to send and receive messages. You can perform actions such as send to queue, send to topic, receive from queue, receive from subscription, etc."),
		// 					BrandColor: to.Ptr("#c4d5ff"),
		// 					Category: to.Ptr(armlogic.APITierStandard),
		// 					DisplayName: to.Ptr("Service Bus"),
		// 					IconURI: to.Ptr("https://powerappsconnectorsdf.blob.core.windows.net/officialicons/servicebus/icon_1.0.1216.1605.png"),
		// 					IntegrationServiceEnvironment: &armlogic.ResourceReference{
		// 						Name: to.Ptr("testIntegrationServiceEnvironment"),
		// 						Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments"),
		// 						ID: to.Ptr("/subscriptions/80d4fe69-c95b-4dd2-a938-9250f1c8ab03/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment"),
		// 					},
		// 				},
		// 				IsNotification: to.Ptr(false),
		// 				IsWebhook: to.Ptr(false),
		// 				Pageable: to.Ptr(false),
		// 				Summary: to.Ptr("Send one or more messages"),
		// 				Visibility: to.Ptr("important"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("GetMessageFromQueue"),
		// 			ID: to.Ptr("/subscriptions/80d4fe69-c95b-4dd2-a938-9250f1c8ab03/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment/managedApis/servicebus/apiOperations/GetMessageFromQueue"),
		// 			Location: to.Ptr("brazilsouth"),
		// 			Properties: &armlogic.APIOperationPropertiesDefinition{
		// 				Description: to.Ptr("This operation triggers a flow when a message is received in a queue and auto completes the message."),
		// 				Annotation: &armlogic.APIOperationAnnotation{
		// 					Family: to.Ptr("GetMessageFromQueue"),
		// 					Revision: to.Ptr[int32](1),
		// 					Status: to.Ptr(armlogic.StatusAnnotationProduction),
		// 				},
		// 				API: &armlogic.APIReference{
		// 					Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments/managedApis"),
		// 					ID: to.Ptr("/subscriptions/80d4fe69-c95b-4dd2-a938-9250f1c8ab03/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment/managedApis/servicebus"),
		// 					Description: to.Ptr("Connect to Azure Service Bus to send and receive messages. You can perform actions such as send to queue, send to topic, receive from queue, receive from subscription, etc."),
		// 					BrandColor: to.Ptr("#c4d5ff"),
		// 					Category: to.Ptr(armlogic.APITierStandard),
		// 					DisplayName: to.Ptr("Service Bus"),
		// 					IconURI: to.Ptr("https://powerappsconnectorsdf.blob.core.windows.net/officialicons/servicebus/icon_1.0.1216.1605.png"),
		// 					IntegrationServiceEnvironment: &armlogic.ResourceReference{
		// 						Name: to.Ptr("testIntegrationServiceEnvironment"),
		// 						Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments"),
		// 						ID: to.Ptr("/subscriptions/80d4fe69-c95b-4dd2-a938-9250f1c8ab03/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment"),
		// 					},
		// 				},
		// 				IsNotification: to.Ptr(false),
		// 				IsWebhook: to.Ptr(false),
		// 				Pageable: to.Ptr(false),
		// 				Summary: to.Ptr("When a message is received in a queue (auto-complete)"),
		// 				Trigger: to.Ptr("single"),
		// 				Visibility: to.Ptr("important"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("GetNewMessageFromQueueWithPeekLock"),
		// 			ID: to.Ptr("/subscriptions/80d4fe69-c95b-4dd2-a938-9250f1c8ab03/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment/managedApis/servicebus/apiOperations/GetNewMessageFromQueueWithPeekLock"),
		// 			Location: to.Ptr("brazilsouth"),
		// 			Properties: &armlogic.APIOperationPropertiesDefinition{
		// 				Description: to.Ptr("The operation triggers a flow when a message received in a queue with peek-lock mode."),
		// 				Annotation: &armlogic.APIOperationAnnotation{
		// 					Family: to.Ptr("GetNewMessageFromQueueWithPeekLock"),
		// 					Revision: to.Ptr[int32](1),
		// 					Status: to.Ptr(armlogic.StatusAnnotationProduction),
		// 				},
		// 				API: &armlogic.APIReference{
		// 					Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments/managedApis"),
		// 					ID: to.Ptr("/subscriptions/80d4fe69-c95b-4dd2-a938-9250f1c8ab03/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment/managedApis/servicebus"),
		// 					Description: to.Ptr("Connect to Azure Service Bus to send and receive messages. You can perform actions such as send to queue, send to topic, receive from queue, receive from subscription, etc."),
		// 					BrandColor: to.Ptr("#c4d5ff"),
		// 					Category: to.Ptr(armlogic.APITierStandard),
		// 					DisplayName: to.Ptr("Service Bus"),
		// 					IconURI: to.Ptr("https://powerappsconnectorsdf.blob.core.windows.net/officialicons/servicebus/icon_1.0.1216.1605.png"),
		// 					IntegrationServiceEnvironment: &armlogic.ResourceReference{
		// 						Name: to.Ptr("testIntegrationServiceEnvironment"),
		// 						Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments"),
		// 						ID: to.Ptr("/subscriptions/80d4fe69-c95b-4dd2-a938-9250f1c8ab03/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment"),
		// 					},
		// 				},
		// 				IsNotification: to.Ptr(false),
		// 				IsWebhook: to.Ptr(false),
		// 				Pageable: to.Ptr(false),
		// 				Summary: to.Ptr("When a message is received in a queue (peek-lock)"),
		// 				Trigger: to.Ptr("single"),
		// 				Visibility: to.Ptr("important"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("GetMessageFromTopic"),
		// 			ID: to.Ptr("/subscriptions/80d4fe69-c95b-4dd2-a938-9250f1c8ab03/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment/managedApis/servicebus/apiOperations/GetMessageFromTopic"),
		// 			Location: to.Ptr("brazilsouth"),
		// 			Properties: &armlogic.APIOperationPropertiesDefinition{
		// 				Description: to.Ptr("This operation triggers a flow when a message is received in a topic subscription and auto completes the message."),
		// 				Annotation: &armlogic.APIOperationAnnotation{
		// 					Family: to.Ptr("GetMessageFromTopic"),
		// 					Revision: to.Ptr[int32](1),
		// 					Status: to.Ptr(armlogic.StatusAnnotationProduction),
		// 				},
		// 				API: &armlogic.APIReference{
		// 					Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments/managedApis"),
		// 					ID: to.Ptr("/subscriptions/80d4fe69-c95b-4dd2-a938-9250f1c8ab03/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment/managedApis/servicebus"),
		// 					Description: to.Ptr("Connect to Azure Service Bus to send and receive messages. You can perform actions such as send to queue, send to topic, receive from queue, receive from subscription, etc."),
		// 					BrandColor: to.Ptr("#c4d5ff"),
		// 					Category: to.Ptr(armlogic.APITierStandard),
		// 					DisplayName: to.Ptr("Service Bus"),
		// 					IconURI: to.Ptr("https://powerappsconnectorsdf.blob.core.windows.net/officialicons/servicebus/icon_1.0.1216.1605.png"),
		// 					IntegrationServiceEnvironment: &armlogic.ResourceReference{
		// 						Name: to.Ptr("testIntegrationServiceEnvironment"),
		// 						Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments"),
		// 						ID: to.Ptr("/subscriptions/80d4fe69-c95b-4dd2-a938-9250f1c8ab03/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment"),
		// 					},
		// 				},
		// 				IsNotification: to.Ptr(false),
		// 				IsWebhook: to.Ptr(false),
		// 				Pageable: to.Ptr(false),
		// 				Summary: to.Ptr("When a message is received in a topic subscription (auto-complete)"),
		// 				Trigger: to.Ptr("single"),
		// 				Visibility: to.Ptr("important"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("GetNewMessageFromTopicWithPeekLock"),
		// 			ID: to.Ptr("/subscriptions/80d4fe69-c95b-4dd2-a938-9250f1c8ab03/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment/managedApis/servicebus/apiOperations/GetNewMessageFromTopicWithPeekLock"),
		// 			Location: to.Ptr("brazilsouth"),
		// 			Properties: &armlogic.APIOperationPropertiesDefinition{
		// 				Description: to.Ptr("The operation triggers a flow when a message received in a topic subscription with peek-lock mode."),
		// 				Annotation: &armlogic.APIOperationAnnotation{
		// 					Family: to.Ptr("GetNewMessageFromTopicWithPeekLock"),
		// 					Revision: to.Ptr[int32](1),
		// 					Status: to.Ptr(armlogic.StatusAnnotationProduction),
		// 				},
		// 				API: &armlogic.APIReference{
		// 					Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments/managedApis"),
		// 					ID: to.Ptr("/subscriptions/80d4fe69-c95b-4dd2-a938-9250f1c8ab03/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment/managedApis/servicebus"),
		// 					Description: to.Ptr("Connect to Azure Service Bus to send and receive messages. You can perform actions such as send to queue, send to topic, receive from queue, receive from subscription, etc."),
		// 					BrandColor: to.Ptr("#c4d5ff"),
		// 					Category: to.Ptr(armlogic.APITierStandard),
		// 					DisplayName: to.Ptr("Service Bus"),
		// 					IconURI: to.Ptr("https://powerappsconnectorsdf.blob.core.windows.net/officialicons/servicebus/icon_1.0.1216.1605.png"),
		// 					IntegrationServiceEnvironment: &armlogic.ResourceReference{
		// 						Name: to.Ptr("testIntegrationServiceEnvironment"),
		// 						Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments"),
		// 						ID: to.Ptr("/subscriptions/80d4fe69-c95b-4dd2-a938-9250f1c8ab03/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment"),
		// 					},
		// 				},
		// 				IsNotification: to.Ptr(false),
		// 				IsWebhook: to.Ptr(false),
		// 				Pageable: to.Ptr(false),
		// 				Summary: to.Ptr("When a message is received in a topic subscription (peek-lock)"),
		// 				Trigger: to.Ptr("single"),
		// 				Visibility: to.Ptr("important"),
		// 			},
		// 	}},
		// }
	}
}
