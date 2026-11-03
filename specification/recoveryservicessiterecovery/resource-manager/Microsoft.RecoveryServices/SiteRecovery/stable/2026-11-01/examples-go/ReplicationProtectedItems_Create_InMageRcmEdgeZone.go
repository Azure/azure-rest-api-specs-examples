package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v4"
)

// Generated from example definition: 2026-11-01/ReplicationProtectedItems_Create_InMageRcmEdgeZone.json
func ExampleReplicationProtectedItemsClient_BeginCreate_enablesProtectionForAnInMageRcmVMInAnExtendedLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationProtectedItemsClient().BeginCreate(ctx, "resourceGroupPS1", "vault1", "vmware-fabric", "vmware-container", "machine1", armrecoveryservicessiterecovery.EnableProtectionInput{
		Properties: &armrecoveryservicessiterecovery.EnableProtectionInputProperties{
			PolicyID:          to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/protectionProfile1"),
			ProtectableItemID: to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/vmware-fabric/replicationProtectionContainers/vmware-container/replicationProtectableItems/machine1"),
			ProviderSpecificDetails: &armrecoveryservicessiterecovery.InMageRcmEnableProtectionInput{
				FabricDiscoveryMachineID: to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/sourceResourceGroup/providers/Microsoft.OffAzure/VMwareSites/site1/machines/machine1"),
				InstanceType:             to.Ptr("InMageRcm"),
				ProcessServerID:          to.Ptr("process-server-1"),
				TargetExtendedLocation: &armrecoveryservicessiterecovery.ExtendedLocation{
					Name: to.Ptr("edge-zone-1"),
					Type: to.Ptr(armrecoveryservicessiterecovery.ExtendedLocationTypeEdgeZone),
				},
				TargetNetworkID:       to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/targetResourceGroup/providers/Microsoft.Network/virtualNetworks/targetVNet"),
				TargetResourceGroupID: to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/targetResourceGroup"),
				TargetSubnetName:      to.Ptr("default"),
				TargetVMName:          to.Ptr("machine1"),
				TargetVMSize:          to.Ptr("Standard_D2s_v3"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecoveryservicessiterecovery.ReplicationProtectedItemsClientCreateResponse{
	// 	ReplicationProtectedItem: armrecoveryservicessiterecovery.ReplicationProtectedItem{
	// 		ID: to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/vmware-fabric/replicationProtectionContainers/vmware-container/replicationProtectedItems/machine1"),
	// 		Name: to.Ptr("machine1"),
	// 		Properties: &armrecoveryservicessiterecovery.ReplicationProtectedItemProperties{
	// 			FriendlyName: to.Ptr("machine1"),
	// 			PolicyFriendlyName: to.Ptr("protectionProfile1"),
	// 			PolicyID: to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/protectionProfile1"),
	// 			ProtectableItemID: to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/vmware-fabric/replicationProtectionContainers/vmware-container/replicationProtectableItems/machine1"),
	// 			ProtectedItemType: to.Ptr("VMwareVirtualMachine"),
	// 			ProtectionState: to.Ptr("Protected"),
	// 			ProviderSpecificDetails: &armrecoveryservicessiterecovery.InMageRcmReplicationDetails{
	// 				FabricDiscoveryMachineID: to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/sourceResourceGroup/providers/Microsoft.OffAzure/VMwareSites/site1/machines/machine1"),
	// 				InstanceType: to.Ptr("InMageRcm"),
	// 				ProcessServerID: to.Ptr("process-server-1"),
	// 				TargetExtendedLocation: &armrecoveryservicessiterecovery.ExtendedLocation{
	// 					Name: to.Ptr("edge-zone-1"),
	// 					Type: to.Ptr(armrecoveryservicessiterecovery.ExtendedLocationTypeEdgeZone),
	// 				},
	// 				TargetLocation: to.Ptr("westus"),
	// 				TargetNetworkID: to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/targetResourceGroup/providers/Microsoft.Network/virtualNetworks/targetVNet"),
	// 				TargetResourceGroupID: to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/targetResourceGroup"),
	// 				TargetVMName: to.Ptr("machine1"),
	// 				TargetVMSize: to.Ptr("Standard_D2s_v3"),
	// 			},
	// 			ReplicationHealth: to.Ptr("Normal"),
	// 		},
	// 		Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems"),
	// 	},
	// }
}
