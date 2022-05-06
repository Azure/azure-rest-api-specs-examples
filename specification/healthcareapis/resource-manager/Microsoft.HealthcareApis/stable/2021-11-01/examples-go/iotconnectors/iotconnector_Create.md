Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhealthcareapis%2Farmhealthcareapis%2Fv0.4.0/sdk/resourcemanager/healthcareapis/armhealthcareapis/README.md) on how to add the SDK to your project and authenticate.

```go
package armhealthcareapis_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/iotconnectors/iotconnector_Create.json
func ExampleIotConnectorsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhealthcareapis.NewIotConnectorsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<iot-connector-name>",
		armhealthcareapis.IotConnector{
			Identity: &armhealthcareapis.ServiceManagedIdentityIdentity{
				Type: to.Ptr(armhealthcareapis.ServiceManagedIdentityTypeSystemAssigned),
			},
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"additionalProp1": to.Ptr("string"),
				"additionalProp2": to.Ptr("string"),
				"additionalProp3": to.Ptr("string"),
			},
			Properties: &armhealthcareapis.IotConnectorProperties{
				DeviceMapping: &armhealthcareapis.IotMappingProperties{
					Content: map[string]interface{}{
						"template": []interface{}{
							map[string]interface{}{
								"template": map[string]interface{}{
									"deviceIdExpression":  "$.deviceid",
									"timestampExpression": "$.measurementdatetime",
									"typeMatchExpression": "$..[?(@heartrate)]",
									"typeName":            "heartrate",
									"values": []interface{}{
										map[string]interface{}{
											"required":        "true",
											"valueExpression": "$.heartrate",
											"valueName":       "hr",
										},
									},
								},
								"templateType": "JsonPathContent",
							},
						},
						"templateType": "CollectionContent",
					},
				},
				IngestionEndpointConfiguration: &armhealthcareapis.IotEventHubIngestionEndpointConfiguration{
					ConsumerGroup:                   to.Ptr("<consumer-group>"),
					EventHubName:                    to.Ptr("<event-hub-name>"),
					FullyQualifiedEventHubNamespace: to.Ptr("<fully-qualified-event-hub-namespace>"),
				},
			},
		},
		&armhealthcareapis.IotConnectorsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
