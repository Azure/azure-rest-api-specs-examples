package armeventgrid_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/Channels_CreateOrUpdate.json
func ExampleChannelsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armeventgrid.NewChannelsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<partner-namespace-name>",
		"<channel-name>",
		armeventgrid.Channel{
			Properties: &armeventgrid.ChannelProperties{
				ChannelType:                     to.Ptr(armeventgrid.ChannelTypePartnerTopic),
				ExpirationTimeIfNotActivatedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-21T22:50:25.410433Z"); return t }()),
				MessageForActivation:            to.Ptr("<message-for-activation>"),
				PartnerTopicInfo: &armeventgrid.PartnerTopicInfo{
					Name:                to.Ptr("<name>"),
					AzureSubscriptionID: to.Ptr("<azure-subscription-id>"),
					ResourceGroupName:   to.Ptr("<resource-group-name>"),
					Source:              to.Ptr("<source>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
