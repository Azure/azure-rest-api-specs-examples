Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmediaservices%2Farmmediaservices%2Fv0.4.0/sdk/resourcemanager/mediaservices/armmediaservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streaming-policies-create-clear.json
func ExampleStreamingPoliciesClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewStreamingPoliciesClient("<subscription-id>", cred, nil)
	_, err = client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<streaming-policy-name>",
		armmediaservices.StreamingPolicy{
			Properties: &armmediaservices.StreamingPolicyProperties{
				NoEncryption: &armmediaservices.NoEncryption{
					EnabledProtocols: &armmediaservices.EnabledProtocols{
						Dash:            to.BoolPtr(true),
						Download:        to.BoolPtr(true),
						Hls:             to.BoolPtr(true),
						SmoothStreaming: to.BoolPtr(true),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
