package armeventgrid_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/Channels_CreateOrUpdate.json
func ExampleChannelsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armeventgrid.NewChannelsClient("5b4b650e-28b9-4790-b3ab-ddbd88d727c4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"examplerg",
		"examplePartnerNamespaceName1",
		"exampleChannelName1",
		armeventgrid.Channel{
			Properties: &armeventgrid.ChannelProperties{
				ChannelType:                     to.Ptr(armeventgrid.ChannelTypePartnerTopic),
				ExpirationTimeIfNotActivatedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-21T22:50:25.410433Z"); return t }()),
				MessageForActivation:            to.Ptr("Example message to approver"),
				PartnerTopicInfo: &armeventgrid.PartnerTopicInfo{
					Name:                to.Ptr("examplePartnerTopic1"),
					AzureSubscriptionID: to.Ptr("5b4b650e-28b9-4790-b3ab-ddbd88d727c4"),
					ResourceGroupName:   to.Ptr("examplerg2"),
					Source:              to.Ptr("ContosoCorp.Accounts.User1"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
