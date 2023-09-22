package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationProtectionContainerMappings_Create.json
func ExampleReplicationProtectionContainerMappingsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationProtectionContainerMappingsClient().BeginCreate(ctx, "vault1", "resourceGroupPS1", "cloud1", "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "cloud1protectionprofile1", armrecoveryservicessiterecovery.CreateProtectionContainerMappingInput{
		Properties: &armrecoveryservicessiterecovery.CreateProtectionContainerMappingInputProperties{
			PolicyID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/protectionprofile1"),
			ProviderSpecificInput: &armrecoveryservicessiterecovery.ReplicationProviderSpecificContainerMappingInput{
				InstanceType: to.Ptr("ReplicationProviderSpecificContainerMappingInput"),
			},
			TargetProtectionContainerID: to.Ptr("Microsoft Azure"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProtectionContainerMapping = armrecoveryservicessiterecovery.ProtectionContainerMapping{
	// 	Name: to.Ptr("cloud1protectionprofile1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectionContainerMappings"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectionContainerMappings/cloud1protectionprofile1"),
	// 	Properties: &armrecoveryservicessiterecovery.ProtectionContainerMappingProperties{
	// 		Health: to.Ptr("Normal"),
	// 		HealthErrorDetails: []*armrecoveryservicessiterecovery.HealthError{
	// 		},
	// 		PolicyFriendlyName: to.Ptr("protectionprofile1"),
	// 		PolicyID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/protectionprofile1"),
	// 		SourceFabricFriendlyName: to.Ptr("cloud1"),
	// 		SourceProtectionContainerFriendlyName: to.Ptr("cloud1"),
	// 		State: to.Ptr("Paired"),
	// 		TargetFabricFriendlyName: to.Ptr("Microsoft Azure"),
	// 		TargetProtectionContainerFriendlyName: to.Ptr("Microsoft Azure"),
	// 		TargetProtectionContainerID: to.Ptr("Microsoft Azure"),
	// 	},
	// }
}
