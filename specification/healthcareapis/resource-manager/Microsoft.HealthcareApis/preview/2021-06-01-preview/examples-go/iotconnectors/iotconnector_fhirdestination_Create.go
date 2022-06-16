package armhealthcareapis_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
)

// x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2021-06-01-preview/examples/iotconnectors/iotconnector_fhirdestination_Create.json
func ExampleIotConnectorFhirDestinationClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhealthcareapis.NewIotConnectorFhirDestinationClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<iot-connector-name>",
		"<fhir-destination-name>",
		armhealthcareapis.IotFhirDestination{
			Location: to.StringPtr("<location>"),
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
				FhirServiceResourceID:          to.StringPtr("<fhir-service-resource-id>"),
				ResourceIdentityResolutionType: armhealthcareapis.IotIdentityResolutionType("Create").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.IotConnectorFhirDestinationClientCreateOrUpdateResult)
}
