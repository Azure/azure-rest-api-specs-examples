package armmediaservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streamingendpoint-create.json
func ExampleStreamingEndpointsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewStreamingEndpointsClient("0a6ec948-5a62-437d-b9df-934dc7c1b722", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"mediaresources",
		"slitestmedia10",
		"myStreamingEndpoint1",
		armmediaservices.StreamingEndpoint{
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
								Expiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2029-12-31T16:00:00-08:00"); return t }()),
								Identifier: to.Ptr("id1"),
							},
							{
								Base64Key:  to.Ptr("dGVzdGlkMQ=="),
								Expiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2030-12-31T16:00:00-08:00"); return t }()),
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
		},
		&armmediaservices.StreamingEndpointsClientBeginCreateOptions{AutoStart: nil})
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
