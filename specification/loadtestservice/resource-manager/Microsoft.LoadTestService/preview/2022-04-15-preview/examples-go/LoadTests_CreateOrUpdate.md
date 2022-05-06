Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Floadtestservice%2Farmloadtestservice%2Fv0.4.0/sdk/resourcemanager/loadtestservice/armloadtestservice/README.md) on how to add the SDK to your project and authenticate.

```go
package armloadtestservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtestservice/armloadtestservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/loadtestservice/resource-manager/Microsoft.LoadTestService/preview/2022-04-15-preview/examples/LoadTests_CreateOrUpdate.json
func ExampleLoadTestsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armloadtestservice.NewLoadTestsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<load-test-name>",
		armloadtestservice.LoadTestResource{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"Team": to.Ptr("Dev Exp"),
			},
			Properties: &armloadtestservice.LoadTestProperties{
				Description: to.Ptr("<description>"),
				Encryption: &armloadtestservice.EncryptionProperties{
					Identity: &armloadtestservice.EncryptionPropertiesIdentity{
						Type:       to.Ptr(armloadtestservice.TypeUserAssigned),
						ResourceID: to.Ptr("<resource-id>"),
					},
					KeyURL: to.Ptr("<key-url>"),
				},
			},
		},
		&armloadtestservice.LoadTestsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
