Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetapp%2Farmnetapp%2Fv0.4.0/sdk/resourcemanager/netapp/armnetapp/README.md) on how to add the SDK to your project and authenticate.

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
		return
	}
	ctx := context.Background()
	client, err := armnetapp.NewSnapshotPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<snapshot-policy-name>",
		armnetapp.SnapshotPolicy{
			Location: to.Ptr("<location>"),
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
					DaysOfMonth:     to.Ptr("<days-of-month>"),
					Hour:            to.Ptr[int32](14),
					Minute:          to.Ptr[int32](15),
					SnapshotsToKeep: to.Ptr[int32](5),
				},
				WeeklySchedule: &armnetapp.WeeklySchedule{
					Day:             to.Ptr("<day>"),
					Hour:            to.Ptr[int32](14),
					Minute:          to.Ptr[int32](45),
					SnapshotsToKeep: to.Ptr[int32](3),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
