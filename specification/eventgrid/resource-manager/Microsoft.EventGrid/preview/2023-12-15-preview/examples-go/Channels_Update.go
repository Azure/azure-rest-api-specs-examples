package armeventgrid_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/Channels_Update.json
func ExampleChannelsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewChannelsClient().Update(ctx, "examplerg", "examplePartnerNamespaceName1", "exampleChannelName1", armeventgrid.ChannelUpdateParameters{
		Properties: &armeventgrid.ChannelUpdateParametersProperties{
			ExpirationTimeIfNotActivatedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T23:06:11.785Z"); return t }()),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
