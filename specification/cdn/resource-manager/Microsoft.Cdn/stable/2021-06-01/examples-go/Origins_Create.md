Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcdn%2Farmcdn%2Fv0.5.0/sdk/resourcemanager/cdn/armcdn/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Origins_Create.json
func ExampleOriginsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcdn.NewOriginsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<endpoint-name>",
		"<origin-name>",
		armcdn.Origin{
			Properties: &armcdn.OriginProperties{
				Enabled:                    to.Ptr(true),
				HostName:                   to.Ptr("<host-name>"),
				HTTPPort:                   to.Ptr[int32](80),
				HTTPSPort:                  to.Ptr[int32](443),
				OriginHostHeader:           to.Ptr("<origin-host-header>"),
				Priority:                   to.Ptr[int32](1),
				PrivateLinkApprovalMessage: to.Ptr("<private-link-approval-message>"),
				PrivateLinkLocation:        to.Ptr("<private-link-location>"),
				PrivateLinkResourceID:      to.Ptr("<private-link-resource-id>"),
				Weight:                     to.Ptr[int32](50),
			},
		},
		&armcdn.OriginsClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
