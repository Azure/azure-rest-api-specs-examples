package armhealthcareapis_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
)

// x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2021-06-01-preview/examples/fhirservices/FhirServices_Patch.json
func ExampleFhirServicesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhealthcareapis.NewFhirServicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<fhir-service-name>",
		"<workspace-name>",
		armhealthcareapis.FhirServicePatchResource{
			Tags: map[string]*string{
				"tagKey": to.StringPtr("tagValue"),
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
	log.Printf("Response result: %#v\n", res.FhirServicesClientUpdateResult)
}
