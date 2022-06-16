package armsignalr_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
)

// x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalRSharedPrivateLinkResources_CreateOrUpdate.json
func ExampleSharedPrivateLinkResourcesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsignalr.NewSharedPrivateLinkResourcesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<shared-private-link-resource-name>",
		"<resource-group-name>",
		"<resource-name>",
		armsignalr.SharedPrivateLinkResource{
			Properties: &armsignalr.SharedPrivateLinkResourceProperties{
				GroupID:               to.StringPtr("<group-id>"),
				PrivateLinkResourceID: to.StringPtr("<private-link-resource-id>"),
				RequestMessage:        to.StringPtr("<request-message>"),
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
	log.Printf("Response result: %#v\n", res.SharedPrivateLinkResourcesClientCreateOrUpdateResult)
}
