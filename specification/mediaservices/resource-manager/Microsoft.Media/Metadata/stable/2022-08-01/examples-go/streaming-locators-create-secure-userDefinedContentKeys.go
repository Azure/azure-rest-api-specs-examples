package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-create-secure-userDefinedContentKeys.json
func ExampleStreamingLocatorsClient_Create_createsAStreamingLocatorWithUserDefinedContentKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewStreamingLocatorsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Create(ctx, "contoso", "contosomedia", "UserCreatedSecureStreamingLocatorWithUserDefinedContentKeys", armmediaservices.StreamingLocator{
		Properties: &armmediaservices.StreamingLocatorProperties{
			AssetName: to.Ptr("ClimbingMountRainier"),
			ContentKeys: []*armmediaservices.StreamingLocatorContentKey{
				{
					ID:                              to.Ptr("60000000-0000-0000-0000-000000000001"),
					LabelReferenceInStreamingPolicy: to.Ptr("aesDefaultKey"),
					Value:                           to.Ptr("1UqLohAfWsEGkULYxHjYZg=="),
				},
				{
					ID:                              to.Ptr("60000000-0000-0000-0000-000000000004"),
					LabelReferenceInStreamingPolicy: to.Ptr("cencDefaultKey"),
					Value:                           to.Ptr("4UqLohAfWsEGkULYxHjYZg=="),
				},
				{
					ID:                              to.Ptr("60000000-0000-0000-0000-000000000007"),
					LabelReferenceInStreamingPolicy: to.Ptr("cbcsDefaultKey"),
					Value:                           to.Ptr("7UqLohAfWsEGkULYxHjYZg=="),
				}},
			StreamingLocatorID:  to.Ptr("90000000-0000-0000-0000-00000000000A"),
			StreamingPolicyName: to.Ptr("secureStreamingPolicy"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
