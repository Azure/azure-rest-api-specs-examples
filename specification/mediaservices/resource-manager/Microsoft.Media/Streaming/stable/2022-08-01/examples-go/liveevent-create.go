package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveevent-create.json
func ExampleLiveEventsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewLiveEventsClient("0a6ec948-5a62-437d-b9df-934dc7c1b722", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "mediaresources", "slitestmedia10", "myLiveEvent1", armmediaservices.LiveEvent{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
		Properties: &armmediaservices.LiveEventProperties{
			Description: to.Ptr("test event 1"),
			Input: &armmediaservices.LiveEventInput{
				AccessControl: &armmediaservices.LiveEventInputAccessControl{
					IP: &armmediaservices.IPAccessControl{
						Allow: []*armmediaservices.IPRange{
							{
								Name:               to.Ptr("AllowAll"),
								Address:            to.Ptr("0.0.0.0"),
								SubnetPrefixLength: to.Ptr[int32](0),
							}},
					},
				},
				KeyFrameIntervalDuration: to.Ptr("PT6S"),
				StreamingProtocol:        to.Ptr(armmediaservices.LiveEventInputProtocolRTMP),
			},
			Preview: &armmediaservices.LiveEventPreview{
				AccessControl: &armmediaservices.LiveEventPreviewAccessControl{
					IP: &armmediaservices.IPAccessControl{
						Allow: []*armmediaservices.IPRange{
							{
								Name:               to.Ptr("AllowAll"),
								Address:            to.Ptr("0.0.0.0"),
								SubnetPrefixLength: to.Ptr[int32](0),
							}},
					},
				},
			},
		},
	}, &armmediaservices.LiveEventsClientBeginCreateOptions{AutoStart: nil})
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
