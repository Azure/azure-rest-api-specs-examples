package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/PartnerTopics_ListByResourceGroup.json
func ExamplePartnerTopicsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPartnerTopicsClient().NewListByResourceGroupPager("examplerg", &armeventgrid.PartnerTopicsClientListByResourceGroupOptions{Filter: nil,
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
		// page.PartnerTopicsListResult = armeventgrid.PartnerTopicsListResult{
		// 	Value: []*armeventgrid.PartnerTopic{
		// 		{
		// 			Name: to.Ptr("examplePartnerTopic1"),
		// 			Type: to.Ptr("Microsoft.EventGrid/partnerTopics"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/amh/providers/Microsoft.EventGrid/partnerTopics/examplePartnerTopic1"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			Properties: &armeventgrid.PartnerTopicProperties{
		// 				ActivationState: to.Ptr(armeventgrid.PartnerTopicActivationStateNeverActivated),
		// 				ProvisioningState: to.Ptr(armeventgrid.PartnerTopicProvisioningStateSucceeded),
		// 				Source: to.Ptr("ContosoCorp.Accounts.User1"),
		// 			},
		// 	}},
		// }
	}
}
