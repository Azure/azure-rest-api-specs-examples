package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/VolumeGroups_Create_SapHana.json
func ExampleVolumeGroupsClient_BeginCreate_volumeGroupsCreateSapHana() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumeGroupsClient().BeginCreate(ctx, "myRG", "account1", "group1", armnetapp.VolumeGroupDetails{
		Location: to.Ptr("westus"),
		Properties: &armnetapp.VolumeGroupProperties{
			GroupMetaData: &armnetapp.VolumeGroupMetaData{
				ApplicationIdentifier: to.Ptr("SH9"),
				ApplicationType:       to.Ptr(armnetapp.ApplicationTypeSAPHANA),
				GroupDescription:      to.Ptr("Volume group"),
			},
			Volumes: []*armnetapp.VolumeGroupVolumeProperties{
				{
					Name: to.Ptr("test-data-mnt00001"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-data-mnt00001"),
						ExportPolicy: &armnetapp.VolumePropertiesExportPolicy{
							Rules: []*armnetapp.ExportPolicyRule{
								{
									AllowedClients:      to.Ptr("0.0.0.0/0"),
									Cifs:                to.Ptr(false),
									HasRootAccess:       to.Ptr(true),
									Kerberos5ReadOnly:   to.Ptr(false),
									Kerberos5ReadWrite:  to.Ptr(false),
									Kerberos5IReadOnly:  to.Ptr(false),
									Kerberos5IReadWrite: to.Ptr(false),
									Kerberos5PReadOnly:  to.Ptr(false),
									Kerberos5PReadWrite: to.Ptr(false),
									Nfsv3:               to.Ptr(false),
									Nfsv41:              to.Ptr(true),
									RuleIndex:           to.Ptr[int32](1),
									UnixReadOnly:        to.Ptr(true),
									UnixReadWrite:       to.Ptr(true),
								},
							},
						},
						ProtocolTypes: []*string{
							to.Ptr("NFSv4.1"),
						},
						ProximityPlacementGroup: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
						ServiceLevel:            to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:                to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps:         to.Ptr[float32](10),
						UsageThreshold:          to.Ptr[int64](107374182400),
						VolumeSpecName:          to.Ptr("data"),
					},
				},
				{
					Name: to.Ptr("test-log-mnt00001"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-log-mnt00001"),
						ExportPolicy: &armnetapp.VolumePropertiesExportPolicy{
							Rules: []*armnetapp.ExportPolicyRule{
								{
									AllowedClients:      to.Ptr("0.0.0.0/0"),
									Cifs:                to.Ptr(false),
									HasRootAccess:       to.Ptr(true),
									Kerberos5ReadOnly:   to.Ptr(false),
									Kerberos5ReadWrite:  to.Ptr(false),
									Kerberos5IReadOnly:  to.Ptr(false),
									Kerberos5IReadWrite: to.Ptr(false),
									Kerberos5PReadOnly:  to.Ptr(false),
									Kerberos5PReadWrite: to.Ptr(false),
									Nfsv3:               to.Ptr(false),
									Nfsv41:              to.Ptr(true),
									RuleIndex:           to.Ptr[int32](1),
									UnixReadOnly:        to.Ptr(true),
									UnixReadWrite:       to.Ptr(true),
								},
							},
						},
						ProtocolTypes: []*string{
							to.Ptr("NFSv4.1"),
						},
						ProximityPlacementGroup: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
						ServiceLevel:            to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:                to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps:         to.Ptr[float32](10),
						UsageThreshold:          to.Ptr[int64](107374182400),
						VolumeSpecName:          to.Ptr("log"),
					},
				},
				{
					Name: to.Ptr("test-shared"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-shared"),
						ExportPolicy: &armnetapp.VolumePropertiesExportPolicy{
							Rules: []*armnetapp.ExportPolicyRule{
								{
									AllowedClients:      to.Ptr("0.0.0.0/0"),
									Cifs:                to.Ptr(false),
									HasRootAccess:       to.Ptr(true),
									Kerberos5ReadOnly:   to.Ptr(false),
									Kerberos5ReadWrite:  to.Ptr(false),
									Kerberos5IReadOnly:  to.Ptr(false),
									Kerberos5IReadWrite: to.Ptr(false),
									Kerberos5PReadOnly:  to.Ptr(false),
									Kerberos5PReadWrite: to.Ptr(false),
									Nfsv3:               to.Ptr(false),
									Nfsv41:              to.Ptr(true),
									RuleIndex:           to.Ptr[int32](1),
									UnixReadOnly:        to.Ptr(true),
									UnixReadWrite:       to.Ptr(true),
								},
							},
						},
						ProtocolTypes: []*string{
							to.Ptr("NFSv4.1"),
						},
						ProximityPlacementGroup: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
						ServiceLevel:            to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:                to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps:         to.Ptr[float32](10),
						UsageThreshold:          to.Ptr[int64](107374182400),
						VolumeSpecName:          to.Ptr("shared"),
					},
				},
				{
					Name: to.Ptr("test-data-backup"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-data-backup"),
						ExportPolicy: &armnetapp.VolumePropertiesExportPolicy{
							Rules: []*armnetapp.ExportPolicyRule{
								{
									AllowedClients:      to.Ptr("0.0.0.0/0"),
									Cifs:                to.Ptr(false),
									HasRootAccess:       to.Ptr(true),
									Kerberos5ReadOnly:   to.Ptr(false),
									Kerberos5ReadWrite:  to.Ptr(false),
									Kerberos5IReadOnly:  to.Ptr(false),
									Kerberos5IReadWrite: to.Ptr(false),
									Kerberos5PReadOnly:  to.Ptr(false),
									Kerberos5PReadWrite: to.Ptr(false),
									Nfsv3:               to.Ptr(false),
									Nfsv41:              to.Ptr(true),
									RuleIndex:           to.Ptr[int32](1),
									UnixReadOnly:        to.Ptr(true),
									UnixReadWrite:       to.Ptr(true),
								},
							},
						},
						ProtocolTypes: []*string{
							to.Ptr("NFSv4.1"),
						},
						ProximityPlacementGroup: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
						ServiceLevel:            to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:                to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps:         to.Ptr[float32](10),
						UsageThreshold:          to.Ptr[int64](107374182400),
						VolumeSpecName:          to.Ptr("data-backup"),
					},
				},
				{
					Name: to.Ptr("test-log-backup"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-log-backup"),
						ExportPolicy: &armnetapp.VolumePropertiesExportPolicy{
							Rules: []*armnetapp.ExportPolicyRule{
								{
									AllowedClients:      to.Ptr("0.0.0.0/0"),
									Cifs:                to.Ptr(false),
									HasRootAccess:       to.Ptr(true),
									Kerberos5ReadOnly:   to.Ptr(false),
									Kerberos5ReadWrite:  to.Ptr(false),
									Kerberos5IReadOnly:  to.Ptr(false),
									Kerberos5IReadWrite: to.Ptr(false),
									Kerberos5PReadOnly:  to.Ptr(false),
									Kerberos5PReadWrite: to.Ptr(false),
									Nfsv3:               to.Ptr(false),
									Nfsv41:              to.Ptr(true),
									RuleIndex:           to.Ptr[int32](1),
									UnixReadOnly:        to.Ptr(true),
									UnixReadWrite:       to.Ptr(true),
								},
							},
						},
						ProtocolTypes: []*string{
							to.Ptr("NFSv4.1"),
						},
						ProximityPlacementGroup: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
						ServiceLevel:            to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:                to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps:         to.Ptr[float32](10),
						UsageThreshold:          to.Ptr[int64](107374182400),
						VolumeSpecName:          to.Ptr("log-backup"),
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
