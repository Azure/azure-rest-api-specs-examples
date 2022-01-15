Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetapp%2Farmnetapp%2Fv0.2.0/sdk/resourcemanager/netapp/armnetapp/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp"
)

// x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-08-01/examples/SnapshotPolicies_Create.json
func ExampleSnapshotPoliciesClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetapp.NewSnapshotPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<snapshot-policy-name>",
		armnetapp.SnapshotPolicy{
			Location: to.StringPtr("<location>"),
			Properties: &armnetapp.SnapshotPolicyProperties{
				DailySchedule: &armnetapp.DailySchedule{
					Hour:            to.Int32Ptr(14),
					Minute:          to.Int32Ptr(30),
					SnapshotsToKeep: to.Int32Ptr(4),
				},
				Enabled: to.BoolPtr(true),
				HourlySchedule: &armnetapp.HourlySchedule{
					Minute:          to.Int32Ptr(50),
					SnapshotsToKeep: to.Int32Ptr(2),
				},
				MonthlySchedule: &armnetapp.MonthlySchedule{
					DaysOfMonth:     to.StringPtr("<days-of-month>"),
					Hour:            to.Int32Ptr(14),
					Minute:          to.Int32Ptr(15),
					SnapshotsToKeep: to.Int32Ptr(5),
				},
				WeeklySchedule: &armnetapp.WeeklySchedule{
					Day:             to.StringPtr("<day>"),
					Hour:            to.Int32Ptr(14),
					Minute:          to.Int32Ptr(45),
					SnapshotsToKeep: to.Int32Ptr(3),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SnapshotPoliciesClientCreateResult)
}
```
