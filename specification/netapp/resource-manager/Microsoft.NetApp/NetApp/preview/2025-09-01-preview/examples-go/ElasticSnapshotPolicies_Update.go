package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/ElasticSnapshotPolicies_Update.json
func ExampleElasticSnapshotPoliciesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewElasticSnapshotPoliciesClient().BeginUpdate(ctx, "myRG", "account1", "snapshotPolicyName", armnetapp.ElasticSnapshotPolicyUpdate{
		Properties: &armnetapp.ElasticSnapshotPolicyUpdateProperties{
			PolicyStatus: to.Ptr(armnetapp.PolicyStatusEnabled),
			HourlySchedule: &armnetapp.ElasticSnapshotPolicyHourlySchedule{
				SnapshotsToKeep: to.Ptr[int32](2),
				Minute:          to.Ptr[int32](50),
			},
			DailySchedule: &armnetapp.ElasticSnapshotPolicyDailySchedule{
				SnapshotsToKeep: to.Ptr[int32](4),
				Hour:            to.Ptr[int32](14),
				Minute:          to.Ptr[int32](30),
			},
			WeeklySchedule: &armnetapp.ElasticSnapshotPolicyWeeklySchedule{
				SnapshotsToKeep: to.Ptr[int32](3),
				Days: []*armnetapp.DayOfWeek{
					to.Ptr(armnetapp.DayOfWeekWednesday),
				},
				Hour:   to.Ptr[int32](14),
				Minute: to.Ptr[int32](45),
			},
			MonthlySchedule: &armnetapp.ElasticSnapshotPolicyMonthlySchedule{
				SnapshotsToKeep: to.Ptr[int32](5),
				DaysOfMonth: []*int32{
					to.Ptr[int32](10),
					to.Ptr[int32](11),
					to.Ptr[int32](12),
				},
				Hour:   to.Ptr[int32](14),
				Minute: to.Ptr[int32](15),
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
	// res = armnetapp.ElasticSnapshotPoliciesClientUpdateResponse{
	// 	ElasticSnapshotPolicy: &armnetapp.ElasticSnapshotPolicy{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/elasticAccounts/account1/elasticSnapshotPolicies/snapshotPolicy1"),
	// 		Name: to.Ptr("account1/snapshotPolicy1"),
	// 		Type: to.Ptr("Microsoft.NetApp/elasticAccounts/elasticSnapshotPolicies"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armnetapp.ElasticSnapshotPolicyProperties{
	// 			PolicyStatus: to.Ptr(armnetapp.PolicyStatusEnabled),
	// 			HourlySchedule: &armnetapp.ElasticSnapshotPolicyHourlySchedule{
	// 				SnapshotsToKeep: to.Ptr[int32](2),
	// 				Minute: to.Ptr[int32](50),
	// 			},
	// 			DailySchedule: &armnetapp.ElasticSnapshotPolicyDailySchedule{
	// 				SnapshotsToKeep: to.Ptr[int32](4),
	// 				Hour: to.Ptr[int32](14),
	// 				Minute: to.Ptr[int32](30),
	// 			},
	// 			WeeklySchedule: &armnetapp.ElasticSnapshotPolicyWeeklySchedule{
	// 				SnapshotsToKeep: to.Ptr[int32](3),
	// 				Days: []*armnetapp.DayOfWeek{
	// 					to.Ptr(armnetapp.DayOfWeekWednesday),
	// 				},
	// 				Hour: to.Ptr[int32](14),
	// 				Minute: to.Ptr[int32](45),
	// 			},
	// 			MonthlySchedule: &armnetapp.ElasticSnapshotPolicyMonthlySchedule{
	// 				SnapshotsToKeep: to.Ptr[int32](5),
	// 				DaysOfMonth: []*int32{
	// 					to.Ptr[int32](10),
	// 					to.Ptr[int32](11),
	// 					to.Ptr[int32](12),
	// 				},
	// 				Hour: to.Ptr[int32](14),
	// 				Minute: to.Ptr[int32](15),
	// 			},
	// 			ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
