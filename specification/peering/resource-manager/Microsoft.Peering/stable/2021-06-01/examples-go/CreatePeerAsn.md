```go
package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/CreatePeerAsn.json
func ExamplePeerAsnsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpeering.NewPeerAsnsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<peer-asn-name>",
		armpeering.PeerAsn{
			Properties: &armpeering.PeerAsnProperties{
				PeerAsn: to.Int32Ptr(65000),
				PeerContactDetail: []*armpeering.ContactDetail{
					{
						Email: to.StringPtr("<email>"),
						Phone: to.StringPtr("<phone>"),
						Role:  armpeering.Role("Noc").ToPtr(),
					},
					{
						Email: to.StringPtr("<email>"),
						Phone: to.StringPtr("<phone>"),
						Role:  armpeering.Role("Policy").ToPtr(),
					},
					{
						Email: to.StringPtr("<email>"),
						Phone: to.StringPtr("<phone>"),
						Role:  armpeering.Role("Technical").ToPtr(),
					}},
				PeerName: to.StringPtr("<peer-name>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PeerAsnsClientCreateOrUpdateResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpeering%2Farmpeering%2Fv0.3.0/sdk/resourcemanager/peering/armpeering/README.md) on how to add the SDK to your project and authenticate.
