package armiotsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotsecurity/armiotsecurity"
)

// x-ms-original-file: specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-07-01-preview/examples/RecommendationTypes/GetRecommendationType.json
func ExampleRecommendationTypesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotsecurity.NewRecommendationTypesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<recommendation-type-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RecommendationTypesClientGetResult)
}
