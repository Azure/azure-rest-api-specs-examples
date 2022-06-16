package armloadtestservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtestservice/armloadtestservice"
)

// x-ms-original-file: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/preview/2021-12-01-preview/examples/LoadTests_Update.json
func ExampleLoadTestsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armloadtestservice.NewLoadTestsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<load-test-name>",
		armloadtestservice.LoadTestResourcePatchRequestBody{
			Properties: &armloadtestservice.LoadTestResourcePatchRequestBodyProperties{
				Description: to.StringPtr("<description>"),
			},
			Tags: map[string]interface{}{
				"Division": "LT",
				"Team":     "Dev Exp",
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LoadTestsClientUpdateResult)
}
