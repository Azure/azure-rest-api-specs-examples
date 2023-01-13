package armmediaservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-create-secure.json
func ExampleStreamingLocatorsClient_Create_createsAStreamingLocatorWithSecureStreaming() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewStreamingLocatorsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Create(ctx, "contoso", "contosomedia", "UserCreatedSecureStreamingLocator", armmediaservices.StreamingLocator{
		Properties: &armmediaservices.StreamingLocatorProperties{
			AssetName:           to.Ptr("ClimbingMountRainier"),
			EndTime:             to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2028-12-31T23:59:59.9999999Z"); return t }()),
			StartTime:           to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-01T00:00:00Z"); return t }()),
			StreamingPolicyName: to.Ptr("secureStreamingPolicy"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
