package armdataboxedge_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/BandwidthSchedulePut.json
func ExampleBandwidthSchedulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewBandwidthSchedulesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<device-name>",
		"<name>",
		"<resource-group-name>",
		armdataboxedge.BandwidthSchedule{
			Properties: &armdataboxedge.BandwidthScheduleProperties{
				Days: []*armdataboxedge.DayOfWeek{
					armdataboxedge.DayOfWeek("Sunday").ToPtr(),
					armdataboxedge.DayOfWeek("Monday").ToPtr()},
				RateInMbps: to.Int32Ptr(100),
				Start:      to.StringPtr("<start>"),
				Stop:       to.StringPtr("<stop>"),
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
	log.Printf("Response result: %#v\n", res.BandwidthSchedulesClientCreateOrUpdateResult)
}
