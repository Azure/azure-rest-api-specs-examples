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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/iotconnectors/iotconnector_fhirdestination_Create.json
func ExampleIotConnectorFhirDestinationClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhealthcareapis.NewIotConnectorFhirDestinationClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<iot-connector-name>",
		"<fhir-destination-name>",
		armhealthcareapis.IotFhirDestination{
			Location: to.Ptr("<location>"),
			Properties: &armhealthcareapis.IotFhirDestinationProperties{
				FhirMapping: &armhealthcareapis.IotMappingProperties{
					Content: map[string]interface{}{
						"template": []interface{}{
							map[string]interface{}{
								"template": map[string]interface{}{
									"codes": []interface{}{
										map[string]interface{}{
											"code":    "8867-4",
											"display": "Heart rate",
											"system":  "http://loinc.org",
										},
									},
									"periodInterval": float64(60),
									"typeName":       "heartrate",
									"value": map[string]interface{}{
										"defaultPeriod": float64(5000),
										"unit":          "count/min",
										"valueName":     "hr",
										"valueType":     "SampledData",
									},
								},
								"templateType": "CodeValueFhir",
							},
						},
						"templateType": "CollectionFhirTemplate",
					},
				},
				FhirServiceResourceID:          to.Ptr("<fhir-service-resource-id>"),
				ResourceIdentityResolutionType: to.Ptr(armhealthcareapis.IotIdentityResolutionTypeCreate),
			},
		},
		&armhealthcareapis.IotConnectorFhirDestinationClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
