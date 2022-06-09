```go
package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-10-01/examples/SnapshotPolicies_Create.json
func ExampleSnapshotPoliciesClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetapp.NewSnapshotPoliciesClient("D633CC2E-722B-4AE1-B636-BBD9E4C60ED9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"myRG",
		"account1",
		"snapshotPolicyName",
		armnetapp.SnapshotPolicy{
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
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetapp%2Farmnetapp%2Fv1.0.0/sdk/resourcemanager/netapp/armnetapp/README.md) on how to add the SDK to your project and authenticate.
