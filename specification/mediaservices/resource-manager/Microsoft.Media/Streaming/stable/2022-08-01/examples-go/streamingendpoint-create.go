package armmediaservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/streamingendpoint-create.json
func ExampleStreamingEndpointsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStreamingEndpointsClient().BeginCreate(ctx, "mediaresources", "slitestmedia10", "myStreamingEndpoint1", armmediaservices.StreamingEndpoint{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
		Properties: &armmediaservices.StreamingEndpointProperties{
			Description: to.Ptr("test event 1"),
			AccessControl: &armmediaservices.StreamingEndpointAccessControl{
				Akamai: &armmediaservices.AkamaiAccessControl{
					AkamaiSignatureHeaderAuthenticationKeyList: []*armmediaservices.AkamaiSignatureHeaderAuthenticationKey{
						{
							Base64Key:  to.Ptr("dGVzdGlkMQ=="),
							Expiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2030-01-01T00:00:00.000Z"); return t }()),
							Identifier: to.Ptr("id1"),
						},
						{
							Base64Key:  to.Ptr("dGVzdGlkMQ=="),
							Expiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2031-01-01T00:00:00.000Z"); return t }()),
							Identifier: to.Ptr("id2"),
						}},
				},
				IP: &armmediaservices.IPAccessControl{
					Allow: []*armmediaservices.IPRange{
						{
							Name:    to.Ptr("AllowedIp"),
							Address: to.Ptr("192.168.1.1"),
						}},
				},
			},
			AvailabilitySetName: to.Ptr("availableset"),
			CdnEnabled:          to.Ptr(false),
			ScaleUnits:          to.Ptr[int32](1),
		},
	}, &armmediaservices.StreamingEndpointsClientBeginCreateOptions{AutoStart: nil})
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
	// res.StreamingEndpoint = armmediaservices.StreamingEndpoint{
	// 	Name: to.Ptr("myStreamingEndpoint1"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/streamingEndpoints"),
	// 	ID: to.Ptr("/subscriptions/0a6ec948-5a62-437d-b9df-934dc7c1b722/resourceGroups/mediaresources/providers/Microsoft.Media/mediaservices/slitestmedia10/streamingendpoints/myStreamingEndpoint1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armmediaservices.StreamingEndpointProperties{
	// 		Description: to.Ptr("test event 1"),
	// 		AccessControl: &armmediaservices.StreamingEndpointAccessControl{
	// 			Akamai: &armmediaservices.AkamaiAccessControl{
	// 				AkamaiSignatureHeaderAuthenticationKeyList: []*armmediaservices.AkamaiSignatureHeaderAuthenticationKey{
	// 					{
	// 						Base64Key: to.Ptr("dGVzdGlkMQ=="),
	// 						Expiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2030-01-01T00:00:00.000Z"); return t}()),
	// 						Identifier: to.Ptr("id1"),
	// 					},
	// 					{
	// 						Base64Key: to.Ptr("dGVzdGlkMQ=="),
	// 						Expiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2031-01-01T00:00:00.000Z"); return t}()),
	// 						Identifier: to.Ptr("id2"),
	// 				}},
	// 			},
	// 			IP: &armmediaservices.IPAccessControl{
	// 				Allow: []*armmediaservices.IPRange{
	// 					{
	// 						Name: to.Ptr("AllowedIp"),
	// 						Address: to.Ptr("192.168.1.1"),
	// 				}},
	// 			},
	// 		},
	// 		AvailabilitySetName: to.Ptr("availableset"),
	// 		CdnEnabled: to.Ptr(false),
	// 		CdnProfile: to.Ptr(""),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:09.756Z"); return t}()),
	// 		CustomHostNames: []*string{
	// 		},
	// 		FreeTrialEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
	// 		HostName: to.Ptr("mystreamingendpoint1-slitestmedia10.streaming.mediaservices.windows.net"),
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:09.756Z"); return t}()),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ResourceState: to.Ptr(armmediaservices.StreamingEndpointResourceStateStopped),
	// 		ScaleUnits: to.Ptr[int32](1),
	// 	},
	// }
}
