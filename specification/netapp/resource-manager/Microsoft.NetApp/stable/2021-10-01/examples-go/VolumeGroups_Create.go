package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-10-01/examples/VolumeGroups_Create.json
func ExampleVolumeGroupsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetapp.NewVolumeGroupsClient("D633CC2E-722B-4AE1-B636-BBD9E4C60ED9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"myRG",
		"account1",
		"group1",
		armnetapp.VolumeGroupDetails{
			Location: to.Ptr("westus"),
			Properties: &armnetapp.VolumeGroupProperties{
				GroupMetaData: &armnetapp.VolumeGroupMetaData{
					ApplicationIdentifier: to.Ptr("DEV"),
					ApplicationType:       to.Ptr(armnetapp.ApplicationTypeSAPHANA),
					DeploymentSpecID:      to.Ptr("fb04dbeb-005d-2703-197e-6208dfadb5d9"),
					GroupDescription:      to.Ptr("Volume group"),
				},
				Volumes: []*armnetapp.VolumeGroupVolumeProperties{
					{
						Name: to.Ptr("testVol1"),
						Properties: &armnetapp.VolumeProperties{
							CapacityPoolResourceID:  to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
							CreationToken:           to.Ptr("testVol1"),
							ProximityPlacementGroup: to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
							ServiceLevel:            to.Ptr(armnetapp.ServiceLevelPremium),
							SubnetID:                to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
							ThroughputMibps:         to.Ptr[float32](10),
							UsageThreshold:          to.Ptr[int64](107374182400),
							VolumeSpecName:          to.Ptr("data"),
						},
					},
					{
						Name: to.Ptr("testVol2"),
						Properties: &armnetapp.VolumeProperties{
							CapacityPoolResourceID:  to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
							CreationToken:           to.Ptr("testVol2"),
							ProximityPlacementGroup: to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
							ServiceLevel:            to.Ptr(armnetapp.ServiceLevelPremium),
							SubnetID:                to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
							ThroughputMibps:         to.Ptr[float32](10),
							UsageThreshold:          to.Ptr[int64](107374182400),
							VolumeSpecName:          to.Ptr("log"),
						},
					},
					{
						Name: to.Ptr("testVol3"),
						Properties: &armnetapp.VolumeProperties{
							CapacityPoolResourceID:  to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
							CreationToken:           to.Ptr("testVol3"),
							ProximityPlacementGroup: to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
							ServiceLevel:            to.Ptr(armnetapp.ServiceLevelPremium),
							SubnetID:                to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
							ThroughputMibps:         to.Ptr[float32](10),
							UsageThreshold:          to.Ptr[int64](107374182400),
							VolumeSpecName:          to.Ptr("shared"),
						},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
