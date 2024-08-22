package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7b98b36e4023331545051284d8500adf98f02fe/specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-03-01/examples/VolumeGroups_Create_Oracle.json
func ExampleVolumeGroupsClient_BeginCreate_volumeGroupsCreateOracle() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumeGroupsClient().BeginCreate(ctx, "myRG", "account1", "group1", armnetapp.VolumeGroupDetails{
		Location: to.Ptr("westus"),
		Properties: &armnetapp.VolumeGroupProperties{
			GroupMetaData: &armnetapp.VolumeGroupMetaData{
				ApplicationIdentifier: to.Ptr("OR2"),
				ApplicationType:       to.Ptr(armnetapp.ApplicationTypeORACLE),
				GroupDescription:      to.Ptr("Volume group"),
			},
			Volumes: []*armnetapp.VolumeGroupVolumeProperties{
				{
					Name: to.Ptr("test-ora-data1"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-ora-data1"),
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
								}},
						},
						ProtocolTypes: []*string{
							to.Ptr("NFSv4.1")},
						ServiceLevel:    to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:        to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps: to.Ptr[float32](10),
						UsageThreshold:  to.Ptr[int64](107374182400),
						VolumeSpecName:  to.Ptr("ora-data1"),
					},
					Zones: []*string{
						to.Ptr("1")},
				},
				{
					Name: to.Ptr("test-ora-data2"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-ora-data2"),
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
								}},
						},
						ProtocolTypes: []*string{
							to.Ptr("NFSv4.1")},
						ServiceLevel:    to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:        to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps: to.Ptr[float32](10),
						UsageThreshold:  to.Ptr[int64](107374182400),
						VolumeSpecName:  to.Ptr("ora-data2"),
					},
					Zones: []*string{
						to.Ptr("1")},
				},
				{
					Name: to.Ptr("test-ora-data3"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-ora-data3"),
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
								}},
						},
						ProtocolTypes: []*string{
							to.Ptr("NFSv4.1")},
						ServiceLevel:    to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:        to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps: to.Ptr[float32](10),
						UsageThreshold:  to.Ptr[int64](107374182400),
						VolumeSpecName:  to.Ptr("ora-data3"),
					},
					Zones: []*string{
						to.Ptr("1")},
				},
				{
					Name: to.Ptr("test-ora-data4"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-ora-data4"),
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
								}},
						},
						ProtocolTypes: []*string{
							to.Ptr("NFSv4.1")},
						ServiceLevel:    to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:        to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps: to.Ptr[float32](10),
						UsageThreshold:  to.Ptr[int64](107374182400),
						VolumeSpecName:  to.Ptr("ora-data4"),
					},
					Zones: []*string{
						to.Ptr("1")},
				},
				{
					Name: to.Ptr("test-ora-data5"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-ora-data5"),
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
								}},
						},
						ProtocolTypes: []*string{
							to.Ptr("NFSv4.1")},
						ServiceLevel:    to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:        to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps: to.Ptr[float32](10),
						UsageThreshold:  to.Ptr[int64](107374182400),
						VolumeSpecName:  to.Ptr("ora-data5"),
					},
					Zones: []*string{
						to.Ptr("1")},
				},
				{
					Name: to.Ptr("test-ora-data6"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-ora-data6"),
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
								}},
						},
						ProtocolTypes: []*string{
							to.Ptr("NFSv4.1")},
						ServiceLevel:    to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:        to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps: to.Ptr[float32](10),
						UsageThreshold:  to.Ptr[int64](107374182400),
						VolumeSpecName:  to.Ptr("ora-data6"),
					},
					Zones: []*string{
						to.Ptr("1")},
				},
				{
					Name: to.Ptr("test-ora-data7"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-ora-data7"),
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
								}},
						},
						ProtocolTypes: []*string{
							to.Ptr("NFSv4.1")},
						ServiceLevel:    to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:        to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps: to.Ptr[float32](10),
						UsageThreshold:  to.Ptr[int64](107374182400),
						VolumeSpecName:  to.Ptr("ora-data7"),
					},
					Zones: []*string{
						to.Ptr("1")},
				},
				{
					Name: to.Ptr("test-ora-data8"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-ora-data8"),
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
								}},
						},
						ProtocolTypes: []*string{
							to.Ptr("NFSv4.1")},
						ServiceLevel:    to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:        to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps: to.Ptr[float32](10),
						UsageThreshold:  to.Ptr[int64](107374182400),
						VolumeSpecName:  to.Ptr("ora-data8"),
					},
					Zones: []*string{
						to.Ptr("1")},
				},
				{
					Name: to.Ptr("test-ora-log"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-ora-log"),
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
								}},
						},
						ProtocolTypes: []*string{
							to.Ptr("NFSv4.1")},
						ServiceLevel:    to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:        to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps: to.Ptr[float32](10),
						UsageThreshold:  to.Ptr[int64](107374182400),
						VolumeSpecName:  to.Ptr("ora-log"),
					},
					Zones: []*string{
						to.Ptr("1")},
				},
				{
					Name: to.Ptr("test-ora-log-mirror"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-ora-log-mirror"),
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
								}},
						},
						ProtocolTypes: []*string{
							to.Ptr("NFSv4.1")},
						ServiceLevel:    to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:        to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps: to.Ptr[float32](10),
						UsageThreshold:  to.Ptr[int64](107374182400),
						VolumeSpecName:  to.Ptr("ora-log-mirror"),
					},
					Zones: []*string{
						to.Ptr("1")},
				},
				{
					Name: to.Ptr("test-ora-binary"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-ora-binary"),
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
								}},
						},
						ProtocolTypes: []*string{
							to.Ptr("NFSv4.1")},
						ServiceLevel:    to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:        to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps: to.Ptr[float32](10),
						UsageThreshold:  to.Ptr[int64](107374182400),
						VolumeSpecName:  to.Ptr("ora-binary"),
					},
					Zones: []*string{
						to.Ptr("1")},
				},
				{
					Name: to.Ptr("test-ora-backup"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-ora-backup"),
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
								}},
						},
						ProtocolTypes: []*string{
							to.Ptr("NFSv4.1")},
						ServiceLevel:    to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:        to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps: to.Ptr[float32](10),
						UsageThreshold:  to.Ptr[int64](107374182400),
						VolumeSpecName:  to.Ptr("ora-backup"),
					},
					Zones: []*string{
						to.Ptr("1")},
				}},
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
