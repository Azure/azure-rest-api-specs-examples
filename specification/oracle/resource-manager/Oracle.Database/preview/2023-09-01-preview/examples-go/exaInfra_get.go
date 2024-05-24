package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/520e274d7d95fc6d1002dd3c1fcaf8d55d27f63e/specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/exaInfra_get.json
func ExampleCloudExadataInfrastructuresClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCloudExadataInfrastructuresClient().Get(ctx, "rg000", "infra1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CloudExadataInfrastructure = armoracledatabase.CloudExadataInfrastructure{
	// 	Type: to.Ptr("Oracle.Database/cloudExadataInfrastructures"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudExadataInfrastructures/infra1"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"tagK1": to.Ptr("tagV1"),
	// 	},
	// 	Properties: &armoracledatabase.CloudExadataInfrastructureProperties{
	// 		ActivatedStorageCount: to.Ptr[int32](1),
	// 		AdditionalStorageCount: to.Ptr[int32](1),
	// 		AvailableStorageSizeInGbs: to.Ptr[int32](1000),
	// 		ComputeCount: to.Ptr[int32](100),
	// 		CPUCount: to.Ptr[int32](10),
	// 		CustomerContacts: []*armoracledatabase.CustomerContact{
	// 			{
	// 				Email: to.Ptr("noreply@oracle.com"),
	// 		}},
	// 		DataStorageSizeInTbs: to.Ptr[int32](10),
	// 		DbNodeStorageSizeInGbs: to.Ptr[int32](10),
	// 		DbServerVersion: to.Ptr("19.0.0.0"),
	// 		DisplayName: to.Ptr("infra 1"),
	// 		EstimatedPatchingTime: &armoracledatabase.EstimatedPatchingTime{
	// 			EstimatedDbServerPatchingTime: to.Ptr[int32](3000),
	// 			EstimatedNetworkSwitchesPatchingTime: to.Ptr[int32](3000),
	// 			EstimatedStorageServerPatchingTime: to.Ptr[int32](3000),
	// 			TotalEstimatedPatchingTime: to.Ptr[int32](3000),
	// 		},
	// 		LastMaintenanceRunID: to.Ptr("ocid1..aaaaa"),
	// 		LifecycleDetails: to.Ptr("none"),
	// 		MaintenanceWindow: &armoracledatabase.MaintenanceWindow{
	// 			CustomActionTimeoutInMins: to.Ptr[int32](120),
	// 			DaysOfWeek: []*armoracledatabase.DayOfWeek{
	// 				{
	// 					Name: to.Ptr(armoracledatabase.DayOfWeekNameMonday),
	// 			}},
	// 			HoursOfDay: []*int32{
	// 				to.Ptr[int32](0)},
	// 				IsCustomActionTimeoutEnabled: to.Ptr(true),
	// 				IsMonthlyPatchingEnabled: to.Ptr(true),
	// 				LeadTimeInWeeks: to.Ptr[int32](0),
	// 				Months: []*armoracledatabase.Month{
	// 					{
	// 						Name: to.Ptr(armoracledatabase.MonthNameJanuary),
	// 				}},
	// 				PatchingMode: to.Ptr(armoracledatabase.PatchingModeRolling),
	// 				Preference: to.Ptr(armoracledatabase.PreferenceNoPreference),
	// 				WeeksOfMonth: []*int32{
	// 					to.Ptr[int32](0)},
	// 				},
	// 				MaxCPUCount: to.Ptr[int32](100),
	// 				MaxDataStorageInTbs: to.Ptr[float64](1000),
	// 				MaxDbNodeStorageSizeInGbs: to.Ptr[int32](10),
	// 				MaxMemoryInGbs: to.Ptr[int32](1000),
	// 				MemorySizeInGbs: to.Ptr[int32](100),
	// 				MonthlyDbServerVersion: to.Ptr("aaaa"),
	// 				MonthlyStorageServerVersion: to.Ptr("aaaa"),
	// 				NextMaintenanceRunID: to.Ptr("ocid1..aaaaaa"),
	// 				OciURL: to.Ptr("https://url"),
	// 				Ocid: to.Ptr("ocid1..aaaaaa"),
	// 				ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
	// 				Shape: to.Ptr("EXADATA.X9M"),
	// 				StorageCount: to.Ptr[int32](10),
	// 				StorageServerVersion: to.Ptr("0.0"),
	// 				TimeCreated: to.Ptr("2023-02-01T01:01:00"),
	// 				TotalStorageSizeInGbs: to.Ptr[int32](1000),
	// 			},
	// 			Zones: []*string{
	// 				to.Ptr("1")},
	// 			}
}
