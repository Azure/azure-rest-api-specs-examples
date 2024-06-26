package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/iotconnectors/iotconnector_fhirdestination_Create.json
func ExampleIotConnectorFhirDestinationClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIotConnectorFhirDestinationClient().BeginCreateOrUpdate(ctx, "testRG", "workspace1", "blue", "dest1", armhealthcareapis.IotFhirDestination{
		Location: to.Ptr("westus"),
		Properties: &armhealthcareapis.IotFhirDestinationProperties{
			FhirMapping: &armhealthcareapis.IotMappingProperties{
				Content: map[string]any{
					"template": []any{
						map[string]any{
							"template": map[string]any{
								"codes": []any{
									map[string]any{
										"code":    "8867-4",
										"display": "Heart rate",
										"system":  "http://loinc.org",
									},
								},
								"periodInterval": float64(60),
								"typeName":       "heartrate",
								"value": map[string]any{
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
	// res.IotFhirDestination = armhealthcareapis.IotFhirDestination{
	// 	Name: to.Ptr("dest1"),
	// 	Type: to.Ptr("Microsoft.HealthcareApis/workspaces/iotconnectors/fhirdestinations"),
	// 	Etag: to.Ptr("00000000-0000-0000-f5ac-912ca49e01d6"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/testRG/providers/Microsoft.HealthcareApis/workspaces/workspace1/iotconnectors/blue/fhirdestinations/dest1"),
	// 	Location: to.Ptr("West US 2"),
	// 	Properties: &armhealthcareapis.IotFhirDestinationProperties{
	// 		ProvisioningState: to.Ptr(armhealthcareapis.ProvisioningStateSucceeded),
	// 		FhirMapping: &armhealthcareapis.IotMappingProperties{
	// 			Content: map[string]any{
	// 				"template":[]any{
	// 					map[string]any{
	// 						"template":map[string]any{
	// 							"codes":[]any{
	// 								map[string]any{
	// 									"code": "8867-4",
	// 									"display": "Heart rate",
	// 									"system": "http://loinc.org",
	// 								},
	// 							},
	// 							"periodInterval": float64(60),
	// 							"typeName": "heartrate",
	// 							"value":map[string]any{
	// 								"defaultPeriod": float64(5000),
	// 								"unit": "count/min",
	// 								"valueName": "hr",
	// 								"valueType": "SampledData",
	// 							},
	// 						},
	// 						"templateType": "CodeValueFhir",
	// 					},
	// 				},
	// 				"templateType": "CollectionFhirTemplate",
	// 			},
	// 		},
	// 		FhirServiceResourceID: to.Ptr("subscriptions/11111111-2222-3333-4444-555566667777/resourceGroups/myrg/providers/Microsoft.HealthcareApis/workspaces/myworkspace/fhirservices/myfhirservice"),
	// 		ResourceIdentityResolutionType: to.Ptr(armhealthcareapis.IotIdentityResolutionTypeCreate),
	// 	},
	// 	SystemData: &armhealthcareapis.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-28T19:26:24.072Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-28T19:26:24.072Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 	},
	// }
}
