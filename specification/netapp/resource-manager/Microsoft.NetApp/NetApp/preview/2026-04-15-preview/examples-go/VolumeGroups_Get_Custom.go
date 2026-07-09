package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v11"
)

// Generated from example definition: 2026-04-15-preview/VolumeGroups_Get_Custom.json
func ExampleVolumeGroupsClient_Get_volumeGroupsGetCustom() {
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
	// 	VolumeGroupDetails: armnetapp.VolumeGroupDetails{
	// 		Name: to.Ptr("group1"),
	// 		Type: to.Ptr("Microsoft.NetApp/netAppAccounts/volumeGroups"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/volumeGroups/group1"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armnetapp.VolumeGroupProperties{
	// 			GroupMetaData: &armnetapp.VolumeGroupMetaData{
	// 				ApplicationIdentifier: to.Ptr("CU2"),
	// 				ApplicationType: to.Ptr(armnetapp.ApplicationType("CUSTOM")),
	// 				GroupDescription: to.Ptr("Volume group"),
	// 				VolumesCount: to.Ptr[int64](12),
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			Volumes: []*armnetapp.VolumeGroupVolumeProperties{
	// 				{
	// 					Name: to.Ptr("account1/pool1/test-cus-data1"),
	// 					Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/test-cus-data1"),
	// 					Properties: &armnetapp.VolumeProperties{
	// 						CreationToken: to.Ptr("test-cus-data1"),
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
	// 						ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
	// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 						ThroughputMibps: to.Ptr[float32](10),
	// 						UsageThreshold: to.Ptr[int64](107374182400),
	// 						VolumeSpecName: to.Ptr("cus-data1"),
	// 					},
	// 					Zones: []*string{
	// 						to.Ptr("1"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("account1/pool1/test-cus-data2"),
	// 					Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/test-cus-data2"),
	// 					Properties: &armnetapp.VolumeProperties{
	// 						CreationToken: to.Ptr("test-cus-data2"),
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
	// 						ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
	// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 						ThroughputMibps: to.Ptr[float32](10),
	// 						UsageThreshold: to.Ptr[int64](107374182400),
	// 						VolumeSpecName: to.Ptr("cus-data2"),
	// 					},
	// 					Zones: []*string{
	// 						to.Ptr("1"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("account1/pool1/test-cus-data3"),
	// 					Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/test-cus-data3"),
	// 					Properties: &armnetapp.VolumeProperties{
	// 						CreationToken: to.Ptr("test-cus-data3"),
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
	// 						ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
	// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 						ThroughputMibps: to.Ptr[float32](10),
	// 						UsageThreshold: to.Ptr[int64](107374182400),
	// 						VolumeSpecName: to.Ptr("cus-data3"),
	// 					},
	// 					Zones: []*string{
	// 						to.Ptr("1"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("account1/pool1/test-cus-data4"),
	// 					Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/test-cus-data4"),
	// 					Properties: &armnetapp.VolumeProperties{
	// 						CreationToken: to.Ptr("test-cus-data4"),
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
	// 						ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
	// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 						ThroughputMibps: to.Ptr[float32](10),
	// 						UsageThreshold: to.Ptr[int64](107374182400),
	// 						VolumeSpecName: to.Ptr("cus-data4"),
	// 					},
	// 					Zones: []*string{
	// 						to.Ptr("1"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("account1/pool1/test-cus-data5"),
	// 					Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/test-cus-data5"),
	// 					Properties: &armnetapp.VolumeProperties{
	// 						CreationToken: to.Ptr("test-cus-data5"),
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
	// 						ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
	// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 						ThroughputMibps: to.Ptr[float32](10),
	// 						UsageThreshold: to.Ptr[int64](107374182400),
	// 						VolumeSpecName: to.Ptr("cus-data5"),
	// 					},
	// 					Zones: []*string{
	// 						to.Ptr("1"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("account1/pool1/test-cus-data6"),
	// 					Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/test-cus-data6"),
	// 					Properties: &armnetapp.VolumeProperties{
	// 						CreationToken: to.Ptr("test-cus-data6"),
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
	// 						ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
	// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 						ThroughputMibps: to.Ptr[float32](10),
	// 						UsageThreshold: to.Ptr[int64](107374182400),
	// 						VolumeSpecName: to.Ptr("cus-data6"),
	// 					},
	// 					Zones: []*string{
	// 						to.Ptr("1"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("account1/pool1/test-cus-data7"),
	// 					Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/test-cus-data7"),
	// 					Properties: &armnetapp.VolumeProperties{
	// 						CreationToken: to.Ptr("test-cus-data7"),
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
	// 						ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
	// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 						ThroughputMibps: to.Ptr[float32](10),
	// 						UsageThreshold: to.Ptr[int64](107374182400),
	// 						VolumeSpecName: to.Ptr("cus-data7"),
	// 					},
	// 					Zones: []*string{
	// 						to.Ptr("1"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("account1/pool1/test-cus-data8"),
	// 					Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/test-cus-data8"),
	// 					Properties: &armnetapp.VolumeProperties{
	// 						CreationToken: to.Ptr("test-cus-data8"),
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
	// 						ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
	// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 						ThroughputMibps: to.Ptr[float32](10),
	// 						UsageThreshold: to.Ptr[int64](107374182400),
	// 						VolumeSpecName: to.Ptr("cus-data8"),
	// 					},
	// 					Zones: []*string{
	// 						to.Ptr("1"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("account1/pool1/test-cus-data9"),
	// 					Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/test-cus-data8"),
	// 					Properties: &armnetapp.VolumeProperties{
	// 						CreationToken: to.Ptr("test-cus-data9"),
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
	// 						ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
	// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 						ThroughputMibps: to.Ptr[float32](10),
	// 						UsageThreshold: to.Ptr[int64](107374182400),
	// 						VolumeSpecName: to.Ptr("cus-data9"),
	// 					},
	// 					Zones: []*string{
	// 						to.Ptr("1"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("account1/pool1/test-cus-data10"),
	// 					Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/test-cus-data8"),
	// 					Properties: &armnetapp.VolumeProperties{
	// 						CreationToken: to.Ptr("test-cus-data10"),
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
	// 						ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
	// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 						ThroughputMibps: to.Ptr[float32](10),
	// 						UsageThreshold: to.Ptr[int64](107374182400),
	// 						VolumeSpecName: to.Ptr("cus-data10"),
	// 					},
	// 					Zones: []*string{
	// 						to.Ptr("1"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("account1/pool1/test-cus-data11"),
	// 					Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/test-cus-data8"),
	// 					Properties: &armnetapp.VolumeProperties{
	// 						CreationToken: to.Ptr("test-cus-data11"),
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
	// 						ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
	// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 						ThroughputMibps: to.Ptr[float32](10),
	// 						UsageThreshold: to.Ptr[int64](107374182400),
	// 						VolumeSpecName: to.Ptr("cus-data11"),
	// 					},
	// 					Zones: []*string{
	// 						to.Ptr("1"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("account1/pool1/test-cus-data12"),
	// 					Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/test-cus-data8"),
	// 					Properties: &armnetapp.VolumeProperties{
	// 						CreationToken: to.Ptr("test-cus-data12"),
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
	// 						ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
	// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 						ThroughputMibps: to.Ptr[float32](10),
	// 						UsageThreshold: to.Ptr[int64](107374182400),
	// 						VolumeSpecName: to.Ptr("cus-data12"),
	// 					},
	// 					Zones: []*string{
	// 						to.Ptr("1"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
