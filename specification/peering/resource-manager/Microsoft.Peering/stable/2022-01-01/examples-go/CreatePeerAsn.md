```go
package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/CreatePeerAsn.json
func ExamplePeerAsnsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpeering.NewPeerAsnsClient("subId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"peerAsnName",
		armpeering.PeerAsn{
			Properties: &armpeering.PeerAsnProperties{
				PeerAsn: to.Ptr[int32](65000),
				PeerContactDetail: []*armpeering.ContactDetail{
					{
						Email: to.Ptr("noc@contoso.com"),
						Phone: to.Ptr("+1 (234) 567-8999"),
						Role:  to.Ptr(armpeering.RoleNoc),
					},
					{
						Email: to.Ptr("abc@contoso.com"),
						Phone: to.Ptr("+1 (234) 567-8900"),
						Role:  to.Ptr(armpeering.RolePolicy),
					},
					{
						Email: to.Ptr("xyz@contoso.com"),
						Phone: to.Ptr("+1 (234) 567-8900"),
						Role:  to.Ptr(armpeering.RoleTechnical),
					}},
				PeerName: to.Ptr("Contoso"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpeering%2Farmpeering%2Fv1.0.0/sdk/resourcemanager/peering/armpeering/README.md) on how to add the SDK to your project and authenticate.
