Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybridnetwork%2Farmhybridnetwork%2Fv0.4.0/sdk/resourcemanager/hybridnetwork/armhybridnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armhybridnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/NetworkFunctionsExecuteRequest.json
func ExampleNetworkFunctionsClient_BeginExecuteRequest() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhybridnetwork.NewNetworkFunctionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginExecuteRequest(ctx,
		"<resource-group-name>",
		"<network-function-name>",
		armhybridnetwork.ExecuteRequestParameters{
			RequestMetadata: &armhybridnetwork.RequestMetadata{
				APIVersion:     to.Ptr("<apiversion>"),
				HTTPMethod:     to.Ptr(armhybridnetwork.HTTPMethodPost),
				RelativePath:   to.Ptr("<relative-path>"),
				SerializedBody: to.Ptr("<serialized-body>"),
			},
			ServiceEndpoint: to.Ptr("<service-endpoint>"),
		},
		&armhybridnetwork.NetworkFunctionsClientBeginExecuteRequestOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
```
