package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2018-05-01-preview/examples/ComponentsPurgeStatus.json
func ExampleComponentsClient_GetPurgeStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewComponentsClient("<subscription-id>", cred, nil)
	res, err := client.GetPurgeStatus(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<purge-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ComponentsClientGetPurgeStatusResult)
}
