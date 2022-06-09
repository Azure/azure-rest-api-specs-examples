```go
package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/NetworkFunctionsExecuteRequest.json
func ExampleNetworkFunctionsClient_BeginExecuteRequest() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhybridnetwork.NewNetworkFunctionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginExecuteRequest(ctx,
		"rg",
		"testNetworkfunction",
		armhybridnetwork.ExecuteRequestParameters{
			RequestMetadata: &armhybridnetwork.RequestMetadata{
				APIVersion:     to.Ptr("apiVersionQueryString"),
				HTTPMethod:     to.Ptr(armhybridnetwork.HTTPMethodPost),
				RelativePath:   to.Ptr("/simProfiles/testSimProfile"),
				SerializedBody: to.Ptr("{\"subscriptionProfile\":\"ChantestSubscription15\",\"permanentKey\":\"00112233445566778899AABBCCDDEEFF\",\"opcOperatorCode\":\"63bfa50ee6523365ff14c1f45f88737d\",\"staticIpAddresses\":{\"internet\":{\"ipv4Addr\":\"198.51.100.1\",\"ipv6Prefix\":\"2001:db8:abcd:12::0/64\"},\"another_network\":{\"ipv6Prefix\":\"2001:111:cdef:22::0/64\"}}}"),
			},
			ServiceEndpoint: to.Ptr("serviceEndpoint"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybridnetwork%2Farmhybridnetwork%2Fv2.0.0-beta.1/sdk/resourcemanager/hybridnetwork/armhybridnetwork/README.md) on how to add the SDK to your project and authenticate.
