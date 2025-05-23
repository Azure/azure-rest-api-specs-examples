package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/AssessedSqlMachinesOperations_Get_MaximumSet_Gen.json
func ExampleAssessedSQLMachinesOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssessedSQLMachinesOperationsClient().Get(ctx, "rgmigrate", "fci-test6904project", "test_fci_hadr", "test_swagger_1", "cc64c9dc-b38e-435d-85ad-d509df5d92c6", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssessedSQLMachine = armmigrationassessment.AssessedSQLMachine{
	// 	Name: to.Ptr("cc64c9dc-b38e-435d-85ad-d509df5d92c6"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/groups/sqlAssessments/assessedSqlMachines"),
	// 	ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/bansalankit-rg/providers/Microsoft.Migrate/assessmentprojects/fci-ankit-test6904project/groups/test_fci_hadr/sqlAssessments/test_swagger_1/assessedSqlMachines/cc64c9dc-b38e-435d-85ad-d509df5d92c6"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 	},
	// 	Properties: &armmigrationassessment.AssessedSQLMachineProperties{
	// 		Type: to.Ptr(armmigrationassessment.AssessedMachineTypeSQLAssessedMachine),
	// 		Description: to.Ptr(""),
	// 		BootType: to.Ptr(armmigrationassessment.MachineBootTypeBios),
	// 		ConfidenceRatingInPercentage: to.Ptr[float32](0),
	// 		CostComponents: []*armmigrationassessment.CostComponent{
	// 			{
	// 				Name: to.Ptr(armmigrationassessment.CostComponentNameMonthlySecurityCost),
	// 				Value: to.Ptr[float32](0),
	// 		}},
	// 		CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-08T05:22:34.841Z"); return t}()),
	// 		DatacenterMachineArmID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourcegroups/bansalankit-rg/providers/microsoft.offazure/vmwaresites/fci-ankit-test6904site/machines/10-150-91-150-202bfa1a-ad9f-414f-91ff-61f5ca0540aa_5002fdf3-5889-c2f8-86e5-3119b251331b"),
	// 		DatacenterManagementServerArmID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/bansalankit-rg/providers/Microsoft.OffAzure/VMwareSites/fci-ankit-test6904site/vcenters/10-150-91-150-202bfa1a-ad9f-414f-91ff-61f5ca0540aa"),
	// 		DatacenterManagementServerName: to.Ptr("10.150.91.150"),
	// 		Disks: map[string]*armmigrationassessment.AssessedDiskData{
	// 			"6000C293-381c-8460-c29b-ef937230e629": &armmigrationassessment.AssessedDiskData{
	// 				Name: to.Ptr("6000C293-381c-8460-c29b-ef937230e629"),
	// 				DisplayName: to.Ptr("scsi0:0"),
	// 				GigabytesProvisioned: to.Ptr[float32](100),
	// 				MegabytesPerSecondOfRead: to.Ptr[float32](0),
	// 				MegabytesPerSecondOfWrite: to.Ptr[float32](0),
	// 				MonthlyStorageCost: to.Ptr[float32](19.71),
	// 				NumberOfReadOperationsPerSecond: to.Ptr[float32](0),
	// 				NumberOfWriteOperationsPerSecond: to.Ptr[float32](0),
	// 				RecommendDiskThroughputInMbps: to.Ptr[float32](100),
	// 				RecommendedDiskIops: to.Ptr[float32](500),
	// 				RecommendedDiskSize: to.Ptr(armmigrationassessment.AzureDiskSizePremiumP10),
	// 				RecommendedDiskSizeGigabytes: to.Ptr[int32](128),
	// 				RecommendedDiskType: to.Ptr(armmigrationassessment.AzureDiskTypePremium),
	// 				Suitability: to.Ptr(armmigrationassessment.CloudSuitabilitySuitable),
	// 				SuitabilityDetail: to.Ptr(armmigrationassessment.AzureDiskSuitabilityDetail("NumberOfReadOperationsPerSecondMissing, NumberOfWriteOperationsPerSecondMissing, MegabytesPerSecondOfReadMissing, MegabytesPerSecondOfWriteMissing")),
	// 				SuitabilityExplanation: to.Ptr(armmigrationassessment.AzureDiskSuitabilityExplanationNotApplicable),
	// 			},
	// 		},
	// 		DisplayName: to.Ptr("SQLHAVM17"),
	// 		MegabytesOfMemory: to.Ptr[float32](4096),
	// 		MigrationGuidelines: []*armmigrationassessment.SQLMigrationGuideline{
	// 			{
	// 				GuidelineID: to.Ptr("MigrateNodeSinceItHasActiveFCIInstance"),
	// 				MigrationGuidelineCategory: to.Ptr(armmigrationassessment.SQLMigrationGuidelineCategoryFailoverCluterInstanceGuideLine),
	// 				MigrationGuidelineContext: []*armmigrationassessment.MigrationGuidelineContext{
	// 				},
	// 			},
	// 			{
	// 				GuidelineID: to.Ptr("UseSharedDiskWitnessForWSFCIfNeeded"),
	// 				MigrationGuidelineCategory: to.Ptr(armmigrationassessment.SQLMigrationGuidelineCategoryFailoverCluterInstanceGuideLine),
	// 				MigrationGuidelineContext: []*armmigrationassessment.MigrationGuidelineContext{
	// 				},
	// 			},
	// 			{
	// 				GuidelineID: to.Ptr("ReconfigureAnyAvailabilityGroupsAfterMigration"),
	// 				MigrationGuidelineCategory: to.Ptr(armmigrationassessment.SQLMigrationGuidelineCategoryAvailabilityGroupGuideline),
	// 				MigrationGuidelineContext: []*armmigrationassessment.MigrationGuidelineContext{
	// 				},
	// 		}},
	// 		MonthlyBandwidthCost: to.Ptr[float32](0),
	// 		MonthlyComputeCost: to.Ptr[float32](32.94432),
	// 		MonthlyStorageCost: to.Ptr[float32](19.71),
	// 		NetworkAdapters: map[string]*armmigrationassessment.SQLAssessedNetworkAdapter{
	// 			"4000": &armmigrationassessment.SQLAssessedNetworkAdapter{
	// 				Name: to.Ptr("4000"),
	// 				DisplayName: to.Ptr("VM Network"),
	// 				IPAddresses: []*string{
	// 					to.Ptr("2404:f801:4800:20:418c:eec9:86c5:aea1"),
	// 					to.Ptr("2404:f801:4800:20:6113:238d:17f0:b246"),
	// 					to.Ptr("10.150.91.119"),
	// 					to.Ptr("10.150.91.138"),
	// 					to.Ptr("10.150.91.139")},
	// 					MacAddress: to.Ptr("00:50:56:82:b7:83"),
	// 					MegabytesPerSecondReceived: to.Ptr[float32](0),
	// 					MegabytesPerSecondTransmitted: to.Ptr[float32](0),
	// 					MonthlyBandwidthCosts: to.Ptr[float32](0),
	// 					NetGigabytesTransmittedPerMonth: to.Ptr[float32](0),
	// 					Suitability: to.Ptr(armmigrationassessment.CloudSuitabilitySuitable),
	// 					SuitabilityDetail: to.Ptr(armmigrationassessment.AzureNetworkAdapterSuitabilityDetail("MegabytesOfDataTransmittedMissing, MegabytesOfDataRecievedMissing")),
	// 					SuitabilityExplanation: to.Ptr(armmigrationassessment.AzureNetworkAdapterSuitabilityExplanationNotApplicable),
	// 				},
	// 			},
	// 			NumberOfCores: to.Ptr[int32](2),
	// 			OperatingSystemArchitecture: to.Ptr(armmigrationassessment.GuestOperatingSystemArchitectureX64),
	// 			OperatingSystemName: to.Ptr("Microsoft Windows Server 2019 (64-bit)"),
	// 			OperatingSystemType: to.Ptr("windowsGuest"),
	// 			PercentageCoresUtilization: to.Ptr[float32](0),
	// 			PercentageMemoryUtilization: to.Ptr[float32](0),
	// 			ProductSupportStatus: &armmigrationassessment.ProductSupportStatus{
	// 				CurrentEsuYear: to.Ptr("Unknown"),
	// 				EsuStatus: to.Ptr("Unknown"),
	// 				Eta: to.Ptr[int32](64),
	// 				ExtendedSecurityUpdateYear1EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 				ExtendedSecurityUpdateYear2EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 				ExtendedSecurityUpdateYear3EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 				ExtendedSupportEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2029-01-09T00:00:00.000Z"); return t}()),
	// 				MainstreamEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-09T00:00:00.000Z"); return t}()),
	// 				ServicePackStatus: to.Ptr("Unknown"),
	// 				SupportStatus: to.Ptr("Mainstream"),
	// 			},
	// 			RecommendedVMFamily: to.Ptr(armmigrationassessment.AzureVMFamilyDasv4Series),
	// 			RecommendedVMSize: to.Ptr(armmigrationassessment.AzureVMSizeStandardD2AsV4),
	// 			RecommendedVMSizeMegabytesOfMemory: to.Ptr[float32](8192),
	// 			RecommendedVMSizeNumberOfCores: to.Ptr[int32](2),
	// 			SecuritySuitability: to.Ptr(armmigrationassessment.CloudSuitabilityUnknown),
	// 			SizingCriterion: to.Ptr(armmigrationassessment.AssessmentSizingCriterionPerformanceBased),
	// 			SQLInstances: []*armmigrationassessment.AssessedSQLInstanceSummary{
	// 				{
	// 					InstanceID: to.Ptr("291313e5-e25f-4b6b-9f21-165a2dd03650"),
	// 					InstanceName: to.Ptr("MSSQLSERVER"),
	// 					IsClustered: to.Ptr(true),
	// 					IsHighAvailabilityEnabled: to.Ptr(true),
	// 					SQLEdition: to.Ptr("Enterprise Edition: Core-based Licensing (64-bit)"),
	// 					SQLFciState: to.Ptr(armmigrationassessment.SQLFCIStateActive),
	// 					SQLInstanceEntityID: to.Ptr("b11d90cc-b528-49e4-aac2-1c9a53209f8e"),
	// 					SQLInstanceSdsArmID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourcegroups/bansalankit-rg/providers/microsoft.offazure/mastersites/fci-ankit-test6065mastersite/sqlsites/fci-ankit-test6065sqlsites/sqlservers/291313e5-e25f-4b6b-9f21-165a2dd03650"),
	// 					SQLVersion: to.Ptr("SQL Server 2014"),
	// 			}},
	// 			Suitability: to.Ptr(armmigrationassessment.CloudSuitabilitySuitable),
	// 			SuitabilityDetail: to.Ptr(armmigrationassessment.AzureVMSuitabilityDetail("PercentageOfCoresUtilizedMissing, PercentageOfMemoryUtilizedMissing")),
	// 			SuitabilityExplanation: to.Ptr(armmigrationassessment.AzureVMSuitabilityExplanationNotApplicable),
	// 			UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-08T05:23:39.174Z"); return t}()),
	// 		},
	// 	}
}
