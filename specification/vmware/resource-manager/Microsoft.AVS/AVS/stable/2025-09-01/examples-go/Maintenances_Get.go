package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/Maintenances_Get.json
func ExampleMaintenancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMaintenancesClient().Get(ctx, "group1", "cloud1", "maintenance1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armavs.MaintenancesClientGetResponse{
	// 	Maintenance: &armavs.Maintenance{
	// 		ID: to.Ptr("/subscriptions/00000001-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/maintenances/maintenance1"),
	// 		Name: to.Ptr("maintenance1"),
	// 		Type: to.Ptr("Microsoft.AVS/privateClouds/maintenances"),
	// 		Properties: &armavs.MaintenanceProperties{
	// 			Component: to.Ptr(armavs.MaintenanceTypeVCSA),
	// 			DisplayName: to.Ptr("vcsa 7.0 upgrade"),
	// 			ClusterID: to.Ptr[int32](1),
	// 			InfoLink: to.Ptr("https://vmwarekb/arcticle"),
	// 			Impact: to.Ptr("This upgrade will update your vcsa to 7.0. Control plance performance will be impacted for the duration"),
	// 			ScheduledStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-12T11:00:11.830Z"); return t}()),
	// 			EstimatedDurationInMinutes: to.Ptr[int64](960),
	// 			State: &armavs.MaintenanceState{
	// 				Message: to.Ptr("CD rom mounted"),
	// 				StartedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-12T11:00:11.830Z"); return t}()),
	// 				EndedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-12T11:00:11.830Z"); return t}()),
	// 				Name: to.Ptr(armavs.MaintenanceStateNameScheduled),
	// 			},
	// 			ProvisioningState: to.Ptr(armavs.MaintenanceProvisioningStateSucceeded),
	// 			ScheduledByMicrosoft: to.Ptr(true),
	// 			Operations: []armavs.MaintenanceManagementOperationClassification{
	// 				&armavs.ScheduleOperation{
	// 					Kind: to.Ptr(armavs.MaintenanceManagementOperationKindSchedule),
	// 					IsDisabled: to.Ptr(true),
	// 					DisabledReason: to.Ptr("Critical upgrade"),
	// 					Constraints: []armavs.ScheduleOperationConstraintClassification{
	// 						&armavs.SchedulingWindow{
	// 							Kind: to.Ptr(armavs.ScheduleOperationConstraintKindSchedulingWindow),
	// 							StartsAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-16T06:21:31.961Z"); return t}()),
	// 							EndsAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-16T06:21:31.961Z"); return t}()),
	// 						},
	// 						&armavs.AvailableWindowForMaintenanceWhileScheduleOperation{
	// 							Kind: to.Ptr(armavs.ScheduleOperationConstraintKindAvailableWindowForMaintenanceWhileScheduleOperation),
	// 							StartsAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-16T06:21:31.961Z"); return t}()),
	// 							EndsAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-16T06:21:31.961Z"); return t}()),
	// 						},
	// 						&armavs.BlockedWhileScheduleOperation{
	// 							Kind: to.Ptr(armavs.ScheduleOperationConstraintKindBlockedWhileScheduleOperation),
	// 							Category: to.Ptr(armavs.BlockedDatesConstraintCategoryHiPriorityEvent),
	// 							TimeRanges: []*armavs.BlockedDatesConstraintTimeRange{
	// 								{
	// 									StartsAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-16T06:21:31.961Z"); return t}()),
	// 									EndsAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-16T06:21:31.961Z"); return t}()),
	// 									Reason: to.Ptr("2024 Summer Opening Ceremony"),
	// 								},
	// 							},
	// 						},
	// 						&armavs.BlockedWhileScheduleOperation{
	// 							Kind: to.Ptr(armavs.ScheduleOperationConstraintKindBlockedWhileScheduleOperation),
	// 							Category: to.Ptr(armavs.BlockedDatesConstraintCategoryQuotaExhausted),
	// 							TimeRanges: []*armavs.BlockedDatesConstraintTimeRange{
	// 								{
	// 									StartsAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-14T17:03:28.609Z"); return t}()),
	// 									EndsAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-15T17:03:28.609Z"); return t}()),
	// 									Reason: to.Ptr("No slots available"),
	// 								},
	// 								{
	// 									StartsAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-13T17:03:28.609Z"); return t}()),
	// 									EndsAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-26T17:03:28.609Z"); return t}()),
	// 									Reason: to.Ptr("No slots available"),
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				&armavs.RescheduleOperation{
	// 					Kind: to.Ptr(armavs.MaintenanceManagementOperationKindReschedule),
	// 					IsDisabled: to.Ptr(true),
	// 					DisabledReason: to.Ptr("Critical Update"),
	// 					Constraints: []armavs.RescheduleOperationConstraintClassification{
	// 						&armavs.AvailableWindowForMaintenanceWhileRescheduleOperation{
	// 							Kind: to.Ptr(armavs.RescheduleOperationConstraintKindAvailableWindowForMaintenanceWhileRescheduleOperation),
	// 							StartsAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-16T06:21:31.961Z"); return t}()),
	// 							EndsAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-16T06:21:31.961Z"); return t}()),
	// 						},
	// 						&armavs.BlockedWhileRescheduleOperation{
	// 							Kind: to.Ptr(armavs.RescheduleOperationConstraintKindBlockedWhileRescheduleOperation),
	// 							Category: to.Ptr(armavs.BlockedDatesConstraintCategoryHiPriorityEvent),
	// 							TimeRanges: []*armavs.BlockedDatesConstraintTimeRange{
	// 								{
	// 									StartsAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-10T06:21:31.961Z"); return t}()),
	// 									EndsAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-19T06:21:31.961Z"); return t}()),
	// 									Reason: to.Ptr("US General Election 2024"),
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				&armavs.MaintenanceReadinessRefreshOperation{
	// 					Kind: to.Ptr(armavs.MaintenanceManagementOperationKindMaintenanceReadinessRefresh),
	// 					IsDisabled: to.Ptr(false),
	// 					Status: to.Ptr(armavs.MaintenanceReadinessRefreshOperationStatusInProgress),
	// 					RefreshedByMicrosoft: to.Ptr(true),
	// 				},
	// 			},
	// 			MaintenanceReadiness: &armavs.MaintenanceReadiness{
	// 				Type: to.Ptr(armavs.MaintenanceCheckTypePrecheck),
	// 				Status: to.Ptr(armavs.MaintenanceReadinessStatusNotReady),
	// 				Message: to.Ptr("Some checks failed"),
	// 				FailedChecks: []*armavs.MaintenanceFailedCheck{
	// 					{
	// 						Name: to.Ptr("HostMaintPrecheck"),
	// 						ImpactedResources: []*armavs.ImpactedMaintenanceResource{
	// 							{
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster2"),
	// 								Errors: []*armavs.ImpactedMaintenanceResourceError{
	// 									{
	// 										ErrorCode: to.Ptr("EVCINTERNALEPR_CDUNMOUNT"),
	// 										Name: to.Ptr("CD-ROM Mounted"),
	// 										Details: to.Ptr("Failed to unmount CD ROM  mounted on vm"),
	// 										ResolutionSteps: []*string{
	// 											to.Ptr("Please remove CDROM /usr/lib/vmware/isoimages/windows.iso mounted in VD000535 VM present in esx10-r06.p01.dummy.com Host"),
	// 										},
	// 										ActionRequired: to.Ptr(true),
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-16T06:21:31.961Z"); return t}()),
	// 			},
	// 		},
	// 	},
	// }
}
