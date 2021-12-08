Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdns%2Farmdns%2Fv0.1.0/sdk/resourcemanager/dns/armdns/README.md) on how to add the SDK to your project and authenticate.

```go
package armdns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
)

// x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/GetDnsResourceReference.json
func ExampleDNSResourceReferenceClient_GetByTargetResources() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdns.NewDNSResourceReferenceClient("<subscription-id>", cred, nil)
	_, err = client.GetByTargetResources(ctx,
		armdns.DNSResourceReferenceRequest{
			Properties: &armdns.DNSResourceReferenceRequestProperties{
				TargetResources: []*armdns.SubResource{
					{
						ID: to.StringPtr("<id>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
