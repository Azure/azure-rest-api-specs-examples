package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ee1eec42dcc710ff88db2d1bf574b2f9afe3d654/specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/PartnerTopics_Deactivate.json
func ExamplePartnerTopicsClient_Deactivate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPartnerTopicsClient().Deactivate(ctx, "examplerg", "examplePartnerTopic1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PartnerTopic = armeventgrid.PartnerTopic{
	// 	Name: to.Ptr("examplePartnerTopic1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/partnerTopics"),
	// 	ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerTopics/examplePartnerTopic1"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Properties: &armeventgrid.PartnerTopicProperties{
	// 		ActivationState: to.Ptr(armeventgrid.PartnerTopicActivationStateDeactivated),
	// 		ProvisioningState: to.Ptr(armeventgrid.PartnerTopicProvisioningStateSucceeded),
	// 		Source: to.Ptr("ContosoCorp.Accounts.User1"),
	// 	},
	// }
}
