package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7b98b36e4023331545051284d8500adf98f02fe/specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-03-01/examples/SnapshotPolicies_List.json
func ExampleSnapshotPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSnapshotPoliciesClient().NewListPager("myRG", "account1", nil)
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
		// page.SnapshotPoliciesList = armnetapp.SnapshotPoliciesList{
		// 	Value: []*armnetapp.SnapshotPolicy{
		// 		{
		// 			Name: to.Ptr("account1/snapshotPolicy1"),
		// 			Type: to.Ptr("Microsoft.NetApp/netAppAccounts/snapshotPolicies"),
		// 			ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/snapshotPolicies/snapshotPolicy1"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armnetapp.SnapshotPolicyProperties{
		// 				DailySchedule: &armnetapp.DailySchedule{
		// 					Hour: to.Ptr[int32](14),
		// 					Minute: to.Ptr[int32](30),
		// 					SnapshotsToKeep: to.Ptr[int32](4),
		// 				},
		// 				Enabled: to.Ptr(true),
		// 				HourlySchedule: &armnetapp.HourlySchedule{
		// 					Minute: to.Ptr[int32](50),
		// 					SnapshotsToKeep: to.Ptr[int32](2),
		// 				},
		// 				MonthlySchedule: &armnetapp.MonthlySchedule{
		// 					DaysOfMonth: to.Ptr("10,11,12"),
		// 					Hour: to.Ptr[int32](14),
		// 					Minute: to.Ptr[int32](15),
		// 					SnapshotsToKeep: to.Ptr[int32](5),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				WeeklySchedule: &armnetapp.WeeklySchedule{
		// 					Day: to.Ptr("Wednesday"),
		// 					Hour: to.Ptr[int32](14),
		// 					Minute: to.Ptr[int32](45),
		// 					SnapshotsToKeep: to.Ptr[int32](3),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
