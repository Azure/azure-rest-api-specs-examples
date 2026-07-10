package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v11"
)

// Generated from example definition: 2026-04-15-preview/VolumeGroups_Create_Custom_SMB.json
func ExampleVolumeGroupsClient_BeginCreate_volumeGroupsCreateCustomSmb() {
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
				ApplicationIdentifier: to.Ptr("CU2"),
				ApplicationType:       to.Ptr(armnetapp.ApplicationType("CUSTOM")),
				GroupDescription:      to.Ptr("Volume group"),
			},
			Volumes: []*armnetapp.VolumeGroupVolumeProperties{
				{
					Name: to.Ptr("test-cus-data1"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-cus-data1"),
						ProtocolTypes: []*string{
							to.Ptr("CIFS"),
						},
						ServiceLevel:              to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:                  to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps:           to.Ptr[float32](10),
						UsageThreshold:            to.Ptr[int64](107374182400),
						VolumeSpecName:            to.Ptr("cus-data1"),
						SmbEncryption:             to.Ptr(false),
						SmbContinuouslyAvailable:  to.Ptr(false),
						SmbNonBrowsable:           to.Ptr(armnetapp.SmbNonBrowsableDisabled),
						SmbAccessBasedEnumeration: to.Ptr(armnetapp.SmbAccessBasedEnumerationDisabled),
						AvsDataStore:              to.Ptr(armnetapp.AvsDataStoreDisabled),
					},
					Zones: []*string{
						to.Ptr("1"),
					},
				},
				{
					Name: to.Ptr("test-cus-data2"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-cus-data2"),
						ProtocolTypes: []*string{
							to.Ptr("CIFS"),
						},
						ServiceLevel:              to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:                  to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps:           to.Ptr[float32](10),
						UsageThreshold:            to.Ptr[int64](107374182400),
						VolumeSpecName:            to.Ptr("cus-data2"),
						SmbEncryption:             to.Ptr(false),
						SmbContinuouslyAvailable:  to.Ptr(false),
						SmbNonBrowsable:           to.Ptr(armnetapp.SmbNonBrowsableDisabled),
						SmbAccessBasedEnumeration: to.Ptr(armnetapp.SmbAccessBasedEnumerationDisabled),
						AvsDataStore:              to.Ptr(armnetapp.AvsDataStoreDisabled),
					},
					Zones: []*string{
						to.Ptr("1"),
					},
				},
				{
					Name: to.Ptr("test-cus-data3"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-cus-data3"),
						ProtocolTypes: []*string{
							to.Ptr("CIFS"),
						},
						ServiceLevel:              to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:                  to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps:           to.Ptr[float32](10),
						UsageThreshold:            to.Ptr[int64](107374182400),
						VolumeSpecName:            to.Ptr("cus-data3"),
						SmbEncryption:             to.Ptr(false),
						SmbContinuouslyAvailable:  to.Ptr(false),
						SmbNonBrowsable:           to.Ptr(armnetapp.SmbNonBrowsableDisabled),
						SmbAccessBasedEnumeration: to.Ptr(armnetapp.SmbAccessBasedEnumerationDisabled),
						AvsDataStore:              to.Ptr(armnetapp.AvsDataStoreDisabled),
					},
					Zones: []*string{
						to.Ptr("1"),
					},
				},
				{
					Name: to.Ptr("test-cus-data4"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-cus-data4"),
						ProtocolTypes: []*string{
							to.Ptr("CIFS"),
						},
						ServiceLevel:              to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:                  to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps:           to.Ptr[float32](10),
						UsageThreshold:            to.Ptr[int64](107374182400),
						VolumeSpecName:            to.Ptr("cus-data4"),
						SmbEncryption:             to.Ptr(false),
						SmbContinuouslyAvailable:  to.Ptr(false),
						SmbNonBrowsable:           to.Ptr(armnetapp.SmbNonBrowsableDisabled),
						SmbAccessBasedEnumeration: to.Ptr(armnetapp.SmbAccessBasedEnumerationDisabled),
						AvsDataStore:              to.Ptr(armnetapp.AvsDataStoreDisabled),
					},
					Zones: []*string{
						to.Ptr("1"),
					},
				},
				{
					Name: to.Ptr("test-cus-data5"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-cus-data5"),
						ProtocolTypes: []*string{
							to.Ptr("CIFS"),
						},
						ServiceLevel:              to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:                  to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps:           to.Ptr[float32](10),
						UsageThreshold:            to.Ptr[int64](107374182400),
						VolumeSpecName:            to.Ptr("cus-data5"),
						SmbEncryption:             to.Ptr(false),
						SmbContinuouslyAvailable:  to.Ptr(false),
						SmbNonBrowsable:           to.Ptr(armnetapp.SmbNonBrowsableDisabled),
						SmbAccessBasedEnumeration: to.Ptr(armnetapp.SmbAccessBasedEnumerationDisabled),
						AvsDataStore:              to.Ptr(armnetapp.AvsDataStoreDisabled),
					},
					Zones: []*string{
						to.Ptr("1"),
					},
				},
				{
					Name: to.Ptr("test-cus-data6"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-cus-data6"),
						ProtocolTypes: []*string{
							to.Ptr("CIFS"),
						},
						ServiceLevel:              to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:                  to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps:           to.Ptr[float32](10),
						UsageThreshold:            to.Ptr[int64](107374182400),
						VolumeSpecName:            to.Ptr("cus-data6"),
						SmbEncryption:             to.Ptr(false),
						SmbContinuouslyAvailable:  to.Ptr(false),
						SmbNonBrowsable:           to.Ptr(armnetapp.SmbNonBrowsableDisabled),
						SmbAccessBasedEnumeration: to.Ptr(armnetapp.SmbAccessBasedEnumerationDisabled),
						AvsDataStore:              to.Ptr(armnetapp.AvsDataStoreDisabled),
					},
					Zones: []*string{
						to.Ptr("1"),
					},
				},
				{
					Name: to.Ptr("test-cus-data7"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-cus-data7"),
						ProtocolTypes: []*string{
							to.Ptr("CIFS"),
						},
						ServiceLevel:              to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:                  to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps:           to.Ptr[float32](10),
						UsageThreshold:            to.Ptr[int64](107374182400),
						VolumeSpecName:            to.Ptr("cus-data7"),
						SmbEncryption:             to.Ptr(false),
						SmbContinuouslyAvailable:  to.Ptr(false),
						SmbNonBrowsable:           to.Ptr(armnetapp.SmbNonBrowsableDisabled),
						SmbAccessBasedEnumeration: to.Ptr(armnetapp.SmbAccessBasedEnumerationDisabled),
						AvsDataStore:              to.Ptr(armnetapp.AvsDataStoreDisabled),
					},
					Zones: []*string{
						to.Ptr("1"),
					},
				},
				{
					Name: to.Ptr("test-cus-data8"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-cus-data8"),
						ProtocolTypes: []*string{
							to.Ptr("CIFS"),
						},
						ServiceLevel:              to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:                  to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps:           to.Ptr[float32](10),
						UsageThreshold:            to.Ptr[int64](107374182400),
						VolumeSpecName:            to.Ptr("cus-data8"),
						SmbEncryption:             to.Ptr(false),
						SmbContinuouslyAvailable:  to.Ptr(false),
						SmbNonBrowsable:           to.Ptr(armnetapp.SmbNonBrowsableDisabled),
						SmbAccessBasedEnumeration: to.Ptr(armnetapp.SmbAccessBasedEnumerationDisabled),
						AvsDataStore:              to.Ptr(armnetapp.AvsDataStoreDisabled),
					},
					Zones: []*string{
						to.Ptr("1"),
					},
				},
				{
					Name: to.Ptr("test-cus-data9"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-cus-data9"),
						ProtocolTypes: []*string{
							to.Ptr("CIFS"),
						},
						ServiceLevel:              to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:                  to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps:           to.Ptr[float32](10),
						UsageThreshold:            to.Ptr[int64](107374182400),
						VolumeSpecName:            to.Ptr("cus-data9"),
						SmbEncryption:             to.Ptr(false),
						SmbContinuouslyAvailable:  to.Ptr(false),
						SmbNonBrowsable:           to.Ptr(armnetapp.SmbNonBrowsableDisabled),
						SmbAccessBasedEnumeration: to.Ptr(armnetapp.SmbAccessBasedEnumerationDisabled),
						AvsDataStore:              to.Ptr(armnetapp.AvsDataStoreDisabled),
					},
					Zones: []*string{
						to.Ptr("1"),
					},
				},
				{
					Name: to.Ptr("test-cus-data10"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-cus-data10"),
						ProtocolTypes: []*string{
							to.Ptr("CIFS"),
						},
						ServiceLevel:              to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:                  to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps:           to.Ptr[float32](10),
						UsageThreshold:            to.Ptr[int64](107374182400),
						VolumeSpecName:            to.Ptr("cus-data10"),
						SmbEncryption:             to.Ptr(false),
						SmbContinuouslyAvailable:  to.Ptr(false),
						SmbNonBrowsable:           to.Ptr(armnetapp.SmbNonBrowsableDisabled),
						SmbAccessBasedEnumeration: to.Ptr(armnetapp.SmbAccessBasedEnumerationDisabled),
						AvsDataStore:              to.Ptr(armnetapp.AvsDataStoreDisabled),
					},
					Zones: []*string{
						to.Ptr("1"),
					},
				},
				{
					Name: to.Ptr("test-cus-data11"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-cus-data11"),
						ProtocolTypes: []*string{
							to.Ptr("CIFS"),
						},
						ServiceLevel:              to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:                  to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps:           to.Ptr[float32](10),
						UsageThreshold:            to.Ptr[int64](107374182400),
						VolumeSpecName:            to.Ptr("cus-data11"),
						SmbEncryption:             to.Ptr(false),
						SmbContinuouslyAvailable:  to.Ptr(false),
						SmbNonBrowsable:           to.Ptr(armnetapp.SmbNonBrowsableDisabled),
						SmbAccessBasedEnumeration: to.Ptr(armnetapp.SmbAccessBasedEnumerationDisabled),
						AvsDataStore:              to.Ptr(armnetapp.AvsDataStoreDisabled),
					},
					Zones: []*string{
						to.Ptr("1"),
					},
				},
				{
					Name: to.Ptr("test-cus-data12"),
					Properties: &armnetapp.VolumeProperties{
						CapacityPoolResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
						CreationToken:          to.Ptr("test-cus-data12"),
						ProtocolTypes: []*string{
							to.Ptr("CIFS"),
						},
						ServiceLevel:              to.Ptr(armnetapp.ServiceLevelPremium),
						SubnetID:                  to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
						ThroughputMibps:           to.Ptr[float32](10),
						UsageThreshold:            to.Ptr[int64](107374182400),
						VolumeSpecName:            to.Ptr("cus-data12"),
						SmbEncryption:             to.Ptr(false),
						SmbContinuouslyAvailable:  to.Ptr(false),
						SmbNonBrowsable:           to.Ptr(armnetapp.SmbNonBrowsableDisabled),
						SmbAccessBasedEnumeration: to.Ptr(armnetapp.SmbAccessBasedEnumerationDisabled),
						AvsDataStore:              to.Ptr(armnetapp.AvsDataStoreDisabled),
					},
					Zones: []*string{
						to.Ptr("1"),
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
		log.Fatalf("failed to poll the result: %v", err)
	}
}
