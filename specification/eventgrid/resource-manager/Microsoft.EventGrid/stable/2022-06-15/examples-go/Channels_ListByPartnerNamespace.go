package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f88928d723133dc392e3297e6d61b7f6d10501fd/specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/Channels_ListByPartnerNamespace.json
func ExampleChannelsClient_NewListByPartnerNamespacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewChannelsClient().NewListByPartnerNamespacePager("examplerg", "examplePartnerNamespaceName1", &armeventgrid.ChannelsClientListByPartnerNamespaceOptions{Filter: nil,
		Top: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ChannelsListResult = armeventgrid.ChannelsListResult{
		// 	Value: []*armeventgrid.Channel{
		// 		{
		// 			Name: to.Ptr("exampleChannelName1"),
		// 			Type: to.Ptr("Microsoft.EventGrid/partnerNamespaces/channels"),
		// 			ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerNamespaces/examplePartnerNamespaceName1/changes/exampleChannelName1"),
		// 			Properties: &armeventgrid.ChannelProperties{
		// 				ChannelType: to.Ptr(armeventgrid.ChannelTypePartnerTopic),
		// 				ExpirationTimeIfNotActivatedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-21T22:50:25.410Z"); return t}()),
		// 				MessageForActivation: to.Ptr("Example message to approver"),
		// 				PartnerTopicInfo: &armeventgrid.PartnerTopicInfo{
		// 					Name: to.Ptr("examplePartnerTopic1"),
		// 					AzureSubscriptionID: to.Ptr("5b4b650e-28b9-4790-b3ab-ddbd88d727c4"),
		// 					ResourceGroupName: to.Ptr("examplerg2"),
		// 					Source: to.Ptr("ContosoCorp.Accounts.User1"),
		// 				},
		// 				ProvisioningState: to.Ptr(armeventgrid.ChannelProvisioningStateSucceeded),
		// 				ReadinessState: to.Ptr(armeventgrid.ReadinessStateNeverActivated),
		// 			},
		// 	}},
		// }
	}
}
