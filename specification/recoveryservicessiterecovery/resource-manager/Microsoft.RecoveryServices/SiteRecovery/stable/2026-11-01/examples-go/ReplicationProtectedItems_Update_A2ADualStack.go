package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v4"
)

// Generated from example definition: 2026-11-01/ReplicationProtectedItems_Update_A2ADualStack.json
func ExampleReplicationProtectedItemsClient_BeginUpdate_updatesTheRecoveryIPConfigsOfADualStackIPv4AndIpv6A2AProtectedItem() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("42195872-7e70-4f8a-837f-84b28ecbb78b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationProtectedItemsClient().BeginUpdate(ctx, "priyanprg", "priyanponeboxvault", "CentralUSCanSite", "CentralUSCancloud", "a2aDualStackVm", armrecoveryservicessiterecovery.UpdateReplicationProtectedItemInput{
		Properties: &armrecoveryservicessiterecovery.UpdateReplicationProtectedItemInputProperties{
			SelectedRecoveryAzureNetworkID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Network/virtualNetworks/dualstack-recovery-vnet"),
			VMNics: []*armrecoveryservicessiterecovery.VMNicInputDetails{
				{
					NicID:         to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Network/networkInterfaces/a2aDualStackVm-nic"),
					SelectionType: to.Ptr("SelectedByUser"),
					IPConfigs: []*armrecoveryservicessiterecovery.IPConfigInputDetails{
						{
							IPConfigName:            to.Ptr("ipconfig1"),
							IsPrimary:               to.Ptr(true),
							IsSeletedForFailover:    to.Ptr(true),
							RecoverySubnetName:      to.Ptr("dualstack-subnet"),
							RecoveryStaticIPAddress: to.Ptr("10.0.2.46"),
						},
						{
							IPConfigName:            to.Ptr("ipconfig-v6"),
							IsPrimary:               to.Ptr(false),
							IsSeletedForFailover:    to.Ptr(true),
							RecoverySubnetName:      to.Ptr("dualstack-subnet"),
							RecoveryStaticIPAddress: to.Ptr("2001:db8:1::4"),
						},
					},
				},
			},
			ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2AUpdateReplicationProtectedItemInput{
				InstanceType:            to.Ptr("A2A"),
				RecoveryResourceGroupID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg"),
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
	// res = armrecoveryservicessiterecovery.ReplicationProtectedItemsClientUpdateResponse{
	// 	ReplicationProtectedItem: armrecoveryservicessiterecovery.ReplicationProtectedItem{
	// 		Name: to.Ptr("a2aDualStackVm"),
	// 		Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems"),
	// 		ID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.RecoveryServices/vaults/priyanponeboxvault/replicationFabrics/CentralUSCanSite/replicationProtectionContainers/CentralUSCancloud/replicationProtectedItems/a2aDualStackVm"),
	// 		Properties: &armrecoveryservicessiterecovery.ReplicationProtectedItemProperties{
	// 			ActiveLocation: to.Ptr("Primary"),
	// 			AllowedOperations: []*string{
	// 				to.Ptr("PlannedFailover"),
	// 				to.Ptr("UnplannedFailover"),
	// 				to.Ptr("DisableProtection"),
	// 				to.Ptr("TestFailover"),
	// 			},
	// 			FriendlyName: to.Ptr("a2aDualStackVm"),
	// 			PolicyFriendlyName: to.Ptr("a2aPolicy"),
	// 			PolicyID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.RecoveryServices/vaults/priyanponeboxvault/replicationPolicies/a2aPolicy"),
	// 			PrimaryFabricFriendlyName: to.Ptr("Central US"),
	// 			PrimaryProtectionContainerFriendlyName: to.Ptr("CentralUSCancloud"),
	// 			ProtectableItemID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.RecoveryServices/vaults/priyanponeboxvault/replicationFabrics/CentralUSCanSite/replicationProtectionContainers/CentralUSCancloud/replicationProtectableItems/a2aDualStackVm"),
	// 			ProtectedItemType: to.Ptr("A2A"),
	// 			ProtectionState: to.Ptr("Protected"),
	// 			ProtectionStateDescription: to.Ptr("Protected"),
	// 			ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2AReplicationDetails{
	// 				InstanceType: to.Ptr("A2A"),
	// 				FabricObjectID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/virtualMachines/a2aDualStackVm"),
	// 				SelectedRecoveryAzureNetworkID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Network/virtualNetworks/dualstack-recovery-vnet"),
	// 				VMNics: []*armrecoveryservicessiterecovery.VMNicDetails{
	// 					{
	// 						NicID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Network/networkInterfaces/a2aDualStackVm-nic"),
	// 						SourceNicArmID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Network/networkInterfaces/a2aDualStackVm-nic"),
	// 						SelectionType: to.Ptr("SelectedByUser"),
	// 						RecoveryVMNetworkID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Network/virtualNetworks/dualstack-recovery-vnet"),
	// 						IPConfigs: []*armrecoveryservicessiterecovery.IPConfigDetails{
	// 							{
	// 								Name: to.Ptr("ipconfig1"),
	// 								IsPrimary: to.Ptr(true),
	// 								IsSeletedForFailover: to.Ptr(true),
	// 								IPVersion: to.Ptr(armrecoveryservicessiterecovery.IPVersionIPv4),
	// 								SubnetName: to.Ptr("dualstack-subnet"),
	// 								StaticIPAddress: to.Ptr("10.0.1.15"),
	// 								IPAddressType: to.Ptr("Static"),
	// 								RecoverySubnetName: to.Ptr("dualstack-subnet"),
	// 								RecoveryStaticIPAddress: to.Ptr("10.0.2.46"),
	// 								RecoveryIPAddressType: to.Ptr("Static"),
	// 							},
	// 							{
	// 								Name: to.Ptr("ipconfig-v6"),
	// 								IsPrimary: to.Ptr(false),
	// 								IsSeletedForFailover: to.Ptr(true),
	// 								IPVersion: to.Ptr(armrecoveryservicessiterecovery.IPVersionIPv6),
	// 								SubnetName: to.Ptr("dualstack-subnet"),
	// 								StaticIPAddress: to.Ptr("2001:db8:0::4"),
	// 								IPAddressType: to.Ptr("Static"),
	// 								RecoverySubnetName: to.Ptr("dualstack-subnet"),
	// 								RecoveryStaticIPAddress: to.Ptr("2001:db8:1::4"),
	// 								RecoveryIPAddressType: to.Ptr("Static"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ReplicationHealth: to.Ptr("Normal"),
	// 			TestFailoverState: to.Ptr("None"),
	// 			TestFailoverStateDescription: to.Ptr("None"),
	// 		},
	// 	},
	// }
}
