package armavs_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/PrivateClouds_CreateOrUpdate_FleetNative.json
func ExamplePrivateCloudsClient_BeginCreateOrUpdate_privateCloudsCreateOrUpdateFleetNative() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateCloudsClient().BeginCreateOrUpdate(ctx, "group1", "cloud1", armavs.PrivateCloud{
		Location: to.Ptr("eastus2"),
		SKU: &armavs.SKU{
			Name: to.Ptr("AV64"),
		},
		Properties: &armavs.PrivateCloudProperties{
			NetworkBlock:     to.Ptr("192.168.48.0/22"),
			VirtualNetworkID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualNetworks/vnet"),
			DNSZoneType:      to.Ptr(armavs.DNSZoneTypePrivate),
			ManagementCluster: &armavs.ManagementCluster{
				ClusterSize: to.Ptr[int32](4),
			},
			VcfLicense: &armavs.Vcf5License{
				Kind:                   to.Ptr(armavs.VcfLicenseKindVcf5),
				LicenseKey:             to.Ptr("12345-12345-12345-12345-12345"),
				EndDate:                to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-31T23:59:59Z"); return t }()),
				Cores:                  to.Ptr[int32](16),
				BroadcomSiteID:         to.Ptr("123456"),
				BroadcomContractNumber: to.Ptr("123456"),
			},
		},
		Tags: map[string]*string{},
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
	// res = armavs.PrivateCloudsClientCreateOrUpdateResponse{
	// 	PrivateCloud: &armavs.PrivateCloud{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1"),
	// 		Location: to.Ptr("eastus2"),
	// 		Name: to.Ptr("cloud1"),
	// 		Properties: &armavs.PrivateCloudProperties{
	// 			Availability: &armavs.AvailabilityProperties{
	// 				Strategy: to.Ptr(armavs.AvailabilityStrategySingleZone),
	// 			},
	// 			DNSZoneType: to.Ptr(armavs.DNSZoneTypePrivate),
	// 			Encryption: &armavs.Encryption{
	// 				Status: to.Ptr(armavs.EncryptionStateDisabled),
	// 			},
	// 			Endpoints: &armavs.Endpoints{
	// 				HcxCloudManager: to.Ptr("https://hcx.c2e20fa95ec249939dc7e3.mockenvavs.azure.com/"),
	// 				HcxCloudManagerIP: to.Ptr("10.31.0.37"),
	// 				NsxtManager: to.Ptr("https://nsx.c2e20fa95ec249939dc7e3.mockenvavs.azure.com/"),
	// 				NsxtManagerIP: to.Ptr("10.31.0.4"),
	// 				VcenterIP: to.Ptr("10.31.0.36"),
	// 				Vcsa: to.Ptr("https://vc.c2e20fa95ec249939dc7e3.mockenvavs.azure.com/"),
	// 			},
	// 			ExternalCloudLinks: []*string{
	// 			},
	// 			IdentitySources: []*armavs.IdentitySource{
	// 			},
	// 			Internet: to.Ptr(armavs.InternetEnumDisabled),
	// 			ManagementCluster: &armavs.ManagementCluster{
	// 				ClusterID: to.Ptr[int32](1),
	// 				ClusterSize: to.Ptr[int32](4),
	// 				Hosts: []*string{
	// 					to.Ptr("fakehost18.nyc1.kubernetes.center"),
	// 					to.Ptr("fakehost19.nyc1.kubernetes.center"),
	// 					to.Ptr("fakehost20.nyc1.kubernetes.center"),
	// 					to.Ptr("fakehost21.nyc1.kubernetes.center"),
	// 				},
	// 				ProvisioningState: to.Ptr(armavs.ClusterProvisioningState("Building")),
	// 				VsanDatastoreName: to.Ptr("vsanDatastore"),
	// 			},
	// 			ManagementNetwork: to.Ptr("10.31.0.0/26"),
	// 			NetworkBlock: to.Ptr("10.31.0.0/22"),
	// 			NsxPublicIPQuotaRaised: to.Ptr(armavs.NsxPublicIPQuotaRaisedEnumDisabled),
	// 			ProvisioningState: to.Ptr(armavs.PrivateCloudProvisioningStateBuilding),
	// 			VirtualNetworkID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup/providers/Microsoft.Network/virtualNetworks/mock-vnet"),
	// 			VmotionNetwork: to.Ptr("10.31.2.0/24"),
	// 			VcfLicense: &armavs.Vcf5License{
	// 				Kind: to.Ptr(armavs.VcfLicenseKindVcf5),
	// 				EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-31T23:59:59Z"); return t}()),
	// 				Cores: to.Ptr[int32](16),
	// 				BroadcomSiteID: to.Ptr("123456"),
	// 				BroadcomContractNumber: to.Ptr("123456"),
	// 				ProvisioningState: to.Ptr(armavs.LicenseProvisioningStateSucceeded),
	// 			},
	// 		},
	// 		SKU: &armavs.SKU{
	// 			Name: to.Ptr("av64"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Type: to.Ptr("Microsoft.AVS/privateClouds"),
	// 	},
	// }
}
