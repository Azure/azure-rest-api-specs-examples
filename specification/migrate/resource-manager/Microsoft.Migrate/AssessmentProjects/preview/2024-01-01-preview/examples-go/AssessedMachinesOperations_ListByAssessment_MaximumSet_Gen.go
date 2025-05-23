package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/AssessedMachinesOperations_ListByAssessment_MaximumSet_Gen.json
func ExampleAssessedMachinesOperationsClient_NewListByAssessmentPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssessedMachinesOperationsClient().NewListByAssessmentPager("rgopenapi", "sloqixzfjk", "kjuepxerwseq", "rhzcmubwrrkhtocsibu", &armmigrationassessment.AssessedMachinesOperationsClientListByAssessmentOptions{Filter: to.Ptr("sbkdovsfqldhdb"),
		PageSize:          to.Ptr[int32](10),
		ContinuationToken: to.Ptr("hbyseetshbplfkjmpjhsiurqgt"),
		TotalRecordCount:  to.Ptr[int32](25),
	})
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
		// page.AssessedMachineListResult = armmigrationassessment.AssessedMachineListResult{
		// 	Value: []*armmigrationassessment.AssessedMachine{
		// 		{
		// 			Name: to.Ptr("riigi"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects/groups/assessments/assessedMachines"),
		// 			ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawrg/providers/Microsoft.Migrate/assessmentprojects/app18700project/groups/kuchatur-test/assessments/asm1/assessedMachines/50fa865c-8c5c-4371-b7eb-5b900d7f9451"),
		// 			SystemData: &armmigrationassessment.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-26T12:46:17.892Z"); return t}()),
		// 				CreatedBy: to.Ptr("jsfqnyqqwykkv"),
		// 				CreatedByType: to.Ptr(armmigrationassessment.CreatedByType("sakanwar")),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-26T12:46:17.892Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("sakanwar"),
		// 				LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmigrationassessment.AssessedMachineProperties{
		// 				Type: to.Ptr(armmigrationassessment.AssessedMachineTypeAssessedMachine),
		// 				BootType: to.Ptr(armmigrationassessment.MachineBootTypeBios),
		// 				ConfidenceRatingInPercentage: to.Ptr[float32](14),
		// 				CostComponents: []*armmigrationassessment.CostComponent{
		// 					{
		// 						Name: to.Ptr(armmigrationassessment.CostComponentNameMonthlyAzureHybridCostSavings),
		// 						Value: to.Ptr[float32](273.792),
		// 					},
		// 					{
		// 						Name: to.Ptr(armmigrationassessment.CostComponentNameMonthlySecurityCost),
		// 						Value: to.Ptr[float32](14.88),
		// 					},
		// 					{
		// 						Name: to.Ptr(armmigrationassessment.CostComponentNameMonthlyPremiumV2StorageCost),
		// 						Value: to.Ptr[float32](25.141248),
		// 				}},
		// 				CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-26T09:14:13.816Z"); return t}()),
		// 				DatacenterMachineArmID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourcegroups/ayagrawrg/providers/microsoft.offazure/vmwaresites/app18700site/machines/idclab-vcen67-fareast-corp-micr-d991d5f4-63f1-41c3-ad6a-40253e24ffa3_50150976-65ec-de99-1b54-a52710c3066e"),
		// 				DatacenterManagementServerArmID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawRG/providers/Microsoft.OffAzure/VMwareSites/app18700site/vcenters/idclab-vcen67-fareast-corp-micr-d991d5f4-63f1-41c3-ad6a-40253e24ffa3"),
		// 				DatacenterManagementServerName: to.Ptr("idclab-vcen67.fareast.corp.microsoft.com"),
		// 				Disks: map[string]*armmigrationassessment.AssessedDisk{
		// 					"6000C297-455f-f81e-37b7-dc17da4433d4": &armmigrationassessment.AssessedDisk{
		// 						Name: to.Ptr("6000C297-455f-f81e-37b7-dc17da4433d4"),
		// 						DisplayName: to.Ptr("scsi0:0"),
		// 						GigabytesForRecommendedDiskSize: to.Ptr[int32](17),
		// 						GigabytesProvisioned: to.Ptr[float32](8),
		// 						MegabytesPerSecondOfRead: to.Ptr[float32](9),
		// 						MegabytesPerSecondOfWrite: to.Ptr[float32](29),
		// 						MonthlyStorageCost: to.Ptr[float32](5),
		// 						NumberOfReadOperationsPerSecond: to.Ptr[float32](12),
		// 						NumberOfWriteOperationsPerSecond: to.Ptr[float32](28),
		// 						RecommendDiskThroughputInMbps: to.Ptr[float32](125),
		// 						RecommendedDiskIops: to.Ptr[float32](3000),
		// 						RecommendedDiskSize: to.Ptr(armmigrationassessment.AzureDiskSizePremiumV2),
		// 						RecommendedDiskType: to.Ptr(armmigrationassessment.AzureDiskTypePremiumV2),
		// 						Suitability: to.Ptr(armmigrationassessment.CloudSuitabilitySuitable),
		// 						SuitabilityDetail: to.Ptr(armmigrationassessment.AzureDiskSuitabilityDetailNumberOfReadOperationsPerSecondMissing),
		// 						SuitabilityExplanation: to.Ptr(armmigrationassessment.AzureDiskSuitabilityExplanationNotApplicable),
		// 					},
		// 				},
		// 				DisplayName: to.Ptr("SQLTestDBVM28"),
		// 				Errors: []*armmigrationassessment.Error{
		// 				},
		// 				HostProcessor: &armmigrationassessment.ProcessorInfo{
		// 					Name: to.Ptr("Intel(R) Xeon(R) Gold 5220R CPU @ 2.20GHz"),
		// 					NumberOfCoresPerSocket: to.Ptr[int32](24),
		// 					NumberOfSockets: to.Ptr[int32](2),
		// 				},
		// 				MegabytesOfMemory: to.Ptr[float32](21),
		// 				MegabytesOfMemoryForRecommendedSize: to.Ptr[float32](27),
		// 				MonthlyBandwidthCost: to.Ptr[float32](4),
		// 				MonthlyComputeCostForRecommendedSize: to.Ptr[float32](1),
		// 				MonthlyPremiumStorageCost: to.Ptr[float32](6),
		// 				MonthlyStandardSsdStorageCost: to.Ptr[float32](4),
		// 				MonthlyStorageCost: to.Ptr[float32](10),
		// 				MonthlyUltraStorageCost: to.Ptr[float32](11),
		// 				NetworkAdapters: map[string]*armmigrationassessment.AssessedNetworkAdapter{
		// 					"4000": &armmigrationassessment.AssessedNetworkAdapter{
		// 						DisplayName: to.Ptr("VM Network"),
		// 						IPAddresses: []*string{
		// 							to.Ptr("2404:f801:4800:1c:c4c2:1cd:7154:a028")},
		// 							MacAddress: to.Ptr("00:50:56:95:98:e8"),
		// 							MegabytesPerSecondReceived: to.Ptr[float32](10),
		// 							MegabytesPerSecondTransmitted: to.Ptr[float32](24),
		// 							MonthlyBandwidthCosts: to.Ptr[float32](4),
		// 							NetGigabytesTransmittedPerMonth: to.Ptr[float32](6),
		// 							Suitability: to.Ptr(armmigrationassessment.CloudSuitabilitySuitable),
		// 							SuitabilityDetail: to.Ptr(armmigrationassessment.AzureNetworkAdapterSuitabilityDetail("NotApplicable")),
		// 							SuitabilityExplanation: to.Ptr(armmigrationassessment.AzureNetworkAdapterSuitabilityExplanation("MegabytesOfDataTransmittedMissing, MegabytesOfDataRecievedMissing")),
		// 						},
		// 					},
		// 					NumberOfCores: to.Ptr[int32](7),
		// 					NumberOfCoresForRecommendedSize: to.Ptr[int32](8),
		// 					OperatingSystemArchitecture: to.Ptr(armmigrationassessment.GuestOperatingSystemArchitectureX64),
		// 					OperatingSystemName: to.Ptr("Microsoft Windows Server 2016 or later (64-bit)"),
		// 					OperatingSystemType: to.Ptr("windowsGuest"),
		// 					PercentageCoresUtilization: to.Ptr[float32](24),
		// 					PercentageMemoryUtilization: to.Ptr[float32](5),
		// 					ProductSupportStatus: &armmigrationassessment.ProductSupportStatus{
		// 						CurrentEsuYear: to.Ptr("Unknown"),
		// 						EsuStatus: to.Ptr("Unknown"),
		// 						Eta: to.Ptr[int32](0),
		// 						ExtendedSecurityUpdateYear1EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-08T00:00:00.000Z"); return t}()),
		// 						ExtendedSecurityUpdateYear2EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-14T00:00:00.000Z"); return t}()),
		// 						ExtendedSecurityUpdateYear3EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-10-13T00:00:00.000Z"); return t}()),
		// 						ExtendedSupportEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-10T00:00:00.000Z"); return t}()),
		// 						MainstreamEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-09T00:00:00.000Z"); return t}()),
		// 						ServicePackStatus: to.Ptr("Unknown"),
		// 						SupportStatus: to.Ptr("Extended"),
		// 					},
		// 					RecommendedSize: to.Ptr(armmigrationassessment.AzureVMSizeStandardF4SV2),
		// 					Suitability: to.Ptr(armmigrationassessment.CloudSuitabilitySuitable),
		// 					SuitabilityDetail: to.Ptr(armmigrationassessment.AzureVMSuitabilityDetailPercentageOfCoresUtilizedMissing),
		// 					SuitabilityExplanation: to.Ptr(armmigrationassessment.AzureVMSuitabilityExplanationNotApplicable),
		// 					UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-26T09:14:13.816Z"); return t}()),
		// 				},
		// 		}},
		// 	}
	}
}
