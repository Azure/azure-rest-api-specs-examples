Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fprivatedns%2Farmprivatedns%2Fv0.3.1/sdk/resourcemanager/privatedns/armprivatedns/README.md) on how to add the SDK to your project and authenticate.

```go
package armprivatedns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
)

// x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetAPut.json
func ExampleRecordSetsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armprivatedns.NewRecordSetsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<private-zone-name>",
		armprivatedns.RecordTypeA,
		"<relative-record-set-name>",
		armprivatedns.RecordSet{
			Properties: &armprivatedns.RecordSetProperties{
				ARecords: []*armprivatedns.ARecord{
					{
						IPv4Address: to.StringPtr("<ipv4address>"),
					}},
				Metadata: map[string]*string{
					"key1": to.StringPtr("value1"),
				},
				TTL: to.Int64Ptr(3600),
			},
		},
		&armprivatedns.RecordSetsClientCreateOrUpdateOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RecordSetsClientCreateOrUpdateResult)
}
```
