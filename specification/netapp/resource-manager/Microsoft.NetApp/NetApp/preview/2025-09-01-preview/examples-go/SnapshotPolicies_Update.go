package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/SnapshotPolicies_Update.json
func ExampleSnapshotPoliciesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSnapshotPoliciesClient().BeginUpdate(ctx, "myRG", "account1", "snapshotPolicyName", armnetapp.SnapshotPolicyPatch{
		Location: to.Ptr("eastus"),
		Properties: &armnetapp.SnapshotPolicyProperties{
			DailySchedule: &armnetapp.DailySchedule{
				Hour:            to.Ptr[int32](14),
				Minute:          to.Ptr[int32](30),
				SnapshotsToKeep: to.Ptr[int32](4),
			},
			Enabled: to.Ptr(true),
			HourlySchedule: &armnetapp.HourlySchedule{
				Minute:          to.Ptr[int32](50),
				SnapshotsToKeep: to.Ptr[int32](2),
			},
			MonthlySchedule: &armnetapp.MonthlySchedule{
				DaysOfMonth:     to.Ptr("10,11,12"),
				Hour:            to.Ptr[int32](14),
				Minute:          to.Ptr[int32](15),
				SnapshotsToKeep: to.Ptr[int32](5),
			},
			WeeklySchedule: &armnetapp.WeeklySchedule{
				Day:             to.Ptr("Wednesday"),
				Hour:            to.Ptr[int32](14),
				Minute:          to.Ptr[int32](45),
				SnapshotsToKeep: to.Ptr[int32](3),
			},
		},
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
	// res = armnetapp.SnapshotPoliciesClientUpdateResponse{
	// 	SnapshotPolicy: &armnetapp.SnapshotPolicy{
	// 		Name: to.Ptr("account1/snapshotPolicy1"),
	// 		Type: to.Ptr("Microsoft.NetApp/netAppAccounts/snapshotPolicies"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/snapshotPolicies/snapshotPolicy1"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armnetapp.SnapshotPolicyProperties{
	// 			DailySchedule: &armnetapp.DailySchedule{
	// 				Hour: to.Ptr[int32](14),
	// 				Minute: to.Ptr[int32](30),
	// 				SnapshotsToKeep: to.Ptr[int32](4),
	// 			},
	// 			Enabled: to.Ptr(true),
	// 			HourlySchedule: &armnetapp.HourlySchedule{
	// 				Minute: to.Ptr[int32](50),
	// 				SnapshotsToKeep: to.Ptr[int32](2),
	// 			},
	// 			MonthlySchedule: &armnetapp.MonthlySchedule{
	// 				DaysOfMonth: to.Ptr("10,11,12"),
	// 				Hour: to.Ptr[int32](14),
	// 				Minute: to.Ptr[int32](15),
	// 				SnapshotsToKeep: to.Ptr[int32](5),
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			WeeklySchedule: &armnetapp.WeeklySchedule{
	// 				Day: to.Ptr("Wednesday"),
	// 				Hour: to.Ptr[int32](14),
	// 				Minute: to.Ptr[int32](45),
	// 				SnapshotsToKeep: to.Ptr[int32](3),
	// 			},
	// 		},
	// 	},
	// }
}
