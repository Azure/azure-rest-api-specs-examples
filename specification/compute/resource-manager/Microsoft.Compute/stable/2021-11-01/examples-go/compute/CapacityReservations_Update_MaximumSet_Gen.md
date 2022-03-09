Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.5.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/CapacityReservations_Update_MaximumSet_Gen.json
func ExampleCapacityReservationsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewCapacityReservationsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<capacity-reservation-group-name>",
		"<capacity-reservation-name>",
		armcompute.CapacityReservationUpdate{
			Tags: map[string]*string{
				"key4974": to.StringPtr("aaaaaaaaaaaaaaaa"),
			},
			Properties: &armcompute.CapacityReservationProperties{
				InstanceView: &armcompute.CapacityReservationInstanceView{
					Statuses: []*armcompute.InstanceViewStatus{
						{
							Code:          to.StringPtr("<code>"),
							DisplayStatus: to.StringPtr("<display-status>"),
							Level:         armcompute.StatusLevelTypesInfo.ToPtr(),
							Message:       to.StringPtr("<message>"),
							Time:          to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
						}},
					UtilizationInfo: &armcompute.CapacityReservationUtilization{},
				},
			},
			SKU: &armcompute.SKU{
				Name:     to.StringPtr("<name>"),
				Capacity: to.Int64Ptr(7),
				Tier:     to.StringPtr("<tier>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CapacityReservationsClientUpdateResult)
}
```
