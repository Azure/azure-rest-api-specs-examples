package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Origins_Create.json
func ExampleOriginsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcdn.NewOriginsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"RG",
		"profile1",
		"endpoint1",
		"www-someDomain-net",
		armcdn.Origin{
			Properties: &armcdn.OriginProperties{
				Enabled:                    to.Ptr(true),
				HostName:                   to.Ptr("www.someDomain.net"),
				HTTPPort:                   to.Ptr[int32](80),
				HTTPSPort:                  to.Ptr[int32](443),
				OriginHostHeader:           to.Ptr("www.someDomain.net"),
				Priority:                   to.Ptr[int32](1),
				PrivateLinkApprovalMessage: to.Ptr("Please approve the connection request for this Private Link"),
				PrivateLinkLocation:        to.Ptr("eastus"),
				PrivateLinkResourceID:      to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.Network/privateLinkServices/pls1"),
				Weight:                     to.Ptr[int32](50),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
