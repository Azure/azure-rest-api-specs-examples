package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationPolicies_Get.json
func ExampleReplicationPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReplicationPoliciesClient().Get(ctx, "resourceGroupPS1", "vault1", "protectionprofile1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Policy = armrecoveryservicessiterecovery.Policy{
	// 	Name: to.Ptr("protectionprofile1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationPolicies"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/protectionprofile1"),
	// 	Properties: &armrecoveryservicessiterecovery.PolicyProperties{
	// 		FriendlyName: to.Ptr("protectionprofile1"),
	// 		ProviderSpecificDetails: &armrecoveryservicessiterecovery.HyperVReplicaAzurePolicyDetails{
	// 			InstanceType: to.Ptr("HyperVReplicaAzure"),
	// 		},
	// 	},
	// }
}
