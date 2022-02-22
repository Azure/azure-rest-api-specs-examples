Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.3.1/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ApplicationGatewayBackendHealthTest.json
func ExampleApplicationGatewaysClient_BeginBackendHealthOnDemand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewApplicationGatewaysClient("<subscription-id>", cred, nil)
	poller, err := client.BeginBackendHealthOnDemand(ctx,
		"<resource-group-name>",
		"<application-gateway-name>",
		armnetwork.ApplicationGatewayOnDemandProbe{
			Path: to.StringPtr("<path>"),
			BackendAddressPool: &armnetwork.SubResource{
				ID: to.StringPtr("<id>"),
			},
			BackendHTTPSettings: &armnetwork.SubResource{
				ID: to.StringPtr("<id>"),
			},
			PickHostNameFromBackendHTTPSettings: to.BoolPtr(true),
			Timeout:                             to.Int32Ptr(30),
			Protocol:                            armnetwork.ApplicationGatewayProtocol("Http").ToPtr(),
		},
		&armnetwork.ApplicationGatewaysClientBeginBackendHealthOnDemandOptions{Expand: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ApplicationGatewaysClientBackendHealthOnDemandResult)
}
```
