package armdevopsinfrastructure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devopsinfrastructure/armdevopsinfrastructure"
)

// Generated from example definition: 2024-10-19/Sku_ListByLocation.json
func ExampleSKUClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevopsinfrastructure.NewClientFactory("a2e95d27-c161-4b61-bda4-11512c14c2c2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSKUClient().NewListByLocationPager("eastus", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armdevopsinfrastructure.SKUClientListByLocationResponse{
		// 	ResourceSKUListResult: armdevopsinfrastructure.ResourceSKUListResult{
		// 		Value: []*armdevopsinfrastructure.ResourceSKU{
		// 			{
		// 				Properties: &armdevopsinfrastructure.ResourceSKUProperties{
		// 					ResourceType: to.Ptr("virtualMachines"),
		// 					Tier: to.Ptr("Basic"),
		// 					Size: to.Ptr("A0"),
		// 					Family: to.Ptr("basicAFamily"),
		// 					Locations: []*string{
		// 						to.Ptr("eastus"),
		// 					},
		// 					LocationInfo: []*armdevopsinfrastructure.ResourceSKULocationInfo{
		// 						{
		// 							Location: to.Ptr("eastus"),
		// 							Zones: []*string{
		// 							},
		// 							ZoneDetails: []*armdevopsinfrastructure.ResourceSKUZoneDetails{
		// 							},
		// 						},
		// 					},
		// 					Capabilities: []*armdevopsinfrastructure.ResourceSKUCapabilities{
		// 						{
		// 							Name: to.Ptr("MaxResourceVolumeMB"),
		// 							Value: to.Ptr("20480"),
		// 						},
		// 						{
		// 							Name: to.Ptr("OSVhdSizeMB"),
		// 							Value: to.Ptr("1047552"),
		// 						},
		// 						{
		// 							Name: to.Ptr("vCPUs"),
		// 							Value: to.Ptr("1"),
		// 						},
		// 						{
		// 							Name: to.Ptr("MemoryPreservingMaintenanceSupported"),
		// 							Value: to.Ptr("True"),
		// 						},
		// 						{
		// 							Name: to.Ptr("HyperVGenerations"),
		// 							Value: to.Ptr("V1"),
		// 						},
		// 						{
		// 							Name: to.Ptr("MemoryGB"),
		// 							Value: to.Ptr("0.75"),
		// 						},
		// 						{
		// 							Name: to.Ptr("MaxDataDiskCount"),
		// 							Value: to.Ptr("1"),
		// 						},
		// 						{
		// 							Name: to.Ptr("CpuArchitectureType"),
		// 							Value: to.Ptr("x64"),
		// 						},
		// 						{
		// 							Name: to.Ptr("LowPriorityCapable"),
		// 							Value: to.Ptr("False"),
		// 						},
		// 						{
		// 							Name: to.Ptr("PremiumIO"),
		// 							Value: to.Ptr("False"),
		// 						},
		// 						{
		// 							Name: to.Ptr("VMDeploymentTypes"),
		// 							Value: to.Ptr("IaaS"),
		// 						},
		// 						{
		// 							Name: to.Ptr("vCPUsAvailable"),
		// 							Value: to.Ptr("1"),
		// 						},
		// 						{
		// 							Name: to.Ptr("ACUs"),
		// 							Value: to.Ptr("50"),
		// 						},
		// 						{
		// 							Name: to.Ptr("vCPUsPerCore"),
		// 							Value: to.Ptr("1"),
		// 						},
		// 						{
		// 							Name: to.Ptr("EphemeralOSDiskSupported"),
		// 							Value: to.Ptr("False"),
		// 						},
		// 						{
		// 							Name: to.Ptr("EncryptionAtHostSupported"),
		// 							Value: to.Ptr("False"),
		// 						},
		// 						{
		// 							Name: to.Ptr("CapacityReservationSupported"),
		// 							Value: to.Ptr("False"),
		// 						},
		// 						{
		// 							Name: to.Ptr("AcceleratedNetworkingEnabled"),
		// 							Value: to.Ptr("False"),
		// 						},
		// 						{
		// 							Name: to.Ptr("RdmaEnabled"),
		// 							Value: to.Ptr("False"),
		// 						},
		// 						{
		// 							Name: to.Ptr("MaxNetworkInterfaces"),
		// 							Value: to.Ptr("2"),
		// 						},
		// 					},
		// 					Restrictions: []*armdevopsinfrastructure.ResourceSKURestrictions{
		// 						{
		// 							Type: to.Ptr(armdevopsinfrastructure.ResourceSKURestrictionsTypeLocation),
		// 							Values: []*string{
		// 								to.Ptr("eastus"),
		// 							},
		// 							RestrictionInfo: &armdevopsinfrastructure.ResourceSKURestrictionInfo{
		// 								Locations: []*string{
		// 									to.Ptr("eastus"),
		// 								},
		// 							},
		// 							ReasonCode: to.Ptr(armdevopsinfrastructure.ResourceSKURestrictionsReasonCodeNotAvailableForSubscription),
		// 						},
		// 						{
		// 							Type: to.Ptr(armdevopsinfrastructure.ResourceSKURestrictionsTypeZone),
		// 							Values: []*string{
		// 								to.Ptr("eastus"),
		// 							},
		// 							RestrictionInfo: &armdevopsinfrastructure.ResourceSKURestrictionInfo{
		// 								Locations: []*string{
		// 									to.Ptr("eastus"),
		// 								},
		// 								Zones: []*string{
		// 									to.Ptr("1"),
		// 									to.Ptr("2"),
		// 									to.Ptr("3"),
		// 								},
		// 							},
		// 							ReasonCode: to.Ptr(armdevopsinfrastructure.ResourceSKURestrictionsReasonCodeNotAvailableForSubscription),
		// 						},
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/a2e95d27-c161-4b61-bda4-11512c14c2c2/providers/Microsoft.DevOpsInfrastructure/locations/eastus/skus/Basic_A0"),
		// 				Name: to.Ptr("Basic_A0"),
		// 			},
		// 			{
		// 				Properties: &armdevopsinfrastructure.ResourceSKUProperties{
		// 					ResourceType: to.Ptr("virtualMachines"),
		// 					Tier: to.Ptr("Standard"),
		// 					Size: to.Ptr("A2_v2"),
		// 					Family: to.Ptr("standardAv2Family"),
		// 					Locations: []*string{
		// 						to.Ptr("eastus"),
		// 					},
		// 					LocationInfo: []*armdevopsinfrastructure.ResourceSKULocationInfo{
		// 						{
		// 							Location: to.Ptr("eastus"),
		// 							Zones: []*string{
		// 								to.Ptr("1"),
		// 								to.Ptr("2"),
		// 								to.Ptr("3"),
		// 							},
		// 							ZoneDetails: []*armdevopsinfrastructure.ResourceSKUZoneDetails{
		// 							},
		// 						},
		// 					},
		// 					Capabilities: []*armdevopsinfrastructure.ResourceSKUCapabilities{
		// 						{
		// 							Name: to.Ptr("MaxResourceVolumeMB"),
		// 							Value: to.Ptr("20480"),
		// 						},
		// 						{
		// 							Name: to.Ptr("OSVhdSizeMB"),
		// 							Value: to.Ptr("1047552"),
		// 						},
		// 						{
		// 							Name: to.Ptr("vCPUs"),
		// 							Value: to.Ptr("2"),
		// 						},
		// 						{
		// 							Name: to.Ptr("MemoryPreservingMaintenanceSupported"),
		// 							Value: to.Ptr("True"),
		// 						},
		// 						{
		// 							Name: to.Ptr("HyperVGenerations"),
		// 							Value: to.Ptr("V1"),
		// 						},
		// 						{
		// 							Name: to.Ptr("MemoryGB"),
		// 							Value: to.Ptr("4"),
		// 						},
		// 						{
		// 							Name: to.Ptr("MaxDataDiskCount"),
		// 							Value: to.Ptr("4"),
		// 						},
		// 						{
		// 							Name: to.Ptr("CpuArchitectureType"),
		// 							Value: to.Ptr("x64"),
		// 						},
		// 						{
		// 							Name: to.Ptr("LowPriorityCapable"),
		// 							Value: to.Ptr("True"),
		// 						},
		// 						{
		// 							Name: to.Ptr("PremiumIO"),
		// 							Value: to.Ptr("False"),
		// 						},
		// 						{
		// 							Name: to.Ptr("VMDeploymentTypes"),
		// 							Value: to.Ptr("PaaS,IaaS"),
		// 						},
		// 						{
		// 							Name: to.Ptr("vCPUsAvailable"),
		// 							Value: to.Ptr("2"),
		// 						},
		// 						{
		// 							Name: to.Ptr("ACUs"),
		// 							Value: to.Ptr("100"),
		// 						},
		// 						{
		// 							Name: to.Ptr("vCPUsPerCore"),
		// 							Value: to.Ptr("1"),
		// 						},
		// 						{
		// 							Name: to.Ptr("CombinedTempDiskAndCachedIOPS"),
		// 							Value: to.Ptr("2000"),
		// 						},
		// 						{
		// 							Name: to.Ptr("CombinedTempDiskAndCachedReadBytesPerSecond"),
		// 							Value: to.Ptr("41943040"),
		// 						},
		// 						{
		// 							Name: to.Ptr("CombinedTempDiskAndCachedWriteBytesPerSecond"),
		// 							Value: to.Ptr("20971520"),
		// 						},
		// 						{
		// 							Name: to.Ptr("UncachedDiskIOPS"),
		// 							Value: to.Ptr("3200"),
		// 						},
		// 						{
		// 							Name: to.Ptr("UncachedDiskBytesPerSecond"),
		// 							Value: to.Ptr("48000000"),
		// 						},
		// 						{
		// 							Name: to.Ptr("EphemeralOSDiskSupported"),
		// 							Value: to.Ptr("False"),
		// 						},
		// 						{
		// 							Name: to.Ptr("EncryptionAtHostSupported"),
		// 							Value: to.Ptr("False"),
		// 						},
		// 						{
		// 							Name: to.Ptr("CapacityReservationSupported"),
		// 							Value: to.Ptr("True"),
		// 						},
		// 						{
		// 							Name: to.Ptr("AcceleratedNetworkingEnabled"),
		// 							Value: to.Ptr("False"),
		// 						},
		// 						{
		// 							Name: to.Ptr("RdmaEnabled"),
		// 							Value: to.Ptr("False"),
		// 						},
		// 						{
		// 							Name: to.Ptr("MaxNetworkInterfaces"),
		// 							Value: to.Ptr("2"),
		// 						},
		// 					},
		// 					Restrictions: []*armdevopsinfrastructure.ResourceSKURestrictions{
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/a2e95d27-c161-4b61-bda4-11512c14c2c2/providers/Microsoft.DevOpsInfrastructure/locations/eastus/skus/Standard_A2_v2"),
		// 				Name: to.Ptr("Standard_A2_v2"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
