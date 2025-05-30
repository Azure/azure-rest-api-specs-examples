package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/AvsAssessedMachinesOperations_Get_MaximumSet_Gen.json
func ExampleAvsAssessedMachinesOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAvsAssessedMachinesOperationsClient().Get(ctx, "ayagrawrg", "app18700project", "kuchatur-test", "asm2", "b6d6fc6f-796f-4c16-96af-a6d22e0f12f7", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AvsAssessedMachine = armmigrationassessment.AvsAssessedMachine{
	// 	Name: to.Ptr("18895660-c5e5-4247-8cfc-cd24e1fe57f3"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/groups/avsAssessments/avsAssessedMachines"),
	// 	ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawrg/providers/Microsoft.Migrate/assessmentprojects/app18700project/groups/kuchatur-test/avsAssessments/asm2/avsAssessedMachines/18895660-c5e5-4247-8cfc-cd24e1fe57f3"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-15T07:22:18.589Z"); return t}()),
	// 		CreatedBy: to.Ptr("bhjfiiwermbzqfoqxtxpjigj"),
	// 		CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-15T07:22:18.589Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("lrmhonmbodwva"),
	// 		LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmigrationassessment.AvsAssessedMachineProperties{
	// 		Type: to.Ptr(armmigrationassessment.AssessedMachineTypeAvsAssessedMachine),
	// 		Description: to.Ptr("Microsoft Azure Migration Image on Windows Server 2016"),
	// 		BootType: to.Ptr(armmigrationassessment.MachineBootTypeBios),
	// 		CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-26T13:35:59.232Z"); return t}()),
	// 		DatacenterMachineArmID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourcegroups/ayagrawrg/providers/microsoft.offazure/vmwaresites/app18700site/machines/idclab-vcen67-fareast-corp-micr-d991d5f4-63f1-41c3-ad6a-40253e24ffa3_501505aa-baaa-af9d-8315-5a45209fd255"),
	// 		DatacenterManagementServerArmID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawRG/providers/Microsoft.OffAzure/VMwareSites/app18700site/vcenters/idclab-vcen67-fareast-corp-micr-d991d5f4-63f1-41c3-ad6a-40253e24ffa3"),
	// 		DatacenterManagementServerName: to.Ptr("idclab-vcen67.fareast.corp.microsoft.com"),
	// 		Disks: map[string]*armmigrationassessment.AvsAssessedDisk{
	// 			"6000C297-455f-f81e-37b7-dc17da4433d4": &armmigrationassessment.AvsAssessedDisk{
	// 				Name: to.Ptr("6000C297-455f-f81e-37b7-dc17da4433d4"),
	// 				AssessedExternalStorageType: to.Ptr(armmigrationassessment.ExternalStorageTypeAnfPremium),
	// 				DisplayName: to.Ptr("scsi0:0"),
	// 				EstimatedDiskSizeInGB: to.Ptr[float32](14),
	// 				GigabytesProvisioned: to.Ptr[float32](80),
	// 				MegabytesPerSecondOfRead: to.Ptr[float32](0),
	// 				MegabytesPerSecondOfWrite: to.Ptr[float32](0),
	// 				NumberOfReadOperationsPerSecond: to.Ptr[float32](0),
	// 				NumberOfWriteOperationsPerSecond: to.Ptr[float32](0),
	// 				SuitabilityDetail: to.Ptr(armmigrationassessment.AzureDiskSuitabilityDetailNone),
	// 			},
	// 		},
	// 		DisplayName: to.Ptr("CustomerAE24Feb"),
	// 		Errors: []*armmigrationassessment.Error{
	// 		},
	// 		MegabytesOfMemory: to.Ptr[float32](32768),
	// 		NetworkAdapters: map[string]*armmigrationassessment.AvsAssessedNetworkAdapter{
	// 			"4000": &armmigrationassessment.AvsAssessedNetworkAdapter{
	// 				DisplayName: to.Ptr("VM Network"),
	// 				IPAddresses: []*string{
	// 					to.Ptr("2404:f801:4800:25:38fe:66de:64e4:5a26"),
	// 					to.Ptr("10.150.9.214")},
	// 					MacAddress: to.Ptr("00:50:56:95:7a:57"),
	// 					MegabytesPerSecondReceived: to.Ptr[float32](0),
	// 					MegabytesPerSecondTransmitted: to.Ptr[float32](0),
	// 				},
	// 			},
	// 			NumberOfCores: to.Ptr[int32](8),
	// 			OperatingSystemArchitecture: to.Ptr(armmigrationassessment.GuestOperatingSystemArchitectureX64),
	// 			OperatingSystemName: to.Ptr("Microsoft Windows Server 2016 or later (64-bit)"),
	// 			OperatingSystemType: to.Ptr("windowsGuest"),
	// 			PercentageCoresUtilization: to.Ptr[float32](0),
	// 			PercentageMemoryUtilization: to.Ptr[float32](0),
	// 			StorageInUseGB: to.Ptr[float32](0),
	// 			Suitability: to.Ptr(armmigrationassessment.CloudSuitabilityConditionallySuitable),
	// 			SuitabilityDetail: to.Ptr(armmigrationassessment.AzureAvsVMSuitabilityDetailNone),
	// 			SuitabilityExplanation: to.Ptr(armmigrationassessment.AzureAvsVMSuitabilityExplanationIPV6NotSupported),
	// 			UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-26T13:35:59.232Z"); return t}()),
	// 		},
	// 	}
}
