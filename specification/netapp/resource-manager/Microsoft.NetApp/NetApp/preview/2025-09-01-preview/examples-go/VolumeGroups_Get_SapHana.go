package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/VolumeGroups_Get_SapHana.json
func ExampleVolumeGroupsClient_Get_volumeGroupsGetSapHana() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVolumeGroupsClient().Get(ctx, "myRG", "account1", "group1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetapp.VolumeGroupsClientGetResponse{
	// 	VolumeGroupDetails: &armnetapp.VolumeGroupDetails{
	// 		Name: to.Ptr("group1"),
	// 		Type: to.Ptr("Microsoft.NetApp/netAppAccounts/volumeGroups"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/volumeGroups/group1"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armnetapp.VolumeGroupProperties{
	// 			GroupMetaData: &armnetapp.VolumeGroupMetaData{
	// 				ApplicationIdentifier: to.Ptr("SH9"),
	// 				ApplicationType: to.Ptr(armnetapp.ApplicationTypeSAPHANA),
	// 				GroupDescription: to.Ptr("Volume group"),
	// 				VolumesCount: to.Ptr[int64](5),
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			Volumes: []*armnetapp.VolumeGroupVolumeProperties{
	// 				{
	// 					Name: to.Ptr("account1/pool1/test-data-mnt00001"),
	// 					Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/test-data-mnt00001"),
	// 					Properties: &armnetapp.VolumeProperties{
	// 						CreationToken: to.Ptr("test-data-mnt00001"),
	// 						ExportPolicy: &armnetapp.VolumePropertiesExportPolicy{
	// 							Rules: []*armnetapp.ExportPolicyRule{
	// 								{
	// 									AllowedClients: to.Ptr("0.0.0.0/0"),
	// 									Cifs: to.Ptr(false),
	// 									HasRootAccess: to.Ptr(true),
	// 									Kerberos5ReadOnly: to.Ptr(false),
	// 									Kerberos5ReadWrite: to.Ptr(false),
	// 									Kerberos5IReadOnly: to.Ptr(false),
	// 									Kerberos5IReadWrite: to.Ptr(false),
	// 									Kerberos5PReadOnly: to.Ptr(false),
	// 									Kerberos5PReadWrite: to.Ptr(false),
	// 									Nfsv3: to.Ptr(false),
	// 									Nfsv41: to.Ptr(true),
	// 									RuleIndex: to.Ptr[int32](1),
	// 									UnixReadOnly: to.Ptr(true),
	// 									UnixReadWrite: to.Ptr(true),
	// 								},
	// 							},
	// 						},
	// 						ProtocolTypes: []*string{
	// 							to.Ptr("NFSv4.1"),
	// 						},
	// 						ProvisioningState: to.Ptr("Succeeded"),
	// 						ProximityPlacementGroup: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
	// 						ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
	// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 						ThroughputMibps: to.Ptr[float32](10),
	// 						UsageThreshold: to.Ptr[int64](107374182400),
	// 						VolumeSpecName: to.Ptr("data"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("account1/pool1/test-log-mnt00001"),
	// 					Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/test-log-mnt00001"),
	// 					Properties: &armnetapp.VolumeProperties{
	// 						CreationToken: to.Ptr("test-log-mnt00001"),
	// 						ExportPolicy: &armnetapp.VolumePropertiesExportPolicy{
	// 							Rules: []*armnetapp.ExportPolicyRule{
	// 								{
	// 									AllowedClients: to.Ptr("0.0.0.0/0"),
	// 									Cifs: to.Ptr(false),
	// 									HasRootAccess: to.Ptr(true),
	// 									Kerberos5ReadOnly: to.Ptr(false),
	// 									Kerberos5ReadWrite: to.Ptr(false),
	// 									Kerberos5IReadOnly: to.Ptr(false),
	// 									Kerberos5IReadWrite: to.Ptr(false),
	// 									Kerberos5PReadOnly: to.Ptr(false),
	// 									Kerberos5PReadWrite: to.Ptr(false),
	// 									Nfsv3: to.Ptr(false),
	// 									Nfsv41: to.Ptr(true),
	// 									RuleIndex: to.Ptr[int32](1),
	// 									UnixReadOnly: to.Ptr(true),
	// 									UnixReadWrite: to.Ptr(true),
	// 								},
	// 							},
	// 						},
	// 						ProtocolTypes: []*string{
	// 							to.Ptr("NFSv4.1"),
	// 						},
	// 						ProvisioningState: to.Ptr("Succeeded"),
	// 						ProximityPlacementGroup: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
	// 						ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
	// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 						ThroughputMibps: to.Ptr[float32](10),
	// 						UsageThreshold: to.Ptr[int64](107374182400),
	// 						VolumeSpecName: to.Ptr("log"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("account1/pool1/test-shared"),
	// 					Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/test-shared"),
	// 					Properties: &armnetapp.VolumeProperties{
	// 						CreationToken: to.Ptr("test-shared"),
	// 						ExportPolicy: &armnetapp.VolumePropertiesExportPolicy{
	// 							Rules: []*armnetapp.ExportPolicyRule{
	// 								{
	// 									AllowedClients: to.Ptr("0.0.0.0/0"),
	// 									Cifs: to.Ptr(false),
	// 									HasRootAccess: to.Ptr(true),
	// 									Kerberos5ReadOnly: to.Ptr(false),
	// 									Kerberos5ReadWrite: to.Ptr(false),
	// 									Kerberos5IReadOnly: to.Ptr(false),
	// 									Kerberos5IReadWrite: to.Ptr(false),
	// 									Kerberos5PReadOnly: to.Ptr(false),
	// 									Kerberos5PReadWrite: to.Ptr(false),
	// 									Nfsv3: to.Ptr(false),
	// 									Nfsv41: to.Ptr(true),
	// 									RuleIndex: to.Ptr[int32](1),
	// 									UnixReadOnly: to.Ptr(true),
	// 									UnixReadWrite: to.Ptr(true),
	// 								},
	// 							},
	// 						},
	// 						ProtocolTypes: []*string{
	// 							to.Ptr("NFSv4.1"),
	// 						},
	// 						ProvisioningState: to.Ptr("Succeeded"),
	// 						ProximityPlacementGroup: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
	// 						ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
	// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 						ThroughputMibps: to.Ptr[float32](10),
	// 						UsageThreshold: to.Ptr[int64](107374182400),
	// 						VolumeSpecName: to.Ptr("shared"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("account1/pool1/test-data-backup"),
	// 					Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/test-data-backup"),
	// 					Properties: &armnetapp.VolumeProperties{
	// 						CreationToken: to.Ptr("test-data-backup"),
	// 						ExportPolicy: &armnetapp.VolumePropertiesExportPolicy{
	// 							Rules: []*armnetapp.ExportPolicyRule{
	// 								{
	// 									AllowedClients: to.Ptr("0.0.0.0/0"),
	// 									Cifs: to.Ptr(false),
	// 									HasRootAccess: to.Ptr(true),
	// 									Kerberos5ReadOnly: to.Ptr(false),
	// 									Kerberos5ReadWrite: to.Ptr(false),
	// 									Kerberos5IReadOnly: to.Ptr(false),
	// 									Kerberos5IReadWrite: to.Ptr(false),
	// 									Kerberos5PReadOnly: to.Ptr(false),
	// 									Kerberos5PReadWrite: to.Ptr(false),
	// 									Nfsv3: to.Ptr(false),
	// 									Nfsv41: to.Ptr(true),
	// 									RuleIndex: to.Ptr[int32](1),
	// 									UnixReadOnly: to.Ptr(true),
	// 									UnixReadWrite: to.Ptr(true),
	// 								},
	// 							},
	// 						},
	// 						ProtocolTypes: []*string{
	// 							to.Ptr("NFSv4.1"),
	// 						},
	// 						ProvisioningState: to.Ptr("Succeeded"),
	// 						ProximityPlacementGroup: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
	// 						ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
	// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 						ThroughputMibps: to.Ptr[float32](10),
	// 						UsageThreshold: to.Ptr[int64](107374182400),
	// 						VolumeSpecName: to.Ptr("data-backup"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("account1/pool1/test-log-backup"),
	// 					Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/test-log-backup"),
	// 					Properties: &armnetapp.VolumeProperties{
	// 						CreationToken: to.Ptr("test-log-backup"),
	// 						ExportPolicy: &armnetapp.VolumePropertiesExportPolicy{
	// 							Rules: []*armnetapp.ExportPolicyRule{
	// 								{
	// 									AllowedClients: to.Ptr("0.0.0.0/0"),
	// 									Cifs: to.Ptr(false),
	// 									HasRootAccess: to.Ptr(true),
	// 									Kerberos5ReadOnly: to.Ptr(false),
	// 									Kerberos5ReadWrite: to.Ptr(false),
	// 									Kerberos5IReadOnly: to.Ptr(false),
	// 									Kerberos5IReadWrite: to.Ptr(false),
	// 									Kerberos5PReadOnly: to.Ptr(false),
	// 									Kerberos5PReadWrite: to.Ptr(false),
	// 									Nfsv3: to.Ptr(false),
	// 									Nfsv41: to.Ptr(true),
	// 									RuleIndex: to.Ptr[int32](1),
	// 									UnixReadOnly: to.Ptr(true),
	// 									UnixReadWrite: to.Ptr(true),
	// 								},
	// 							},
	// 						},
	// 						ProtocolTypes: []*string{
	// 							to.Ptr("NFSv4.1"),
	// 						},
	// 						ProvisioningState: to.Ptr("Succeeded"),
	// 						ProximityPlacementGroup: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
	// 						ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
	// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 						ThroughputMibps: to.Ptr[float32](10),
	// 						UsageThreshold: to.Ptr[int64](107374182400),
	// 						VolumeSpecName: to.Ptr("log-backup"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
