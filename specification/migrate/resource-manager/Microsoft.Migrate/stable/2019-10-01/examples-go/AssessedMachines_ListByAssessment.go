package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/AssessedMachines_ListByAssessment.json
func ExampleAssessedMachinesClient_NewListByAssessmentPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssessedMachinesClient().NewListByAssessmentPager("abgoyal-westEurope", "abgoyalWEselfhostb72bproject", "Test1", "assessment_5_9_2019_16_22_14", nil)
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
		// page.AssessedMachineResultList = armmigrate.AssessedMachineResultList{
		// 	Value: []*armmigrate.AssessedMachine{
		// 		{
		// 			Name: to.Ptr("f57fe432-3bd2-486a-a83a-6f4d99f1a952"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects/groups/assessments/assessedMachines"),
		// 			ETag: to.Ptr("\"b300e5dd-0000-0d00-0000-5cd4065f0000\""),
		// 			ID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westeurope/providers/Microsoft.Migrate/assessmentprojects/abgoyalWEselfhostb72bproject/groups/Test1/assessments/assessment_5_9_2019_16_22_14/assessedMachines/f57fe432-3bd2-486a-a83a-6f4d99f1a952"),
		// 			Properties: &armmigrate.AssessedMachineProperties{
		// 				Description: to.Ptr("Microsoft Azure Migration Image on Windows Server 2016"),
		// 				BootType: to.Ptr(armmigrate.MachineBootTypeBIOS),
		// 				ConfidenceRatingInPercentage: to.Ptr[float64](0),
		// 				CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T10:52:15.778Z"); return t}()),
		// 				DatacenterMachineArmID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourcegroups/abgoyal-westeurope/providers/microsoft.offazure/vmwaresites/portalvcenterbc2fsite/machines/idclab-a360-fareast-corp-micros-86617dcf-effe-59ad-8c3a-cdd3ea7300d3_52e1f68c-bea5-19ff-d0ad-6a94b79a286f"),
		// 				DatacenterManagementServerArmID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westEurope/providers/Microsoft.OffAzure/VMwareSites/PortalvCenterbc2fsite/vcenters/idclab-a360-fareast-corp-micros-86617dcf-effe-59ad-8c3a-cdd3ea7300d3"),
		// 				DatacenterManagementServerName: to.Ptr("IDCLAB-A360.fareast.corp.microsoft.com"),
		// 				Disks: map[string]*armmigrate.AssessedDisk{
		// 					"6000C299-02f5-d137-9bab-8a8ee7b192a0": &armmigrate.AssessedDisk{
		// 						Name: to.Ptr("6000C299-02f5-d137-9bab-8a8ee7b192a0"),
		// 						DisplayName: to.Ptr("scsi0:0"),
		// 						GigabytesForRecommendedDiskSize: to.Ptr[int32](128),
		// 						GigabytesProvisioned: to.Ptr[float64](80),
		// 						MegabytesPerSecondOfRead: to.Ptr[float64](0),
		// 						MegabytesPerSecondOfWrite: to.Ptr[float64](0),
		// 						MonthlyStorageCost: to.Ptr[float64](5.888),
		// 						NumberOfReadOperationsPerSecond: to.Ptr[float64](0),
		// 						NumberOfWriteOperationsPerSecond: to.Ptr[float64](0),
		// 						RecommendedDiskSize: to.Ptr(armmigrate.AzureDiskSizeStandardS10),
		// 						RecommendedDiskType: to.Ptr(armmigrate.AzureDiskTypeStandard),
		// 						Suitability: to.Ptr(armmigrate.CloudSuitabilitySuitable),
		// 						SuitabilityDetail: to.Ptr(armmigrate.AzureDiskSuitabilityDetail("NumberOfReadOperationsPerSecondMissing, NumberOfWriteOperationsPerSecondMissing, MegabytesPerSecondOfReadMissing, MegabytesPerSecondOfWriteMissing")),
		// 						SuitabilityExplanation: to.Ptr(armmigrate.AzureDiskSuitabilityExplanationNotApplicable),
		// 					},
		// 				},
		// 				DisplayName: to.Ptr("SHubhamVMNew"),
		// 				MegabytesOfMemory: to.Ptr[float64](16384),
		// 				MegabytesOfMemoryForRecommendedSize: to.Ptr[float64](16384),
		// 				MonthlyBandwidthCost: to.Ptr[float64](0),
		// 				MonthlyComputeCostForRecommendedSize: to.Ptr[float64](101.138616),
		// 				MonthlyPremiumStorageCost: to.Ptr[float64](0),
		// 				MonthlyStandardSSDStorageCost: to.Ptr[float64](0),
		// 				MonthlyStorageCost: to.Ptr[float64](5.888),
		// 				NetworkAdapters: map[string]*armmigrate.AssessedNetworkAdapter{
		// 					"4000": &armmigrate.AssessedNetworkAdapter{
		// 						DisplayName: to.Ptr("VM Network"),
		// 						IPAddresses: []*string{
		// 						},
		// 						MacAddress: to.Ptr("00:0c:29:ad:13:d3"),
		// 						MegabytesPerSecondReceived: to.Ptr[float64](0),
		// 						MegabytesPerSecondTransmitted: to.Ptr[float64](0),
		// 						MonthlyBandwidthCosts: to.Ptr[float64](0),
		// 						NetGigabytesTransmittedPerMonth: to.Ptr[float64](0),
		// 						Suitability: to.Ptr(armmigrate.CloudSuitabilitySuitable),
		// 						SuitabilityDetail: to.Ptr(armmigrate.AzureNetworkAdapterSuitabilityDetailMegabytesOfDataTransmittedMissing),
		// 						SuitabilityExplanation: to.Ptr(armmigrate.AzureNetworkAdapterSuitabilityExplanationNotApplicable),
		// 					},
		// 				},
		// 				NumberOfCores: to.Ptr[int32](8),
		// 				NumberOfCoresForRecommendedSize: to.Ptr[int32](8),
		// 				OperatingSystemName: to.Ptr("Microsoft Windows Server 2016 (64-bit)"),
		// 				OperatingSystemType: to.Ptr("windowsGuest"),
		// 				PercentageCoresUtilization: to.Ptr[float64](0),
		// 				PercentageMemoryUtilization: to.Ptr[float64](0),
		// 				RecommendedSize: to.Ptr(armmigrate.AzureVMSizeStandardF8SV2),
		// 				Suitability: to.Ptr(armmigrate.CloudSuitabilitySuitable),
		// 				SuitabilityDetail: to.Ptr(armmigrate.AzureVMSuitabilityDetail("PercentageOfCoresUtilizedMissing, PercentageOfMemoryUtilizedMissing")),
		// 				SuitabilityExplanation: to.Ptr(armmigrate.AzureVMSuitabilityExplanationNotApplicable),
		// 				UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T10:52:15.778Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("3b4a34a6-c729-46d2-bfd1-bcb52cc4935e"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects/groups/assessments/assessedMachines"),
		// 			ETag: to.Ptr("\"b300e6dd-0000-0d00-0000-5cd4065f0000\""),
		// 			ID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westeurope/providers/Microsoft.Migrate/assessmentprojects/abgoyalWEselfhostb72bproject/groups/Test1/assessments/assessment_5_9_2019_16_22_14/assessedMachines/3b4a34a6-c729-46d2-bfd1-bcb52cc4935e"),
		// 			Properties: &armmigrate.AssessedMachineProperties{
		// 				Description: to.Ptr("Microsoft Azure Migration Image on Windows Server 2016"),
		// 				BootType: to.Ptr(armmigrate.MachineBootTypeBIOS),
		// 				ConfidenceRatingInPercentage: to.Ptr[float64](0),
		// 				CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T10:52:15.778Z"); return t}()),
		// 				DatacenterMachineArmID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourcegroups/abgoyal-westeurope/providers/microsoft.offazure/vmwaresites/portalvcenterbc2fsite/machines/idclab-a360-fareast-corp-micros-86617dcf-effe-59ad-8c3a-cdd3ea7300d3_50296915-8b4b-5c82-79a1-adf3966acb6b"),
		// 				DatacenterManagementServerArmID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westEurope/providers/Microsoft.OffAzure/VMwareSites/PortalvCenterbc2fsite/vcenters/idclab-a360-fareast-corp-micros-86617dcf-effe-59ad-8c3a-cdd3ea7300d3"),
		// 				DatacenterManagementServerName: to.Ptr("IDCLAB-A360.fareast.corp.microsoft.com"),
		// 				Disks: map[string]*armmigrate.AssessedDisk{
		// 					"6000C298-8305-5635-e618-3a8675c42495": &armmigrate.AssessedDisk{
		// 						Name: to.Ptr("6000C298-8305-5635-e618-3a8675c42495"),
		// 						DisplayName: to.Ptr("scsi0:0"),
		// 						GigabytesForRecommendedDiskSize: to.Ptr[int32](128),
		// 						GigabytesProvisioned: to.Ptr[float64](80),
		// 						MegabytesPerSecondOfRead: to.Ptr[float64](0),
		// 						MegabytesPerSecondOfWrite: to.Ptr[float64](0),
		// 						MonthlyStorageCost: to.Ptr[float64](5.888),
		// 						NumberOfReadOperationsPerSecond: to.Ptr[float64](0),
		// 						NumberOfWriteOperationsPerSecond: to.Ptr[float64](0),
		// 						RecommendedDiskSize: to.Ptr(armmigrate.AzureDiskSizeStandardS10),
		// 						RecommendedDiskType: to.Ptr(armmigrate.AzureDiskTypeStandard),
		// 						Suitability: to.Ptr(armmigrate.CloudSuitabilitySuitable),
		// 						SuitabilityDetail: to.Ptr(armmigrate.AzureDiskSuitabilityDetail("NumberOfReadOperationsPerSecondMissing, NumberOfWriteOperationsPerSecondMissing, MegabytesPerSecondOfReadMissing, MegabytesPerSecondOfWriteMissing")),
		// 						SuitabilityExplanation: to.Ptr(armmigrate.AzureDiskSuitabilityExplanationNotApplicable),
		// 					},
		// 				},
		// 				DisplayName: to.Ptr("testfpl1"),
		// 				MegabytesOfMemory: to.Ptr[float64](16384),
		// 				MegabytesOfMemoryForRecommendedSize: to.Ptr[float64](16384),
		// 				MonthlyBandwidthCost: to.Ptr[float64](0),
		// 				MonthlyComputeCostForRecommendedSize: to.Ptr[float64](101.138616),
		// 				MonthlyPremiumStorageCost: to.Ptr[float64](0),
		// 				MonthlyStandardSSDStorageCost: to.Ptr[float64](0),
		// 				MonthlyStorageCost: to.Ptr[float64](5.888),
		// 				NetworkAdapters: map[string]*armmigrate.AssessedNetworkAdapter{
		// 					"4000": &armmigrate.AssessedNetworkAdapter{
		// 						DisplayName: to.Ptr("VM Network"),
		// 						IPAddresses: []*string{
		// 						},
		// 						MacAddress: to.Ptr("00:50:56:a9:35:ca"),
		// 						MegabytesPerSecondReceived: to.Ptr[float64](0),
		// 						MegabytesPerSecondTransmitted: to.Ptr[float64](0),
		// 						MonthlyBandwidthCosts: to.Ptr[float64](0),
		// 						NetGigabytesTransmittedPerMonth: to.Ptr[float64](0),
		// 						Suitability: to.Ptr(armmigrate.CloudSuitabilitySuitable),
		// 						SuitabilityDetail: to.Ptr(armmigrate.AzureNetworkAdapterSuitabilityDetailMegabytesOfDataTransmittedMissing),
		// 						SuitabilityExplanation: to.Ptr(armmigrate.AzureNetworkAdapterSuitabilityExplanationNotApplicable),
		// 					},
		// 				},
		// 				NumberOfCores: to.Ptr[int32](8),
		// 				NumberOfCoresForRecommendedSize: to.Ptr[int32](8),
		// 				OperatingSystemName: to.Ptr("Microsoft Windows Server 2016 (64-bit)"),
		// 				OperatingSystemType: to.Ptr("windowsguest"),
		// 				PercentageCoresUtilization: to.Ptr[float64](0),
		// 				PercentageMemoryUtilization: to.Ptr[float64](0),
		// 				RecommendedSize: to.Ptr(armmigrate.AzureVMSizeStandardF8SV2),
		// 				Suitability: to.Ptr(armmigrate.CloudSuitabilitySuitable),
		// 				SuitabilityDetail: to.Ptr(armmigrate.AzureVMSuitabilityDetail("PercentageOfCoresUtilizedMissing, PercentageOfMemoryUtilizedMissing")),
		// 				SuitabilityExplanation: to.Ptr(armmigrate.AzureVMSuitabilityExplanationNotApplicable),
		// 				UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T10:52:15.778Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
