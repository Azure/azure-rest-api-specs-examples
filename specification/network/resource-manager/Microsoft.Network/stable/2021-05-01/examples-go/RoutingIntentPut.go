package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/RoutingIntentPut.json
func ExampleRoutingIntentClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewRoutingIntentClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<virtual-hub-name>",
		"<routing-intent-name>",
		armnetwork.RoutingIntent{
			Properties: &armnetwork.RoutingIntentProperties{
				RoutingPolicies: []*armnetwork.RoutingPolicy{
					{
						Name: to.Ptr("<name>"),
						Destinations: []*string{
							to.Ptr("Internet")},
						NextHop: to.Ptr("<next-hop>"),
					},
					{
						Name: to.Ptr("<name>"),
						Destinations: []*string{
							to.Ptr("PrivateTraffic")},
						NextHop: to.Ptr("<next-hop>"),
					}},
			},
		},
		&armnetwork.RoutingIntentClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
