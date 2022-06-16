package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
)

// x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2021-06-01-preview/examples/OperationResultsGet.json
func ExampleOperationResultsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhealthcareapis.NewOperationResultsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<location-name>",
		"<operation-result-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.OperationResultsClientGetResult)
}
