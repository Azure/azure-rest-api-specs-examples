package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/iotconnectors/iotconnector_List.json
func ExampleIotConnectorsClient_NewListByWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIotConnectorsClient().NewListByWorkspacePager("testRG", "workspace1", nil)
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
		// page.IotConnectorCollection = armhealthcareapis.IotConnectorCollection{
		// 	Value: []*armhealthcareapis.IotConnector{
		// 		{
		// 			Identity: &armhealthcareapis.ServiceManagedIdentityIdentity{
		// 				Type: to.Ptr(armhealthcareapis.ServiceManagedIdentityTypeSystemAssigned),
		// 			},
		// 			Name: to.Ptr("blue"),
		// 			Type: to.Ptr("Microsoft.HealthcareApis/workspaces/iotconnectors"),
		// 			Etag: to.Ptr("00000000-0000-0000-f5ac-912ca49e01d6"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testRG/providers/Microsoft.HealthcareApis/workspaces/workspace1/iotconnectors/blue"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"additionalProp1": to.Ptr("string"),
		// 				"additionalProp2": to.Ptr("string"),
		// 				"additionalProp3": to.Ptr("string"),
		// 			},
		// 			Properties: &armhealthcareapis.IotConnectorProperties{
		// 				DeviceMapping: &armhealthcareapis.IotMappingProperties{
		// 					Content: map[string]any{
		// 						"template":[]any{
		// 							map[string]any{
		// 								"template":map[string]any{
		// 									"deviceIdExpression": "$.deviceid",
		// 									"timestampExpression": "$.measurementdatetime",
		// 									"typeMatchExpression": "$..[?(@heartrate)]",
		// 									"typeName": "heartrate",
		// 									"values":[]any{
		// 										map[string]any{
		// 											"required": "true",
		// 											"valueExpression": "$.heartrate",
		// 											"valueName": "hr",
		// 										},
		// 									},
		// 								},
		// 								"templateType": "JsonPathContent",
		// 							},
		// 						},
		// 						"templateType": "CollectionContent",
		// 					},
		// 				},
		// 				IngestionEndpointConfiguration: &armhealthcareapis.IotEventHubIngestionEndpointConfiguration{
		// 					ConsumerGroup: to.Ptr("ConsumerGroupA"),
		// 					EventHubName: to.Ptr("MyEventHubName"),
		// 					FullyQualifiedEventHubNamespace: to.Ptr("myeventhub.servicesbus.windows.net"),
		// 				},
		// 				ProvisioningState: to.Ptr(armhealthcareapis.ProvisioningStateSucceeded),
		// 			},
		// 			SystemData: &armhealthcareapis.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-28T19:26:24.072Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-28T19:26:24.072Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 			},
		// 		},
		// 		{
		// 			Identity: &armhealthcareapis.ServiceManagedIdentityIdentity{
		// 				Type: to.Ptr(armhealthcareapis.ServiceManagedIdentityTypeSystemAssigned),
		// 			},
		// 			Name: to.Ptr("red"),
		// 			Type: to.Ptr("Microsoft.HealthcareApis/workspaces/iotconnectors"),
		// 			Etag: to.Ptr("00000000-0000-0000-f6ac-912ca49e01d6"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testRG/providers/Microsoft.HealthcareApis/workspaces/workspace1/iotconnectors/red"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"additionalProp1": to.Ptr("string"),
		// 				"additionalProp2": to.Ptr("string"),
		// 				"additionalProp3": to.Ptr("string"),
		// 			},
		// 			Properties: &armhealthcareapis.IotConnectorProperties{
		// 				DeviceMapping: &armhealthcareapis.IotMappingProperties{
		// 					Content: map[string]any{
		// 						"template":[]any{
		// 							map[string]any{
		// 								"template":map[string]any{
		// 									"deviceIdExpression": "$.deviceid",
		// 									"timestampExpression": "$.measurementdatetime",
		// 									"typeMatchExpression": "$..[?(@steps)]",
		// 									"typeName": "steps",
		// 									"values":[]any{
		// 										map[string]any{
		// 											"required": "true",
		// 											"valueExpression": "$.steps",
		// 											"valueName": "steps",
		// 										},
		// 									},
		// 								},
		// 								"templateType": "JsonPathContent",
		// 							},
		// 						},
		// 						"templateType": "CollectionContent",
		// 					},
		// 				},
		// 				IngestionEndpointConfiguration: &armhealthcareapis.IotEventHubIngestionEndpointConfiguration{
		// 					ConsumerGroup: to.Ptr("ConsumerGroupA"),
		// 					EventHubName: to.Ptr("MyEventHubName"),
		// 					FullyQualifiedEventHubNamespace: to.Ptr("myeventhub.servicesbus.windows.net"),
		// 				},
		// 				ProvisioningState: to.Ptr(armhealthcareapis.ProvisioningStateSucceeded),
		// 			},
		// 			SystemData: &armhealthcareapis.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-27T19:26:24.072Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-27T19:26:24.072Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 			},
		// 	}},
		// }
	}
}
