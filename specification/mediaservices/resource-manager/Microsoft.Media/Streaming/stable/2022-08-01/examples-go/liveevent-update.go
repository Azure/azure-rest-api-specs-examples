package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveevent-update.json
func ExampleLiveEventsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLiveEventsClient().BeginUpdate(ctx, "mediaresources", "slitestmedia10", "myLiveEvent1", armmediaservices.LiveEvent{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
			"tag3": to.Ptr("value3"),
		},
		Properties: &armmediaservices.LiveEventProperties{
			Description: to.Ptr("test event updated"),
			Input: &armmediaservices.LiveEventInput{
				AccessControl: &armmediaservices.LiveEventInputAccessControl{
					IP: &armmediaservices.IPAccessControl{
						Allow: []*armmediaservices.IPRange{
							{
								Name:    to.Ptr("AllowOne"),
								Address: to.Ptr("192.1.1.0"),
							}},
					},
				},
				KeyFrameIntervalDuration: to.Ptr("PT6S"),
				StreamingProtocol:        to.Ptr(armmediaservices.LiveEventInputProtocolFragmentedMP4),
			},
			Preview: &armmediaservices.LiveEventPreview{
				AccessControl: &armmediaservices.LiveEventPreviewAccessControl{
					IP: &armmediaservices.IPAccessControl{
						Allow: []*armmediaservices.IPRange{
							{
								Name:    to.Ptr("AllowOne"),
								Address: to.Ptr("192.1.1.0"),
							}},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LiveEvent = armmediaservices.LiveEvent{
	// 	Name: to.Ptr("myLiveEvent1"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/liveevents"),
	// 	ID: to.Ptr("/subscriptions/0a6ec948-5a62-437d-b9df-934dc7c1b722/resourceGroups/mediaresources/providers/Microsoft.Media/mediaservices/slitestmedia10/liveevents/myLiveEvent1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 		"tag3": to.Ptr("value3"),
	// 	},
	// 	Properties: &armmediaservices.LiveEventProperties{
	// 		Description: to.Ptr("test event updated"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
	// 		Encoding: &armmediaservices.LiveEventEncoding{
	// 			EncodingType: to.Ptr(armmediaservices.LiveEventEncodingTypeNone),
	// 		},
	// 		Input: &armmediaservices.LiveEventInput{
	// 			AccessControl: &armmediaservices.LiveEventInputAccessControl{
	// 				IP: &armmediaservices.IPAccessControl{
	// 					Allow: []*armmediaservices.IPRange{
	// 						{
	// 							Name: to.Ptr("AllowOne"),
	// 							Address: to.Ptr("192.1.1.0"),
	// 					}},
	// 				},
	// 			},
	// 			AccessToken: to.Ptr("<accessToken>"),
	// 			Endpoints: []*armmediaservices.LiveEventEndpoint{
	// 			},
	// 			KeyFrameIntervalDuration: to.Ptr("PT6S"),
	// 			StreamingProtocol: to.Ptr(armmediaservices.LiveEventInputProtocolFragmentedMP4),
	// 		},
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
	// 		Preview: &armmediaservices.LiveEventPreview{
	// 			AccessControl: &armmediaservices.LiveEventPreviewAccessControl{
	// 				IP: &armmediaservices.IPAccessControl{
	// 					Allow: []*armmediaservices.IPRange{
	// 						{
	// 							Name: to.Ptr("AllowOne"),
	// 							Address: to.Ptr("192.1.1.0"),
	// 					}},
	// 				},
	// 			},
	// 			Endpoints: []*armmediaservices.LiveEventEndpoint{
	// 			},
	// 			PreviewLocator: to.Ptr("c10ea3fc-587f-4daf-b2b2-fa8f647a9ed2"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ResourceState: to.Ptr(armmediaservices.LiveEventResourceStateRunning),
	// 		StreamOptions: []*armmediaservices.StreamOptionsFlag{
	// 		},
	// 		UseStaticHostname: to.Ptr(false),
	// 	},
	// }
}
