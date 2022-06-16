package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/iotconnectors/iotconnector_fhirdestination_Create.json
func ExampleIotConnectorFhirDestinationClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhealthcareapis.NewIotConnectorFhirDestinationClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"testRG",
		"workspace1",
		"blue",
		"dest1",
		armhealthcareapis.IotFhirDestination{
			Location: to.Ptr("westus"),
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
				FhirServiceResourceID:          to.Ptr("subscriptions/11111111-2222-3333-4444-555566667777/resourceGroups/myrg/providers/Microsoft.HealthcareApis/workspaces/myworkspace/fhirservices/myfhirservice"),
				ResourceIdentityResolutionType: to.Ptr(armhealthcareapis.IotIdentityResolutionTypeCreate),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
