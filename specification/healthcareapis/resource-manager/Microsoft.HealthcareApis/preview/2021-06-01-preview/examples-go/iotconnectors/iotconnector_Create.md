Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhealthcareapis%2Farmhealthcareapis%2Fv0.2.1/sdk/resourcemanager/healthcareapis/armhealthcareapis/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2021-06-01-preview/examples/iotconnectors/iotconnector_Create.json
func ExampleIotConnectorsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhealthcareapis.NewIotConnectorsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<iot-connector-name>",
		armhealthcareapis.IotConnector{
			Identity: &armhealthcareapis.ServiceManagedIdentityIdentity{
				Type: armhealthcareapis.ManagedServiceIdentityType("SystemAssigned").ToPtr(),
			},
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"additionalProp1": to.StringPtr("string"),
				"additionalProp2": to.StringPtr("string"),
				"additionalProp3": to.StringPtr("string"),
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
					ConsumerGroup:                   to.StringPtr("<consumer-group>"),
					EventHubName:                    to.StringPtr("<event-hub-name>"),
					FullyQualifiedEventHubNamespace: to.StringPtr("<fully-qualified-event-hub-namespace>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IotConnectorsClientCreateOrUpdateResult)
}
```
