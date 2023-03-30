package armeventgrid_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f88928d723133dc392e3297e6d61b7f6d10501fd/specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/PartnerTopics_CreateOrUpdate.json
func ExamplePartnerTopicsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPartnerTopicsClient().CreateOrUpdate(ctx, "examplerg", "examplePartnerTopicName1", armeventgrid.PartnerTopic{
		Location: to.Ptr("westus2"),
		Properties: &armeventgrid.PartnerTopicProperties{
			ExpirationTimeIfNotActivatedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T23:06:13.109Z"); return t }()),
			MessageForActivation:            to.Ptr("Example message for activation"),
			PartnerRegistrationImmutableID:  to.Ptr("6f541064-031d-4cc8-9ec3-a3b4fc0f7185"),
			PartnerTopicFriendlyDescription: to.Ptr("Example description"),
			Source:                          to.Ptr("ContosoCorp.Accounts.User1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PartnerTopic = armeventgrid.PartnerTopic{
	// 	Name: to.Ptr("examplePartnerTopicName1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/partnerTopics"),
	// 	ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerTopics/examplePartnerTopicName1"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armeventgrid.PartnerTopicProperties{
	// 		ActivationState: to.Ptr(armeventgrid.PartnerTopicActivationStateNeverActivated),
	// 		ExpirationTimeIfNotActivatedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T23:06:13.109Z"); return t}()),
	// 		MessageForActivation: to.Ptr("Example message for activation"),
	// 		PartnerRegistrationImmutableID: to.Ptr("6f541064-031d-4cc8-9ec3-a3b4fc0f7185"),
	// 		PartnerTopicFriendlyDescription: to.Ptr("Example description"),
	// 		ProvisioningState: to.Ptr(armeventgrid.PartnerTopicProvisioningStateSucceeded),
	// 		Source: to.Ptr("ContosoCorp.Accounts.User1"),
	// 	},
	// }
}
