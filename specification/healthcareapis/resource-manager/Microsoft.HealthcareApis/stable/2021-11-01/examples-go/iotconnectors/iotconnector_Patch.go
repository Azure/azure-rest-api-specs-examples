package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/iotconnectors/iotconnector_Patch.json
func ExampleIotConnectorsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIotConnectorsClient().BeginUpdate(ctx, "testRG", "blue", "workspace1", armhealthcareapis.IotConnectorPatchResource{
		Tags: map[string]*string{
			"additionalProp1": to.Ptr("string"),
			"additionalProp2": to.Ptr("string"),
			"additionalProp3": to.Ptr("string"),
		},
		Identity: &armhealthcareapis.ServiceManagedIdentityIdentity{
			Type: to.Ptr(armhealthcareapis.ServiceManagedIdentityTypeSystemAssigned),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IotConnector = armhealthcareapis.IotConnector{
	// 	Identity: &armhealthcareapis.ServiceManagedIdentityIdentity{
	// 		Type: to.Ptr(armhealthcareapis.ServiceManagedIdentityTypeSystemAssigned),
	// 	},
	// 	Name: to.Ptr("blue"),
	// 	Type: to.Ptr("Microsoft.HealthcareApis/workspaces/iotconnectors"),
	// 	Etag: to.Ptr("00000000-0000-0000-f5ac-912ca49e01d6"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/testRG/providers/Microsoft.HealthcareApis/workspaces/workspace1/iotconnectors/blue"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"additionalProp1": to.Ptr("string"),
	// 		"additionalProp2": to.Ptr("string"),
	// 		"additionalProp3": to.Ptr("string"),
	// 	},
	// 	Properties: &armhealthcareapis.IotConnectorProperties{
	// 		DeviceMapping: &armhealthcareapis.IotMappingProperties{
	// 			Content: map[string]any{
	// 				"template":[]any{
	// 					map[string]any{
	// 						"template":map[string]any{
	// 							"deviceIdExpression": "$.deviceid",
	// 							"timestampExpression": "$.measurementdatetime",
	// 							"typeMatchExpression": "$..[?(@heartrate)]",
	// 							"typeName": "heartrate",
	// 							"values":[]any{
	// 								map[string]any{
	// 									"required": "true",
	// 									"valueExpression": "$.heartrate",
	// 									"valueName": "hr",
	// 								},
	// 							},
	// 						},
	// 						"templateType": "JsonPathContent",
	// 					},
	// 				},
	// 				"templateType": "CollectionContent",
	// 			},
	// 		},
	// 		IngestionEndpointConfiguration: &armhealthcareapis.IotEventHubIngestionEndpointConfiguration{
	// 			ConsumerGroup: to.Ptr("ConsumerGroupA"),
	// 			EventHubName: to.Ptr("MyEventHubName"),
	// 			FullyQualifiedEventHubNamespace: to.Ptr("myeventhub.servicesbus.windows.net"),
	// 		},
	// 		ProvisioningState: to.Ptr(armhealthcareapis.ProvisioningStateSucceeded),
	// 	},
	// 	SystemData: &armhealthcareapis.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-28T19:26:24.072Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-28T19:26:24.072Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 	},
	// }
}
