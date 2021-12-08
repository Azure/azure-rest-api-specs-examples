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

// x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/CreateOrUpdateZone.json
func ExampleZonesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdns.NewZonesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<zone-name>",
		armdns.Zone{
			Resource: armdns.Resource{
				Location: to.StringPtr("<location>"),
				Tags: map[string]*string{
					"key1": to.StringPtr("value1"),
				},
			},
		},
		&armdns.ZonesCreateOrUpdateOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Zone.ID: %s\n", *res.ID)
}
```
