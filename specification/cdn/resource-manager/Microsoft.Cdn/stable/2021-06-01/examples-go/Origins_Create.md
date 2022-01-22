Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcdn%2Farmcdn%2Fv0.3.0/sdk/resourcemanager/cdn/armcdn/README.md) on how to add the SDK to your project and authenticate.

```go
package armcdn_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Origins_Create.json
func ExampleOriginsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewOriginsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<endpoint-name>",
		"<origin-name>",
		armcdn.Origin{
			Properties: &armcdn.OriginProperties{
				Enabled:                    to.BoolPtr(true),
				HostName:                   to.StringPtr("<host-name>"),
				HTTPPort:                   to.Int32Ptr(80),
				HTTPSPort:                  to.Int32Ptr(443),
				OriginHostHeader:           to.StringPtr("<origin-host-header>"),
				Priority:                   to.Int32Ptr(1),
				PrivateLinkApprovalMessage: to.StringPtr("<private-link-approval-message>"),
				PrivateLinkLocation:        to.StringPtr("<private-link-location>"),
				PrivateLinkResourceID:      to.StringPtr("<private-link-resource-id>"),
				Weight:                     to.Int32Ptr(50),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.OriginsClientCreateResult)
}
```
