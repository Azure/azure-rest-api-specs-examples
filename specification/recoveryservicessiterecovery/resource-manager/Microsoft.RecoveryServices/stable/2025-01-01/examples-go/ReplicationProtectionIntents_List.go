package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationProtectionIntents_List.json
func ExampleReplicationProtectionIntentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicationProtectionIntentsClient().NewListPager("resourceGroupPS1", "2007vttp", &armrecoveryservicessiterecovery.ReplicationProtectionIntentsClientListOptions{SkipToken: nil,
		TakeToken: nil,
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
		// page.ReplicationProtectionIntentCollection = armrecoveryservicessiterecovery.ReplicationProtectionIntentCollection{
		// 	Value: []*armrecoveryservicessiterecovery.ReplicationProtectionIntent{
		// 		{
		// 			Name: to.Ptr("vm1"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationProtectionIntents"),
		// 			ID: to.Ptr("/Subscriptions/509099b2-9d2c-4636-b43e-bd5cafb6be69/resourceGroups/2007vttp/providers/Microsoft.RecoveryServices/vaults/tp2007vt/replicationProtectionIntents/vm1"),
		// 			Properties: &armrecoveryservicessiterecovery.ReplicationProtectionIntentProperties{
		// 				FriendlyName: to.Ptr("vm1"),
		// 				JobID: to.Ptr("/Subscriptions/d90d145a-4cdd-45a3-b2c4-971d69775278/resourceGroups/a2acl-rg-vault-prod-gip-ccy/providers/Microsoft.RecoveryServices/vaults/a2acl-vault-prod-gip-ccy/replicationJobs/02004ea7-d498-4bb4-bdeb-cdb611706867"),
		// 				JobState: to.Ptr("InProgress"),
		// 				ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2AReplicationIntentDetails{
		// 					InstanceType: to.Ptr("A2A"),
		// 					FabricObjectID: to.Ptr("/subscriptions/509099b2-9d2c-4636-b43e-bd5cafb6be69/resourceGroups/removeOne/providers/Microsoft.Compute/virtualMachines/vmPpgAv5"),
		// 					PrimaryLocation: to.Ptr("eastUs2"),
		// 					RecoveryAvailabilityType: to.Ptr("Single"),
		// 					RecoveryLocation: to.Ptr("westus2"),
		// 					RecoverySubscriptionID: to.Ptr("ed5bcdf6-d61e-47bd-8ea9-f2bd379a2640"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
