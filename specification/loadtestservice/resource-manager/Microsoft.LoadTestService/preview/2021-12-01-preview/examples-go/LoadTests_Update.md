Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Floadtestservice%2Farmloadtestservice%2Fv0.1.0/sdk/resourcemanager/loadtestservice/armloadtestservice/README.md) on how to add the SDK to your project and authenticate.

```go
package armloadtestservice_test

import (
	"context"
	"log"

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
		armloadtestservice.LoadTestResourcePatchRequestBody{},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("LoadTestResource.ID: %s\n", *res.ID)
}
```
